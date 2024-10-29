package lisgo

import "strconv"

type LisgoString struct {
	Value string
	DType int
}

func NewLisgoString(value string) *LisgoString {
	return &LisgoString{Value: value, DType: LisgoTypeString}
}
func (data *LisgoString) ToString() string {
	return data.Value
}

func (data *LisgoString) ToBoolean() bool {
	return len(data.Value) > 0
}

func (data *LisgoString) ToInteger() int64 {
	value, err := strconv.ParseInt(data.Value, 10, 64)
	if err != nil {
		panic("Cant convert string to int")
	}
	return value
}

func (data *LisgoString) ToFloat() float64 {
	value, err := strconv.ParseFloat(data.Value, 64)
	if err != nil {
		panic("Cant convert string to float")
	}
	return value
}

func (data *LisgoString) GetType() int {
	return data.DType
}

func (data *LisgoString) GetTypeName() string {
	return "string"
}

func (data *LisgoString) GetValue() any {
	return data.Value
}

func (data *LisgoString) Equals(other LisgoData) bool {
	return other.GetType() == data.DType && other.ToString() == data.Value
}
