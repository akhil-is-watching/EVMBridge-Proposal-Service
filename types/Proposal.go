package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Proposal struct {
	TxHash        [32]byte
	Receiver      common.Address
	Amount        *big.Int
	ProposalNonce *big.Int
	ChainId       *big.Int
}
