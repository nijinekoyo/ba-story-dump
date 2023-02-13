/*
 * @Author: nijineko
 * @Date: 2023-02-13 20:20:45
 * @LastEditTime: 2023-02-13 20:41:15
 * @LastEditors: nijineko
 * @Description: read all files
 * @FilePath: \StoryDump\fileRead.go
 */
package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	FilesPath = "./data_file"
)

type OriginalFile struct {
	DataList []struct {
		GroupId        int    `json:"GroupId"`
		SelectionGroup int    `json:"SelectionGroup"`
		BGMId          int    `json:"BGMId"`
		Sound          string `json:"Sound"`
		Transition     int    `json:"Transition"`
		BGName         int    `json:"BGName"`
		BGEffect       int    `json:"BGEffect"`
		PopupFileName  string `json:"PopupFileName"`
		ScriptKr       string `json:"ScriptKr"`
		TextJp         string `json:"TextJp"`
		VoiceJp        string `json:"VoiceJp"`
	} `json:"DataList"`
}

/**
 * @description: 读取FilesPath下的所有文件并写入内存
 * @return {OriginalFile} OriginalFile 读取的文件内容
 * @return {error} error 错误信息
 */
func FilesRead() (OriginalFile, error) {
	FileList, err := GetFilePaths(FilesPath)
	if err != nil {
		return OriginalFile{}, err
	}

	var FileData OriginalFile

	for _, FilePath := range FileList {
		File, err := os.OpenFile(FilePath, os.O_RDONLY, 0666)
		if err != nil {
			return OriginalFile{}, err
		}

		FileContent, err := ioutil.ReadAll(File)
		if err != nil {
			return OriginalFile{}, err
		}

		// 解析Json
		var Original OriginalFile
		err = json.Unmarshal(FileContent, &Original)
		if err != nil {
			return OriginalFile{}, err
		}

		// 追加进DataList
		FileData.DataList = append(FileData.DataList, Original.DataList...)
	}

	return FileData, nil
}

/**
 * @description: 遍历出所有文件路径
 * @param {string} DirPth
 * @return {*}
 */
func GetFilePaths(DirPth string) ([]string, error) {
	DirPth = filepath.Clean(DirPth)
	var dirs []string
	dir, err := ioutil.ReadDir(DirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	var Files []string
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, filepath.Clean(DirPth+PthSep+fi.Name()))
			GetFilePaths(DirPth + PthSep + fi.Name())
		} else {
			Files = append(Files, filepath.Clean(DirPth+PthSep+fi.Name()))
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetFilePaths(table)
		for _, temp1 := range temp {
			Files = append(Files, filepath.Clean(temp1))
		}
	}

	return Files, nil
}
