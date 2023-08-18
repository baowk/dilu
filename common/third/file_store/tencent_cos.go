package file_store

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"dilu/common/third/file_store/config"

	"github.com/baowk/dilu-core/core"

	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

func NewCOS(cfg *config.FSCfg) *TencentCOS {
	return &TencentCOS{
		cfg: cfg,
	}
}

type TencentCOS struct {
	cfg *config.FSCfg
}

// UploadFile upload file to COS
func (e *TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient(e.cfg)
	f, openError := file.Open()
	if openError != nil {
		core.Log.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), e.cfg.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return e.cfg.BaseURL + "/" + e.cfg.PathPrefix + "/" + fileKey, fileKey, nil
}

// DeleteFile delete file form COS
func (e *TencentCOS) DeleteFile(key string) error {
	client := NewClient(e.cfg)
	name := e.cfg.PathPrefix + "/" + key
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		core.Log.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// NewClient init COS client
func NewClient(cfg *config.FSCfg) *cos.Client {
	urlStr, _ := url.Parse("https://" + cfg.Bucket + ".cos." + cfg.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cfg.SecretID,
			SecretKey: cfg.SecretKey,
		},
	})
	return client
}
