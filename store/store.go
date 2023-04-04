package store

import (
	"github.com/reapchain/reapchain-core/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/reapchain/cosmos-sdk/store/cache"
	"github.com/reapchain/cosmos-sdk/store/rootmulti"
	"github.com/reapchain/cosmos-sdk/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db, log.NewNopLogger())
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
