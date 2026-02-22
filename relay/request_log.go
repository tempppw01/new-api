package relay

import (
	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/logger"
	"github.com/QuantumNous/new-api/setting/model_setting"
	"github.com/gin-gonic/gin"
)

func logUserRequestBody(c *gin.Context, message string) {
	if common.DebugEnabled {
		logger.LogDebug(c, message)
		return
	}
	if model_setting.IsRequestBodyLogEnabled() {
		logger.LogInfo(c, message)
	}
}
