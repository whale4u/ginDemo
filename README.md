## 简介
基于gin+casbin+jwt实现REST接口。

Todo：
-[ ] 基于rbac的casbin中间件
-[x] 基于acl的casbin中间件
-[x] jwt中间件

代码架构：
```
ginDemo
├── common 公共组件
├── config  配置
├── controller 路由实现
│   └── v1
├── entity  对象实体
├── middleware  中间件
│   └── jwt
├── router  路由配置
```

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

### 新增密码接口

```
curl --location --request POST 'http://127.0.0.1:8080/passwd/addpasswd' \
--header 'Content-Type: application/json' \
--data-raw '{"type":"passwd", "name":"saas", "url":"http://www.baidu.com", "username":"admin", "password":"pass", "note":"this_is_note"}'
```

### add user

```
curl --location --request POST 'http://127.0.0.1:8080/user/add' \
--header 'Content-Type: application/json' \
--data-raw '{"username":"root", "password":"root123"}'
```

### del user

```
curl --location --request DELETE 'http://127.0.0.1:8080/user/delete/1'
```

### find user

```
curl --location --request POST 'http://127.0.0.1:8080/user/find?username="admin"'
```

### update user

```
curl --location --request POST 'http://127.0.0.1:8080/user/update' \
--header 'Content-Type: application/json' \
--data-raw '{"username":"root", "password":"root234"}'
```

### get all user

```
curl --location --request GET 'http://127.0.0.1:8080/user/getall'
```

## docker部署

1、生成docker images

```shell script
docker build -t gindemo1 .
```

2、运行容器

```shell script
docker run -p 8080:8080 gindemo1
```

## 参考文章

- https://github.com/kuangshp/gin-admin-api/
- https://learnku.com/articles/23548/gingormrouter-quickly-build-crud-restful-api-interface