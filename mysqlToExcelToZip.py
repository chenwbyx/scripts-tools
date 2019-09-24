#!/usr/bin/env python2.7
#coding=utf-8
 
# __Desc__ = 从数据库中导出数据到excel数据表
import zipfile #引入zip管理模块
import time, os, xlwt, pymysql, datetime
import sys #引入sys模块，获取脚本所在目录
import shutil, json, requests
from pathlib import Path

 
class MYSQL:
    def __init__(self):
        pass
 
    def __del__(self):
        self._cursor.close()
        self._connect.close()
 
    def connectDB(self, db_name):
        """
        连接数据库
        :return:
        """
        try:
            self._connect = pymysql.Connect(
                host='192.168.1.1',
                port=3306,
                user='root',
                passwd='*****',
                db=db_name,
                charset='utf8'
            )
 
            return 0
        except:
            return -1
 
    def export(self, table_name, output_path):
        self._cursor = self._connect.cursor()
        count = self._cursor.execute('select * from '+ table_name)
        # print(self._cursor.lastrowid)
        # 重置游标的位置
        self._cursor.scroll(0, mode='absolute')
        # 搜取所有结果
        results = self._cursor.fetchall()
 
        # 获取MYSQL里面的数据字段名称
        fields = self._cursor.description
        workbook = xlwt.Workbook()
 
        # 注意: 在add_sheet时, 置参数cell_overwrite_ok=True, 可以覆盖原单元格中数据。
        # cell_overwrite_ok默认为False, 覆盖的话, 会抛出异常.
        sheet = workbook.add_sheet('table_'+ table_name, cell_overwrite_ok=True)
 
        # 写上字段信息
        for field in range(0, len(fields)):
            sheet.write(0, field, fields[field][0])
 
        # 获取并写入数据段信息
        row = 1
        col = 0
        for row in range(1,len(results)+1):
            for col in range(0, len(fields)):
                sheet.write(row, col, u'%s' % results[row-1][col])
 
        workbook.save(output_path)

    #直接执行SQL语句的一些操作
    def sqlTask(self, sql):
        self._cursor = self._connect.cursor()
        try:
            count = self._cursor.execute(sql)
            # 重置游标的位置
            self._cursor.scroll(0, mode='absolute')
            results = self._cursor.fetchall()
            fields = self._cursor.description

            for row in range(0,len(results)):
                for col in range(0, len(fields)):
                    print(u'%s' % results[row][col]),
                print("")
        except:
            print("Error: unable to fecth data")
 
 
def zipDir(dirpath,outFullName):
    """
    压缩指定文件夹
    :param dirpath: 目标文件夹路径
    :param outFullName: 压缩文件保存路径+xxxx.zip
    :return: 无
    """
    zip = zipfile.ZipFile(outFullName,"w",zipfile.ZIP_DEFLATED)
    for path,dirnames,filenames in os.walk(dirpath):
        # 去掉目标跟路径，只对目标文件夹下边的文件及文件夹进行压缩
        fpath = path.replace(dirpath,'')

        for filename in filenames:
            zip.write(os.path.join(path,filename),os.path.join(fpath,filename))
    zip.close()


if __name__ == '__main__':
    mysql = MYSQL()
    dbList = ['garden_game','garden_main'] #两个数据库
    dbMainTable = ['base_data','retention_data'] #主数据库table表
    dbGameTable = ['module_data','conversion_data','production_sales_list'] #游戏数据库table表
    for x in dbList:
        if x == 'garden_main' :
            flag = mysql.connectDB(x)
            if flag == 0:
                for y in dbMainTable:
                    mysql.export(y, 'data/data_' + y + '.xls')
        if x == 'garden_game' :
            flag = mysql.connectDB(x)
            if flag == 0 :
                for y in dbGameTable:
                    mysql.export(y, 'data/data_' + y + '.xls')


    print('导出数据成功...')
    nowTime=datetime.datetime.now().strftime('%Y-%m-%d-%H-%M')
    filename=nowTime+'_data.zip'
    zipDir('./data',filename)
    print('压缩成功...')