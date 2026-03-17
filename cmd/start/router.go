package start

import "github.com/gin-gonic/gin"

// RouterRegistry 路由注册表，替代全局 AppRouters slice + init() 模式
type RouterRegistry struct {
	routers []func(*gin.Engine)
}

// Register 注册一个路由初始化函数
func (r *RouterRegistry) Register(fn func(*gin.Engine)) {
	r.routers = append(r.routers, fn)
}

// Init 执行所有已注册的路由初始化
func (r *RouterRegistry) Init(engine *gin.Engine) {
	for _, fn := range r.routers {
		fn(engine)
	}
}

// Routers 全局路由注册表
var Routers = &RouterRegistry{}
