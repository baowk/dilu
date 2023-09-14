package middleware

import (
	"github.com/baowk/dilu-core/config"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(r *gin.Engine, cfg *config.AppCfg) {
	r.Use(LoggerToFile())
	if cfg.Cors.Enable {
		r.Use(CorsByRules(&cfg.Cors))
	}
	r.Use(CustomError)
	r.Use(ReqId)
	//r.Use(NoCache)

}
