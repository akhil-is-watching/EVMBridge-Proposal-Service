package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/akhil-is-watching/E2E-Bridge/Proposal-Indexer/service"
	"github.com/akhil-is-watching/E2E-Bridge/Proposal-Indexer/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	listenRpcUrl := flag.String("listenRPC", "", "")
	foreignRpcUrl := flag.String("foreignRPC", "", "")
	listenChainID := flag.Int64("listenChain", 0, "")
	foreignChainID := flag.Int64("foreignChain", 0, "")
	listenContract := flag.String("listenAddress", "", "")
	foreignContract := flag.String("foreignAddress", "", "")
	privateKey := flag.String("privateKey", "", "")
	startBlock := flag.Int64("block", 0, "")

	flag.Parse()
	if *listenRpcUrl == "" {
		log.Fatalf("Missing listenRPC")
		os.Exit(0)
	}
	if *foreignRpcUrl == "" {
		log.Fatalf("Missing foreignRPC")
		os.Exit(0)
	}
	if *listenChainID == 0 {
		log.Fatalf("Missing listenChain")
		os.Exit(0)
	}
	if *foreignChainID == 0 {
		log.Fatalf("Missing foreignChain")
		os.Exit(0)
	}
	if *listenContract == "" {
		log.Fatalf("Missing listenAddress")
		os.Exit(0)
	}
	if *foreignContract == "" {
		log.Fatalf("Missing foreignAddress")
		os.Exit(0)
	}
	if *privateKey == "" {
		log.Fatalf("Missing privateKey")
		os.Exit(0)
	}

	listenClient, err := ethclient.DialContext(context.Background(), *listenRpcUrl)
	if err != nil {
		log.Fatal("Error initializing client")
	}
	foreignClient, err := ethclient.DialContext(context.Background(), *foreignRpcUrl)
	if err != nil {
		log.Fatal("Error initializing client")
	}

	proposer := types.NewProposer(
		*privateKey,
		*foreignContract,
		*foreignChainID,
		foreignClient,
	)

	proposal_service := service.NewProposalService(
		listenClient,
		foreignClient,
		*listenChainID,
		*foreignChainID,
		*listenContract,
		*foreignContract,
		proposer,
	)

	if *startBlock != 0 {
		proposal_service.ListenForDeposit(uint64(*startBlock))
	} else {
		proposal_service.ListenForDeposit(0)
	}
}
