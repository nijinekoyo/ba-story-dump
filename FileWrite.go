/*
 * @Author: nijineko
 * @Date: 2023-02-13 21:10:02
 * @LastEditTime: 2023-02-20 23:23:00
 * @LastEditors: nijineko
 * @Description: 写入文件
 * @FilePath: \StoryDump\FileWrite.go
 */
package main

import (
	"os"
)

var (
	StorysPath = "./storys/"
)

/**
 * @description: 写入剧情文本
 * @param {[]StoryData} Datas
 * @return {*}
 */
func WriteStorysData(Datas []StoryData) error {
	CreateFolder(StorysPath)

	for _, Data := range Datas {
		var Content string
		for _, DialogueText := range Data.DialogueText {
			Content += DialogueText + "\r"
		}

		err := os.WriteFile(StorysPath+Data.Title+".txt", []byte(Content), 0666)
		if err != nil {
			return err
		}
	}

	return nil
}

/**
 * @description: 创建文件夹
 * @param {string} Path 文件夹路径
 * @return {error} 错误信息
 */
func CreateFolder(Path string) error {
	_, err := os.Stat(Path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err := os.MkdirAll(Path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
