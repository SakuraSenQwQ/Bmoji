package method

import (
	"encoding/json"
	"io"
	"main/config"
	"net/http"
	"os"
	"strings"
)

var list struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data struct {
		Packages []struct {
			ID      int    `json:"id"`
			Text    string `json:"text"`
			Type    int    `json:"type"`
			Url     string `json:"url"`
			UrlType int
		} `json:"all_packages"`
	} `json:"data"`
}

func DownloadEmoteList() {
	baseurl := "https://api.bilibili.com/x/emote/setting/panel?business=reply"
	client := http.Client{}
	req, err := http.NewRequest("GET", baseurl, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Cookie", "SESSDATA="+config.Cfg.BiliBili.SESSDATA)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &list)
	for i, a := range list.Data.Packages {
		surl, urltype := geturl(a.Url)
		list.Data.Packages[i].Url = surl
		list.Data.Packages[i].UrlType = urltype
	}
	file, err := json.Marshal(list)
	if err != nil {
		panic(err)
	}
	os.WriteFile("resp.json", file, 0775)
}
func geturl(url string) (string, int) {

	//神人BiliBili,有四个域名作为图片域名，甚至还有一个是自定义域名
	urls := []string{"http://i0.hdslb.com/bfs/emote/", "https://i0.hdslb.com/bfs/emote/", "https://i0.hdslb.com/bfs/garb/", "http://i0.hdslb.com/bfs/garb/"}
	for i, v := range urls {
		arr := strings.Split(url, v)
		if len(arr) > 1 {
			return arr[len(arr)-1], i
		}
	}
	return url, -1
}
