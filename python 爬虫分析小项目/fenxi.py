import numpy as np
import pandas as pd
from pyecharts import options as opts
from pyecharts.charts import Bar,Pie,Page

# 去哪儿热门景点excel文件保存路径
PLACE_EXCEL_PATH = 'sightplace.xlsx'
# 读取数据
DF = pd.read_excel(PLACE_EXCEL_PATH)


def grid_vertical():
    global DF
    df = DF.copy()
    # 1、生成一个名称和销量的透视表
    place_sale = df.pivot_table(index='name', values='sale')
    # 2、根据销量排序
    place_sale.sort_values('sale', inplace=True, ascending=True)
    # print(place_sale)
    # 3、生成柱状图
    place_sale_bar = (
        Bar()
            # place_sale中行号0-20的数据
            .add_xaxis(place_sale.index.tolist()[-20:])
            .add_yaxis("", list(map(int, np.ravel(place_sale.values)))[-20:])
            # 翻转XY轴，将柱状图转换为条形图
            .reversal_axis()
            # 主要是对图元、文字、等内容进行配置 右对齐
            .set_series_opts(label_opts=opts.LabelOpts(position="right"))
            .set_global_opts(
            title_opts=opts.TitleOpts(title="国庆旅游热门景点门票销量TOP20"),
            yaxis_opts=opts.AxisOpts(name="景点名称"),
            xaxis_opts=opts.AxisOpts(name="销量")
        )
    )

    place_sale_pie = (
        # 设置长宽
        Pie(init_opts=opts.InitOpts(width="100%"))
            # .add("饼状图", [list(z) for z in zip(place_sale.index.tolist()[-20:], list(map(int, np.ravel(place_sale.values)))[-20:])],radius=["40%", "75%"])
            .add("", [list(z) for z in zip(place_sale.index.tolist()[-20:], list(map(int, np.ravel(place_sale.values)))[-20:])])
            .set_global_opts(title_opts=opts.TitleOpts(title="饼状图"),
                             legend_opts=opts.LegendOpts(orient="vertical", pos_top="15%", pos_left="1%")
                             )
            # .set_series_opts(label_opts=opts.LabelOpts(formatter="{b}:{c}%"))
        )

    # grid = (
    #     Grid()
    #         .add(place_sale_bar, grid_opts=opts.GridOpts(pos_top="15%"))
    #         .add(place_sale_pie, grid_opts=opts.GridOpts(pos_top='55%'))
    # )

    page = Page()
    page.add(place_sale_bar,place_sale_pie)

    # place_sale_pie.render('place_sale_pie.html')
    # place_sale_bar.render('place-sale-bar.html')
    page.render('bar_pie.html')

if __name__ == '__main__':
    grid_vertical()