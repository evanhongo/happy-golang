package array

// Public visibility
func Sum(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}
	return total
}

func Map[T1, T2 any](arr []T1, fn func(T1) T2) []T2 {
	newArr := make([]T2, len(arr))
	for i := 0; i < len(arr); i++ {
		newArr[i] = fn(arr[i])
	}
	return newArr
}

func Filter[T any](arr []T, fn func(T) bool) []T {
	newArr := make([]T, 0)
	for i := 0; i < len(arr); i++ {
		if fn(arr[i]) {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

func Any[T any](arr []T, fn func(T) bool) bool {
	for i := 0; i < len(arr); i++ {
		if fn(arr[i]) {
			return true
		}
	}
	return false
}

func All[T any](arr []T, fn func(T) bool) bool {
	for i := 0; i < len(arr); i++ {
		if !fn(arr[i]) {
			return false
		}
	}
	return true
}
