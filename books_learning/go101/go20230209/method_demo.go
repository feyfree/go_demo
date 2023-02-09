package go20230209

// Age Age和int是两个不同的类型。我们不能为int和*int
// 类型声明方法，但是可以为Age和*Age类型声明方法。
type Age int

func (age Age) LargerThan(a Age) bool {
	return age > a
}
func (age *Age) Increase() {
	*age++
}

// FilterFunc 为自定义的函数类型FilterFunc声明方法。
type FilterFunc func(in int) bool

func (ff FilterFunc) Filter(in int) bool {
	return ff(in)
}

// StringSet 为自定义的映射类型StringSet声明方法。
type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, present := ss[key]
	return present
}
func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}
func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

// 为自定义的结构体类型Book和它的指针类型*Book声明方法。

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}
