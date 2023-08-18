package dto

import "github.com/baowk/dilu-core/core/base"

type DemoDto struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required,min=2,max=64"`
	// Password   string `json:"password" binding:"required,min=6"`
	// RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	// Email      string `json:"email" binding:"omitempty,email"`
	// Age        int    `json:"age" binding:"omitempty,gt=18,lt=60"`
}

type DemePageReq struct {
	base.ReqPage
}
