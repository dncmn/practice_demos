package main

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"log"
)

func main(){
	tt:=gorm.Logger{}
	log.Println(tt)
	beego.Run()
}

