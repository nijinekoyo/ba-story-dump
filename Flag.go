/*
 * @Author: nijineko
 * @Date: 2023-02-21 21:19:54
 * @LastEditTime: 2023-03-09 13:27:52
 * @LastEditors: nijineko
 * @Description: Flag
 * @FilePath: \StoryDump\Flag.go
 */
package main

import "flag"

type Flag struct {
	FilterTag        bool // 特殊Tag过滤器（开启后将过滤带有特殊Tag的文本）
	AddRuby          bool // 添加Ruby语句（开启后Ruby语句以括号的方式写入文本）
	AddCharacterName bool // 添加角色名字
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
	flag.Parse()

	// 参数写入变量
	Flags.FilterTag = *FilterTag
	Flags.AddRuby = *AddRuby
	Flags.AddCharacterName = *AddCharacterName

	return nil
}
