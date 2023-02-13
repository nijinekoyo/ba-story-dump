package main

import "os"

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
		File, err := os.OpenFile(StorysPath+Data.Title+".txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			return err
		}

		var Content string
		for _, DialogueText := range Data.DialogueText {
			Content += DialogueText + "\r"
		}

		ret, _ := File.Seek(0, os.SEEK_END)
		_, err = File.WriteAt([]byte(Content), ret)
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
