package main

import (
	"fmt"
	"os"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os/exec"
	"strings"
)

func getCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}
func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func listbucket() {
	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", "LTAIhKBvrxkYoAy6", "YzL66Ubh0iwOOD3KJdp4TqvuBVpDRp")
	if err != nil {
		handleError(err)
	}
	lsRes, err := client.ListBuckets()
	if err != nil {
		handleError(err)
	}
	for _, bucket := range lsRes.Buckets {
		fmt.Println("bucket:", bucket.Name)
	}
}
func listobjKey() {
	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", "LTAIhKBvrxkYoAy6", "YzL66Ubh0iwOOD3KJdp4TqvuBVpDRp")
	if err != nil {
		handleError(err)
	}
	bucket, err := client.Bucket("jhg-apk")
	if err != nil {
		handleError(err)
	}
	lsRes, err := bucket.ListObjects()
	if err != nil {
		handleError(err)
	}
	for _, object := range lsRes.Objects {
		fmt.Println(object.Key, object.LastModified)
	}
}

// AccessKeyID：LTAIhKBvrxkYoAy6 AccessKeySecret：YzL66Ubh0iwOOD3KJdp4TqvuBVpDRp
func upfile(filename, path string) {

	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", "LTAIhKBvrxkYoAy6", "YzL66Ubh0iwOOD3KJdp4TqvuBVpDRp")
	if err != nil {
		// HandleError(err)
	}

	bucket, err := client.Bucket("jhg-apk")
	if err != nil {
		// HandleError(err)
	}
	err = bucket.PutObjectFromFile(filename, path)
	if err != nil {
		handleError(err)
	}
	fmt.Printf("%v上传成功", filename)
}

//var list  *string = flag.String("l", "jhg-apk", "输入要查看的目录，默认为jhg-apk")
//var path  *string = flag.String("o", "pwd","上传文件路径， 默认为当前路径")
func main() {
	// flag.Parse()
	// listobjKey()
	reader := os.Args
	if len(os.Args) <2{
		panic("输入错误")
	}
	switch reader[1] {
	case "list":
		listobjKey()
	case "-o":
		if len(os.Args) ==3{
			path := reader[2]
			filename :=strings.Split(path,`\`)
			upfile(filename[len(filename)-1],path)
		}
	default:
		fmt.Println("list参数列出服务器当前文件")
		fmt.Println("-o 指定上传文件, 例如: oss.exe -o aaaa.txt")
		fmt.Println("其他未知错误，请联系管理员！")
	}
}