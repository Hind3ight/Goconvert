package fileUtils

import (
	"github.com/hind3ight/Goconvert/consts"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func GetFilesFromParams(arguments []string) []string {
	ret := make([]string, 0)

	for _, arg := range arguments {
		if strings.Index(arg, "-") != 0 {
			if arg == "." {
				arg = AbsolutePath(".")
			} else if strings.Index(arg, "."+consts.PthSep) == 0 {
				arg = AbsolutePath(".") + arg[2:]
			} else if !IsAbosutePath(arg) {
				arg = AbsolutePath(".") + arg
			}

			ret = append(ret, arg)
		} else {
			break
		}
	}

	return ret
}

func AbsolutePath(pth string) string {
	if !IsAbosutePath(pth) {
		pth, _ = filepath.Abs(pth)
	}

	pth = AddPathSepIfNeeded(pth)

	return pth
}

func IsAbosutePath(pth string) bool {
	return path.IsAbs(pth) ||
		strings.Index(pth, ":") == 1 // windows
}

func AddPathSepIfNeeded(pth string) string {
	sep := consts.PthSep

	if strings.LastIndex(pth, sep) < len(pth)-1 {
		pth += sep
	}
	return pth
}

func GetWorkDir() string { // where we run file in
	dir, _ := os.Getwd()

	dir, _ = filepath.Abs(dir)
	dir = AddSepIfNeeded(dir)

	return dir
}

func AddSepIfNeeded(pth string) string {
	if strings.LastIndex(pth, consts.PthSep) < len(pth)-1 {
		pth += consts.PthSep
	}
	return pth
}

func MkDirIfNeeded(dir string) error {
	if !FileExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		return err
	}

	return nil
}

func FileExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func WriteFile(filePath string, content string) {
	dir := filepath.Dir(filePath)
	MkDirIfNeeded(dir)

	var d1 = []byte(content)
	err2 := ioutil.WriteFile(filePath, d1, 0666) //写入文件(字节数组)
	check(err2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
