package main

import (
	"dilu/cmd"
	_ "dilu/common/consts"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.io,direct
//go:generate go mod tidy
//go:generate go mod download
//go:generate swag init --parseDependency --parseDepth=6

// @title Dilu API
// @version V0.0.1
// @description 致力于做一个开发快速，运行稳定的框架
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @BasePath /api/v1
func main() {
	cmd.Execute()
}
