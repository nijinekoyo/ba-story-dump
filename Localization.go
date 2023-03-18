/*
 * @Author: nijineko
 * @Date: 2023-02-23 20:43:13
 * @LastEditTime: 2023-03-18 15:44:35
 * @LastEditors: nijineko
 * @Description: 本地化部分
 * @FilePath: \StoryDump\Localization.go
 */
package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/pierrec/xxHash/xxHash32"
)

type CharacterNameExcel struct {
	DataList []CharacterInfo `json:"DataList"`
}

type CharacterInfo struct {
	CharacterName   uint32 `json:"CharacterName"`
	NameJP          string `json:"NameJP"`     // 日文名字
	NameKR          string `json:"NameKR"`     // 韩文名字
	NicknameJP      string `json:"NicknameJP"` // 日文所属
	NicknameKR      string `json:"NicknameKR"` // 韩文所属
	ProductionStep  int    `json:"ProductionStep"`
	Shape           int    `json:"Shape"`
	SmallPortrait   string `json:"SmallPortrait"`   // 小头像
	SpinePrefabName string `json:"SpinePrefabName"` // Spine动画
}

var CharacterInfos map[uint32]CharacterInfo // 角色名字，Key为角色xxhash

var CharacterNameLocalizationJsonPath = "./localization/ScenarioCharacterNameExcelTable.json" // 角色名字本地化文件路径

/**
 * @description: 初始化角色信息本地化
 * @return {error} 错误
 */
func InitCharacterInfoLocalization() error {
	// 读取文件
	File, err := os.OpenFile(CharacterNameLocalizationJsonPath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	JsonFile, err := io.ReadAll(File)
	if err != nil {
		return err
	}

	// 解析Json
	var CharacterNameExcel CharacterNameExcel
	err = json.Unmarshal(JsonFile, &CharacterNameExcel)
	if err != nil {
		return err
	}

	// 生成本地化数据
	CharacterInfos = make(map[uint32]CharacterInfo)

	for _, CharacterNameExcelData := range CharacterNameExcel.DataList {
		CharacterInfos[CharacterNameExcelData.CharacterName] = CharacterNameExcelData
	}

	return err
}

/**
 * @description: 将韩文角色名字转换为日文角色名字
 * @param {string} NameKR 韩文角色名字
 * @return {string} 日文角色名字
 * @return {string} 日文角色所属
 */
func CharacterNameKRToJP(NameKR string) (string, string) {
	CharacterName := xxHash32.Checksum([]byte(NameKR), 0)

	CharacterInfos, ok := CharacterInfos[CharacterName]
	if !ok {
		return "", ""
	}

	return CharacterInfos.NameJP, CharacterInfos.NicknameJP
}
