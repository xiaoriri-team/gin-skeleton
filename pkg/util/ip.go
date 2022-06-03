package util

import "github.com/xiaoriri-team/gin-skeleton/pkg/util/iploc"

func GetIPLoc(ip string) string {
	country, _ := iploc.Find(ip)
	return country
}
