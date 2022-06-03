package api

import (
	"bytes"
	"encoding/base64"
	"image/color"
	"image/png"
	"time"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/xiaoriri-team/gin-skeleton/global"
	"github.com/xiaoriri-team/gin-skeleton/pkg/app"
	"github.com/xiaoriri-team/gin-skeleton/pkg/util"
)

func Version(c *gin.Context) {
	response := app.NewResponse(c)
	response.ToResponse(gin.H{
		"version": "PaoPao Service v1.0",
	})
}

func GetCaptcha(c *gin.Context) {
	cap := captcha.New()

	if err := cap.SetFont("assets/comic.ttf"); err != nil {
		panic(err.Error())
	}

	cap.SetSize(160, 64)
	cap.SetDisturbance(captcha.MEDIUM)
	cap.SetFrontColor(color.RGBA{0, 0, 0, 255})
	cap.SetBkgColor(color.RGBA{218, 240, 228, 255})
	img, password := cap.Create(6, captcha.NUM)
	emptyBuff := bytes.NewBuffer(nil)
	_ = png.Encode(emptyBuff, img)

	key := util.EncodeMD5(uuid.Must(uuid.NewV4()).String())

	// 五分钟有效期
	global.Redis.SetEX(c, "PaoPaoCaptcha:"+key, password, time.Minute*5)

	response := app.NewResponse(c)
	response.ToResponse(gin.H{
		"id":   key,
		"b64s": "data:image/png;base64," + base64.StdEncoding.EncodeToString(emptyBuff.Bytes()),
	})
}
