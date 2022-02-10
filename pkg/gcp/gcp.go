package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/base64"
	"github.com/junminhong/member-center-service/pkg/logger"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"time"
)

var zapLogger = logger.Setup()
var storageClient *storage.Client

const (
	bucket = "file-center"
)

func init() {
	c := context.Background()
	tmp, err := storage.NewClient(c, option.WithCredentialsFile("file-center.json"))
	if err != nil {
		zapLogger.Error(err.Error())
	}
	storageClient = tmp
}

func GetFileUrlHaveExpired(fileName string) string {
	keyFile := "file-center.json"
	key, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatalln(err)
	}
	config, err := google.JWTConfigFromJSON(key)
	if err != nil {
		log.Fatalln(err)
	}
	method := "GET"
	expires := time.Now().Add(time.Second * 60)
	url, err := storage.SignedURL(bucket, fileName, &storage.SignedURLOptions{
		GoogleAccessID: config.Email,
		PrivateKey:     config.PrivateKey,
		Method:         method,
		Expires:        expires,
	})
	return url
}

func InsertFileToGCS(dir string, fileName string, uploadFile *multipart.FileHeader, file multipart.File) bool {
	c := context.Background()
	sw := storageClient.Bucket(bucket).Object(dir + fileName).NewWriter(c)
	if _, err := io.Copy(sw, file); err != nil {
		zapLogger.Info(err.Error())
		return false
	}
	if err := sw.Close(); err != nil {
		zapLogger.Info(err.Error())
		return false
	}
	return true
}

func GetFileForGCS(filePath string) string {
	c := context.Background()
	rc, err := storageClient.Bucket(bucket).Object(filePath).NewReader(c)
	if err != nil {
		zapLogger.Info(err.Error())
		return ""
	}
	defer rc.Close()

	slurp, err := ioutil.ReadAll(rc)
	if err != nil {
		zapLogger.Info(err.Error())
		return ""
	}

	return "data:" + rc.Attrs.ContentType + ";base64, " + base64.StdEncoding.EncodeToString(slurp)
}
