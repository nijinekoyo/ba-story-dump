/*
 * @Author: nijineko
 * @Date: 2023-12-27 23:10:03
 * @LastEditTime: 2023-12-27 23:58:21
 * @LastEditors: nijineko
 * @Description: 数据库文件读取
 * @FilePath: \StoryDump\SQLRead.go
 */
package main

import (
	"StoryDump/Flatbuf/flat/flat"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type ScenarioScriptDB struct {
	GroupId int
	Bytes   []byte
}

type ScenarioScript struct {
	GroupId        int    `json:"group_id"`
	SelectionGroup int    `json:"selection_group"`
	BGMId          int    `json:"bgm_id"`
	Sound          string `json:"sound"`
	Transition     uint   `json:"transition"`
	BGName         uint   `json:"bg_name"`
	BGEffect       uint   `json:"bg_effect"`
	PopupFileName  string `json:"popup_file_name"`
	ScriptKr       string `json:"script_kr"`
	TextJp         string `json:"text_jp"`
	VoiceId        uint   `json:"voice_id"`
}

type ScenarioCharacterNameDB struct {
	CharacterName int
	Bytes         []byte
}

type ScenarioCharacterName struct {
	CharacterName   uint   `json:"character_name"`
	ProductionSte   int    `json:"production_ste"`
	NameKR          string `json:"name_kr"`
	NicknameKR      string `json:"nickname_kr"`
	NameJP          string `json:"name_jp"`
	NicknameJP      string `json:"nickname_jp"`
	Shape           int    `json:"shape"`
	SpinePrefabName string `json:"spine_prefab_name"`
	SmallPortrait   string `json:"small_portrait"`
}

const (
	DBFile = "./db/ExcelDB.db"
)

/**
 * @description: 读取数据库文件
 * @return {*}
 */
func SQLRead() ([]ScenarioScript, []ScenarioCharacterName, error) {
	db, err := sql.Open("sqlite3", DBFile)
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	// 读取剧情文本
	var ScenarioScriptDBSchema []ScenarioScriptDB
	rows, err := db.Query("SELECT * FROM ScenarioScriptDBSchema")
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ScenarioScriptDBSchemaTemp ScenarioScriptDB
		err = rows.Scan(&ScenarioScriptDBSchemaTemp.GroupId, &ScenarioScriptDBSchemaTemp.Bytes)
		if err != nil {
			return nil, nil, err
		}
		ScenarioScriptDBSchema = append(ScenarioScriptDBSchema, ScenarioScriptDBSchemaTemp)
	}

	// 解析剧情文本
	var ScenarioScriptSchema []ScenarioScript
	for _, ScenarioScriptDBSchemaTemp := range ScenarioScriptDBSchema {
		ScenarioScriptSchema = append(ScenarioScriptSchema, ReadScenarioScript(ScenarioScriptDBSchemaTemp))
	}

	// 读取角色名字本地化
	var ScenarioCharacterNameDBSchema []ScenarioCharacterNameDB
	rows, err = db.Query("SELECT * FROM ScenarioCharacterNameDBSchema")
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var ScenarioCharacterNameDBSchemaTemp ScenarioCharacterNameDB
		err = rows.Scan(&ScenarioCharacterNameDBSchemaTemp.CharacterName, &ScenarioCharacterNameDBSchemaTemp.Bytes)
		if err != nil {
			return nil, nil, err
		}
		ScenarioCharacterNameDBSchema = append(ScenarioCharacterNameDBSchema, ScenarioCharacterNameDBSchemaTemp)
	}

	// 解析角色名字本地化
	var ScenarioCharacterNameSchema []ScenarioCharacterName
	for _, ScenarioCharacterNameDBSchemaTemp := range ScenarioCharacterNameDBSchema {
		ScenarioCharacterNameSchema = append(ScenarioCharacterNameSchema, ReadScenarioCharacterName(ScenarioCharacterNameDBSchemaTemp))
	}

	return ScenarioScriptSchema, ScenarioCharacterNameSchema, nil
}

/**
 * @description: 解析ScenarioScript数据
 * @param {ScenarioScriptDB} ScenarioScriptData ScenarioScript数据
 * @return {ScenarioScript} ScenarioScript解析后的数据
 */
func ReadScenarioScript(ScenarioScriptData ScenarioScriptDB) ScenarioScript {
	ScenarioScriptFBData := flat.GetRootAsScenarioScriptExcel(ScenarioScriptData.Bytes, 0)

	var ScenarioScriptTemp ScenarioScript = ScenarioScript{
		GroupId:        int(ScenarioScriptFBData.GroupId()),
		SelectionGroup: int(ScenarioScriptFBData.SelectionGroup()),
		BGMId:          int(ScenarioScriptFBData.Bgmid()),
		Sound:          string(ScenarioScriptFBData.Sound()),
		Transition:     uint(ScenarioScriptFBData.Transition()),
		BGName:         uint(ScenarioScriptFBData.Bgname()),
		BGEffect:       uint(ScenarioScriptFBData.Bgeffect()),
		PopupFileName:  string(ScenarioScriptFBData.PopupFileName()),
		ScriptKr:       string(ScenarioScriptFBData.ScriptKr()),
		TextJp:         string(ScenarioScriptFBData.TextJp()),
		VoiceId:        uint(ScenarioScriptFBData.VoiceId()),
	}

	return ScenarioScriptTemp
}

/**
 * @description: 解析ScenarioCharacterName数据
 * @param {ScenarioCharacterNameDB} ScenarioCharacterNameData ScenarioCharacterName数据
 * @return {ScenarioCharacterName} ScenarioCharacterName解析后的数据
 */
func ReadScenarioCharacterName(ScenarioCharacterNameData ScenarioCharacterNameDB) ScenarioCharacterName {
	ScenarioCharacterNameFBData := flat.GetRootAsScenarioCharacterNameExcel(ScenarioCharacterNameData.Bytes, 0)

	var ScenarioCharacterNameTemp ScenarioCharacterName = ScenarioCharacterName{
		CharacterName:   uint(ScenarioCharacterNameFBData.CharacterName()),
		ProductionSte:   int(ScenarioCharacterNameFBData.ProductionStep()),
		NameKR:          string(ScenarioCharacterNameFBData.NameKr()),
		NicknameKR:      string(ScenarioCharacterNameFBData.NicknameKr()),
		NameJP:          string(ScenarioCharacterNameFBData.NameJp()),
		NicknameJP:      string(ScenarioCharacterNameFBData.NicknameJp()),
		Shape:           int(ScenarioCharacterNameFBData.Shape()),
		SpinePrefabName: string(ScenarioCharacterNameFBData.SpinePrefabName()),
		SmallPortrait:   string(ScenarioCharacterNameFBData.SmallPortrait()),
	}

	return ScenarioCharacterNameTemp
}
