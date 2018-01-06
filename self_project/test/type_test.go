package test

import (
	"testing"
	"log"
)


func Test_read_Map(T *testing.T){
	data:=make(map[int32]int32)
	data[1]=2
	v,ok:=data[2]
	log.Println(v,ok)
	tmp:=make(map[int]bool)
	tmp[1]=true

	val,oka:=tmp[1]
	log.Println(val,oka)
}

type  MyType func(int,map[int]string)

type Mod struct {
	ipc []MyType
}

func Test_new_type( T *testing.T){
	t:=&Mod{}

	src:=make([]MyType,1024)

	src[0]=demo
	t.ipc=src
	t.ipc[0](1,map[int]string{1:"tomcat"})


}


func demo(index int,reward map[int]string){

	log.Printf("index=%d",index)
	log.Println("reward=",reward)

}
