package common

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"remote-part-job-back/config"
	"strconv"
	"strings"
	"time"
)

// 获取随机数 纯文字
func GetRandomString(n int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// sha1加密
func Sha1En(data string) string {
	t := sha1.New() ///产生一个散列值得方式
	_, _ = io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func SplitImgUrls(urls string) []string {
	return strings.Split(urls, ",")
}

func MergeImgUrls(urls []string) string {
	return strings.Join(urls, ",")
}

func GenFilePathAndUrl(imageName string, options ...string) (path, url string) {
	fileName := strings.Join(options, "-")
	if fileName != "" {
		fileName += "-"
	}
	if imageName != "" {
		fileName += imageName
	} else {
		fileName += strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	fileName += "-" + GetRandomString(5) + ".jpg"
	path = config.ConfigHolder.ImageRoot + "/" + fileName
	url = config.ConfigHolder.ServiceDomain + "/remote_part_job/api/v1/image/download?image=" + fileName
	return
}

func GetPathByFileName(fileName string) string {
	return config.ConfigHolder.ImageRoot + "/" + fileName
}

// ConvertImageSize 图片压缩函数
// targetSize单位为KB
func ConvertImageSize(inputFile string) {
	outputFile := inputFile
	// 获取原始图片大小
	fileInfo, err := os.Stat(inputFile)
	if err != nil {
		return
	}
	originalSize := fileInfo.Size()
	sizeLimit := config.ConfigHolder.ImageSizeLimit

	top := 0
	for originalSize > int64(sizeLimit*1024) && top < 5 {
		// 根据压缩比例调整图片质量
		quality := "90"

		// 使用 convert 命令进行压缩
		cmd := exec.Command("convert", inputFile, "-quality", quality, outputFile)
		err = cmd.Run()
		if err != nil {
			fmt.Println("convert error: ", err)
		}
		fmt.Println("top=", top, quality, originalSize)
		fileInfo, _ = os.Stat(inputFile)
		originalSize = fileInfo.Size()
		fmt.Println("top=", top, quality, originalSize)
		top += 1
	}
}
