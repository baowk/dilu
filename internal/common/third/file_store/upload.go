package file_store

import (
	"mime/multipart"
)

// OSS 对象存储接口
type OSSHandler interface {
	UploadFile(file *multipart.FileHeader) (filePath string, fileKey string, err error)
	DeleteFile(key string) error
}

// Deprecated: 使用 container.Global().OSS 替代
var Oss OSSHandler

func Setup(oss OSSHandler) {
	Oss = oss
}
