package lisgo

func MathReduce(items []LisgoData, reducer func(int64, LisgoData, int) int64, initial int64) int64 {
	accumulator := initial

	for index, item := range items {
		accumulator = reducer(accumulator, item, index)
	}
	return accumulator
}

func Reduce[T any, R any](items []T, reducer func(acc R, item T, index int) R, initial R) R {
	acc := initial

	for i, item := range items {
		acc = reducer(acc, item, i)
	}
	return acc
}

func Every[T any](items []T, predicate func(T, int) bool) bool {
	for index, item := range items {
		if !predicate(item, index) {
			return false
		}
	}
	return true
}

func Some[T any](items []T, predicate func(T, int) bool) bool {
	for index, item := range items {
		if predicate(item, index) {
			return true
		}
	}
	return false
}

func Map[T any, R any](items []T, transform func(T) R) []R {
	result := make([]R, len(items))
	for index, item := range items {
		result[index] = transform(item)
	}
	return result
}

func ParamsSomeAreType(params []LisgoData, lisgoType int) bool {
	return Some(params, func(param LisgoData, _ int) bool {
		return param.GetType() == lisgoType
	})
}

func ParamsEveryAreType(params []LisgoData, lisgoType int) bool {
	return Every(params, func(param LisgoData, _ int) bool {
		return param.GetType() == lisgoType
	})
}

/*


func Filter(vs []LisgoData, f func(LisgoData) bool) []LisgoData {
	vsf := make([]LisgoData, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

*/
