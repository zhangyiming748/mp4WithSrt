package merge

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Merge(dir, pattern string) {

	files, err := getFilesWithExtension(dir, pattern)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		srt := strings.Replace(file, ".mp4", ".srt", 1)
		if isExist(srt) {
			subtitles := strings.Join([]string{"subtitles=", srt, ",scale=-1:1080"}, "")
			output := strings.Replace(file, ".mp4", "_with_subtitle.mp4", 1)
			cmd := exec.Command("ffmpeg", "-i", file, "-vf", subtitles, "-c:v", "libx265", "-c:a", "aac", "-ac", "1", "-tag:v", "hvc1", output)
			fmt.Printf("生成的命令: %s\n", cmd.String())
			combinedOutput, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("命令执行失败: %s\n", err.Error())
				continue
			} else {
				fmt.Printf("命令成功执行: %s\n", string(combinedOutput))
				//os.Remove(file)
			}
		}
	}
}
func isExist(fp string) bool {
	_, err := os.Stat(fp)
	if os.IsNotExist(err) {
		fmt.Printf("%s 对应的字幕文件不存在\n", fp)
		return false
	} else {
		fmt.Printf("%s 对应的字幕文件存在\n", fp)
		return true
	}
}
func getFilesWithExtension(folderPath string, extension string) ([]string, error) {
	var files []string
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), extension) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
