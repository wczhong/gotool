package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 程序离不开业务的需求
// 文件处理
// option： recursion:""
// option： extension:[]string
// default no recursion
func ReadAllFileInDir(dir string, options map[string]interface{}) []string {
	DIRECTORY_RECURSION := "recursion"
	FILE_EXTENSION := "extension"
	filePathSlice := []string{}
	var f func(string, map[string]interface{})
	f = func(dir string, options map[string]interface{}) {
		_, err := os.Stat(dir)
		if err != nil {
			return
		}
		eles, err := ioutil.ReadDir(dir)
		if err != nil {
			return
		}
		for _, ele := range eles {
			if ele.IsDir() {
				if _, ok := options[DIRECTORY_RECURSION]; ok {
					f(filepath.Join(dir, ele.Name()), options)
				}
				continue
			} else {
				if ext, ok := options[FILE_EXTENSION]; ok {
					file_ext := filepath.Ext(ele.Name())
					ext_a, _ := ext.([]string)
					for _, value := range ext_a {
						if file_ext == value {
							filePathSlice = append(filePathSlice, filepath.Join(dir, ele.Name()))
						} else {
							continue
						}
					}
				} else {
					filePathSlice = append(filePathSlice, filepath.Join(dir, ele.Name()))
				}
			}
		}
	}
	f(dir, options)
	return filePathSlice
}

func main() {
	dir := "C:\\detections-master"
	options := map[string]interface{}{}
	options["recursion"] = ""
	options["extension"] = []string{".csv"}
	filePathSlice := ReadAllFileInDir(dir, options)
	for _, v := range filePathSlice {
		fmt.Println(v)
	}

}
