package v1

import (
	"net/http"

	"aioc/internal/api/response"
	"aioc/internal/service"
	"aioc/pkg/e"

	"github.com/gin-gonic/gin"
)

func GetResourcePools(c *gin.Context) {
	data := service.GetResourcePools()
	response.Resp(c, http.StatusOK, e.SUCCESS, data)
}
