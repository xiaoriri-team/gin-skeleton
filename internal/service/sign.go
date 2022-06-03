package service

import (
	"fmt"
	"sort"

	"github.com/xiaoriri-team/gin-skeleton/global"
	"github.com/xiaoriri-team/gin-skeleton/pkg/util"
)

func (svc *Service) GetParamSign(param map[string]interface{}, secretKey string) string {
	signRaw := ""

	rawStrs := []string{}
	for k, v := range param {
		if k != "sign" {
			rawStrs = append(rawStrs, k+"="+fmt.Sprintf("%v", v))
		}
	}

	sort.Strings(rawStrs)
	for _, v := range rawStrs {
		signRaw += v
	}

	if global.ServerSetting.RunMode == "debug" {
		global.Logger.Info(map[string]string{
			"signRaw": signRaw,
			"sysSign": util.EncodeMD5(signRaw + secretKey),
		})
	}

	return util.EncodeMD5(signRaw + secretKey)
}
