package cache

import (
	"math/big"
	"net/netip"
	"sync"

	"github.com/hashicorp/golang-lru" // Importing the lru package for cache implementation
)

// BiDirectionalCache struct maintains two caches: one for mapping IP to big int,
// and another for mapping big int to IP.
type BiDirectionalCache struct {
	ipToBigIntCache *lru.Cache // Cache for mapping IP to big int
	bigIntToIPCache *lru.Cache // Cache for mapping big int to IP
	lock            sync.RWMutex
}

// NewBiDirectionalCache function initializes and returns a new BiDirectionalCache
// instance with the given size. It creates two caches of the specified size using
// the lru package.
func NewBiDirectionalCache(size int) (*BiDirectionalCache, error) {
	ipToBigInt, err := lru.New(size)
	if err != nil {
		return nil, err
	}

	bigIntToIP, err := lru.New(size)
	if err != nil {
		return nil, err
	}

	return &BiDirectionalCache{
	
