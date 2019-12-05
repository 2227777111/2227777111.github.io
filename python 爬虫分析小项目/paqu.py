# -*- coding:utf-8 -*-

import requests,json,os,re,random,time
import pandas as pd


class guoQing:
    def __init__(self):
        # 让目标网站误以为本程序是浏览器，并非爬虫。
        self.headers = {
            "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36"
        }
        # 设置Excel路径
        self.PLACE_EXCEL_PATH = 'sightplace.xlsx'
        self.stopFlag = True

    # ip代理池
    def getIpList(self):
        # 从别人的网址爬取代理
        url = 'http://www.66ip.cn/mo.php?sxb=&tqsl=20'  # 10为提取代理数量
        # 获取网页内容，contents转化为16进制，decode解码，windows的控制台是gb2312，避免乱码
        res = requests.get(url).content.decode('gb2312')
        # 正则表达式，爬取页面ip
        ipOringinList = re.findall('\t\t(.*?)<br />',res)
        ipList = []
        for ip in ipOringinList:
            # 测ip可用性，移除不可用ip
            proxy = {
                'http': '{}'.format(ip)
            }
            try:
                res = requests.get('https://www.baidu.com',proxies=proxy).status_code
                # 返回200，可用，加入ipList
                if res == 200:
                    ipList.append(ip)
            except Exception:
                pass
        return ipList

    # 页面属性获取
    def getPageContent(self,url):
        try:
            # contents转化为16进制
            res = requests.get(url,headers = self.headers).content.decode()
            # 转成python数据类型
            res_json = json.loads(res)
            guo_Qing.getSightData(res_json)
        except Exception:
            pass

    # 界面爬取所需要的内容
    def getSightData(self,res_json):
        try:
            # json数组，将内容分别放入data和sightList中
            sightList = res_json['data']['sightList']
            if sightList!=[]:
                placeList = []
                for sight in sightList:
                    place = {
                        'id':sight['sightId'],
                        'name':sight['sightName'],
                        'star':sight.get('star','无'),
                        'score':sight.get('score',0),
                        'price':sight.get('qunarPrice','免费'),
                        'sale':sight.get('saleCount',0),
                        'districts':sight['districts'],
                        'point':sight['point'],
                        'intro':sight.get('intro','暂无简介')
                    }
                    placeList.append(place)
                guo_Qing.saveExcel(placeList)
            else:
                self.stopFlag = False
        except Exception:
            pass

    # Excel存储数据
    def saveExcel(self,placeList):
        if os.path.exists(self.PLACE_EXCEL_PATH):
            df = pd.read_excel(self.PLACE_EXCEL_PATH)
            df = df.append(placeList)
        else:
            df = pd.DataFrame(placeList)
        writer = pd.ExcelWriter(self.PLACE_EXCEL_PATH)
        # columns参数用于指定生成的excel中列的顺序
        df.to_excel(excel_writer=writer,
                    columns=['id', 'name', 'star', 'score', 'price', 'sale', 'districts', 'point', 'intro'],
                    index=False,
                    encoding='utf-8', sheet_name='去哪儿热门景点')
        writer.save()
        writer.close()


    def main(self,keyword):
        # 判断括号里的文件是否存在
        if os.path.exists(self.PLACE_EXCEL_PATH):
            os.remove(self.PLACE_EXCEL_PATH)
        else:
            # 拼接路径
            os.path.join(self.PLACE_EXCEL_PATH)
        startUrl = 'https://piao.qunar.com/ticket/list.json?' \
                   'keyword={}&region=&from=mpl_search_suggest&page={}'
        for i in range(500):
            url = startUrl.format(keyword,i)
            guo_Qing.getPageContent(url)
            if not self.stopFlag:
                break
            print('正在爬取'+keyword+'第'+str(i)+'页')
            time.sleep(random.randint(2, 5))

        print("爬取完成！")


if __name__ == '__main__':
    keyword = input('请输入搜索关键词：\n')
    guo_Qing = guoQing()
    guo_Qing.main(keyword)