package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/reapchain/cosmos-sdk/store/cache"
	"github.com/reapchain/cosmos-sdk/store/rootmulti"
	"github.com/reapchain/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
