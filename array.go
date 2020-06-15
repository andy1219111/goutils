package utils

import (
	"reflect"
	"strconv"
)

//Contains 判断切片中是否包含某元素
func Contains(arrayOrMap interface{}, val interface{}) (isExist bool) {
	isExist = false
	switch reflect.TypeOf(arrayOrMap).Kind() {
	case reflect.Slice:
		{
			s := reflect.ValueOf(arrayOrMap)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(val, s.Index(i).Interface()) {
					return true
				}
			}
		}
	case reflect.Map:
		s := reflect.ValueOf(arrayOrMap).MapRange()
		for s.Next() {
			v := s.Value()
			if reflect.DeepEqual(val, v.Interface()) {
				return true
			}
		}
	}
	return
}

//GetDiffElement 计算两个数组或者切片的差集
func GetDiffElement(arr1, arr2 []string) []string {
	var shortOne []string
	var longOne []string
	var different []string
	if len(arr1) >= len(arr2) {
		shortOne = arr2
		longOne = arr1
	} else {
		shortOne = arr1
		longOne = arr2
	}

	for _, one := range longOne {
		if !Contains(shortOne, one) {
			different = append(different, one)
		}
	}

	return different
}

//StringArray2Int 将字符串数组转化为整形数组
func StringArray2Int(data []string) []int {

	var result []int
	for _, i := range data {
		num, err := strconv.Atoi(i)
		if err != nil {
			continue
		}
		result = append(result, num)
	}

	return result
}
