package response

import (
	"aioc/pkg/e"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `example:"20001"`
	Msg  string `example:"token验证错误"`
	Data any
}

type CommonList struct {
	Page     int64 `example:"1"`
	PageSize int64 `example:"10"`
	Total    int64 `example:"100"`
}

func GetAccountId(c *gin.Context) int64 {
	// userInfo := c.Value("account")
	// userService, ok := userInfo.(*service.User)
	// if !ok {
	return 0
	// }
	// return userService.GetUser().ID
}

func Resp(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})
}

func RespError(c *gin.Context, httpCode, errCode int, err error) {
	c.JSON(httpCode, gin.H{
		"code": errCode,
		"msg":  err.Error(),
		"data": nil,
	})
}
