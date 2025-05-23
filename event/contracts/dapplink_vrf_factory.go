package contracts

import (
	"math/big"
	"time"

	"github.com/google/uuid"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3-contracts/vrf-node/bindings/vrf"
	"github.com/the-web3-contracts/vrf-node/database"
	"github.com/the-web3-contracts/vrf-node/database/event"
	"github.com/the-web3-contracts/vrf-node/database/worker"
)

type DappLinkVrfFactory struct {
	DlVrfFactoryAbi    *abi.ABI
	DlVrfFactoryFilter *vrf.DappLinkVRFFactoryFilterer
}

func NewDappLinkVrfFactory() (*DappLinkVrfFactory, error) {
	dappLinkVrfFactoryAbi, err := vrf.DappLinkVRFFactoryMetaData.GetAbi()
	if err != nil {
		log.Error("get dapplink vrf factory abi fail", "err", err)
		return nil, err
	}
	dappLinkVrfFactoryFilterer, err := vrf.NewDappLinkVRFFactoryFilterer(common.Address{}, nil)
	if err != nil {
		log.Error("new dapplink vrf factory filter fail", "err", err)
		return nil, err
	}
	return &DappLinkVrfFactory{
		DlVrfFactoryAbi:    dappLinkVrfFactoryAbi,
		DlVrfFactoryFilter: dappLinkVrfFactoryFilterer,
	}, nil
}

func (dvff *DappLinkVrfFactory) ProcessDappLinkVrfFactoryEvent(db *database.DB, dappLinkVrfFactoryAddres string, startBlock, endBlock *big.Int) ([]worker.PoxyCreated, error) {
	var proxyCreatedList []worker.PoxyCreated
	contactFilter := event.ContractEvent{ContractAddress: common.HexToAddress(dappLinkVrfFactoryAddres)}
	contractEventList, err := db.ContractEvent.ContractEventsWithFilter(contactFilter, startBlock, endBlock)
	if err != nil {
		log.Error("query contacts event fail", "err", err)
		return proxyCreatedList, err
	}
	for _, contractEvent := range contractEventList {
		if contractEvent.EventSignature.String() == dvff.DlVrfFactoryAbi.Events["ProxyCreated"].ID.String() {
			proxyCreated, err := dvff.DlVrfFactoryFilter.ParseProxyCreated(*contractEvent.RLPLog)
			if err != nil {
				log.Error("proxy created fail", "err", err)
				return proxyCreatedList, err
			}
			log.Info("proxy created event", "MintProxyAddress", proxyCreated.MintProxyAddress)
			pc := worker.PoxyCreated{
				GUID:         uuid.New(),
				ProxyAddress: proxyCreated.MintProxyAddress,
				Timestamp:    uint64(time.Now().Unix()),
			}
			proxyCreatedList = append(proxyCreatedList, pc)
		}
	}
	return proxyCreatedList, nil
}
