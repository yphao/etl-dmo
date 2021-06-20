package routers

import (
	"archive/tar"
	"etl-demo/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)
import "net/http"

// @Description  etl-demo
// @Author playclouds
// @Update    2021/6/20 14:44

func UploadFileHandler(c *gin.Context) {
	config, err := utils.InitConfig()
	if err != nil {
		fmt.Println(err)
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	filename := filepath.Base(file.Filename)
	fmt.Println(filename)
	upLoadFileName := fmt.Sprintf(config.Upload.Tmpdir + filename)

	fmt.Println(config.Upload.Dir)
	uploadFileNamePath := fmt.Sprintf(strings.TrimSuffix(filename, path.Ext(filename)))
	fmt.Println(uploadFileNamePath)
	if err != nil {
		panic(err)
	}
	if err := c.SaveUploadedFile(file, upLoadFileName); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	//计算md5值
	fileMd5, err := utils.CalcFileMd5(upLoadFileName)

	//解压tar包
	fileReader, err := os.Open(upLoadFileName)
	if err != nil {
		panic(err)
	}
	defer fileReader.Close()
	tarReader := tar.NewReader(fileReader)
	for hdr, err := tarReader.Next(); err != io.EOF; hdr, err = tarReader.Next() {
		if err != nil {
			log.Println(err)
			continue
		}
		// 读取文件信息
		fi := hdr.FileInfo()

		// 创建一个空文件，用来写入解包后的数据
		fmt.Println(fmt.Sprintf(uploadFileNamePath + "/" + fi.Name()))
		fw, err := os.Create(fmt.Sprintf(uploadFileNamePath + fi.Name())) //fi.Name()
		if err != nil {
			log.Println(err)
			continue
		}

		if _, err := io.Copy(fw, tarReader); err != nil {
			log.Println(err)
		}
		os.Chmod(fi.Name(), fi.Mode().Perm())
		fw.Close()
	}

	c.String(http.StatusOK,
		fmt.Sprintf("File %s uploaded successfully,File md5 %s,Store in path %s ", filename, fileMd5, upLoadFileName))
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/upload", "views/upload")
	r.POST("/uploadFile", UploadFileHandler)
	return r
}
