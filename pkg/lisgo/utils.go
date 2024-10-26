package lisgo

/*
	func Reduce(items []LisgoData, reducer func(LisgoData, LisgoData, int, []LisgoData) LisgoData, initial LisgoData) LisgoData {
		accumulator := initial

		for index, item := range items {
			accumulator = reducer(accumulator, item, index, items)
		}
		return accumulator
	}

	func RReduce(items []interface{}, reducer func(interface{}, interface{}, int, []interface{}) interface{}, initial interface{}) interface{} {
		accumulator := initial

		for index, item := range items {
			accumulator = reducer(accumulator, item, index, items)
		}
		return accumulator
	}
*/
func MathReduce(items []LisgoData, reducer func(int64, LisgoData, int) int64, initial int64) int64 {
	accumulator := initial

	for index, item := range items {
		accumulator = reducer(accumulator, item, index)
	}
	return accumulator
}

func Every(items []LisgoData, f func(LisgoData, int) bool) bool {
	for index, v := range items {
		if !f(v, index) {
			return false
		}
	}
	return true
}

/*
func Any(vs []LisgoData, f func(LisgoData) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []LisgoData, f func(LisgoData) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []LisgoData, f func(LisgoData) bool) []LisgoData {
	vsf := make([]LisgoData, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []LisgoData, f func(LisgoData) LisgoData) []LisgoData {
	vsm := make([]LisgoData, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
*/
