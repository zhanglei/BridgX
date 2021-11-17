package helper

import (
	"github.com/galaxy-future/BridgX/cmd/api/middleware/authorization"
	"github.com/galaxy-future/BridgX/config"
	"github.com/gin-gonic/gin"
)

func GetUserClaims(ctx *gin.Context) *authorization.CustomClaims {
	claim, exist := ctx.Get(config.GlobalConfig.JwtToken.BindContextKeyName)
	if exist {
		ret, ok := claim.(*authorization.CustomClaims)
		if ok {
			return ret
		}
	}
	return nil
}
