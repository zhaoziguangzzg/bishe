package controller

import (
	"bishe/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ImageUploadHandler(c *gin.Context) {
	var req struct {
		Type int `form:"type"`
	}
	if c.ShouldBind(&req) != nil || req.Type == 0 {
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

	imgPath, err := service.ImgSave(file, header, req.Type, time.Now())
	if err != nil {
		service.Logger.Error("ImgSave err", zap.Error(err))
		MakeApiResponseErrorDefault(c)
		return
	}

	MakeApiResponseSuccess(c, gin.H{"url": imgPath})
}
