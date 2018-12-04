package config

import (
	"github.com/ipfs/go-ipfs-api"
	"github.com/kooksee/dmeta/internal/utils"
	"github.com/patrickmn/go-cache"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"sync"
	"time"
)

type config struct {
	cache  *cache.Cache
	Debug  bool
	id     string
	dbPath string
	db     *leveldb.DB

	ipfsAddr string
	ipfs     *shell.Shell
}

func (t *config) IsDebug() bool {
	return t.Debug
}

func (t *config) GetDb() *leveldb.DB {
	if t.db == nil {
		panic("please init db")
	}
	return t.db
}

func (t *config) Init() {
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"

	var ll = zerolog.DebugLevel
	if !t.Debug {
		ll = zerolog.ErrorLevel
	}
	zerolog.SetGlobalLevel(ll)

	ip := utils.IpAddress()
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().
		Str("service_name", "dmeta").
		Str("service_id", ip).
		Str("service_ip", ip).
		Caller().
		Logger()

	db, err := leveldb.OpenFile(t.dbPath, nil)
	utils.MustNotError(err)
	t.db = db

	t.ipfs = shell.NewShell(t.ipfsAddr)
}

var cfg *config
var once sync.Once

func DefaultConfig() *config {
	once.Do(func() {
		cfg = &config{
			Debug:    true,
			cache:    cache.New(time.Minute*10, time.Minute*30),
			dbPath:   "kdata",
			ipfsAddr: "localhost:5001",
		}

		if e := env("debug"); e != "" {
			cfg.Debug = e == "true"
		}

		if e := env("db_path"); e != "" {
			cfg.dbPath = e
		}

		if e := env("ipfs_addr"); e != "" {
			cfg.dbPath = e
		}

	})
	return cfg
}
