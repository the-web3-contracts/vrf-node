package synchronizer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"

	"github.com/the-web3-contracts/vrf-node/common/tasks"
	"github.com/the-web3-contracts/vrf-node/config"
	"github.com/the-web3-contracts/vrf-node/database"
	common2 "github.com/the-web3-contracts/vrf-node/database/common"
	"github.com/the-web3-contracts/vrf-node/database/event"
	"github.com/the-web3-contracts/vrf-node/database/utils"
	"github.com/the-web3-contracts/vrf-node/synchronizer/node"
	"github.com/the-web3-contracts/vrf-node/synchronizer/retry"
)

type Synchronizer struct {
	ethClient       node.EthClient
	db              *database.DB
	headerTraversal *node.HeaderTraversal
	headers         []types.Header
	latestHeader    *types.Header
	chainCfg        *config.ChainConfig
	resourceCtx     context.Context
	resourceCancel  context.CancelFunc
	tasks           tasks.Group
}

func NewSynchronizer(cfg *config.Config, db *database.DB, client node.EthClient, shutdown context.CancelCauseFunc) (*Synchronizer, error) {
	latestHeader, err := db.Blocks.LatestBlockHeader()
	if err != nil {
		log.Error("query latest block header fail", "err", err)
		return nil, err
	}
	var fromHeader *types.Header
	if latestHeader != nil {
		fromHeader = latestHeader.RLPHeader.Header()
	} else if cfg.Chain.StartingHeight > 0 {
		header, err := client.BlockHeaderByNumber(big.NewInt(int64(cfg.Chain.StartingHeight)))
		if err != nil {
			log.Error("get block from chain fail", "err", err)
			return nil, err
		}
		fromHeader = header
	} else {
		log.Info("no eth block indexed state")
	}
	headerTraversal := node.NewHeaderTraversal(client, fromHeader, big.NewInt(0), cfg.Chain.ChainId)

	resCtx, resCancel := context.WithCancel(context.Background())

	return &Synchronizer{
		ethClient:       client,
		db:              db,
		headerTraversal: headerTraversal,
		latestHeader:    fromHeader,
		chainCfg:        &cfg.Chain,
		resourceCtx:     resCtx,
		resourceCancel:  resCancel,
		tasks: tasks.Group{
			HandleCrit: func(err error) {
				shutdown(fmt.Errorf("critical error in Synchronizer %w", err))
			},
		},
	}, err
}

func (syncer *Synchronizer) Start() error {
	tickerSyncer := time.NewTicker(syncer.chainCfg.MainLoopInterval)
	syncer.tasks.Go(func() error {
		for range tickerSyncer.C {
			newHeaders, err := syncer.headerTraversal.NextHeaders(syncer.chainCfg.BlockStep)
			if err != nil {
				log.Error("error querying for header", "err", err)
				continue
			} else if len(newHeaders) == 0 {
				log.Warn("no new header, sync at head")
			} else {
				syncer.headers = newHeaders
			}

			latestHeader := syncer.headerTraversal.LatestHeader()
			if latestHeader != nil {
				log.Info("Latest header", "latestHeader", latestHeader.Number)
			}
			err = syncer.processBatch(syncer.headers, syncer.chainCfg)
			if err == nil {
				syncer.headers = nil
			}
		}
		return nil
	})
	return nil
}

func (syncer *Synchronizer) processBatch(headers []types.Header, chainCfg *config.ChainConfig) error {
	if len(headers) == 0 {
		return nil
	}
	firstHeader, lastHead := headers[0], headers[len(headers)-1]
	log.Info("sync batch", "size", len(headers), "startBlock", firstHeader.Number, "endBlock", lastHead.Number)

	headerMap := make(map[common.Hash]*types.Header, len(headers))
	for i := range headers {
		header := headers[i]
		headerMap[header.Hash()] = &header
	}

	var addressList []common.Address

	filterQuery := ethereum.FilterQuery{FromBlock: firstHeader.Number, ToBlock: lastHead.Number, Addresses: addressList}

	logs, err := syncer.ethClient.FilterLogs(filterQuery)
	if err != nil {
		log.Error("filter logs fail", "err", err)
		return err
	}

	if logs.ToBlockHeader.Number.Cmp(lastHead.Number) != 0 {
		return errors.New("mismatch in filter#toBlock numer")
	} else if logs.ToBlockHeader.Hash() != lastHead.Hash() {
		return errors.New("mismatch in filter#toBlock hash")
	}

	if len(logs.Logs) > 0 {
		log.Info("detected logs", "size", len(logs.Logs))
	}

	blockHeaders := make([]common2.BlockHeader, 0, len(headers))

	for i := range headers {
		if headers[i].Number == nil {
			continue
		}
		bHeader := common2.BlockHeader{
			Hash:       headers[i].Hash(),
			ParentHash: headers[i].ParentHash,
			Number:     headers[i].Number,
			Timestamp:  headers[i].Time,
			RLPHeader:  (*utils.RLPHeader)(&headers[i]),
		}
		blockHeaders = append(blockHeaders, bHeader)
	}

	chainContractEvent := make([]event.ContractEvent, 0, len(headers))
	for i := range logs.Logs {
		logEvent := logs.Logs[i]
		if _, ok := headerMap[logEvent.BlockHash]; !ok {
			continue
		}
		timestamp := headerMap[logEvent.BlockHash].Time
		chainContractEvent[i] = event.ContractEventFromLog(&logs.Logs[i], timestamp)
	}

	retryStrategy := &retry.ExponentialStrategy{Min: 1000, Max: 20_000, MaxJitter: 250}
	if _, err := retry.Do[interface{}](syncer.resourceCtx, 10, retryStrategy, func() (interface{}, error) {
		if err := syncer.db.Transaction(func(tx *database.DB) error {
			if err := tx.Blocks.StoreBlockHeaders(blockHeaders); err != nil {
				return err
			}
			if err := tx.ContractEvent.StoreContractEvents(chainContractEvent); err != nil {
				return err
			}
			return nil
		}); err != nil {
			log.Info("unable to persist batch", err)
			return nil, fmt.Errorf("unable to persist batch: %w", err)
		}
		return nil, nil
	}); err != nil {
		return err
	}
	return nil
}

func (syncer *Synchronizer) Close() error {
	syncer.resourceCancel()
	return nil
}
