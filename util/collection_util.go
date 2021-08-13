package util

func ConvertStrArr2Inter(strArr []string) []interface{} {
	s := make([]interface{}, len(strArr))
	for i, v := range strArr {
		s[i] = v
	}
	return s
}

func DiffStrArr(arr1, arr2 []string) []string {
	for _, arr2Ele := range arr2 {
		for arr1Index, arr1Ele := range arr1 {
			if arr1Ele == arr2Ele {
				arr1 = append(arr1[:arr1Index], arr1[arr1Index+1:]...)
			}
		}
	}
	return arr1
}
