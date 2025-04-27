package vrf_node

import (
	"context"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3-contracts/vrf-node/config"
	"github.com/the-web3-contracts/vrf-node/database"
	"github.com/the-web3-contracts/vrf-node/synchronizer"
	"github.com/the-web3-contracts/vrf-node/synchronizer/node"
)

type VrfNode struct {
	db *database.DB

	synchronizer *synchronizer.Synchronizer

	shutdown context.CancelCauseFunc
	stopped  atomic.Bool
}

func NewVrfNode(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*VrfNode, error) {
	ethClient, err := node.DialEthClient(ctx, cfg.Chain.ChainRpcUrl)
	if err != nil {
		log.Error("new eth client fail", "err", err)
		return nil, err
	}

	db, err := database.NewDB(ctx, cfg.MasterDB)
	if err != nil {
		log.Error("new database fail", "err", err)
		return nil, err
	}

	syncer, err := synchronizer.NewSynchronizer(cfg, db, ethClient, shutdown)
	if err != nil {
		log.Error("new synchronizer fail", "err", err)
		return nil, err
	}

	return &VrfNode{
		db:           db,
		synchronizer: syncer,
		shutdown:     shutdown,
	}, nil
}

func (vn *VrfNode) Start(ctx context.Context) error {
	err := vn.synchronizer.Start()
	if err != nil {
		return err
	}
	return nil
}

func (vn *VrfNode) Stop(ctx context.Context) error {
	err := vn.synchronizer.Close()
	if err != nil {
		return err
	}
	return nil
}

func (vn *VrfNode) Stopped() bool {
	return vn.stopped.Load()
}
