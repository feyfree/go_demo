package pointer

func f() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}

//  The expression new(T) creates an unnamed variable of typ e T,
//  initializes it to the zero value of T, and return it's
//  address, which is a value of typ e *T
func newInt() *int {
	return new(int)
}
