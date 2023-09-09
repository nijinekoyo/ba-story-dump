/*
 * @Author: nijineko
 * @Date: 2023-02-21 21:19:54
 * @LastEditTime: 2023-09-09 16:17:50
 * @LastEditors: nijineko
 * @Description: Flag
 * @FilePath: \StoryDump\Flag.go
 */
package main

import "flag"

type Flag struct {
	FilterTag             bool // 特殊Tag过滤器（开启后将过滤带有特殊Tag的文本）
	AddRuby               bool // 添加Ruby语句（开启后Ruby语句以括号的方式写入文本）
	AddCharacterName      bool // 添加角色名字
	AddFontSizeTips       bool // 添加字体大小提示
	AddSpineIntervalStyle bool // 添加回忆大厅间隔样式
}

var Flags Flag // 全局参数变量

/**
 * @description: 初始化参数
 * @return {error} 错误
 */
func InitFlag() error {
	// 参数解析
	FilterTag := flag.Bool("filter_tag", false, "特殊Tag过滤器（开启后将过滤带有特殊Tag的文本）")
	AddRuby := flag.Bool("add_ruby", false, "添加Ruby语句（开启后Ruby语句以括号的方式写入文本）")
	AddCharacterName := flag.Bool("add_character_name", false, "添加角色名字")
	AddFontSizeTips := flag.Bool("add_font_size_tips", false, "添加字体大小提示")
	AddSpineIntervalStyle := flag.Bool("add_spine_interval_style", false, "添加回忆大厅间隔样式")
	flag.Parse()

	// 参数写入变量
	Flags.FilterTag = *FilterTag
	Flags.AddRuby = *AddRuby
	Flags.AddCharacterName = *AddCharacterName
	Flags.AddFontSizeTips = *AddFontSizeTips
	Flags.AddSpineIntervalStyle = *AddSpineIntervalStyle

	return nil
}
