package utils

import (
	"context"
	"fmt"
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
	// set your own ak sk here
	accessKey := ""
	secretKey := ""

	localFile := localPath
	// storage space name
	bucket := "douyin-demo"
	// file path in storage space
	key := "root/" + localPath

	// use returnBody define response format:key/hash/...
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// set zone region
	cfg.Zone = &storage.ZoneBeimei
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("bucket: "+ret.Bucket, "key: "+ret.Key, ret.Fsize, "Hash: "+ret.Hash, "Name: "+ret.Name)
	return &ret
}
