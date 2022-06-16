package store

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/mycodeku/transtionhelper/store/cache"
	"github.com/mycodeku/transtionhelper/store/rootmulti"
	"github.com/mycodeku/transtionhelper/store/types"
)

func NewCommitMultiStore(db dbm.DB) types.CommitMultiStore {
	return rootmulti.NewStore(db)
}

func NewCommitKVStoreCacheManager() types.MultiStorePersistentCache {
	return cache.NewCommitKVStoreCacheManager(cache.DefaultCommitKVStoreCacheSize)
}
