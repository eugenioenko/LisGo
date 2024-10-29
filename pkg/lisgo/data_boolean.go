package lisgo

type LisgoBoolean struct {
	Value bool
	DType int
}

func NewLisgoBoolean(value bool) *LisgoBoolean {
	return &LisgoBoolean{Value: value, DType: LisgoTypeBoolean}
}
func (data *LisgoBoolean) ToString() string {
	if data.Value {
		return "true"
	}
	return "false"
}

func (data *LisgoBoolean) ToBoolean() bool {
	return data.Value
}

func (data *LisgoBoolean) ToInteger() int64 {
	if data.Value {
		return 1
	}
	return 0
}

func (data *LisgoBoolean) ToFloat() float64 {
	if data.Value {
		return 1
	}
	return 0
}

func (data *LisgoBoolean) GetType() int {
	return data.DType
}

func (data *LisgoBoolean) GetTypeName() string {
	return "boolean"
}

func (data *LisgoBoolean) GetValue() any {
	return data.Value
}

func (data *LisgoBoolean) Equals(other LisgoData) bool {
	return data.GetType() == other.GetType() && data.Value == other.GetValue()
}
