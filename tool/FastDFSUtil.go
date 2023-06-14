package tool

import (
	"bufio"
	"github.com/tedcy/fdfs_client"
	"log"
	"os"
	"strings"
)

// 用代码实现fastDFS进行文件上传，上传到fastDFS分布式文件系统中
func UploadFile(fileName string) string {
	client, err := fdfs_client.NewClientWithConfig("./config/fdfs.conf")
	defer client.Destory()
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	fileId, err := client.UploadByFilename(fileName)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return fileId
}

/*
*
从配置文件中读取分布式文件系统服务器的ip和端口等内容
*/
func FileServerAddr() string {
	file, err := os.Open("./config/fastdfs.conf")
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2)
		switch str[0] {
		case "http_server_port":
			return str[1]
		}
		if err != nil {
			return ""
		}
	}
}
