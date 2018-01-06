package stack


/*
	利用go設計一個簡單的堆棧
	堆棧的特點就是先進後廚
*/

type (
	node struct{
		data interface{} // 要保存的數據
		next * node // 下一個節點的位置s
	}

	Stack struct{
		count int // 當前堆棧中元素的數量
		head *node
	}

)

// 創建一個新的節點
func New()*Stack{
	s:=&Stack{}
	return s
}

// 返回節點的長度
func Len(s *Stack)int{
	return s.count
}

// 入棧
func(s *Stack)Push(item interface{}){
	n:=&node{data:item}
	if s.head==nil{  // 判断头结点是否为空
		s.head=n
	}else{
		n.next=s.head  // 新添加的节点指向之前的头结点
		s.head=n      // 让头结点指向新添加的节点
	}
	s.count++
}

// 出棧
func(s *Stack)Pop() interface{}{

	var n *node
	if s.head!=nil{
		n=s.head  // 将头结点赋值给一个临时节点
		s.head=n.next // 将指针指向新的头结点
		s.count--
	}

	if n==nil{
		return nil
	}

	return n.data
}

// 获取头结点的值
func(s *Stack)Peek()interface{}{
	n:=s.head
	if n==nil || n.data==nil{
		return nil
	}
	return n.data
}










