package lisgo

import "strconv"

type LisgoFloat struct {
	Value float64
	DType int
}

func NewLisgoFloat(value float64) *LisgoFloat {
	return &LisgoFloat{Value: value, DType: LisgoTypeFloat}
}

func (data *LisgoFloat) ToString() string {
	return strconv.FormatFloat(data.Value, 'E', -1, 64)
}

func (data *LisgoFloat) ToBoolean() bool {
	return data.Value != 0
}

func (data *LisgoFloat) ToInteger() int64 {
	return int64(data.Value)
}

func (data *LisgoFloat) ToFloat() float64 {
	return data.Value
}

func (data *LisgoFloat) GetType() int {
	return data.DType
}

func (data *LisgoFloat) GetTypeName() string {
	return "float"
}

func (data *LisgoFloat) GetValue() interface{} {
	return data.Value
}

func (data *LisgoFloat) Equals(other LisgoData) bool {
	return data.GetType() == other.GetType() && data.Value == other.GetValue()
}
