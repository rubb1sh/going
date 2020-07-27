package test

// ces

import (
	"fmt"
	"sync"
	"testing"
)

func TestMut(t *testing.T) {
	type Node struct {
		K int
	}
	var mu sync.Mutex // var初始化的变量是有零值的。
	a := 10
	fmt.Printf("%x\n", &a)
	mu.Lock()
	mu.Unlock()

}

type node struct {
	aa []int
}

func (this *node) setA(a []int) {
	this.aa = a
}

// 对于数组、map的赋值是指针赋值
func TestChangeArr(t *testing.T) {

	arr := make([]int, 0)
	arr = append(arr, 1)
	var n node
	n.setA(arr)
	fmt.Println(n.aa)
	fmt.Println("after change it.")
	arr[0] = 2
	fmt.Println(n.aa)
	arr = append(arr, 1, 2, 2, 3, 3, 4)
	fmt.Println("append more value it.")
	fmt.Println(n.aa) // 当arr扩容，arr的指针指向新的地址，而n.aa的指针还是指向原来。
}

// 可以通过实现Error() string方法定义一个错误，
// 便于上层判断
type DBConnectError struct {
	S string
}

func (this DBConnectError) Error() string {
	return this.S
}

func TestErr(t *testing.T) {
	str := `string'":":"';'"'` // 使用`可以避免转义
	fmt.Println(str)
}
