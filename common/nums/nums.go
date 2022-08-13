package nums

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//from https://github.com/newbee-ltd/newbee-mall-api-go

// GenValidateCode 随机6位数
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, err := fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
		if err != nil {
			return ""
		}
	}
	return sb.String()
}

// StrToInt "2,3" 转换为[2,3]
func StrToInt(strNum string) ([]int64, error) {
	strNums := strings.Split(strNum, ",")
	var nums []int64
	for _, s := range strNums {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nums, errors.New("Atoi失败" + err.Error())
		}
		nums = append(nums, int64(i))
	}
	return nums, nil
}

// IntToStr [2,3] 转换为 "(2,3)"
func IntToStr(nums []int64) string {
	idStr := "("
	for i := 0; i < len(nums); i++ {
		idStr += strconv.FormatInt(nums[i], 10)
		if i != len(nums)-1 {
			idStr += ","
		}
	}
	idStr += ")"
	return idStr
}
