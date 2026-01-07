package config

import (
	"encoding/json"
	"os"
)

var Cfg struct {
	App struct {
		Cors     []string `json:"cors"`
		SaveType string   `json:"savetype"`
	} `json:"app"`
	BiliBili struct {
		SESSDATA string `json:"sessdata"`
	} `json:"bilibili"`
}

func GetConfig() {
	wd, err := os.Getwd()
	path := wd + "/config.json"
	if err != nil {
		panic(err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if err == err.(*os.PathError) {
			f, e := json.Marshal(Cfg)
			if e != nil {
				panic(e)
			}
			os.WriteFile(path, f, 0775)
		} else {
			panic(err)
		}
	}
	json.Unmarshal(data, &Cfg)

}
