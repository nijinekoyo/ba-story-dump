/*
 * @Author: nijineko
 * @Date: 2023-12-28 00:23:43
 * @LastEditTime: 2024-02-21 18:31:01
 * @LastEditors: nijineko
 * @Description: 数据转换为json
 * @FilePath: \StoryDump\ToJson.go
 */
package main

import (
	"encoding/json"
	"os"
)

func ScenarioScriptToJson(ScenarioScriptData []ScenarioScript) error {
	jsonBytes, err := json.MarshalIndent(ScenarioScriptData, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile("./json/ScenarioScript.json", jsonBytes, 0644)
}

func ScenarioCharacterNameToJson(ScenarioCharacterNameData []ScenarioCharacterName) error {
	jsonBytes, err := json.MarshalIndent(ScenarioCharacterNameData, "", "    ")
	if err != nil {
		return err
	}

	return os.WriteFile("./json/ScenarioCharacterName.json", jsonBytes, 0644)
}
