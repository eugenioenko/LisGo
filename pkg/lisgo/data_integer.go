package lisgo

import "strconv"

type LisgoInteger struct {
	Value int64
	DType int
}

func NewLisgoInteger(value int64) *LisgoInteger {
	return &LisgoInteger{Value: value, DType: LisgoTypeInteger}
}

func (data *LisgoInteger) ToString() string {
	return strconv.FormatInt(data.Value, 10)
}

func (data *LisgoInteger) ToBoolean() bool {
	return data.Value != 0
}

func (data *LisgoInteger) ToInteger() int64 {
	return data.Value
}

func (data *LisgoInteger) ToFloat() float64 {
	return float64(data.Value)
}

func (data *LisgoInteger) GetType() int {
	return data.DType
}

func (data *LisgoInteger) GetTypeName() string {
	return "integer"
}

func (data *LisgoInteger) GetValue() interface{} {
	return data.Value
}

func (data *LisgoInteger) Equals(other LisgoData) bool {
	return data.GetType() == other.GetType() && data.Value == other.GetValue()
}
