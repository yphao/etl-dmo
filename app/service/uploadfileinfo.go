package service

import (
	"archive/tar"
	"etl-demo/app/models"
	"etl-demo/app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// @Description  etl-demo
// @Author playclouds
// @Update    2021/6/26 17:30

func UpLoad(c *gin.Context) {
	config := utils.GetConfig()

	//获取post方法参数
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	upLoadTmpPath := config.Upload.Tmpdir
	isFileExit(upLoadTmpPath)
	filename := filepath.Base(file.Filename)
	upLoadFileFullPath := filepath.Join(upLoadTmpPath + filename)
	uploadFileName := fmt.Sprintf(strings.TrimSuffix(filename, path.Ext(filename)))

	//将文件保存在临时路径
	if err != nil {
		panic(err)
	}
	if err := c.SaveUploadedFile(file, upLoadFileFullPath); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	//todo 使用md5值进行去重
	//计算md5值
	fileMd5, err := utils.CalcFileMd5(upLoadFileFullPath)

	//通过md5值去重
	var recMes string
	_, errMd5 := models.FindUFIMd5(fileMd5)
	if errMd5 != nil {
		ufi := &models.UpLoadFile{Filename: filename, FullPath: upLoadFileFullPath, Md5sum: fileMd5, Extract: 0, Delete: 0}
		models.AddUpLoadFileInfo(ufi)
		recMes = fmt.Sprintf("File %s uploaded successfully,File md5 %s,Store in path %s ", filename, fileMd5, upLoadFileFullPath)

		//解压tar包
		fileReader, err := os.Open(upLoadFileFullPath)
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

			//判断解压路径是存在，不存在则新建路径
			tarXfPath := filepath.Join(config.Upload.Dir, uploadFileName)
			isFileExit(tarXfPath)
			// 创建一个空文件，用来写入解包后的数据
			newFileName := filepath.Join(tarXfPath, fi.Name())
			//fis,err:= os.Stat(newFileName)
			if err == nil {
				if fi.Size() == 0 {
					//去掉解压后的同名文件夹
					if fi.Name() == uploadFileName {
						continue
					}
					//去除空文件
					log.Printf("file %s is nil", fi.Name())
					continue
				}
			}
			fw, err := os.Create(newFileName)
			if err != nil {
				log.Println(err)
				continue
			}

			if _, err := io.Copy(fw, tarReader); err != nil {
				log.Println(err)
			}
			os.Chmod(fi.Name(), fi.Mode().Perm())
			fw.Close()

			efi := &models.ExtractFile{Filename: fi.Name(), FromFile: filename, FullPath: newFileName, JobExec: 0, JobStatus: 0}
			models.AddExtractFileInfo(efi)

		}

	} else {
		recMes = "该文件已存在，请要重复上传，请检查上传文件"

	}

	c.HTML(http.StatusOK, "uploadres.html", gin.H{
		"rec": recMes,
	})
}

// isFileExit  判断文件路径是否存在，不存在则创建该路径，传入文件路径
func isFileExit(filepath string) {
	_, err := os.Stat(filepath)
	if err != nil {
		os.Mkdir(filepath, os.ModePerm)
	}
}
