/*
 * @Author: nijineko
 * @Date: 2023-02-23 20:43:13
 * @LastEditTime: 2023-02-23 20:56:17
 * @LastEditors: nijineko
 * @Description: 本地化部分
 * @FilePath: \StoryDump\Localization.go
 */
package main

import (
	"encoding/json"
	"io"
	"os"
)

type CharacterNameExcel struct {
	DataList []struct {
		CharacterName   int64  `json:"CharacterName"`
		NameJP          string `json:"NameJP"`     // 日文名字
		NameKR          string `json:"NameKR"`     // 韩文名字
		NicknameJP      string `json:"NicknameJP"` // 日文所属
		NicknameKR      string `json:"NicknameKR"` // 韩文所属
		ProductionStep  int    `json:"ProductionStep"`
		Shape           int    `json:"Shape"`
		SmallPortrait   string `json:"SmallPortrait"`   // 小头像
		SpinePrefabName string `json:"SpinePrefabName"` // Spine动画
	} `json:"DataList"`
}

type CharacterNameLocalization struct {
	NameJP     string // 日文名字
	NicknameJP string // 日文所属
}

var CharacterName map[string]CharacterNameLocalization // 角色名字，Key为韩文

var CharacterNameLocalizationJsonPath = "./localization/ScenarioCharacterNameExcelTable.json" // 角色名字本地化文件路径

/**
 * @description: 初始化角色名字本地化
 * @return {error} 错误
 */
func InitCharacterNameLocalization() error {
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
	CharacterName = make(map[string]CharacterNameLocalization)

	for _, CharacterNameExcelData := range CharacterNameExcel.DataList {
		// 判断是否存在重复的Key
		if _, ok := CharacterName[CharacterNameExcelData.NameKR]; ok {
			// 如果存在则判断所属是否为空
			if CharacterName[CharacterNameExcelData.NameKR].NicknameJP == "" {
				// 如果为空则覆盖
				CharacterName[CharacterNameExcelData.NameKR] = CharacterNameLocalization{
					NameJP:     CharacterNameExcelData.NameJP,
					NicknameJP: CharacterNameExcelData.NicknameJP,
				}
			}
		} else {
			// 如果不存在则直接写入
			CharacterName[CharacterNameExcelData.NameKR] = CharacterNameLocalization{
				NameJP:     CharacterNameExcelData.NameJP,
				NicknameJP: CharacterNameExcelData.NicknameJP,
			}
		}
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
	// 判断是否存在
	if CharacterNameLocalization, ok := CharacterName[NameKR]; ok {
		return CharacterNameLocalization.NameJP, CharacterNameLocalization.NicknameJP
	} else {
		return "", ""
	}
}
