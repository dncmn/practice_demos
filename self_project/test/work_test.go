package test

import (
	"github.com/XanthusL/zset"
	"testing"

	"github.com/xytd/cm/misc"
	"log"
	"reflect"
	"time"
	//"unsafe"
	"fmt"
	"math"
	"sort"
)

func Test_transLate(T *testing.T) {
	var data float32 = 0.58
	log.Println(data)
	log.Println(int64(data * float32(100)))
}

func Test_coll(T *testing.T) {
	s := zset.New()
	s.Set(50, 1001, "English")
	s.Set(80, 1002, "math")
	s.Set(95, 1003, "Chinese")
	s.Set(65, 1004, "PE")

	id, score, value := s.GetDataByRank(0, false)
	log.Println("id=", id, "  score=", score, "  value=", value)

	//fmt.Println(s.GetData(1004))
	////log.Printf("collection's length=%d",s.Length())
	//log.Println(s.GetRank(1001,true)) // true:降顺,false:升顺
	//log.Println(s.GetData(1001)) //  返回对应的数据,和布尔值,判断是否存在
	//log.Println(s.GetDataByRank(1003,false))// returns the id,score and extra data of an element
	//if ok:=s.Delete(1001);ok{
	//	log.Println("delete success")
	//}
	//
	//log.Println(s.GetData(1001))
	//s.Set(66,1002,"mathes")
	//log.Println("collection'length=========",s.Length())
	//log.Println(s.GetRank(1002,false))

}

func Test_Get_TimeDemo(T *testing.T) {
	log.Println(misc.GetTimestamp())
}

func Test_for_demo(T *testing.T) {
	for {
		time.Sleep(time.Second * 2)
		log.Println(time.Now())
	}
}

func Test_return_val(T *testing.T) {
	var param int = 2
	Selfprint(param)
	log.Println("end")
}

func Selfprint(param int) {
	if param == 0 {
		return
	}

	log.Print("hello world")
}

func Test_basic2(T *testing.T) {
	//hitLvl := int32(0x7fffffff)
	//log.Println(hitLvl)

	data := make(map[int32]map[int32]string)

	tmp := make(map[int32]string)
	tmp[2] = "world"
	data[3] = tmp

	for v := range data {
		log.Println(v)
	}
}

func Test_TransLate(T *testing.T) {
	//var num1 float32= 0.58
	//var num2 float64 = 0.58
	//
	//log.Println(int64(num1*100))
	//log.Println(int64(num2*100))
	//
	//fmt.Println("hello")
	//
	//type MyInt int
	//var num3 MyInt =2
	//var num4 int =4
	//log.Println(num3,num4)
	//
	//
	//type Name1 map[string]string
	//type Name2 map[string]string
	//type Alias = map[string]string

	type MyInt = int // 定义一个新类型
	var val MyInt = 2
	var num int = val
	log.Println("num=", num, "  reflect.TypeOf(val)=", reflect.TypeOf(val))

}

func Test_daily1(T *testing.T) {
	//var src = map[int32]string{
	//	1:"one",
	//	2:"two",
	//}
	//log.Println("before=",src)
	//changeValue(src)
	//log.Println("after=",src)
	//
	//val,ok:=src[3]
	//log.Println("value=",val, "  isExists=",ok)
	//
	//for k,v:=range src{
	//	log.Printf("key=%d,value=%s",k,v)
	//}
	//var test map[int32]int32
	//test:=make(map[int32]int32)
	////test[1]+=2
	//log.Println(test)
	//
	//test2:=new([2]int)
	//test2[0]+=1
	//test2[1]+=2
	//
	//
	//log.Println("test2=",test2)
	//
	//
	//log.Println(unsafe.Pointer(&test),"  ",unsafe.Sizeof(&test))
	//
	//
	//arr:=make([]int,10,5)
	//log.Println("len(arr)=",len(arr)," cap(arr)=",cap(arr))

	arr := new([5]int)
	log.Println(arr)

	data := make([]int, 5)
	log.Println(data)

	t1 := make(chan int, 3)
	log.Println(t1)
	t2 := make(map[int]int)
	log.Println(t2)

}

func changeValue(data map[int32]string) {
	for key := range data {
		data[key] = "hello world"
	}
}

func Test_Zero(T *testing.T) {
	//data:=make(map[int32]bool)
	//log.Println(data[1])

	var m map[int32]string
	log.Println(m[1] == "") // read from blank map
	//m[1]="string"
	log.Println(m)

	c := map[int]string{}
	log.Println("c=", c)
}

func Test_rang_map(T *testing.T) {
	data := make(map[int]bool)
	data[1] = true
	data[2] = false
	data[3] = true

	//l := len(data)
	//for i := 1; i <= l; i++ {
	//
	//	if data[i] { // 表示这个值是存在的
	//		log.Printf("key %d  exists", i)
	//		continue
	//	}
	//
	//	data[i] = true
	//	log.Printf("key %d already add in map!", i)
	//
	//}

	for k, v := range data {
		if v {
			log.Printf("key= %d is exists,value=%v", k, v)
			continue
		}

		log.Printf("key=%d is not exists", k)
	}

}

func Test_sort_demo(T *testing.T) {

	log.Println(1001 * 1000)
	data := make([]int, 0)

	data = append(data, 2)
	data = append(data, 1)
	data = append(data, 9)
	data = append(data, 5)
	data = append(data, 100)
	sort.Slice(data, func(i, j int) bool {
		return i <= j
	})

	log.Println(data)
}

func Test_basic_variable(T *testing.T) {
	var num int32
	log.Println("num=", num)
}

func Test_YearDay(T *testing.T) {
	//log.Println(time.Now().YearDay())
	log.Println(misc.GetTime().Add(-24 * time.Hour).YearDay())
	log.Println(misc.GetTime().YearDay())
	log.Println(misc.GetTime().Add(24 * time.Hour).YearDay())

}

func Test_range_slice_demo2(T *testing.T) {
	data := make([]int, 0)

	for i := 1; i <= 4; i++ {
		data = append(data, i)
	}

	for v := range data {
		log.Println(v)
	}
}

const (
	AA int = iota
	BB
	CC
	DD int = 20
	EE int = iota
	SUM
)

func Test_itoa_demo(T *testing.T) {
	fmt.Println(AA)
	fmt.Println(BB)
	fmt.Println(CC)
	fmt.Println(DD)
	fmt.Println(EE)
	fmt.Println(SUM)
}

func Test_math_value(T *testing.T) {
	fmt.Println(math.Ceil(float64(14) / float64(30)))
}

func Test_slice_demo1(T *testing.T) {
	b := []int{1, 2, 3, 4, 5}
	log.Printf("b[:]=%v,reflect.TypeOf([b[:])=%v", b[:], reflect.TypeOf(b[:]).Kind())
}

func Test_append_slice(T *testing.T) {
	a := make([]int, 0)
	b := make([]int, 0)

	for i := 1; i <= 6; i++ {
		a = append(a, i)
	}

	a = append(a, b...)
	log.Println("a=", len(a))
}

// 测试在切片中三个冒号的用法

func Test_three_colons(T *testing.T) {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	v := data[6:8]
	fmt.Println(v)
	fmt.Println(len(v))
	fmt.Println(cap(v))

	v = data[:6:8]
	fmt.Println(v)
	fmt.Println(len(v))
	fmt.Println(cap(v))

	v = append(v, 6, 7, 8, 9)
	fmt.Println(v)
	fmt.Println(len(v))
	fmt.Println(cap(v))
}

func Test_slice_demo2(T *testing.T) {
	//data := make([]int, 0)
	//
	//for i := 1; i <= 4; i++ {
	//	data = append(data, i)
	//}

	data := [...]int{1, 2, 3, 4}
	v := data[:2:3]
	fmt.Println(v)
	fmt.Println(len(v))
	fmt.Println(cap(v))
}

func Test_slice_demo3(T *testing.T) {
	var array [10]int
	for i := 1; i <= 10; i++ {
		array[i-1] = i
	}

	fmt.Printf("origin-array=%v\n", array)

	slice := array[2:4]
	fmt.Printf("slice1=%v\n", slice) // 2 3

	fmt.Printf("len(slice)=%d,cap(slice)=%d\n", len(slice), cap(slice)) // 2  8

	slice = array[2:4:7]
	fmt.Printf("slice1=%v\n", slice)                                    // 2 3
	fmt.Printf("len(slice)=%d,cap(slice)=%d\n", len(slice), cap(slice)) // 2   5
}
