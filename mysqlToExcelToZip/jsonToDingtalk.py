#!/usr/bin/env python2.7
#coding=utf-8
 
# __Desc__ = 利用钉钉机器人发送消息到钉钉群
import requests
import json

#  请求地址
post_url = "https://oapi.dingtalk.com/robot/send?access_token=xxxxxx"

#  消息头部
headers = {'Content-Type': 'application/json'}

# 消息主体
data =  {
"msgtype": "text",
 "text": {
     "content": " Are you E~T?"
 }
}

# 使用post请求推送消息
requests.post(post_url, data=json.dumps(data), headers=headers)