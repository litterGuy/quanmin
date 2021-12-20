package utils

import (
	"bufio"
	"errors"
	"github.com/BurntSushi/toml"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 获取的当前运行路径
func GetWorkDir() string {
	workDir, _ := os.Getwd()
	return workDir
}

// 判断文件或文件夹否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

// 创建文件夹（递归）
// @param perm 权限 0777（读4写2执行1）
func CreateDir(dirPath string, perms ...os.FileMode) error {
	if dirPath == "" {
		return errors.New("path can not be empty")
	}
	ok := IsExist(dirPath)
	if ok {
		return nil
	}
	var perm os.FileMode
	if len(perms) > 0 {
		perm = perms[0]
	} else {
		perm = 0777
	}
	err := os.MkdirAll(dirPath, perm)
	if err != nil {
		return err
	}
	return nil
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

//写入文本文件内容
// @param force 文件夹不存在时自动创建
func WriteFile(filePath string, body string, forces ...bool) error {
	if len(forces) > 0 && forces[0] {
		dir := strings.Replace(filePath, `\`, "/", -1)
		index := strings.LastIndex(dir, "/")
		dir = dir[:index]
		err := CreateDir(dir)
		if err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filePath, []byte(body), 0777)
}

func WriteToml(path string, bean interface{}) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	err = toml.NewEncoder(w).Encode(bean)
	if err != nil {
		return err
	}
	return nil
}

//读取文本文件内容
func ReadFile(filePath string) (string, error) {
	if filePath == "" {
		return "", errors.New("path can not be empty")
	}
	ok := IsExist(filePath)
	if !ok {
		return "", errors.New("file does not exist, path：" + filePath)
	}
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// 删除文件或文件夹
// @param force 强制删除非空的文件夹
func Delete(path string, forces ...bool) error {
	if path == "" {
		return errors.New("path can not be empty")
	}
	ok := IsExist(path)
	if !ok {
		return nil
	}
	if len(forces) > 0 && forces[0] {
		return os.RemoveAll(path)
	} else {
		return os.Remove(path)
	}
}
