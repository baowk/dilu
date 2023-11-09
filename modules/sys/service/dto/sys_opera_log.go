package dto

import (
	"dilu/modules/sys/models"
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type SysOperaLogGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int    `json:"status" query:"column:status"` //操作状态 1:成功 2:失败

}

func (SysOperaLogGetPageReq) TableName() string {
	return models.TBSysOperaLog
}

// 操作日志
type SysOperaLogDto struct {
	Id            int       `json:"id"`            //主键
	Title         string    `json:"title"`         //操作模块
	BusinessType  string    `json:"businessType"`  //操作类型
	BusinessTypes string    `json:"businessTypes"` //BusinessTypes
	Method        string    `json:"method"`        //函数
	RequestMethod string    `json:"requestMethod"` //请求方式 GET POST PUT DELETE
	OperatorType  string    `json:"operatorType"`  //操作类型
	OperName      string    `json:"operName"`      //操作者
	DeptName      string    `json:"deptName"`      //部门名称
	OperUrl       string    `json:"operUrl"`       //访问地址
	OperIp        string    `json:"operIp"`        //客户端ip
	OperLocation  string    `json:"operLocation"`  //访问位置
	OperParam     string    `json:"operParam"`     //请求参数
	Status        int       `json:"status"`        //操作状态 1:成功 2:失败
	OperTime      time.Time `json:"operTime"`      //操作时间
	JsonResult    string    `json:"jsonResult"`    //返回数据
	Remark        string    `json:"remark"`        //备注
	LatencyTime   string    `json:"latencyTime"`   //耗时
	UserAgent     string    `json:"userAgent"`     //ua
}
