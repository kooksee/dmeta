package config

import (
	"io"
	"io/ioutil"
)

func (t *config) PutObject(r io.Reader) (string, error) {
	return t.ipfs.Add(r)
}

func (t *config) GetObject(hash string) ([]byte, error) {
	dt, err := t.ipfs.Cat(hash)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(dt)
}
