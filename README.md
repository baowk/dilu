# dilu 

## 简介
dilu 已经发布一段时间，使用还是比较稳定，也收到一些优化反馈，冗余的有点多，为了方便开发，现把当前版本改为简化版本。
去掉不必要的功能，保留快速开发的一个核心。以后使用过程不必再删除冗余的东西。其他库以后在[dilu-plugin](https://github.com/baowk/dilu-plugin)中维护。去掉了默认的数据库和demo，去掉的库在[dilu-all](https://github.com/baowk/dilu-all)中维护。想要快速开始的可以参考[dilu-all](https://github.com/baowk/dilu-all)。代码生成改为命令行提供，此版本没有基础数据库，所以也跟[dili-admin](https://github.com/baowk/dilu-admin)无关，dilu-admin要跟dilu-all一起使用。

dilu 是一套基于gin+gorm封装的web快速框架，系统基于gin+gorm封装，并实现了代码自动生成。
前端基于vue3+element-plus[点我查看](https://github.com/baowk/dilu-admin)

## 特性
- 基于gin+gorm封装，并实现了代码自动生成
- 前端基于vue3+element-plus[点我查看](https://github.com/baowk/dilu-admin)
- 系统做了saas化设计，主账号直接dilu登录，团队

的卢，是golang下的一套web快速框架，系统基于gin+gorm封装，并实现了代码自动生成。
- 前端基于vue3+element-plus[点我查看](https://github.com/baowk/dilu-admin)
- [演示地址](http://dilu.youwan.art),系统做了saas化设计，主账号直接dilu登录，团队账户为tangtang，密码默认。

## 安装使用

- 获取项目代码
```bash
git clone https://github.com/baowk/dilu.git
```

- 安装依赖
```bash
cd dilu
go mod tidy
```

- 安装及使用
mysql 创建自己的数据库

修改数据库配置
resources\config.dev.yaml
```yaml
dbcfg: # 数据库配置
  driver: mysql  
  dns: root:12345678@tcp(127.0.0.1:3306)/dilu-db?charset=utf8&parseTime=True&loc=Local&timeout=1000ms  # 数据库连接字符串
```
默认库为 sys

- 代码生成
假设 自己的数据库里面已经有了 sys_user表，生成代码
```bash
go run main.go gen -c resources/config.dev.yaml -d sys -t sys_user -f false
```
gen 为生成代码命令，c为配置文件，d为数据库（默认为sys），t为表，f为是否强制生成。

- 运行
```bash
go run main.go start -c resources/config.dev.yaml
```

- 交流群

    欢迎家人们star，微信群二维码一周就过期，有问题或者想进行技术交流的先加我微信(注明 的卢 )，我拉你进群!![微信图片_20240308215943](https://github.com/baowk/dilu/assets/142554979/29a6863c-4bdc-4963-99c2-0c400e132f6f)

前端使用请[跳至前端](https://github.com/baowk/dilu-admin)

![d7f4b2513f7440d6c4c9bd932b4800f](https://github.com/baowk/dilu/assets/142554979/ee341fb7-f98e-4f18-9658-f89b4f7d466f)


