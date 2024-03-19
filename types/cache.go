package types

import "github.com/ethereum/go-ethereum/common"

type BoundedCache struct {
	maxSize  int
	keys     []common.Hash
	hashSet  map[common.Hash]bool
	position int // Tracks the position of the oldest item in the cache
}

func NewBoundedCache(maxSize int) *BoundedCache {
	return &BoundedCache{
		maxSize:  maxSize,
		keys:     make([]common.Hash, 0, maxSize),
		hashSet:  make(map[common.Hash]bool),
		position: 0,
	}
}

func (c *BoundedCache) Add(hash common.Hash) {
	if c.hashSet[hash] {
		return
	}

	if len(c.keys) == c.maxSize {
		// Remove the oldest item
		delete(c.hashSet, c.keys[c.position])
		c.keys[c.position] = hash
		c.hashSet[hash] = true
		c.position = (c.position + 1) % c.maxSize
	} else {
		c.keys = append(c.keys, hash)
		c.hashSet[hash] = true
	}
}

func (c *BoundedCache) Contains(hash common.Hash) bool {
	return c.hashSet[hash]
}
