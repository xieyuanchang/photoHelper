// photoHelper project main.go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func showHeaderInfo() {
	fmt.Println("★★★★★★★★★★★★★★★★★★★★★★★★★")
	fmt.Println(" create by : xieyuanchang at 20150204")
	fmt.Println(" use for   : bakup photos by camera  ")
	fmt.Println("★★★★★★★★★★★★★★★★★★★★★★★★★")
}
func main() {
	showHeaderInfo()
	if len(os.Args) == 3 {
		fmt.Println("——————————————照片拷贝开始——————————————")
		fmt.Println("U盘文件夹：" + GetSrcPlace())
		fmt.Println("备份文件夹：" + GetMovePlace())
		MoveJPG(GetSrcPlace())
		fmt.Println("——————————————照片拷贝结束——————————————")
	} else {
		fmt.Println("请按下列方式指定参数！")
		fmt.Print(os.Args[0], " 【U盘文件夹】", " 【备份文件夹】", "\n")
	}
}

func GetSrcPlace() string {
	return os.Args[1]
}

func GetMovePlace() string {
	return os.Args[2]
}

func MoveJPG(dir string) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, info := range infos {
			if info.IsDir() {
				MoveJPG(dir + "\\" + info.Name())
			} else {
				if fileName := info.Name(); strings.Contains(fileName, "JPG") {
					fmt.Println(fileName)
					Copy(dir, info)
				}
			}
		}
	}
}

func Copy(fdir string, f os.FileInfo) {
	dir := GetMovePlace()
	tm := f.ModTime()
	year := tm.Format("2006")
	dirName := tm.Format("01月02日")
	dir = dir + "\\" + year + "\\" + dirName
	if _, err := os.Open(dir); err != nil {
		os.MkdirAll(dir, os.ModePerm)
	}
	fileName := dir + "\\" + f.Name()
	if _, err := os.Open(fileName); err != nil {
		srcFile, err := os.Open(fdir + "\\" + f.Name())
		if err != nil {
			fmt.Println(err)
		}
		defer srcFile.Close()
		desFile, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
		}
		defer desFile.Close()
		io.Copy(desFile, srcFile)
	}

}
