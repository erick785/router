package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	. "github.com/erick785/router/importJson"
	. "github.com/erick785/router/route"
)

func main() {
	// 1. 导入json数据
	importJsonData := NewImportData()
	readfile(importJsonData)

	// 2.添加到route
	route := NewRoute()
	addRouter(importJsonData, route)

	// 3.获取路由

	// 4.输出路由到文件
}

func readfile(importJsonData *ImportData) {
	// 设置要读取的文件夹路径
	folderPath := "./pairData"

	// 获取文件夹下所有文件
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println("读取文件夹失败:", err)
		return
	}

	// 遍历文件
	for _, file := range files {
		// 检查文件是否为.json后缀
		if filepath.Ext(file.Name()) == ".json" {
			// 读取文件内容
			filePath := folderPath + "/" + file.Name()
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println("读取文件失败:", err)
				continue
			}
			if err := importJsonData.Import(content); err != nil {
				fmt.Println("导入数据失败:", err)
				continue
			}
		}
	}
}

func addRouter(importJsonData *ImportData, route *Route) {
	for token0token1, _ := range importJsonData.Pairs {
		tokens := strings.Split(token0token1, "_")
		link := NewNodeLink(tokens[0], []string{tokens[1]})
		route.UpdateNetworkTopology(link)
	}
}
