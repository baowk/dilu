package file_store

import (
	"mime/multipart"
)

// OSS 对象存储接口
type OSSHandler interface {
	UploadFile(file *multipart.FileHeader) (filePath string, fileKey string, err error)
	DeleteFile(key string) error
}

var Oss OSSHandler

func Setup(oss OSSHandler) {
	Oss = oss
}
