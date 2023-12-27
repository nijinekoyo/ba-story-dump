/*
 * @Author: nijineko
 * @Date: 2023-02-23 20:43:13
 * @LastEditTime: 2023-12-28 00:04:39
 * @LastEditors: nijineko
 * @Description: 本地化部分
 * @FilePath: \StoryDump\Localization.go
 */
package main

import (
	"github.com/pierrec/xxHash/xxHash32"
)

var CharacterInfos map[uint]ScenarioCharacterName // 角色名字，Key为角色xxhash

/**
 * @description: 初始化角色信息本地化
 * @return {error} 错误
 */
func InitCharacterInfoLocalization(ScenarioCharacterNameData []ScenarioCharacterName) {
	// 生成本地化数据
	CharacterInfos = make(map[uint]ScenarioCharacterName)

	for _, CharacterNameExcelData := range ScenarioCharacterNameData {
		CharacterInfos[CharacterNameExcelData.CharacterName] = CharacterNameExcelData
	}
}

/**
 * @description: 将韩文角色名字转换为日文角色名字
 * @param {string} NameKR 韩文角色名字
 * @return {string} 日文角色名字
 * @return {string} 日文角色所属
 */
func CharacterNameKRToJP(NameKR string) (string, string) {
	CharacterName := xxHash32.Checksum([]byte(NameKR), 0)

	CharacterInfos, ok := CharacterInfos[uint(CharacterName)]
	if !ok {
		return "", ""
	}

	return CharacterInfos.NameJP, CharacterInfos.NicknameJP
}
