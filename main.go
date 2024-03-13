package main

import (
	"mp4WithSrt/merge"
)

func main() {
	folderPath := "/mnt/f/large/GirlFriend4ever/G4E" // 指定文件夹路径
	extension := ".mp4"                              // 指定扩展名

	merge.Merge(folderPath, extension)
}
