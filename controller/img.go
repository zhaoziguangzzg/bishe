package controller

import (
	"bishe/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ImageUploadHandler(c *gin.Context) {
	imgType := c.PostForm("type")
	if imgType == "" {
		MakeApiResponseErrorParams(c)
		return
	}

	file, header, err := c.Request.FormFile("img")
	if err != nil {
		if err == http.ErrMissingFile {
			MakeApiResponseErrorParams(c)
			return
		}
		service.Logger.Error("FormFile err", zap.Error(err))
		MakeApiResponseErrorParams(c)
		return
	}

	if header.Size == 0 {
		MakeApiResponseErrorParams(c)
		return
	}

	imgPath, err := service.ImageSave(file, header, imgType, time.Now())
	if err != nil {
		service.Logger.Error("ImageSave err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, gin.H{"url": imgPath})
}
