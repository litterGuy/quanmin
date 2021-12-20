package utils

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

var execPath string
var pwd string

func init() {
	pwd = GetWorkDir()
	execPath = filepath.Join(pwd, "libs", "ffmpeg.exe")
}

// VideoHstack 两个视频左右分屏
func VideoHstack(out, v1, v2 string) error {
	pams := []string{
		"-i",
		v1,
		"-i",
		v2,
		"-filter_complex",
		"hstack",
		"-y",
		out,
	}
	rt, err := Command(execPath, pams...)
	fmt.Println(rt)
	return err
}

// VideoJoin 视频连接
func VideoJoin(out string, videos ...string) error {
	var err error
	if len(videos) == 0 {
		return errors.New("videos length is empty")
	}
	CreateDir(filepath.Join(pwd, "out", "tmp"))
	defer Delete(filepath.Join(pwd, "out", "tmp"))
	tsPath := []string{}
	for i, s := range videos {
		out_ts := filepath.Join(pwd, "out", "tmp", fmt.Sprintf("%v-%v.ts", time.Now().Unix(), i))
		//out_ts = filepath.ToSlash(out_ts)
		tsPath = append(tsPath, out_ts)
		err = video2ts(s, out_ts)
		if err != nil {
			return err
		}
	}
	//out = filepath.ToSlash(out)

	pams := []string{
		"-i",
		`concat:` + strings.Join(tsPath, "|") + ``,
		"-acodec",
		"copy",
		"-vcodec",
		"copy",
		"-absf",
		"aac_adtstoasc",
		"-y",
		out,
	}
	rt, err := Command(execPath, pams...)
	fmt.Println(rt)
	return err
}

// video2ts 视频转化成ts格式，后续处理视频速度快，文件小
func video2ts(v1 string, out string) error {
	pams := []string{
		// -i 2.mp4 -vcodec copy -acodec copy -vbsf h264_mp4toannexb 2.ts
		"-i",
		v1,
		"-vcodec",
		"copy",
		"-acodec",
		"copy",
		"-vbsf",
		"h264_mp4toannexb",
		out,
	}
	rt, err := Command(execPath, pams...)
	fmt.Println(rt)
	return err
}
