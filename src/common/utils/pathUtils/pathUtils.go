package pathUtils

import (
    "os"
    "path/filepath"
    "runtime"
    "strings"
)

func GetExecPath() string {
    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    return dir
}

func GetFilePath() string {
    // runtime.Caller()用户获取栈帧的信息, 参数为0表示当前函数, 参数为1表示上层调用该函数的函数.
    _, filename, _, _ := runtime.Caller(1)
    //return filepath.Dir(filename)
    return strings.Replace(filepath.Dir(filename), "\\", "/", -1)
}

func GetParentPath(path string) string {
    return substrCh(path, 0, strings.LastIndex(path, "/"))
}

// 中文不能使用string切片的方法来截取
func substrCh(s string, pos, length int) string {
    runes := []rune(s)
    l := pos + length
    if l > len(runes) {
        l = len(runes)
    }
    return string(runes[pos:l])
}

