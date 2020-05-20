## 项目结构
```
|-- api 路由访问的方法
    |-- book.go
    |-- section.go
    |-- syn.go
|-- database 数据层的配置
    |-- elasticsearch.go
    |-- mysql.go
|-- models 对数据的定义与封装
    |-- book.go
    |-- section.go
|-- result 对返回值的定义封装
    |-- result.go
|-- router 请求路由的封装与转发
    |-- router.go
|-- go.mod 项目管理工具
|-- main.go 程序入口
|-- README.md 我自己
```
## 环境依赖
```
goland  1.14.3
elasticsearch 7.1.0
mysql 8.1.0
```