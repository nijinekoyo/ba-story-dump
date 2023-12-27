/*
 * @Author: nijineko
 * @Date: 2023-02-13 20:15:07
 * @LastEditTime: 2023-12-28 00:02:52
 * @LastEditors: nijineko
 * @Description: main file
 * @FilePath: \StoryDump\main.go
 */
package main

import (
	"fmt"
)

func main() {
	// 初始化参数
	err := InitFlag()
	if err != nil {
		panic(err)
	}

	// 读取数据库文件
	ScenarioScriptData, ScenarioCharacterNameData, err := SQLRead()
	if err != nil {
		panic(err)
	}

	// 初始化角色信息本地化
	InitCharacterInfoLocalization(ScenarioCharacterNameData)

	// 剧情文本筛选
	StorysData, err := StoryDataFiltering(ScenarioScriptData)
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
