# goods_srv

## 目录分层

不是固定不变的，也不是死板的
每个公司或者每个人的喜好都不一样
关键是你自己能够清晰，别人看了也能清晰就可以了。


handler       -->   biz     -->   dao  -->  MySQL\redis\ES
入口（参数处理）     业务逻辑       数据操作          数据

proto: 参数和响应已经通过`.proto`文件定义好了。
model: 模型（VO、DTO...）
errno: 错误码
sql: 建表语句
third_pkg: 第三方库
middleware: 中间件
config: 配置文件解析
logger: 日志实例
