package util

// ConvertStrArr2Inter
//  @Description:
//  @param strArr
//  @return []interface{}
func ConvertStrArr2Inter(strArr []string) []interface{} {
	s := make([]interface{}, len(strArr))
	for i, v := range strArr {
		s[i] = v
	}
	return s
}

// DiffStrArr
//  @Description: Get the different in arr1
//  @param arr1
//  @param arr2
//  @return []string
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
