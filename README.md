# translateTool

## 界面设计

## 后端设计

### 日志配置

```
func GetJson() {
    resp, err := http.Get(beego.AppConfig.String("url"))
    if err != nil {
        beego.Critical("http get info error")
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(body, &AllInfo)
    if err != nil {
        beego.Critical("error:", err)
    }
}
```

## Tools

### 调用Google翻译API方式

有道翻译：http://api.microsofttranslator.com/v2/Http.svc/Translate?appId=AFC76A66CF4F434ED080D245C30CF1E71C22959C&from=&to=en&text=老季博客
``` XML
<string xmlns="http://schemas.microsoft.com/2003/10/Serialization/">Old season blog</string>
```

### 调用Google翻译API方式

有道翻译：http://fanyi.youdao.com/translate?&doctype=json&type=ZH_CN2EN&i=你吃饭了吗
``` JSON
{
    "type": "ZH_CN2EN",
    "errorCode": 0,
    "elapsedTime": 0,
    "translateResult": [
        [
            {
                "src": "你吃饭了吗",
                "tgt": "have you had meals"
            }
        ]
    ]
}
```

type的类型有：
``` 
ZH_CN2EN 中文　 　英语 
ZH_CN2JA 中文　 　日语 
ZH_CN2KR 中文　 　韩语 
ZH_CN2FR 中文　 　法语 
ZH_CN2RU 中文　 　俄语 
ZH_CN2SP 中文　 　西语 
EN2ZH_CN 英语　 　中文 
JA2ZH_CN 日语　 　中文 
KR2ZH_CN 韩语　 　中文 
FR2ZH_CN 法语　 　中文 
RU2ZH_CN 俄语　 　中文 
SP2ZH_CN 西语　 　中文
```

## 数据库设计

### 简单的数据库信息

```
CREATE TABLE entries (
    id INT AUTO_INCREMENT,
    word_zh_cn TEXT,
    word_zh_tw TEXT,
    word_en TEXT,
    word_de TEXT,
    created DATETIME,
    primary key (id)
);
```