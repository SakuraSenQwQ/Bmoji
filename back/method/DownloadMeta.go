package method

import (
	"encoding/json"
	"fmt"
	"io"
	"main/config"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func DownloadEmoteMeta() {
	var download []struct {
		ID     int    `json:"id"`
		Text   string `json:"text"`
		Emotes []struct {
			Text   string `json:"text"`
			Url    string `json:"url"`
			GifUrl string `json:"gif_url"`
			Type   int    `json:"type"`
			Meta   struct {
				Alias string `json:"alias"`
			} `json:"meta"`
		} `json:"emote"`
	}
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.ReadFile(wd + "/resp.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(file, &list)

	//获取所有id字符串
	var idlist []string
	for _, v := range list.Data.Packages {
		idlist = append(idlist, strconv.Itoa(v.ID))
	}

	//50个为一组分组
	var ids [][]string
	lent := int(math.Floor(float64(len(idlist) / 50)))
	for i := range lent {
		s := i * 50
		o := (i + 1) * 50
		if i == lent-1 {
			ids = append(ids, idlist[o:])
		} else {
			ids = append(ids, idlist[s:o])
		}

	}
	//获取
	for i, v := range ids {
		var metalist struct {
			Code int    `json:"code"`
			Msg  string `json:"message"`
			Data struct {
				Package []struct {
					ID     int    `json:"id"`
					Text   string `json:"text"`
					Emotes []struct {
						Text   string `json:"text"`
						Url    string `json:"url"`
						GifUrl string `json:"gif_url"`
						Type   int    `json:"type"`
						Meta   struct {
							Alias string `json:"alias"`
						} `json:"meta"`
					} `json:"emote"`
				} `json:"packages"`
			} `json:"data"`
		}

		client := http.Client{}
		baseurl := "https://api.bilibili.com/x/emote/package"
		var str strings.Builder
		for _, val := range v {
			str.WriteString(val + ",")
		}
		resq, err := http.NewRequest("GET", baseurl+"?business=reply&ids="+str.String(), nil)
		resq.Header.Set("Cookie", "SESSDATA="+config.Cfg.BiliBili.SESSDATA)
		resp, err := client.Do(resq)
		if err != nil {
			panic(err)
		}
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		json.Unmarshal(data, &metalist)
		download = append(download, metalist.Data.Package...)

		//显示进度，如果太快害怕被封可以自己加sleep
		fmt.Println(i, len(ids))

	}
	file, err = json.Marshal(download)
	if err != nil {
		panic(err)
	}
	//写入
	os.WriteFile("download.json", file, 0775)

}
