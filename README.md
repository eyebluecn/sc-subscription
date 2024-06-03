
## 安装依赖
```shell
go get github.com/eyebluecn/sc-misc-idl@master
go get github.com/eyebluecn/sc-subscription-idl@master

```

## 首次启动

### 导入数据库
首次启动需要先导入数据库。参考[misc项目的README](https://github.com/eyebluecn/sc-misc)

### 指定启动参数
数据库连接信息在`src/common/config/mysql_config.go#getDefaultMysqlUrl()`中修改即可。

如果你希望从启动参数中配置数据库信息，可参考下面的配置：
```shell
smart.classroom.subscription "your_username:your_password@tcp(your_host:3306)/your_schema?charset=utf8mb4&parseTime=True&loc=Local"
```


## 日常启动
执行`main.go`中的main函数即可启动。
