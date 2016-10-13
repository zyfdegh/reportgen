# reportgen
Analyse and generate Excel report.

# 条件
* 在Windows 64上运行，推荐Windows 7，其他版本未经测试；
* 安装有Microsoft Excel，推荐2013版，其他版本未经测试；
* 数据格式与019132.xls高度相似。

# 过程
对当前目录下所有xls文件（report.xls除外），程序将作如下处理

1. 取sheet名中号段，开始分析；
1. 以“已导出明细”为结束符，扫描出数据行数；
1. 排序并取出最大值数组，即“业务代码”列重复次数最高的30个；
1. 忽略9-15点之外频次，计算各个时段频率；
1. 追加写入数据到报表report.xls；
1. 继续分析下一个文件，重复步骤1-5。

# 使用

1. 将所有待处理的Excel表格放入同一个目录；
1. 将reportgen.exe放入与Excel相同目录；
1. 双击执行report.exe，在出现的控制台中留意信息；
1. 生成的报表文件为report.xls。
