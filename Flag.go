/*
 * @Author: nijineko
 * @Date: 2023-02-21 21:19:54
 * @LastEditTime: 2023-02-23 21:33:18
 * @LastEditors: nijineko
 * @Description: Flag
 * @FilePath: \StoryDump\Flag.go
 */
package main

import "flag"

type Flag struct {
	Filter           bool // 启用过滤器
	FilterRuby       bool // 启用Ruby过滤器（开启后Ruby语句以括号的方式写入文本）
	AddCharacterName bool // 添加角色名字
}

var Flags Flag // 全局参数变量

/**
 * @description: 初始化参数
 * @return {error} 错误
 */
func InitFlag() error {
	// 参数解析
	Filter := flag.Bool("filter", false, "启用过滤器")
	FilterRuby := flag.Bool("filter_ruby", false, "启用Ruby过滤器")
	AddCharacterName := flag.Bool("add_character_name", false, "添加角色名字")
	flag.Parse()

	// 参数写入变量
	Flags.Filter = *Filter
	Flags.FilterRuby = *FilterRuby
	Flags.AddCharacterName = *AddCharacterName

	return nil
}
