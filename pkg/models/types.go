package models

import (
	"encoding/json"
	"strconv"
)

// FlexString 可以解析为字符串或数组的类型
// 高德API中某些字段可能返回字符串或空数组
type FlexString string

// UnmarshalJSON 实现自定义 JSON 解析
func (fs *FlexString) UnmarshalJSON(data []byte) error {
	// 尝试解析为字符串
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		*fs = FlexString(s)
		return nil
	}

	// 尝试解析为数组
	var arr []string
	if err := json.Unmarshal(data, &arr); err == nil {
		if len(arr) > 0 {
			*fs = FlexString(arr[0])
		} else {
			*fs = ""
		}
		return nil
	}

	*fs = ""
	return nil
}

// IntOrString 可以解析为整数或字符串的类型
// 高德API中某些数字字段可能返回字符串格式
type IntOrString int

// UnmarshalJSON 实现自定义 JSON 解析
func (is *IntOrString) UnmarshalJSON(data []byte) error {
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*is = IntOrString(i)
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		val, _ := strconv.Atoi(s)
		*is = IntOrString(val)
		return nil
	}
	*is = 0
	return nil
}

// Float64OrString 可以解析为浮点数、字符串或数组的类型
// 高德API中某些浮点数字段可能返回字符串或数组格式
type Float64OrString float64

// UnmarshalJSON 实现自定义 JSON 解析
func (fs *Float64OrString) UnmarshalJSON(data []byte) error {
	var f float64
	if err := json.Unmarshal(data, &f); err == nil {
		*fs = Float64OrString(f)
		return nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		val, _ := strconv.ParseFloat(s, 64)
		*fs = Float64OrString(val)
		return nil
	}
	var arr []interface{}
	if err := json.Unmarshal(data, &arr); err == nil && len(arr) > 0 {
		if val, ok := arr[0].(float64); ok {
			*fs = Float64OrString(val)
			return nil
		}
	}
	*fs = 0
	return nil
}
