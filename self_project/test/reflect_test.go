package test

import (
	"strconv"
	"testing"
	"reflect"
	"log"
)

type myInt int

func(this *myInt)ParseInt(str string)(result int ,err error){

	result,err=strconv.Atoi(str)
	return
}

func (this *myInt)Int2Str(num int,tt string)(result string){

	result=strconv.Itoa(num)
	return
}


func Test_reult_demo(T *testing.T){
	var s myInt

	v:=reflect.TypeOf(&s)
	for i:=0;i<v.NumMethod();i++{
		m:=v.Method(i)
		log.Println(m.Name)
	}

	t:=reflect.ValueOf(&s)
	//
	for i:=0;i<t.NumMethod();i++{
		m:=t.Method(i)

		for j:=0;j<m.Type().NumIn();j++{

			log.Println(m.Type().In(j).String())
		}
	}



}