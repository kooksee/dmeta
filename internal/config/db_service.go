package config

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"math/big"
)

func (t *config) SaveObject() {

}

func (t *config) Ranger(startTime, limit int) (dt []string) {
	fmt.Println(startTime, limit)

	iter := t.db.NewIterator(util.BytesPrefix([]byte("")), nil)
	iter.Seek(append([]byte("t:"),big.NewInt(int64(startTime)).Bytes()...))
	for iter.Next() && (len(dt) < limit) {
		dt = append(dt, string(iter.Value()))
	}
	return
}
