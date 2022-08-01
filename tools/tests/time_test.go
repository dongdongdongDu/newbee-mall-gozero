package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	var time1 int64 = 1659340240
	var time2 int64 = 1659512689

	//时间戳 to 时间
	tm1 := time.Unix(time1, 0)
	tm2 := time.Unix(time2, 0)

	fmt.Println(tm1)
	fmt.Println(tm2)
}
