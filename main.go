package main

import (
	"path/filepath"
	"quanmin/utils"
)

func main() {
	path := utils.GetWorkDir()
	//err := utils.VideoHstack(filepath.Join(path, "out", "hstack.mp4"), filepath.Join(path, "data", "1.mp4"), filepath.Join(path, "data", "2.mp4"))
	//if err != nil {
	//	panic(err)
	//}

	err := utils.VideoJoin(filepath.Join(path, "out", "join.mp4"), filepath.Join(path, "data", "1.mp4"), filepath.Join(path, "data", "2.mp4"))
	if err != nil {
		panic(err)
	}
}
