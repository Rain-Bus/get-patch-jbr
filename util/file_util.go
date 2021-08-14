package util

import "os"

// Exist
//  @Description: Validate the file exist
//  @param filename
//  @return bool
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
