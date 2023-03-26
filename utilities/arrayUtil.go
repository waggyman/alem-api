package utilities

import "reflect"

func InArray(val interface{}, array interface{}) (index int) {
	values := reflect.ValueOf(array)

	if reflect.TypeOf(array).Kind() == reflect.Slice || values.Len() > 0 {
		for i := 0; i < values.Len(); i++ {
			if reflect.DeepEqual(val, values.Index(i).Interface()) {
				return i
			}
		}
	}

	return -1
}

func RemoveByIndex(arr []string, pos int) []string {
	new_arr := make([]string, (len(arr) - 1))
	k := 0
	for i := 0; i < (len(arr) - 1); {
		if i != pos {
			new_arr[i] = arr[k]
			k++
		} else {
			k++
		}
		i++
	}

	return new_arr
}
