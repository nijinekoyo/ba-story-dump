/*
 * @Author: nijineko
 * @Date: 2023-02-13 20:44:35
 * @LastEditTime: 2023-02-13 21:06:58
 * @LastEditors: nijineko
 * @Description: 数据筛选
 * @FilePath: \StoryDump\DataFiltering.go
 */
package main

import (
	"strings"
)

type StoryData struct {
	Title        string   `json:"title"`
	DialogueText []string `json:"dialogueText"`
}

/**
 * @description: 筛选出剧情文本
 * @param {OriginalFile} OriginalFile
 * @return {*}
 */
func StoryDataFiltering(OriginalData OriginalFile) ([]StoryData, error) {
	var StorysData []StoryData

	var OneStoryData StoryData
	for _, Data := range OriginalData.DataList {
		ScriptData := strings.SplitN(Data.ScriptKr, ";", -1)
		switch ScriptData[0] {
		case "#title": // 剧情标题
			// 判断OneStoryData是否存在标题，不存在则将标题存入OneStoryData，存在则表示上一话已完成，将OneStoryData存入StorysData
			if OneStoryData.Title == "" {
				OneStoryData.Title = Data.TextJp
			} else {
				StorysData = append(StorysData, OneStoryData)
				// 清空OneStoryData
				OneStoryData = StoryData{}
				OneStoryData.Title = Data.TextJp
			}
		case "#hidemenu": // 隐藏菜单
		case "#wait": // 等待
		case "#all": // 全部
		case "#showmenu": // 显示菜单
		case "#place": // 场景文本
		default: // 判断为剧情文本
			if Data.TextJp != "" {
				OneStoryData.DialogueText = append(OneStoryData.DialogueText, Data.TextJp)
			}
		}
	}

	return StorysData, nil
}
