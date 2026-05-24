package service

import (
	"fmt"
	"io"
	"math/rand/v2"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

const (
	FileTypeUserAvatar  = 1 //用户头像
	FileTypeAdvertImg   = 2 //广告图片
	FileTypePayImg      = 3 //收款码图片
	FileTypeCircleImg   = 4 //圈子图片
	FileTypeCourseImg   = 5 //课程图片
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

func ImgSave(file multipart.File, header *multipart.FileHeader, imgType int, timeNow time.Time) (string, error) {
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".jpg"
	}

	prefix, ok := imgTypePrefix[imgType]
	if !ok {
		prefix = "default"
	}
	name := fmt.Sprintf("%s_%d_%d%s", prefix, timeNow.Unix(), rand.IntN(1000), ext)

	uploadPath := filepath.Join("web", "img", name)

	out, err := os.Create(uploadPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	if _, err = io.Copy(out, file); err != nil {
		return "", err
	}

	return "/img/" + name, nil
}
