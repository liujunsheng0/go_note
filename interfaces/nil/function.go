package nil

import "reflect"

// 判断接口是否为空指针的正确方法
func IsNil(i interface{}) bool {
	if i == nil {
		return true
	}
	vi := reflect.ValueOf(i)
	return vi.Kind() == reflect.Ptr && vi.IsNil()
}
