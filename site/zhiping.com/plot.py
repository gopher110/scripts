import pandas
import seaborn
import matplotlib.pyplot as plt

jobs = pandas.read_csv('zhiping-jobs-628001851.csv')

# seaborn中文乱码解决方案
plt.rcParams['font.family'] = ['Arial Unicode MS']  # 用来正常显示中文标签
plt.rcParams['axes.unicode_minus'] = False  # 用来正常显示负号
seaborn.set_style('whitegrid', {'font.sans-serif': ['Arial Unicode MS', 'Arial']})

# clean data
jobs = jobs[~jobs['经验'].isnull()]
jobs = jobs[~jobs['经验'].str.contains('负责')]
# print(jobs.columns)

# plot pie
forPie = jobs['经验'].value_counts()
print(forPie)
plt.pie(forPie.values, radius=0.9, labels=forPie.index, startangle=160, autopct='%.1f%%',shadow=False,colors=seaborn.color_palette('hls',5))
plt.savefig("./analysis-result/pie-experience.png",dpi=180,bbox_inches='tight')

# plot bar
company_num = 20
companyList = jobs['公司名称'].value_counts(ascending=False)
companyList = companyList.head(company_num).rename_axis('招聘公司').reset_index(name='职位数')
print(companyList)
companyPlot = seaborn.catplot(x='职位数', y='招聘公司', data=companyList, kind='bar')
companyPlot.savefig("./analysis-result/bar-company.png")


# plot = seaborn.scatterplot(x='薪资',y='经验',data=jobs)
plt.show()

