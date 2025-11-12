package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"bosh-admin/core/exception"
	"bosh-admin/global"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/duke-git/lancet/v2/cryptor"
)

func Upload(src multipart.File, filename, where string) (string, string, error) {
	switch global.Config.Server.OssType {
	case "local":
		return LocalUpload(src, filename, where)
	case "aliyun":
		return AliyunOssUpload(src, filename, where)
	default:
		return LocalUpload(src, filename, where)
	}
}

func LocalUpload(src multipart.File, filename, where string) (string, string, error) {
	if where == "" {
		where = "default"
	}
	date := time.Now().Format(time.DateOnly)
	dirPath := filepath.Join(global.Config.Local.StorePath, where, date)
	ext := filepath.Ext(filename)
	rename := cryptor.Md5String(fmt.Sprintf("%d%s", time.Now().Unix(), filename)) + ext
	storePath := filepath.Join(dirPath, rename)
	fullPath := fmt.Sprintf("%s/%s/%s/%s", global.Config.Local.Path, where, date, rename)
	if global.Config.Server.Env == global.DEV {
		fullPath = fmt.Sprintf("http://%s:%d/%s", global.Config.Server.Url, global.Config.Server.Port, fullPath)
	} else {
		fullPath = fmt.Sprintf("%s/%s", global.Config.Server.Url, fullPath)
	}
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return "", "", exception.NewException("创建目录失败", err)
	}
	out, err := os.Create(storePath)
	if err != nil {
		return "", "", exception.NewException("创建文件失败", err)
	}
	defer func(out *os.File) {
		_ = out.Close()
	}(out)
	_, err = io.Copy(out, src)
	if err != nil {
		return "", "", exception.NewException("写入文件失败", err)
	}
	return storePath, fullPath, nil
}

func NewBucket() (*oss.Bucket, error) {
	// 创建oss client实例
	client, err := oss.New(global.Config.AliyunOss.Endpoint, global.Config.AliyunOss.AccessKeyId, global.Config.AliyunOss.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	// 获取存储桶
	bucket, err := client.Bucket(global.Config.AliyunOss.BucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func AliyunOssUpload(src multipart.File, filename, where string) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		return "", "", exception.NewException("创建存储桶失败", err)
	}
	if where == "" {
		where = "default"
	}
	dirPath := fmt.Sprintf("%s/%s/%s", global.Config.AliyunOss.BasePath, where, time.Now().Format(time.DateOnly))
	ext := filepath.Ext(filename)
	rename := cryptor.Md5String(fmt.Sprintf("%d%s", time.Now().Unix(), filename)) + ext
	storePath := fmt.Sprintf("%s/%s", dirPath, rename)
	fullPath := fmt.Sprintf("%s/%s", global.Config.AliyunOss.BucketUrl, storePath)
	if err = bucket.PutObject(storePath, src); err != nil {
		return "", "", exception.NewException("上传文件失败", err)
	}
	return storePath, fullPath, nil
}
