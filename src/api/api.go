package api

import (
	"encoding/json"
	"iad/contract"
	"iad/provider"
	"iad/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jfcarter2358/go-logger"
)

func AWSPost(ctx *gin.Context) {
	var data []map[string]interface{}
	if err := ctx.ShouldBindJSON(&data); err != nil {
		utils.Error(err, ctx, http.StatusInternalServerError)
		return
	}
	bs, _ := json.Marshal(data)

	var reply contract.Comm
	args := contract.Comm{
		Body: bs,
	}

	if err := provider.Providers["aws"].Client.Call("Provider.Post", args, &reply); err != nil {
		logger.Fatalf("", "error communicating with provider 'aws':%s", err.Error())
	}
}
