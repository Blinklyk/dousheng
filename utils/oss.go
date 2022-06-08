package utils

import (
	"context"
	"fmt"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func UploadFile(localPath string) *MyPutRet {
	// get from platform

	accessKey := global.App.DY_CONFIG.Qiniu.AccessKey
	secretKey := global.App.DY_CONFIG.Qiniu.SecretKey

	localFile := localPath
	// storage space name
	putPolicy := storage.PutPolicy{
		Scope:      global.App.DY_CONFIG.Qiniu.Bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	// set configuration
	cfg := qiniuConfig()

	formUploader := storage.NewFormUploader(cfg)

	// file path in storage space
	key := "root/" + localPath

	// use returnBody define response format:key/hash/...

	ret := MyPutRet{}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("bucket: "+ret.Bucket, "key: "+ret.Key, ret.Fsize, "Hash: "+ret.Hash, "Name: "+ret.Name)
	return &ret
}

func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      global.App.DY_CONFIG.Qiniu.UseHTTPS,
		UseCdnDomains: global.App.DY_CONFIG.Qiniu.UseCdnDomains,
	}
	switch global.App.DY_CONFIG.Qiniu.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
