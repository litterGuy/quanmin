package utils

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"runtime"
	"time"
)

/**
go 执行命令时，存在管道符无法处理的问题
*/

//设置超时时间秒
var timeout = 600 * time.Second

//执行命令并添加超时检测
func Command(name string, arg ...string) (string, error) {

	ctxt, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() //releases resources if slowOperation completes before timeout elapses

	bash := "/bin/bash"
	arg_type := "-c"
	if runtime.GOOS == "windows" {
		bash = "cmd"
		arg_type = "/C"
	}
	cmdline := []string{}
	cmdline = append(cmdline, arg_type)
	cmdline = append(cmdline, name)
	for _, s := range arg {
		cmdline = append(cmdline, s)
	}

	cmd := exec.CommandContext(ctxt, bash, cmdline...)
	fmt.Println(cmd.String())
	//当经过Timeout时间后，程序依然没有运行完，则会杀掉进程，ctxt也会有err信息
	var out bytes.Buffer
	var errbuf bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errbuf

	//当经过Timeout时间后，程序依然没有运行完，则会杀掉进程，ctxt也会有err信息
	if err := cmd.Run(); err != nil {
		//检测报错是否是因为超时引起的
		if ctxt.Err() != nil && ctxt.Err() == context.DeadlineExceeded {
			return "", errors.New("command timeout")
		}
		if errbuf.Len() == 0 {
			return ConvertByte2String(out.Bytes(), GB18030), err
		}
		return ConvertByte2String(out.Bytes(), GB18030), errors.New(ConvertByte2String(errbuf.Bytes(), GB18030))
	} else {
		return ConvertByte2String(out.Bytes(), GB18030), nil
	}
}

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
