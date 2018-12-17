package main

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"fmt"
	"os"
	"io/ioutil"
)

func stream(client oss.Client) {
	// 创建OSSClient实例。
	// 获取存储空间。
	bucket, err := client.Bucket("6by-data")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	lsRes, err := bucket.ListObjects(oss.MaxKeys(1000))
	fmt.Printf("开始遍历*********************")
	for _, object := range lsRes.Objects {
		body, err := bucket.GetObject(object.Key)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
		defer body.Close()
		data, err := ioutil.ReadAll(body)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
		fmt.Println("data:", string(data))
	}
	// 下载文件到流。

}