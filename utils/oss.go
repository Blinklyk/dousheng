package utils

import (
	"context"
<<<<<<< HEAD
<<<<<<< HEAD
=======

>>>>>>> upstream/gzh
	"encoding/base64"
	"fmt"
	"path/filepath"
	"strings"

<<<<<<< HEAD
=======
>>>>>>> a5ad9421cddcb4c71a3ebda7d6ed77f835c4b828
=======
>>>>>>> upstream/gzh
	"github.com/RaymondCode/simple-demo/global"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
)

type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

func UploadFile(localPath string) (string, error) {
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
		global.App.DY_LOG.Error("put file error ", zap.Error(err))
		return "", err
	}
	global.App.DY_LOG.Info("bucket: " + ret.Bucket + "key: " + ret.Key + "Hash: " + ret.Hash + "Name: " + ret.Name)

	// get url: oss domain + ret.key
	VideoUrl := global.App.DY_CONFIG.Qiniu.Domain + "/" + ret.Key
	global.App.DY_LOG.Info("Publish video url: " + VideoUrl)
	return VideoUrl, nil
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
