package test

import (
	"testing"
	"sync"
	"log"
	"time"
	"github.com/xytd/cm/misc"
	"github.com/xytd/cm/misc/rng"
)

func Test_lock(T *testing.T){
	sync:=&sync.Mutex{}
	sync.Lock()
	log.Println("hello")
	sync.Unlock()
}


func Test_timeDemo(T *testing.T){
	now:=misc.GetTime().Unix()

	before:=time.Now().Add(-time.Hour*5).Unix()
	log.Println((now-before)/3600/1000000)
}


func Test_rand_range(T *testing.T){

	log.Println(rng.IntRange(uint64(1),uint64(20)))
	log.Println(rng.IntRange(uint64(1),uint64(20)))
	log.Println(rng.IntRange(uint64(1),uint64(20)))
	log.Println(rng.IntRange(uint64(1),uint64(20)))
	log.Println(rng.IntRange(uint64(1),uint64(20)))

}