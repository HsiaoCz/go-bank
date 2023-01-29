# go-bank

youtube json API study_note

1、 使用makefile有一个问题
makefile的缩进必须使用tab键
不能使用空格代替

2、使用@符号可以隐藏make操作的具体内容

## postgreSQL

docker :
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres

安装库
github.com/lib/pq

这玩意儿又不会，不用了，有点傻逼
换成mysql
这里注意一下的是，连接的ip为:127.0.0.1：3306

charset=utf8mb4&parseTime=True&loc=Local 在mysql连接里面一定要加上这个
sql的多行查询，一是记得关闭*sql.Rows，二是res.Scan()这个函数要扫描指针
另外，数据库里有值为null的，扫描会报错
