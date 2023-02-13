/*
 * @Author: nijineko
 * @Date: 2023-02-13 20:15:07
 * @LastEditTime: 2023-02-13 20:18:13
 * @LastEditors: nijineko
 * @Description: main file
 * @FilePath: \StoryDump\main.go
 */
package main

import "fmt"

func main() {
	OriginalFileData, err := FilesRead()
	if err != nil {
		panic(err)
	}

	// 剧情文本筛选
	StorysData, err := StoryDataFiltering(OriginalFileData)
	if err != nil {
		panic(err)
	}

	// 写入剧情文本
	err = WriteStorysData(StorysData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")
}
