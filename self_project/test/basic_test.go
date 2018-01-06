package test

import (
	"testing"
	"log"
	"fmt"
	"strconv"
	"math"
	"math/rand"
)


type User struct{
	data map[int32]bool
}

func Test_user_demo(T *testing.T){

	tmp:=make(map[int32]bool)
	var user=&User{
		data:tmp,
	}

	user.data[1]=true
	user.data[2]=false
	log.Println(user)

	dst:=user.data
	dst[1]=false
	dst[2]=true
	log.Println(user.data)
}



func Test_math_demo(T *testing.T){

	fmt.Println(math.Ceil(1.5))
	fmt.Println(math.Floor(3.1))

	//data:=make(map[int]bool)
	//
	//data[-1]=true
	//data[-2]=true
	//data[1]=false
	//data[2]=false
	//fmt.Println(data)

	}

type Admin struct{
	Name int
	Sum []int
}


func Test_read_map(T *testing.T){
	data:=make(map[int]Admin)
	v,_:=data[1]
	log.Println(v.Name)

	for k,v:=range v.Sum{
		log.Printf("key=%d,value=%d",k,v)
	}

	log.Println("end....................end")
}




func Test_range_array(T *testing.T){
 data:=[]int32{1,2,3,4,5,6}

 fmt.Println(getIndexSum(data,3))
}

func getIndexSum(rate[]int32,index int32)(sum int32){

	rageLen:=int32(len(rate))
	for i:=index;i<rageLen;i++{
		var sum int32=0
		if i==0{
			return rate[i]
		}

		for tmp:=int32(0);tmp<=index;tmp++{
			sum+=rate[tmp]
		}
		return sum
	}
	return
}

func Test_map_demo2(T *testing.T){
	data:=make(map[int32]map[int32]int)

	data[0]=make(map[int32]int)
	data[0][0]+=2


	for _,v:=range data{



		for _,value:=range v{
			log.Println(value)
		}
	}
}


func Test_breakFor(T *testing.T){
	var count int32=0
	for count<30{
		for {
			count ++
			log.Printf("count=%d",count)
		}
	}

	log.Println("count=",count)
}

func Test(T *testing.T){
	data:=[]int{1,3,4,5,6,7,8}

	for k,v:=range data{

		if k==0{
			break
		}
		fmt.Println("key=",k)
		fmt.Println("value=",v)

	}
}


func Test_rand_next(T *testing.T){
	log.Println(rand.Int31n(10000))
	log.Println(rand.Int31n(10000))
	log.Println(rand.Int31n(10000))
	log.Println(rand.Int31n(10000))
}


// 当前章节(id-1)*9+当前关卡id

func Test_dive(T *testing.T){
	log.Println(2%9) // 章节id  1   1		当前


}



/*
	判断宝箱是否已经领取
	box map[int32]bool

	key:=当前章节id*100+关卡id

*/

func Test_translate(T *testing.T){
	//val,err:=strconv.Atoi("as")
	//
	//if err!=nil{
	//	log.Println(err)
	//	//return
	//}
	//
	//log.Println(val)
}

func Test_map_demo(T *testing.T){
	data:=make(map[int32]bool)
	log.Println(data)
	//_,ok:=data[100]
	//
	//
	//log.Println(ok)
	//data[1]=true
	//log.Println("len(data)=",len(data))



}



func Str2int64(str string)(result int64,err error){
	result,err=strconv.ParseInt(str,10,64)
	return

}


func Test_double_array(T *testing.T){
	fmt.Println(math.Pow(float64(10), 1.5))
	fmt.Println(float64(10)*1.5)

//var data [][]int
//	//data:=make([][]int,0)
//	data=append(data,[]int{1,2,3,4})
//	data=append(data,[]int{5,6,7,8,9})
//
//	for k,v:=range data[1]{
//
//		fmt.Println(k,v)
//	}


	//arr:=make([][]int32,6)

	//
	//for k,v:=range arr{
	//	log.Println(k,v)
	//}


	//data:=make(map[int32]bool,0)
	//data[1]=false
	//
	//for k,v:=range data{
	//	log.Println(k,v)
	//}
}


func Test_Slice_spread(T *testing.T){
	src:=make([]int,0)
	data:=make([]int,0)
	data=append(data,src...)
	log.Println(data)
}


type admin struct {
	src map[int32]bool
}

 const name  = iota+1800

 func Test_demo1(T *testing.T){
 	log.Println(name)
 }

func Test_map_demo3(T *testing.T){
	data:=make(map[int32]string)
	data[1]="hello"
	v,ok:=data[1]
	log.Println(v,ok)

	//a :=new(admin)
	////a.src[1]
	//log.Println(a.src[1])
}

type student struct{
	name string
	sex string
}

func Test_demo2(T *testing.T){

	arr:=make(map[int]map[int32]*student)
	data:=make(map[int32]*student)

	for i:=int32(1);i<3;i++{
		data[i]=&student{name:fmt.Sprint("tom",i),sex:fmt.Sprint("male",i)}
	}
	arr[1027]=data

	log.Println("data=",data)
	log.Println(arr)

}


func Test_read_from_map(T *testing.T){
	data:=make(map[int]string)
	data[1]="hello"

	k,v:=data[2]
	log.Println(k,v)


}


func Test_Place_Option_demo1(T *testing.T){

	a:=15
	b:=14

	fmt.Println(a&b)

	c:=4
	fmt.Println(^c)


}

func Test_int(T *testing.T){
	var a uint8=12
	var b uint8=30

	log.Println("a=",a)
	log.Println("b=",b)
}

func Test_if_else(T *testing.T){
	if true{
		defer fmt.Printf("1")
	}else{
		defer fmt.Printf("2")
	}
	fmt.Printf("3")
}