package some_code

import "testing"

// 当一个协程的栈的大小（因为栈增长或者收缩而）改变时，
// 一个新的内存段将申请给此栈使用。
// 原先已经开辟在老的内存段上的内存块将很有可能被转移到新的内存段上，
// 或者说这些内存块的地址将改变。
// 相应地，引用着这些开辟在此栈上的内存块的指针（它们同样开辟在此栈上）
// 中存储的地址也将得到刷新

// 下面这行是为了防止f函数的调用被内联。
//
//go:noinline
func f(i int) byte {
	var a [1 << 20]byte // 使栈增长
	return a[i]
}

func TestMove(t *testing.T) {
	var x int
	println(&x)
	f(100)
	println(&x)
}
