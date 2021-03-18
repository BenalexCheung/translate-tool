package controllers

var transMap = map[string]string{
	"auto":  "自动检测",
	"en":    "英语",
	"fr":    "法语",
	"de":    "德语",
	"it":    "意大利语",
	"ja":    "日语",
	"ko":    "韩语",
	"pt-PT": "葡萄牙语",
	"ru":    "俄语",
	"es":    "西班牙语",
	"ar":    "阿拉伯语",
	"he":    "希伯来语",
	"tr":    "土耳其语",
	"pl":    "波兰语",
	"ro-RO": "罗马尼亚语",
	"bg-BG": "保加利亚语",
	"zh-CN": "中文简体",
	"zh-TW": "中文繁体",
}

type Translate interface {
	Text(sl, tl, text string) (string, error)
	File() string
}
