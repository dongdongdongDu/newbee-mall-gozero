package tests

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	goodsCategoryId := 1
	keyword := "衣服"
	conditions := fmt.Sprintf(" where `goods_category_id` = %d and (`goods_name` like \"%%%s%%\" or `goods_intro` like \"%%%s%%\")", goodsCategoryId, keyword, keyword)
	fmt.Println(conditions)
}
