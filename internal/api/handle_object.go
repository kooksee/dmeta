package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kooksee/dmeta/internal/config"
	"github.com/kooksee/dmeta/internal/kts"
	"math/big"
	"net/http"
	"time"
)

func PutObject(ctx *gin.Context) {
	cfg := config.DefaultConfig()
	id, err := cfg.PutObject(ctx.Request.Body)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	t := time.Now()
	if err := cfg.GetDb().Put(append([]byte("t:"), big.NewInt(int64(t.UnixNano())).Bytes()...), []byte(id), nil); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	md := &kts.Metadata{
		ID:        id,
		TimeStamp: t.UnixNano(),
		Time:      t.Format("2006-01-02 15:03:04"),
	}
	if err := cfg.GetDb().Put(append([]byte("md:"), id...), md.Encode(), nil); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	ctx.String(http.StatusOK, id)
}

func GetObject(ctx *gin.Context) {
	id := ctx.Param("id")
	cfg := config.DefaultConfig()
	dt, err := cfg.GetObject(id)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Data(http.StatusOK, "", dt)
}
