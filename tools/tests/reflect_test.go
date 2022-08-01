package tests

import (
	"fmt"
	"reflect"
	"testing"
)

type RegisterRequest struct {
	LoginName string `protobuf:"bytes,1,opt,name=loginName,proto3" json:"loginName,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func TestReflect(t *testing.T) {
	st := &RegisterRequest{LoginName: "15012341234", Password: "123456"}
	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st) // 获取reflect.Type类型

	t_kd := typ.Kind()
	v_kd := val.Kind() // 获取到st对应的类别
	//num := val.NumField()
	//fmt.Printf("%d\n", num)
	if v_kd != reflect.Struct {
		fmt.Println("v_kd不是struct", v_kd)
	} else {
		fmt.Println("v_kd是struct", v_kd)
	}
	if t_kd != reflect.Struct {
		fmt.Println("t_kd不是struct", t_kd)
	} else {
		fmt.Println("t_kd是struct", t_kd)
	}
}
