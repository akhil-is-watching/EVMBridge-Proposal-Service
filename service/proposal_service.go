package service

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/akhil-is-watching/E2E-Bridge/Proposal-Indexer/bridge"
	"github.com/akhil-is-watching/E2E-Bridge/Proposal-Indexer/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ProposalService struct {
	listenClient   *ethclient.Client
	foreignClient  *ethclient.Client
	listenChainID  *big.Int
	foreignChainID *big.Int
	listenAddress  common.Address
	foreignAddress common.Address
	listenBridge   *bridge.Bridge
	foreignBridge  *bridge.Bridge
	proposer       *types.Proposer
	cache          *types.BoundedCache
}

func NewProposalService(
	listenClient *ethclient.Client,
	foreignClient *ethclient.Client,
	listenChainIDInt int64,
	foreignChainIDInt int64,
	listenAddressHex string,
	foreignAddressHex string,
	proposer *types.Proposer,
) *ProposalService {

	listenChainID := big.NewInt(listenChainIDInt)
	listenAddress := common.HexToAddress(listenAddressHex)
	listenBridge, err := bridge.NewBridge(listenAddress, listenClient)
	if err != nil {
		log.Fatal(err)
	}

	foreignChainID := big.NewInt(foreignChainIDInt)
	foreignAddress := common.HexToAddress(foreignAddressHex)
	foreignBridge, err := bridge.NewBridge(foreignAddress, foreignClient)
	if err != nil {
		log.Fatal(err)
	}

	return &ProposalService{
		listenClient:   listenClient,
		foreignClient:  foreignClient,
		listenChainID:  listenChainID,
		foreignChainID: foreignChainID,
		listenAddress:  listenAddress,
		foreignAddress: foreignAddress,
		listenBridge:   listenBridge,
		foreignBridge:  foreignBridge,
		proposer:       proposer,
		cache:          types.NewBoundedCache(15),
	}
}

func (service *ProposalService) ListenForDeposit(startBlock uint64) {

	for {
		latestBlock, err := service.listenClient.BlockNumber(context.Background())
		if err != nil {
			log.Println("Error getting latest block:", err)
			continue
		}

		if startBlock == 0 {
			startBlock = latestBlock
		}

		for blockNumber := startBlock; blockNumber <= latestBlock; {
			endBlock := blockNumber + 5
			if endBlock > latestBlock {
				endBlock = latestBlock
			}

			it, err := service.listenBridge.FilterDepositEvent(&bind.FilterOpts{
				Start: blockNumber,
				End:   &endBlock,
			})

			if err != nil {
				log.Println("Error getting latest block:", err)
				continue
			}

			for it.Next() {
				event := it.Event

				if service.cache.Contains(event.Raw.TxHash) {
					continue // Skip processing if the event has been processed
				}

				fmt.Println("[LISTEN-CHAIN] DEPOSIT EVENT FOUND")
				fmt.Println("[LISTEN-CHAIN] Block:", it.Event.Raw.BlockNumber)
				fmt.Println("[LISTEN-CHAIN] TransactionHash:", it.Event.Raw.TxHash.Hex())
				fmt.Println("[SERVICE] INITIATING FOREIGN CHAIN TRANSACTION")

				proposal := types.Proposal{
					TxHash:        [32]byte(event.Raw.TxHash.Bytes()),
					Receiver:      event.Recepient,
					Amount:        event.Amount,
					ProposalNonce: event.Nonce,
					ChainId:       service.listenChainID,
				}

				go service.proposer.SendProposal(proposal)
				service.cache.Add(event.Raw.TxHash)
			}
			blockNumber = endBlock + 1 // Move to the next block after the last processed block
			time.Sleep(5 * time.Second)
		}

		startBlock = latestBlock
		time.Sleep(20 * time.Second)
	}
}
