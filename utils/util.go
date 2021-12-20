package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tcolgate/mp3"
)

// RandString 生成随机字符串
func RandString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//PathExists 判断文件是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

//下载视频或者音频
func DownloadVideoOrAudio(ResourcesURL string, SavePath string) (string, error) {

	if find := strings.Contains(ResourcesURL, "local"); find {
		//该资源资源连接为本地文件
		fileNameTotalSplit := strings.SplitAfterN(ResourcesURL, "/", 4) //斜杠切割字符串 ，4为切割前3个斜杠，分成四个部分，返回slice,例子切割结果:[https:/, /,***********-shenzhen.aliyuncs.com/,voice/data/202105130005542EaJbdY7aH.wav]
		fileNameTotal := fileNameTotalSplit[len(fileNameTotalSplit)-1]  //取slice最后一个元素,去掉域名部分
		return "uploads/" + fileNameTotal, nil                          //返回文件路径和文件名
	}

	ext := path.Ext(ResourcesURL) //获取资源的后缀
	filePath := "uploads/" + SavePath + "/" + time.Now().Format("20060102") + "/"
	CreateDir(filePath)
	fileNameStr := time.Now().Format("20060102150405") + RandString(10)
	fileName := fileNameStr + ext
	localFile := filePath + fileName //生成文件夹和文件名
	//通过http请求获取图片的流文件
	resp, err := http.Get(ResourcesURL)
	if err != nil {
		return "资源请求失败", err
	}
	body, err := ioutil.ReadAll(resp.Body) //获取文件流
	if err != nil {
		return "资源数据流读取失败", err
	}
	out, err := os.Create(localFile) //创建文件
	if err != nil {
		return "保存资源文件创建失败", err
	}
	io.Copy(out, bytes.NewReader(body)) //写入文件

	return filePath + fileName, nil //返回文件路径和文件名
}

//获取MP3的时长（秒）
func Getmp3len(mp3Path string) (string, error) {
	t := 0.0
	r, err := os.Open(mp3Path)
	if err != nil {
		fmt.Println(err)
		return "mp3路径无效", err
	}
	d := mp3.NewDecoder(r)
	var f mp3.Frame
	skipped := 0
	for {

		if err := d.Decode(&f, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return "mp3文件无效", err
		}
		t = t + f.Duration().Seconds()
	}
	return strconv.Itoa(int(math.Ceil(t))), nil //浮点数向上取整转字符串
}

//，
func GetRandArray(origin []string, count int) []string {
	tmpArray := make([]string, len(origin)) //生成一个临时的切片，数组必须要指定长度，值类型
	copy(tmpArray, origin)                  //把传进来的切片复制到这个切片，切片是属于引用类型
	//一定要seed
	rand.Seed(time.Now().Unix()) //以当前时间为随机因子
	rand.Shuffle(len(tmpArray), func(i int, j int) { //把临时数组打乱
		tmpArray[i], tmpArray[j] = tmpArray[j], tmpArray[i]
	})

	resultArray := make([]string, 0, count) //生成一个长度为传入个数的结果切片
	for index, value := range tmpArray {    //把临时切片复制到结果切片，达到传入个数退出循环，返回结果
		if index == count {
			break
		}
		resultArray = append(resultArray, value)
	}
	return resultArray
}
