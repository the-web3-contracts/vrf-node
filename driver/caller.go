package driver

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3-contracts/vrf-node/bindings/vrf"
	"github.com/the-web3-contracts/vrf-node/txmgr"
)

var (
	errMaxPriorityFeePerGasNotFound = errors.New(
		"Method eth_maxPriorityFeePerGas not found",
	)
	FallbackGasTipCap = big.NewInt(1500000000)
)

type CallerConfig struct {
	ChainClient               *ethclient.Client
	ChainId                   *big.Int
	DappLinkVrfManagerAddress common.Address
	CallerAddress             common.Address
	PrivateKey                *ecdsa.PrivateKey
	NumConfirmations          uint64
	SafeAbortNonceToLowCount  uint64
}

type Caller struct {
	Ctx                     context.Context
	Cfg                     *CallerConfig
	DappLinkVrfContracts    *vrf.DappLinkVRFManager
	RawDappLinkVrfContracts *bind.BoundContract
	DappLinkVrfContractsAbi *abi.ABI
	TxMrg                   txmgr.TxManager
}

func NewCaller(Ctx context.Context, Cfg *CallerConfig) (*Caller, error) {
	dappLinkVrfContracts, err := vrf.NewDappLinkVRFManager(Cfg.DappLinkVrfManagerAddress, Cfg.ChainClient)
	if err != nil {
		log.Error("New DappLink Vrf Manager Fail", "err", err)
		return nil, err
	}

	parsed, err := abi.JSON(strings.NewReader(vrf.DappLinkVRFManagerMetaData.ABI))
	if err != nil {
		log.Error("abi parsed fail", "err", err)
		return nil, err
	}

	dappLinkVrfContractsAbi, err := vrf.DappLinkVRFManagerMetaData.GetAbi()
	if err != nil {
		log.Error("get abi fail", "err", err)
		return nil, err
	}

	rawDappLinkVrfContracts := bind.NewBoundContract(Cfg.DappLinkVrfManagerAddress, parsed, Cfg.ChainClient, Cfg.ChainClient, Cfg.ChainClient)

	txManagerConfig := txmgr.Config{
		ResubmissionTimeout:       time.Second * 5,
		ReceiptQueryInterval:      time.Second,
		NumConfirmations:          Cfg.NumConfirmations,
		SafeAbortNonceTooLowCount: Cfg.SafeAbortNonceToLowCount,
	}

	txManager := txmgr.NewSimpleTxManager(txManagerConfig, Cfg.ChainClient)

	return &Caller{
		Ctx:                     Ctx,
		Cfg:                     Cfg,
		DappLinkVrfContracts:    dappLinkVrfContracts,
		RawDappLinkVrfContracts: rawDappLinkVrfContracts,
		DappLinkVrfContractsAbi: dappLinkVrfContractsAbi,
		TxMrg:                   txManager,
	}, nil
}

func (caller *Caller) isMaxPriorityFeePerGasNotFoundError(err error) bool {
	return strings.Contains(err.Error(), errMaxPriorityFeePerGasNotFound.Error())
}

func (caller *Caller) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return caller.Cfg.ChainClient.SendTransaction(ctx, tx)
}

func (caller *Caller) UpdateGasPrice(ctx context.Context, tx *types.Transaction) (*types.Transaction, error) {
	var opts *bind.TransactOpts
	var err error
	opts, err = bind.NewKeyedTransactorWithChainID(caller.Cfg.PrivateKey, caller.Cfg.ChainId)
	if err != nil {
		log.Error("new keyed transactor with chain id fail", "err", err)
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(tx.Nonce())
	opts.NoSend = true

	finalTx, err := caller.RawDappLinkVrfContracts.RawTransact(opts, tx.Data())
	switch {
	case err == nil:
		return finalTx, nil

	case caller.isMaxPriorityFeePerGasNotFoundError(err):
		log.Info("Don't support priority fee")
		opts.GasTipCap = FallbackGasTipCap
		return caller.RawDappLinkVrfContracts.RawTransact(opts, tx.Data())

	default:
		return nil, err
	}
}

func (caller *Caller) fulfillRandomWords(ctx context.Context, requestId *big.Int, randomList []*big.Int) (*types.Transaction, error) {
	nonce, err := caller.Cfg.ChainClient.NonceAt(ctx, caller.Cfg.CallerAddress, nil)
	if err != nil {
		log.Error("get eth nonce fail", "err", err)
		return nil, err
	}

	opts, err := bind.NewKeyedTransactorWithChainID(caller.Cfg.PrivateKey, caller.Cfg.ChainId)
	if err != nil {
		log.Error("new keyed transactor with chain id fail", "err", err)
		return nil, err
	}
	opts.Context = ctx
	opts.Nonce = new(big.Int).SetUint64(nonce)
	opts.NoSend = true

	var msgHash [32]byte
	var blsParam vrf.IBLSApkRegistryVrfNoSignerAndSignature

	tx, err := caller.DappLinkVrfContracts.FulfillRandomWords(opts, requestId, randomList, msgHash, big.NewInt(100), blsParam)
	if err != nil {
		log.Error("fulfill random words fail", "err", err)
		return nil, err
	}

	switch {
	case err == nil:
		return tx, nil

	case caller.isMaxPriorityFeePerGasNotFoundError(err):
		log.Info("Don't support priority fee")
		opts.GasTipCap = FallbackGasTipCap
		return caller.DappLinkVrfContracts.FulfillRandomWords(opts, requestId, randomList, msgHash, big.NewInt(100), blsParam)
	default:
		return nil, err
	}
}

func (caller *Caller) FulfillRandomWords(requestId *big.Int, randomList []*big.Int) (*types.Receipt, error) {
	tx, err := caller.fulfillRandomWords(caller.Ctx, requestId, randomList)
	if err != nil {
		log.Error("build request random words tx fail", "err", err)
		return nil, err
	}
	updateGasPrice := func(ctx context.Context) (*types.Transaction, error) {
		return caller.UpdateGasPrice(ctx, tx)
	}
	receipt, err := caller.TxMrg.Send(caller.Ctx, updateGasPrice, caller.SendTransaction)
	if err != nil {
		log.Error("send tx fail", "err", err)
		return nil, err
	}
	return receipt, nil
}
