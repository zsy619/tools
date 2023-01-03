package lists

type List[T any] []T

func (l List[T]) Filter(f func(T) bool) List[T] {
	var dst []T
	for _, v := range l {
		if f(v) {
			dst = append(dst, v)
		}
	}
	return List[T](dst)
}
