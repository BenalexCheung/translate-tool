# -*- coding: utf-8 -*-
import requests
import json
import sys

class JinShan:
    def __init__(self):
        self.url = 'http://fanyi.youdao.com/translate?doctype=json'
        self.headers = {
            'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36',
            'X-Requested-With': 'XMLHttpRequest'  ## 表示这是一个 ajax 请求
        }


    def get_data(self, word):
        post_data = {
            'type':'auto',
            'i':word
        }
        r = requests.post(self.url, data=post_data, headers=self.headers)
        return r.text


    def search(self, word):
        rs = self.get_data(word)
        # print(rs)
        json_data = json.loads(rs, encoding='utf-8')
        # print(json_data)
        if json_data['errorCode'] == 0:
            print(json_data['translateResult'][0])
        if json_data['errorCode'] == 1:
            print(json_data['translateResult'][0])

if __name__ == '__main__':
    word = sys.argv[1]
    js = JinShan()
    js.search(word)
