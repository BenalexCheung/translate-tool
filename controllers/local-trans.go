package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type transInfo struct {
	Sentences []struct {
		Trans   string `json:"trans"`
		Orig    string `json:"orig"`
		Backend int    `json:"backend"`
	} `json:"sentences"`
	Src        string `json:"src"`
	Confidence int    `json:"confidence"`
}

type LocalTranslate struct {
}

func (lt *LocalTranslate) Text(sl, tl, text string) (*transInfo, error) {
	params := url.Values{}

	Url, err := url.Parse("http://translate.google.cn/translate_a/single")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	params.Set("client", "gtx")
	params.Set("dt", "t")
	params.Set("dj", "1")
	params.Set("ie", "UTF-8")
	params.Set("sl", sl)
	params.Set("tl", tl)
	params.Set("q", text)

	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var ts transInfo
	err = json.Unmarshal(res, &ts)
	fmt.Println(ts)
	return &ts, err
}
