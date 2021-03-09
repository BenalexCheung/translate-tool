package controllers

var transMap = map[string]string{
	"auto":  "自动检测",
	"zh_cn": "中文简体",
	"zh_tw": "中文繁体",
	"en":    "英文",
	"fr":    "法语",
	"de":    "德语",
	"ko":    "韩语",
}

type Translate interface {
	Text(sl, tl, text string) (string, error)
	File() string
}
