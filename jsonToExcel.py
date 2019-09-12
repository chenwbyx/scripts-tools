#!/usr/bin/env python2.7
#coding=utf-8
 
# __Desc__ = 读Excel中的Json数据处理后放入新的Excel文件
###########################
#源格式：{"174":169,"175":251,"185":126,"193":39,"167":62,"171":140}
#目标格式：每一个K，V单独填入一个单元格并根据K排序
###########################

import json,xlwt
import xlrd
from datetime import date,datetime
from xlrd import xldate_as_tuple
import datetime,operator


file = 'conversion_data.xlsx'
g_hang = 0

def readExcel(file):
    with open(file,'r',encoding='utf8') as fr:
        data = json.load(fr) # 用json中的load方法，将json串转换成字典
    return data


def writeM(sheet, data, dataName, dataTime):
    timedatastyle = xlwt.XFStyle()
    timedatastyle.num_format_str = 'yyyy-mm-dd'
    dataTime = xldate_as_tuple(dataTime, 0)
    dataTime = datetime.datetime(*dataTime)
    global g_hang
    title = ["ID", dataName]
    for i in range(len(title)): # 循环列
        sheet.write(g_hang,i,title[i]) # 将title数组中的字段写入到0行i列中
    data = json.loads(data)
    keys = []
    for line in data: #　循环字典
        keys.append(int(line))
    keys.sort()
    for line in keys: #　循环字典
        # print('line:',line," ",data[line])
        line = str(line)
        g_hang = g_hang + 1
        sheet.write(g_hang,0,line) #　将line写入到第int(line)行，第0列中
        sheet.write(g_hang,1,data[line])
    g_hang = g_hang + 1
    sheet.write(g_hang,0,dataTime,timedatastyle)
    g_hang = g_hang + 3
	

def read_excel():

    wb = xlrd.open_workbook(filename=file)#打开文件
    # sheet1 = wb.sheet_by_index(0)#通过索引获取表格
    sheet2 = wb.sheet_by_name('conversion_data')#通过名字获取表格
    print(sheet2.name,sheet2.nrows,sheet2.ncols)
    book = xlwt.Workbook() # 创建一个excel对象
    sheet = book.add_sheet('Sheet1',cell_overwrite_ok=True) # 添加一个sheet页
    for i in range(1,sheet2.nrows):
        for j in range(4,8):
            writeM(sheet,sheet2.cell_value(i,j),sheet2.cell_value(0,j),sheet2.cell_value(i,9))
	
    book.save('demo.xls')
    


if __name__ == '__main__':
    read_excel()