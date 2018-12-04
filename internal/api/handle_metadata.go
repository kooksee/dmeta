package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kooksee/dmeta/internal/config"
	"github.com/tidwall/gjson"
	"net/http"
	"strconv"
)

func PutMetadata(ctx *gin.Context) {
	cfg := config.DefaultConfig()

	dt, _ := ctx.GetRawData()
	id := gjson.GetBytes(dt, "id").String()
	if err := cfg.GetDb().Put(append([]byte("md:"), id...), []byte(dt), nil); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.String(http.StatusOK, "ok")
}

func GetMetadata(ctx *gin.Context) {
	cfg := config.DefaultConfig()
	id := ctx.Param("id")
	if v, err := cfg.GetDb().Get(append([]byte("md:"), id...), nil); err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
	} else {
		ctx.Data(http.StatusOK, "application/json",v)
	}
}

func Ranger(ctx *gin.Context) {
	startTime := ctx.Query("start_time")
	limit := ctx.Query("limit")

	cfg := config.DefaultConfig()
	st, _ := strconv.Atoi(startTime)
	lt, _ := strconv.Atoi(limit)
	ctx.JSON(http.StatusOK, cfg.Ranger(st, lt))
}
