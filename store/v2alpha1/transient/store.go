package transient

import (
	dbm "github.com/mycodeku/transtionhelper/db"
	"github.com/mycodeku/transtionhelper/db/memdb"
	"github.com/mycodeku/transtionhelper/store/types"
	"github.com/mycodeku/transtionhelper/store/v2alpha1/dbadapter"
)

var (
	_ types.KVStore   = (*Store)(nil)
	_ types.Committer = (*Store)(nil)
)

// Store is a wrapper for a memory store which does not persist data.
type Store struct {
	dbadapter.Store
	conn dbm.DBConnection
}

// NewStore constructs a new transient store.
func NewStore() *Store {
	db := memdb.NewDB()
	return &Store{
		Store: dbadapter.Store{DB: db.ReadWriter()},
		conn:  db,
	}
}

// Implements Store.
func (ts *Store) GetStoreType() types.StoreType {
	return types.StoreTypeTransient
}

// Implements CommitStore
// Commit cleans up Store.
func (ts *Store) Commit() (id types.CommitID) {
	ts.DB.Discard()
	ts.Store = dbadapter.Store{DB: ts.conn.ReadWriter()}
	return
}

func (ts *Store) SetPruning(types.PruningOptions)  {}
func (ts *Store) GetPruning() types.PruningOptions { return types.PruningOptions{} }

func (ts *Store) LastCommitID() (id types.CommitID) { return }
