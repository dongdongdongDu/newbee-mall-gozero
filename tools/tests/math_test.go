package tests

import (
	"fmt"
	"math"
	"testing"
)

func TestMath(t *testing.T) {
	var totalCount int64 = 100
	var PageSize int64 = 33
	res1 := totalCount / PageSize
	res2 := float64(totalCount) / float64(PageSize)
	res3 := math.Ceil(res2)
	fmt.Println("res1=", res1)
	fmt.Println("res2=", res2)
	fmt.Println("res3=", res3)
}
