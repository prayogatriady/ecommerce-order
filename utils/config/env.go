package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type EnvVal struct {
	App    *App    `json:"app"`
	MySql  *MySql  `json:"mysql"`
	Logger *Logger `json:"logger"`
}

type App struct {
	Port     int    `json:"port"`
	Mode     string `json:"mode"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Timezone string `json:"timezone"`
}

type Logger struct {
	Dir   string `json:"dir"`
	Level string `json:"level"`
}

type MySql struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
}

var Env *EnvVal

// func init() {

// 	env := InitEnv(constant.DIR_ENV)
// 	Env = env

// }

func InitEnv(dir string) *EnvVal {

	raw, err := os.ReadFile(dir)
	if err != nil {
		fmt.Printf("error reading env.json file: %v \n", err)
		return nil
	}

	var env *EnvVal
	err = json.Unmarshal(raw, &env)
	if err != nil {
		fmt.Printf("error unmarshalling env.json: %v \n", err)
		return nil
	}

	return env
}
