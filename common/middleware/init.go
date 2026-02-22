package middleware

import (
	"dilu/common/config"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine, cfg *config.Config) {
	r.Use(LoggerToFile())
	if cfg.AccessLimit.Enable {
		r.Use(AccessLimitfunc())
	}
	if cfg.Cors.Enable {
		r.Use(CorsByRules(&cfg.Cors))
	}
	r.Use(CustomError)
	r.Use(ReqId)
	//r.Use(NoCache)

}
