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

// , error
const (
	FILE_TYPE_UAER_AVATAR    int = 1 //用户头像
	FILE_TYPE_ADVERT_IMG     int = 2 //广告图片
	FILE_TYPE_PAY_IMG        int = 3 //收款码图片
	FILE_TYPE_CIRCLE_IMG     int = 4 //圈子图片
	FILE_TYPE_COURSE_IMG     int = 5 //课程图片
	FILE_TYPE_COURSE_PAY_IMG int = 6 //课程收款码图片
)

func ImageSave(file multipart.File, header *multipart.FileHeader, imgType string, timeNow time.Time) (imgPath string, err error) {
	defer file.Close()

	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".jpg"
	}

	name := fmt.Sprintf("%d_%d%s", timeNow.Unix(), rand.IntN(1000), ext)

	dirPath := filepath.Join("web", "img", imgType)
	if err = os.MkdirAll(dirPath, 0755); err != nil {
		return
	}

	uploadPath := filepath.Join(dirPath, name)

	out, err := os.Create(uploadPath)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return
	}
	imgPath = "/img/" + imgType + "/" + name
	return
}

func FileSave(file multipart.File, header *multipart.FileHeader, fileType int, timeNow time.Time) (avatarPath string, err error) {
	defer file.Close()

	// 获取文件扩展名
	ext := filepath.Ext(header.Filename)
	if ext == "" {
		ext = ".jpg"
	}
	var name string

	switch fileType {
	case FILE_TYPE_UAER_AVATAR:
		name = fmt.Sprintf("user_%d_%d.%s", timeNow.Unix(), rand.IntN(1000), ext)
	case FILE_TYPE_ADVERT_IMG:
		name = fmt.Sprintf("advert_%d_%d.%s", timeNow.Unix(), rand.IntN(1000), ext)
	case FILE_TYPE_PAY_IMG:
		name = fmt.Sprintf("pay_%d_%d.%s", timeNow.Unix(), rand.IntN(1000), ext)
	case FILE_TYPE_CIRCLE_IMG:
		name = fmt.Sprintf("circle_%d_%d.%s", timeNow.Unix(), rand.IntN(1000), ext)
	case FILE_TYPE_COURSE_IMG:
		name = fmt.Sprintf("course_%d_%d.%s", timeNow.Unix(), rand.IntN(1000), ext)
	case FILE_TYPE_COURSE_PAY_IMG:
		name = fmt.Sprintf("course_pay_%d_%d.%s", timeNow.Unix(), rand.IntN(1000), ext)
	default:
		name = fmt.Sprintf("default_%d_%d.%s", timeNow.Unix(), rand.IntN(1000), ext)
	}

	// 图像名，用户名+文件扩展名
	uploadPath := filepath.Join("web", "img", name)

	// 保存文件
	out, err := os.Create(uploadPath)
	if err != nil {
		return
	}
	defer out.Close()
	//将用户上传的file复制到out文件里
	_, err = io.Copy(out, file)
	if err != nil {
		return
	}
	avatarPath = "/img/" + name
	return
}
