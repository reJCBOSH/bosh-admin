package basic

import (
	"errors"
	"io"
	"mime/multipart"

	"bosh-admin/core/db"
	"bosh-admin/core/exception"
	"bosh-admin/model"
	"bosh-admin/util"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/h2non/filetype"
)

type BasicSvc struct{}

func NewBasicSvc() *BasicSvc {
	return &BasicSvc{}
}

func (svc *BasicSvc) UploadFile(file *multipart.FileHeader, where, source, ip string) (*model.Resource, error) {
	src, err := file.Open()
	if err != nil {
		return nil, exception.NewException("打开文件失败", err)
	}
	defer func() {
		_ = src.Close()
	}()
	buf, _ := io.ReadAll(src)
	kind, _ := filetype.Match(buf)
	// 计算MD5校验值
	checkSum := cryptor.Md5Byte(buf)
	// 重置文件指针到开头
	if _, err = src.Seek(0, 0); err != nil {
		return nil, exception.NewException("重置文件指针失败", err)
	}
	var resource *model.Resource
	err = db.GormDB().Model(&model.Resource{}).Where("ip = ?", ip).Where("check_sum = ?", checkSum).First(&resource).Error
	if err == nil {
		return resource, nil
	} else {
		if !errors.Is(err, db.NotFound) {
			return nil, exception.NewException("查询资源记录失败", err)
		}
	}
	storePath, fullPath, err := util.Upload(src, file.Filename, where)
	if err != nil {
		return nil, err
	}
	resource = &model.Resource{
		Source:    source,
		IP:        ip,
		FileName:  file.Filename,
		FileSize:  file.Size,
		FileType:  kind.Extension,
		MimeType:  kind.MIME.Value,
		StorePath: storePath,
		FullPath:  fullPath,
		CheckSum:  checkSum,
	}
	err = db.Create(resource)
	if err != nil {
		return nil, exception.NewException("创建资源记录失败", err)
	}
	return resource, nil
}
