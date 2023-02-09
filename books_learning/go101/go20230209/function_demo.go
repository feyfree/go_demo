package go20230209

func HalfAndNegative(n int) (int, int) {
	return n / 2, -n
}

func AddSub(a, b int) (int, int) {
	return a + b, a - b
}

func Dummy(values ...int) {}

func main() {
	// 这几行编译没问题。
	AddSub(HalfAndNegative(6))
	AddSub(AddSub(AddSub(7, 5)))
	AddSub(AddSub(HalfAndNegative(6)))
	Dummy(HalfAndNegative(6))
	_, _ = AddSub(7, 5)

	// 下面这几行编译不通过。
	/*
		_, _, _ = 6, AddSub(7, 5)
		Dummy(AddSub(7, 5), 9)
		Dummy(AddSub(7, 5), HalfAndNegative(6))
	*/
}

func Double(n int) int {
	return n + n
}

func Apply(n int, f func(int) int) int {
	return f(n) // f的类型为"func(int) int"
}
