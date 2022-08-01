package jwt

import (
	"newbee-mall-gozero/common/md5"
	"newbee-mall-gozero/common/nums"
	"strconv"
	"strings"
)

//from https://github.com/newbee-ltd/newbee-mall-api-go

func GetNewToken(timeInt int64, userId int) (token string) {
	var build strings.Builder
	build.WriteString(strconv.FormatInt(timeInt, 10))
	build.WriteString(strconv.Itoa(userId))
	build.WriteString(nums.GenValidateCode(6))
	return md5.MD5V([]byte(build.String()))
}
