/*
 * @Author: nijineko
 * @Date: 2023-12-28 00:23:43
 * @LastEditTime: 2023-12-28 00:25:59
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
	jsonBytes, err := json.Marshal(ScenarioScriptData)
	if err != nil {
		return err
	}

	return os.WriteFile("./json/ScenarioScript.json", jsonBytes, 0644)
}

func ScenarioCharacterNameToJson(ScenarioCharacterNameData []ScenarioCharacterName) error {
	jsonBytes, err := json.Marshal(ScenarioCharacterNameData)
	if err != nil {
		return err
	}

	return os.WriteFile("./json/ScenarioCharacterName.json", jsonBytes, 0644)
}
