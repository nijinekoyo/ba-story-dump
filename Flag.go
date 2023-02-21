/*
 * @Author: nijineko
 * @Date: 2023-02-21 21:19:54
 * @LastEditTime: 2023-02-21 21:20:00
 * @LastEditors: nijineko
 * @Description: Flag
 * @FilePath: \StoryDump\Flag.go
 */
package main

import "flag"

type Flag struct {
	Filter bool // 启用过滤器
}

var Flags Flag // 全局参数变量

/**
 * @description: 初始化参数
 * @return {error} 错误
 */
func InitFlag() error {
	// 参数解析
	Filter := flag.Bool("filter", true, "启用过滤器")
	flag.Parse()

	// 参数写入变量
	Flags.Filter = *Filter

	return nil
}
