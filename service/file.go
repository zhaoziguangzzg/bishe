package service

import (
	"fmt"
	"io"
	"math/rand/v2"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	qrcode "github.com/skip2/go-qrcode"
)

const (
	FileTypeUserAvatar   = 1 //用户头像
	FileTypeAdvertImg    = 2 //广告图片
	FileTypePayImg       = 3 //收款码图片
	FileTypeCircleImg    = 4 //圈子图片
	FileTypeCourseImg    = 5 //课程图片
	FileTypeCoursePayImg = 6 //课程收款码图片
)

var imgTypePrefix = map[int]string{
	FileTypeUserAvatar:   "user",
	FileTypeAdvertImg:    "advert",
	FileTypePayImg:       "pay",
	FileTypeCircleImg:    "circle",
	FileTypeCourseImg:    "course",
	FileTypeCoursePayImg: "course_pay",
}

func saveFile(imgType int, timeNow time.Time, ext string, data []byte) (string, error) {
	prefix, ok := imgTypePrefix[imgType]
	if !ok {
		prefix = "default"
	}
	dateDir := timeNow.Format("2006-01-02")
	name := fmt.Sprintf("%s_%d_%d%s", prefix, timeNow.Unix(), rand.IntN(1000), ext)

	uploadDir := filepath.Join("web", "img", prefix, dateDir)
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", err
	}
	uploadPath := filepath.Join(uploadDir, name)

	if err := os.WriteFile(uploadPath, data, 0644); err != nil {
		return "", err
	}

	return "/img/" + prefix + "/" + dateDir + "/" + name, nil
}

func ImgSave(file multipart.File, header *multipart.FileHeader, imgType int, timeNow time.Time) (string, error) {
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".jpg"
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return saveFile(imgType, timeNow, ext, data)
}

func QrcodeImgSave(content string, size int, imgType int, timeNow time.Time) (string, error) {
	png, err := qrcode.Encode(content, qrcode.Medium, size)
	if err != nil {
		return "", err
	}

	return saveFile(imgType, timeNow, ".png", png)
}
