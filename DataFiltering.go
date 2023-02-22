/*
 * @Author: nijineko
 * @Date: 2023-02-13 20:44:35
 * @LastEditTime: 2023-02-22 19:03:43
 * @LastEditors: nijineko
 * @Description: 数据筛选
 * @FilePath: \StoryDump\DataFiltering.go
 */
package main

import (
	"regexp"
	"strconv"
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
	var TitleNum int = 1
	for _, Data := range OriginalData.DataList {
		ScriptData := strings.SplitN(Data.ScriptKr, ";", -1)
		switch ScriptData[0] {
		case "#title": // 剧情标题
			// 判断OneStoryData是否存在标题，不存在则将标题存入OneStoryData，存在则表示上一话已完成，将OneStoryData存入StorysData
			if OneStoryData.Title == "" {
				OneStoryData.Title = FilterLabelData(strconv.Itoa(TitleNum) + " " + Data.TextJp)
				TitleNum++
			} else {
				StorysData = append(StorysData, OneStoryData)
				// 清空OneStoryData
				OneStoryData = StoryData{}
				OneStoryData.Title = FilterLabelData(strconv.Itoa(TitleNum) + " " + Data.TextJp)
				TitleNum++
			}
		case "#hidemenu": // 隐藏菜单
		case "#wait": // 等待
		case "#all": // 全部
		case "#showmenu": // 显示菜单
		case "#place": // 场景文本
		case "#video": // 视频
		default: // 判断为剧情文本
			if Data.TextJp != "" {
				// 清理文本
				CleanString := func(Text string) string {
					if len(FindRubyLabel(Text)) != 0 {
						// 替换文本的ruby标签
						for _, RubyLabel := range FindRubyLabel(Text) {
							Text = strings.Replace(Text, RubyLabel, "（"+FindLabelParameter(RubyLabel)+"）", -1)
						}
					}

					Text = FilterLabelData(Text)                   // 去除文本中的标签
					Text = strings.Replace(Text, "#n", "", -1) // 去除换行符

					return Text
				}

				// 判断是否启用过滤器
				if Flags.Filter {
					// 过滤文本中带特殊标签的文本
					Label := []string{
						"[s]",
						"[s1]",
						"[s2]",
						"[s3]",
						"[s4]",
						"[s5]",
						"[s6]",
						"[s7]",
						"[s8]",
						"[s9]",
						"[ns]",
						"[ns1]",
						"[ns2]",
						"[ns3]",
						"[ns4]",
						"[ns5]",
						"[ns6]",
						"[ns7]",
						"[ns8]",
						"[ns9]",
					} // 标签列表
					if Find := CheckArray(FindLabel(Data.TextJp), Label); !Find {
						Text := CleanString(Data.TextJp)
						OneStoryData.DialogueText = append(OneStoryData.DialogueText, Text)
					}
				} else {
					Text := CleanString(Data.TextJp)
					OneStoryData.DialogueText = append(OneStoryData.DialogueText, Text)
				}
			}
		}
	}

	return StorysData, nil
}

/**
 * @description: 去除文本中的标签
 * @param {string} Content 文本内容
 * @return {string} 去除标签后的文本
 */
func FilterLabelData(Content string) string {
	Regexp := regexp.MustCompile(`\[[^\]]*\]`)
	Replaced := Regexp.ReplaceAllString(Content, "")

	return Replaced
}

/**
 * @description: 提取文本中的标签
 * @param {string} Content
 * @return {[]string} 标签列表
 */
func FindLabel(Content string) []string {
	Regexp := regexp.MustCompile(`\[[^\]]*\]`)
	Finds := Regexp.FindAllString(Content, -1)

	return Finds
}

/**
 * @description: 提取文本中的Ruby标签
 * @param {string} Content
 * @return {[]string} 标签列表
 */
func FindRubyLabel(Content string) []string {
	Regexp := regexp.MustCompile(`\[ruby=[^\]]*\]`)
	Finds := Regexp.FindAllString(Content, -1)

	return Finds
}

/**
 * @description: 提取提取标签参数
 * @param {string} Content 文本内容
 * @return {string} 标签参数
 */
func FindLabelParameter(Content string) string {
	ContentSplit := strings.SplitN(Content, "=", -1)
	Parameter := strings.Replace(ContentSplit[1], "]", "", -1)

	return Parameter
}

/**
 * @description: 利用map判断源数组中是否存在目标数组的某个值
 * @param {[]string} SourceArray 源数组
 * @param {[]string} TargetArray 目标数组
 * @return {bool} 是否存在
 */
func CheckArray(SourceArray []string, TargetArray []string) bool {
	Map := make(map[string]struct{})

	for _, Value := range SourceArray {
		Map[Value] = struct{}{}
	}

	for _, Value := range TargetArray {
		if _, ok := Map[Value]; ok {
			return true
		} else {
			continue
		}
	}

	return false
}
