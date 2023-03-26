/*
 * @Author: nijineko
 * @Date: 2023-02-13 20:44:35
 * @LastEditTime: 2023-03-26 20:17:53
 * @LastEditors: nijineko
 * @Description: 数据筛选
 * @FilePath: \StoryDump\DataFiltering.go
 */
package main

import (
	"fmt"
	"regexp"
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
func StoryDataFiltering(OriginalData OriginalFile) (map[int]StoryData, error) {
	StorysDatas := make(map[int]StoryData)

	for _, Data := range OriginalData.DataList {
		ScriptData := strings.SplitN(Data.ScriptKr, "\n", 2)
		if len(ScriptData) == 2 {
			ScriptData = strings.SplitN(ScriptData[1], ";", -1)
		} else {
			ScriptData = strings.SplitN(Data.ScriptKr, ";", -1)
		}
		switch ScriptData[0] {
		case "#title": // 剧情标题
			// 将NowGroupID与标题文本写入对应GroupId的Title中
			StorysDatas[Data.GroupId] = StoryData{
				Title: FilterLabelData(Data.TextJp),
			}
		case "#Title": // 大写的也是剧情标题
			// 将NowGroupID与标题文本写入对应GroupId的Title中
			StorysDatas[Data.GroupId] = StoryData{
				Title: FilterLabelData(Data.TextJp),
			}
		case "#hidemenu": // 隐藏菜单
		case "#wait": // 等待
		case "#showmenu": // 显示菜单
		case "#place": // 场景文本
		case "#video": // 视频
		default: // 判断为剧情文本
			if Data.TextJp != "" {
				// 清理文本
				CleanString := func(Text string) string {
					if Flags.AddRuby {
						if len(FindRubyLabel(Text)) != 0 {
							// 替换文本的ruby标签
							for _, RubyLabel := range FindRubyLabel(Text) {
								Text = strings.Replace(Text, RubyLabel, "（"+FindLabelParameter(RubyLabel)+"）", -1)
							}
						}
					}

					Text = FilterLabelData(Text)               // 去除文本中的标签
					Text = strings.Replace(Text, "#n", "", -1) // 去除换行符

					// 按照换行符分割,去除开头的空格
					TextArray := strings.Split(Text, "\n")
					if len(TextArray) != 0 {
						var TextTrimSpace string
						for Index, Value := range TextArray {
							TextTrimSpace += strings.TrimLeft(Value, " ")
							if Index != len(TextArray)-1 {
								TextTrimSpace += "\n"
							}
						}
						Text = TextTrimSpace
					} else {
						Text = strings.TrimLeft(Text, " ")
					}

					return Text
				}

				var Text string

				if Flags.AddCharacterName {
					// 匹配获取角色名字，反转数组获取说话人
					for _, Value := range InversionArray(ScriptData) {
						// 如果匹配到#则表示是对话文本，忽略
						if Find := strings.Contains(Value, "#"); Find {
							continue
						}

						// 尝试匹配角色名字
						NameJP, NicknameJP := CharacterNameKRToJP(Value)
						if NameJP != "" {
							// 如果匹配到角色名字，则将角色名字添加到文本前面
							if NicknameJP == "" {
								Text = fmt.Sprintf("%s：", NameJP)
								break
							}
							Text = fmt.Sprintf("%s（%s）：", NameJP, NicknameJP)
							break
						}
					}
				}

				// 样式提示
				var StyleTips string

				// 判断是否启用字体大小提示
				if Flags.AddFontSizeTips {
					for Index, Value := range ScriptData {
						// 判断是否存在字体大小标签
						if Find := strings.Contains(Value, "#fontsize"); Find {
							// 如果存在，则提取下一个元素为字体大小
							FontSize := ScriptData[Index+1]

							// 写入样式提示
							StyleTips += fmt.Sprintf("FontSize:%s,", FontSize)
						}
					}
				}

				// 清理样式提示的最后一个逗号
				if StyleTips != "" {
					StyleTips = strings.TrimRight(StyleTips, ",")
				}

				// 判断是否启用Tag过滤器
				if Flags.FilterTag {
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
						// 清理文本
						Text += CleanString(Data.TextJp)
						// 追加样式提示
						if StyleTips != "" {
							Text += " (" + StyleTips + ")"
						}

						// 将GroupId的最后一位数修正为0后判断是否存在
						if _, Find := StorysDatas[Data.GroupId-(Data.GroupId%10)]; Find {
							// 如果存在，则将文本追加到GroupId的最后一位数修正为0的对话文本中
							StorysDatas[Data.GroupId-(Data.GroupId%10)] = StoryData{
								Title:        StorysDatas[Data.GroupId-(Data.GroupId%10)].Title,
								DialogueText: append(StorysDatas[Data.GroupId-(Data.GroupId%10)].DialogueText, Text),
							}
						} else {
							// 如果不存在，则将文本追加到GroupId的对话文本中
							StorysDatas[Data.GroupId] = StoryData{
								Title:        StorysDatas[Data.GroupId].Title,
								DialogueText: append(StorysDatas[Data.GroupId].DialogueText, Text),
							}
						}
					}
				} else {
					// 清理文本
					Text += CleanString(Data.TextJp)
					// 追加样式提示
					if StyleTips != "" {
						Text += " (" + StyleTips + ")"
					}

					// 将GroupId的最后一位数修正为0后判断是否存在
					if _, Find := StorysDatas[Data.GroupId-(Data.GroupId%10)]; Find {
						// 如果存在，则将文本追加到GroupId的最后一位数修正为0的对话文本中
						StorysDatas[Data.GroupId-(Data.GroupId%10)] = StoryData{
							Title:        StorysDatas[Data.GroupId-(Data.GroupId%10)].Title,
							DialogueText: append(StorysDatas[Data.GroupId-(Data.GroupId%10)].DialogueText, Text),
						}
					} else {
						// 如果不存在，则将文本追加到GroupId的对话文本中
						StorysDatas[Data.GroupId] = StoryData{
							Title:        StorysDatas[Data.GroupId].Title,
							DialogueText: append(StorysDatas[Data.GroupId].DialogueText, Text),
						}
					}
				}
			}
		}
	}

	return StorysDatas, nil
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

/**
 * @description: 反转数组
 * @return {*}
 */
func InversionArray(Array []string) []string {
	var InversionArray []string

	for Index := len(Array) - 1; Index >= 0; Index-- {
		InversionArray = append(InversionArray, Array[Index])
	}

	return InversionArray
}
