package kts

import (
	"encoding/json"
	"github.com/kooksee/dmeta/internal/utils"
)

type Metadata struct {
	ID        string `json:"id,omitempty"`
	TimeStamp int64  `json:"timestamp,omitempty"`
	Time      string `json:"time,omitempty"`
}

func (t *Metadata) Encode() []byte {
	dt, err := json.Marshal(t)
	utils.MustNotError(err)
	return dt
}
