package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

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

	fmt.Println("tokens 数量: ", importJsonData.GetTokensLen())

	fmt.Println("pairs 数量: ", importJsonData.GetPairsLen())

	fmt.Println("route 数量: ", len(route.GetNetworkTopology()))

	// 3.获取路由
	result := getRouter(importJsonData, route)

	fmt.Println("生成路由完毕开始写入文件...")

	// 4.输出路由到文件
	writeRouterfile(result)

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
	importJsonData.PairsForEach(func(key, value interface{}) bool {
		token0token1 := key.(string)
		tokens := strings.Split(token0token1, "_")
		link := NewNodeLink(tokens[0], []string{tokens[1]})
		route.UpdateNetworkTopology(link)
		return true
	})
}

// 获取路由
func getRouter(importJsonData *ImportData, route *Route) []interface{} {
	var wg sync.WaitGroup
	var result []interface{}
	importJsonData.TokensForEach(func(key, value interface{}) bool {
		token := key.(string)
		wg.Add(1)
		go func(result []interface{}, importJsonData *ImportData, route *Route) {
			importJsonData.TokensForEach(func(key, value interface{}) bool {
				anotherToken := key.(string)
				var routeResult []interface{}

				if token == anotherToken {
					return true
				}

				routers := route.GetRoutes(token, anotherToken)
				if len(routers) == 0 {
					return true
				}

				for _, router := range routers {
					pairs := []interface{}{}
					for i := 0; i < len(router)-1; i++ {
						pair := importJsonData.GetPair(router[i], router[i+1])
						pairs = append(pairs, pair)
					}
					routeResult = append(routeResult, pairs)
					routeResult = append(routeResult, router)
				}

				result = append(result, routeResult)
				return true
			})
			wg.Done()
		}(result, importJsonData, route)

		return true
	})

	wg.Wait()

	// for token := range importJsonData.Tokens {
	// 	for anotherToken := range importJsonData.Tokens {
	// 		var routeResult []interface{}

	// 		if token == anotherToken {
	// 			continue
	// 		}
	// 		routers := route.GetRoutes(token, anotherToken)
	// 		if len(routers) == 0 {
	// 			continue
	// 		}
	// 		for _, router := range routers {
	// 			pairs := []interface{}{}
	// 			for i := 0; i < len(router)-1; i++ {
	// 				pair := importJsonData.GetPair(router[i], router[i+1])
	// 				pairs = append(pairs, pair)
	// 			}
	// 			routeResult = append(routeResult, pairs)
	// 			routeResult = append(routeResult, router)
	// 		}

	// 		result = append(result, routeResult)
	// 	}
	// }

	return result
}

// 输出路由到文件
func writeRouterfile(res []interface{}) {
	file, err := os.OpenFile("merge_route.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, r := range res {
		data, err := json.Marshal(r)
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = writer.WriteString(string(data) + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}

}
