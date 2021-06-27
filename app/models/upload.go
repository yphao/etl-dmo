package models

// @Description  etl-demo
// @Author playclouds
// @Update    2021/6/27 11:40

type UpLoadFile struct {
	Num      int    `gorm:"column:id; parimary_key", json:"num"`
	Filename string `gorm:"column:filename; ", json:"filename"`
	FullPath string `gorm:"column:fullpath; ", json:"fullpath"`
	Md5sum   string `gorm:"column:md5sum; ", json:"md5sum"`
	Extract  int    `gorm:"column:extract; ", json:"extract"`
	Delete   int    `gorm:"column:delete; ", json:"delete"`
}

type ExtractFile struct {
	Num       int    `gorm:"column:id; parimary_key", json:"num"`
	Filename  string `gorm:"column:filename; ", json:"filename"`
	FromFile  string `gorm:"column:fromfile; ", json:"fromfile"`
	FullPath  string `gorm:"column:fullpath; ", json:"fullpath"`
	JobExec   int    `gorm:"column:jobexec; ", json:"jobexec"`
	JobStatus int    `gorm:"column:jobstatus; ", json:"jobstatus"`
}

// ZipFile

//AddUpLoadFileInfo 往数据库插入上传文件信息
func AddUpLoadFileInfo(ufi *UpLoadFile) error {
	if err := DB.Create(ufi).Error; err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

func FindUFIMd5(md5 string) (*UpLoadFile, error) {
	var ufi UpLoadFile
	err := DB.First(&ufi, "md5sum = ?", md5).Error
	return &ufi, err
}

func QueryAllZipInfo() (*[]UpLoadFile, error) {
	ufi := make([]UpLoadFile, 100)
	err := DB.Find(&ufi).Error
	return &ufi, err
}

// ExecExtractFile

func AddExtractFileInfo(efi *ExtractFile) error {
	if err := DB.Create(efi).Error; err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

func QueryALLExtractFileInfo() (*[]ExtractFile, error) {
	ufi := make([]ExtractFile, 100)
	err := DB.Find(&ufi).Error
	return &ufi, err
}
