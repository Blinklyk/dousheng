package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"

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

	filename := filepath.Base(localFile)

	indexOfDot := strings.LastIndex(filename, ".")
	prevfix := filename[0 : indexOfDot-1]
	coverName := prevfix + "." + "jpg"

	photoKey := "root/cover" + coverName //封面的访问路径，我们通过此路径在七牛云空间中定位封面
	entry := global.App.DY_CONFIG.Qiniu.Bucket + ":" + photoKey
	encodedEntryURI := base64.StdEncoding.EncodeToString([]byte(entry))

	// storage space name
	putPolicy := storage.PutPolicy{
		Scope:         global.App.DY_CONFIG.Qiniu.Bucket,
		PersistentOps: "vframe/jpg/offset/1|saveas/" + encodedEntryURI, //取视频第1秒的截图,
		ReturnBody:    `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
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
