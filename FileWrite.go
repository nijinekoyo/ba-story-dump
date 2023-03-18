/*
 * @Author: nijineko
 * @Date: 2023-02-13 21:10:02
 * @LastEditTime: 2023-03-18 16:12:47
 * @LastEditors: nijineko
 * @Description: 写入文件
 * @FilePath: \StoryDump\FileWrite.go
 */
package main

import (
	"os"
	"path"
	"strconv"
	"strings"
)

var (
	StorysPath = "./storys/"
)

/**
 * @description: 写入剧情文本
 * @param {[]StoryData} Datas
 * @return {*}
 */
func WriteStorysData(Datas map[int]StoryData) error {
	CreateFolder(StorysPath)

	for Index, Data := range Datas {
		var Content string
		for _, DialogueText := range Data.DialogueText {
			Content += DialogueText + "\r"
		}

		// 去除标题中不能存在的符号
		Data.Title = strings.Replace(Data.Title, "/", "", -1)

		err := os.WriteFile(path.Join(StorysPath, strconv.Itoa(Index)+" "+Data.Title+".txt"), []byte(Content), 0666)
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
