package lisgo

type Callable func(*Interpreter, []Expression) LisgoData

type LisgoCallable struct {
	function Callable
	name     string
	DType    int
}

func NewLisgoCallable(name string, function Callable) *LisgoCallable {
	return &LisgoCallable{DType: LisgoTypeCallable, name: name, function: function}
}

func (data *LisgoCallable) ToString() string {
	return data.name
}

func (data *LisgoCallable) ToBoolean() bool {
	return true
}

func (data *LisgoCallable) ToInteger() int64 {
	return 0
}

func (data *LisgoCallable) ToFloat() float64 {
	return 0
}

func (data *LisgoCallable) GetType() int {
	return data.DType
}

func (data *LisgoCallable) GetTypeName() string {
	return "function"
}

func (data *LisgoCallable) GetValue() any {
	return data.function
}

func (data *LisgoCallable) Equals(other LisgoData) bool {
	return other.GetType() == LisgoTypeFunction && data.GetValue() == other.GetValue()
}
