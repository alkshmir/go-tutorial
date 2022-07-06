package foo

const (
	Max = 100 // 先頭大文字は外部から参照可
	min = 1   // 先頭小文字は外部から参照不可
)

func ReturnMin() int {
	return min
}
