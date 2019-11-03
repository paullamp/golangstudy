package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", dirHandleFunc)
	err := http.ListenAndServe("0.0.0.0:9981", nil)
	if err != nil {
		fmt.Println("Listen in http service failed")
	}
}

func dirHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(walkDirToString("/opt")))

}

func walkDirToString(dir string) (fileNameList string) {
	dirHandler, openerr := os.Open(dir)
	if openerr != nil {
		fmt.Println("open dir failed")
		return
	}
	filelist, readdirerr := dirHandler.Readdirnames(-1)
	if readdirerr != nil {
		fmt.Println("readdir err:", readdirerr)
		return
	}
	for _, filename := range filelist {
		fileNameList = fileNameList + packFileInfo(dir+string(os.PathSeparator)+filename)
	}
	return fileNameList
}

func packFileInfo(filename string) (filecontent string) {
	fi, fierr := os.Lstat(filename)
	if fierr != nil {
		fmt.Println("fi err:", fierr)
		return
	}
	return fmt.Sprintf("%-30s\t\t%-15d byte\t %-10s %s\n", fi.Name(), fi.Size(),
		getType(fi), fi.ModTime().Format("2006/1/2 15:04:05"))
}

func walkFolder(dirName string) {
	dirHandler, openerr := os.Open(dirName)
	if openerr != nil {
		fmt.Println("Open dir err:", openerr)
		return
	}
	filenames, readdirerr := dirHandler.Readdirnames(-1)
	if readdirerr != nil {
		fmt.Println("read dir err:", readdirerr)
		return
	}
	fmt.Printf("%-30s\t\t%-15s byte\t %-10s %s\n", "FileName", "FileSize",
		"FileType", "FileModifyTime")
	for _, filename := range filenames {
		showFileInfo(dirName + string(os.PathSeparator) + filename)
	}

	// dirHandler.Readdir()
}

func showFileInfo(filename string) {
	fi, err := os.Lstat(filename)
	if err != nil {
		fmt.Println("get Lstat failed:", err)
		return
	}

	fmt.Printf("%-30s\t\t%-15d byte\t %-10s %s\n", fi.Name(), fi.Size(),
		getType(fi), fi.ModTime().Format("2006/1/2 15:04:05"))
}

func getType(fi os.FileInfo) string {
	s := fi.IsDir()
	switch s {
	case true:
		return "Folder"
	case false:
		return "File"
	}
	return ""
}
