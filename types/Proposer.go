package types

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/akhil-is-watching/E2E-Bridge/Proposal-Indexer/bridge"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Proposer struct {
	mu             sync.Mutex
	client         *ethclient.Client
	address        string
	foreignAddress common.Address
	foreignChainID *big.Int
	privateKey     *ecdsa.PrivateKey
	nonce          uint64
	nonceMu        sync.Mutex // Mutex for nonce management
}

func NewProposer(
	privateKeyHex string,
	foreignAddressHex string,
	foreignChainIDInt int64,
	client *ethclient.Client,
) *Proposer {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("PRIVATE KEY DERIVATION FAILED")
	}

	foreignAddress := common.HexToAddress(foreignAddressHex)
	foreignChainID := big.NewInt(foreignChainIDInt)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("PUBLIC KEY DERIVATION FAILED")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, _ := client.PendingNonceAt(context.Background(), address)

	return &Proposer{
		client:         client,
		foreignAddress: foreignAddress,
		foreignChainID: foreignChainID,
		address:        address.Hex(),
		privateKey:     privateKey,
		nonce:          nonce,
	}
}

func (p *Proposer) Address() string {
	return p.address
}

func (p *Proposer) GetNonce() uint64 {
	p.nonceMu.Lock()
	defer p.nonceMu.Unlock()
	return p.nonce
}

func (p *Proposer) IncrementNonce() {
	p.nonceMu.Lock()
	defer p.nonceMu.Unlock()
	p.nonce++
}

func (p *Proposer) SendProposal(proposal Proposal) {
	p.mu.Lock()
	defer p.mu.Unlock()

	gasPrice, err := p.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	auth, _ := bind.NewKeyedTransactorWithChainID(p.privateKey, p.foreignChainID)
	auth.Nonce = big.NewInt(int64(p.GetNonce()))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(0)
	auth.GasPrice = gasPrice

	instance, err := bridge.NewBridge(p.foreignAddress, p.client)
	if err != nil {
		log.Println(err)
		return
	}

	tx, err := instance.Propose(
		auth,
		proposal.TxHash,
		proposal.Receiver,
		proposal.Amount,
		proposal.ProposalNonce,
		proposal.ChainId,
	)

	if err != nil {
		log.Println(err)
		return
	}
	p.IncrementNonce()
	fmt.Println("[FOREIGN-CHAIN]TransactionHash: ", tx.Hash().Hex())
	fmt.Println(" ")
}
