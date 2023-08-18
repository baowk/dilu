package file_store

import (
	"fmt"
	"mime/multipart"

	"dilu/common/third/file_store/config"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	DeleteFile(key string) error
}

var fsCfg config.FSCfg

func InitCfg() {
	if fsCfg.SecretID == "" {
		//TODO
		fmt.Println("这里需要配置")
	}
}

// NewOss OSS的实例化方法
func NewFs(ossType string) OSS {
	InitCfg()
	switch ossType {
	case "local":
		return NewLocal(&fsCfg)
	case "qiniu":
		return NewQiniu(&fsCfg)
	case "tencent":
		return NewCOS(&fsCfg)
	case "aliyun":
		return NewOss(&fsCfg)
	case "huawei":
		return NewOBS(&fsCfg)
	case "aws":
		return NewS3(&fsCfg)
	default:
		return NewLocal(&fsCfg)
	}
}
