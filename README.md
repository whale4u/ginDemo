## 简介
基于gin+casbin+jwt实现REST接口。

代码架构：
ginDemo
├── common 公共组件
├── config  配置
├── controller 路由实现
│   └── v1
├── entity  对象实体
├── middleware  中间件
│   └── jwt
├── router  路由配置


### 登录接口
```
curl --location --request POST 'http://127.0.0.1:8080/home/login' \
--header 'Content-Type: application/json' \
--data-raw '{"username":"admin", "password":"pass"}'
```

### 主页接口
```
curl --location --request GET 'http://127.0.0.1:8080/v1/home/index' \
--header 'Authorization:Bearer xxx'
```