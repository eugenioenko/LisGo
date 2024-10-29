package lisgo

type LisgoNull struct {
	DType int
}

func NewLisgoNull() *LisgoNull {
	return &LisgoNull{DType: LisgoTypeNull}
}

func (data *LisgoNull) ToString() string {
	return "null"
}

func (data *LisgoNull) ToBoolean() bool {
	return false
}

func (data *LisgoNull) ToInteger() int64 {
	return 0
}

func (data *LisgoNull) ToFloat() float64 {
	return 0
}

func (data *LisgoNull) GetType() int {
	return data.DType
}

func (data *LisgoNull) GetTypeName() string {
	return "null"
}

func (data *LisgoNull) GetValue() any {
	panic("Cant GetValue of Null")
}

func (data *LisgoNull) Equals(other LisgoData) bool {
	return data.GetType() == other.GetType()
}
