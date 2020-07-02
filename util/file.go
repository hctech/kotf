package util

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreatePath(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func MoveFileToPath(sourcePath string, filename string, targetPath string) error {

	fileRead, err := os.Open(path.Join(sourcePath, filename))
	if err != nil {
		fmt.Println("Open err:", err)
		return err
	}
	defer fileRead.Close()

	fileWrite, err := os.Create(path.Join(targetPath, filename))
	if err != nil {
		fmt.Println("Create err:", err)
		return err
	}
	defer fileWrite.Close()

	buf := make([]byte, 4096)
	for {
		n, err := fileRead.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			return nil
		}
		_, err = fileWrite.Write(buf[:n])
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			return nil
		}
	}
}

func CoverFileVars(filePath string, vars map[string]interface{}, targetPath string) error {
	f, err := template.ParseFiles(filePath)
	if err != nil {
		return err
	}

	t, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	err = f.Execute(t, vars)
	if err != nil {
		return err
	}
	t.Close()
	return nil
}
