package util

import (
	"crypto/md5"
	"encoding/hex"
)

func GenGroupKey(systemId, groupName string) string {
	return systemId + ":" + groupName
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//判断数组种是否包含元素
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
