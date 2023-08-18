package main

import (
	"dilu/cmd"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download
//go:generate swag init --parseDependency --parseDepth=6

// @title dilu API
// @version V0.0.1
// @description 一个简单的脚手
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
