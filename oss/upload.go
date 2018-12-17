package main

import (
	"fmt"
	"os"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"strings"
	"sync"
)

var wg sync.WaitGroup
//var logger = logzap.Initloger()
//func init(){
//	log.SetPrefix("[oss]")
//	log.SetFlags(log.LstdFlags | log.Lshortfile |log.LUTC)
//}
func handleError(err error) {
	//fmt.Println("Error:", err)
	//logger.Info("i'm ok")
	os.Exit(-1)
}
func listbucket(client oss.Client) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		handleError(err)
	}
	for _, bucket := range lsRes.Buckets {
		fmt.Println("bucket:", bucket.Name)
	}
}
func listdir(client oss.Client, bucket oss.Bucket) {

}
func listobjKey(client oss.Client, bucket oss.Bucket) {

	lsRes, err := bucket.ListObjects()
	if err != nil {
		handleError(err)
	}
	for _, object := range lsRes.Objects {
		fmt.Println(object.Key, object.LastModified)
	}
}

func down(client oss.Client, lsRes oss.ListObjectsResult, bucket oss.Bucket) {
	for _, object := range lsRes.Objects {
		err := bucket.GetObjectToFile(object.Key, object.Key)
		if err != nil {
			fmt.Println("DownLoad Error:", err)
			//os.Exit(-1)
			return
		}
	}
}

// AccessKeyID：vLTAIhKBvrxkYoAy61v AccessKeySecret：vYzL66Ubh0iwOOD3KJdp4TqvuBVpDRp1v
//del1v
func upfile(filename, path string) {

	//client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", "LTAIhKBvrxkYoAy6", "YzL66Ubh0iwOOD3KJdp4TqvuBVpDRp")
	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", "accessKeyID=vLTAIhhNw61bkGd1k1v", "accessKeySecret=co8NC33EeWaxYTllPkrNgePj4fRMQXw1v")

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

	client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", "FLTAIhhNw61bkGd1k1v", "KVo8NC33EeWaxYTllPkrNgePj4fRMQXw1v")
	bucket, err := client.Bucket("6by-data")
	if err != nil {
		panic("初始化oss失败.")
	}
	reader := os.Args
	if len(os.Args) < 2 {
		panic("输入错误")
	}
	switch reader[1] {
	case "list":
		listobjKey(*client, *bucket)
	case "download":
		fmt.Printf("******你选择了下载************\n")
		task := make(chan oss.ListObjectsResult, 10000)
		marker := oss.Marker("")

		go func() {
			for {
				lsRes, err := bucket.ListObjects(oss.MaxKeys(50), marker)
				if err != nil {
					fmt.Println("读取bucket错误,Error:", err)
					os.Exit(-1)
				}
				task <- lsRes
				marker = oss.Marker(lsRes.NextMarker)
				if !lsRes.IsTruncated {
					close(task)
					break
				}
			}
		}()
		fmt.Printf("******开始下载************\n")
		//logger.Info("开始下载")
		wg.Add(8)
		for i := 0; i < 8; i++ {
			go func() {
				defer wg.Done()
				for obj := range task {
					down(*client, obj, *bucket)
				}
			}()
		}
		wg.Wait()

	case "-o":
		if len(os.Args) == 3 {
			path := reader[2]
			filename := strings.Split(path, `\`)
			upfile(filename[len(filename)-1], path)
		}
	default:
		fmt.Println("list参数列出服务器当前文件")
		fmt.Println("-o 指定上传文件, 例如: oss.exe -o aaaa.txt")
		fmt.Println("其他未知错误，请联系管理员！")
	}
}
