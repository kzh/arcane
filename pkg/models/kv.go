package models

import (
	"database/sql"
	"github.com/gobuffalo/pop/v6"
	"github.com/gofrs/uuid"
)

type KV struct {
	ID    uuid.UUID `db:"id"`
	Key   string    `db:"key"`
	Value string    `db:"value"`
}

func (kv *KV) Create(conn *pop.Connection) error {
	return conn.Transaction(func(tx *pop.Connection) error {
		res, err := KVs.ByKey(tx, kv.Key)
		if err == nil {
			res.Value = kv.Value
			return tx.Update(res)
		} else if err == sql.ErrNoRows {
			return tx.Create(kv)
		}
		return err
	})
}

type KVList []KV

type kvRepo struct{}

var KVs kvRepo

func (kvRepo) ByKey(conn *pop.Connection, key string) (*KV, error) {
	var results KVList
	err := conn.Where("key = ?", key).All(&results)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, sql.ErrNoRows
	}
	return &results[0], err
}

func (kvRepo) All(conn *pop.Connection) (KVList, error) {
	var results KVList
	err := conn.All(&results)
	return results, err
}
