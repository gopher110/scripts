import pandas as pd
import seaborn
import matplotlib.pyplot as plt

jobs = pd.read_csv('zhiping-jobs-628001851.csv')

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
plt.pie(forPie.values, radius=0.9, labels=forPie.index, startangle=160, autopct='%.1f%%', shadow=False, colors=seaborn.color_palette('hls', 5))
plt.savefig("./analysis-result/pie-experience.png", dpi=180, bbox_inches='tight')

# plot bar
company_num = 20
companyList = jobs['公司名称'].value_counts(ascending=False)
companyList = companyList.head(company_num).rename_axis('招聘公司').reset_index(name='职位数')
# print(companyList)
companyPlot = seaborn.catplot(x='职位数', y='招聘公司', data=companyList, kind='bar')
companyPlot.savefig("./analysis-result/bar-company.png")

# salary hist
# print(jobs.loc[2:5, ['职位描述']])
# jobs.set_index("公司名称").iloc[2:5, 3]
# print(jobs.loc[(jobs.公司名称.str.contains('腾讯')) & (jobs.城市.notnull())])
# print(jobs.薪资.unique())
jobs['salary'] = jobs['薪资'].str.title().replace(regex=True, to_replace='\\·\\d+薪', value='')
minSalary = jobs['salary'].str.split('-').str[0].str.rstrip('k').str.rstrip('K').astype('float64')*1000
maxSalary = jobs['salary'].str.split('-').str[1].str.rstrip('k').str.rstrip('K').astype('float64')*1000
jobs['salary'] = 0.5*(minSalary+maxSalary)
salaryPlot = plt.figure().add_subplot(1, 1, 1)
salaryPlot.hist(jobs['salary'], bins=10, color=seaborn.color_palette('hls', 5)[3])
plt.xlabel('工资（(n1+n2)2）')
plt.ylabel('职位数')
plt.title('岗位工资分布')
plt.savefig("./analysis-result/hist-salary.png", dpi=180)

plt.show()

