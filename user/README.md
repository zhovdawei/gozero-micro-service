# User服务模块

## 数据模型
```sql

CREATE DATABASE go_user DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE `user` (
    `user_id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户唯一id',
    `name` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名称',
    `id_card` varchar(18) NOT NULL DEFAULT '' COMMENT '身份证',
    `email` varchar(64) NOT NULL DEFAULT '' COMMENT '邮箱',
    `phone` varchar(11) NOT NULL DEFAULT '' COMMENT '手机号',
    `password` varchar(64) NOT NULL DEFAULT '' COMMENT '用户密码',
    `gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '男-1｜女-2｜未公开-0',
    `birthday` bigint NOT NULL DEFAULT 0 COMMENT '出生年月日',
    `province_code` int NOT NULL DEFAULT 000000 COMMENT '省/直辖市编号',
    `province` varchar(32) NOT NULL DEFAULT '' COMMENT '省/直辖市名称',
    `city_code` int NOT NULL DEFAULT 000000 COMMENT '城市编号',
    `city` varchar(32) NOT NULL DEFAULT '' COMMENT '城市名称',
    `area_code` int NOT NULL DEFAULT 000000 COMMENT '区域编号',
    `area` varchar(32) NOT NULL DEFAULT '' COMMENT '区域名称',
    `create_at` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_at` bigint NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `phone_unique` (`phone`),
    UNIQUE KEY `id_card_unique` (`id_card`),
    UNIQUE KEY `email_unique` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户注册表';

CREATE TABLE `user_post` (
    `post_id` bigint NOT NULL AUTO_INCREMENT COMMENT '收件地址表',
    `user_id` bigint NOT NULL COMMENT '用户唯一id',
    `user_name` varchar(64) NOT NULL DEFAULT '' COMMENT '收货人姓名',
    `user_phone` varchar(11) NOT NULL DEFAULT '' COMMENT '收货人手机号',
    `province_code` int NOT NULL DEFAULT 000000 COMMENT '省/直辖市编号',
    `province` varchar(32) NOT NULL DEFAULT '' COMMENT '省/直辖市名称',
    `city_code` int NOT NULL DEFAULT 000000 COMMENT '城市编号',
    `city` varchar(32) NOT NULL DEFAULT '' COMMENT '城市名称',
    `area_code` int NOT NULL DEFAULT 000000 COMMENT '区域编号',
    `area` varchar(32) NOT NULL DEFAULT '' COMMENT '区域名称',
    `address` varchar(256) NOT NULL DEFAULT '' COMMENT '收货地址',
    `status` tinyint NOT NULL DEFAULT 1 COMMENT '0-未启用 | 1-已启用',
    `create_at` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_at` bigint NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`post_id`),
    UNIQUE KEY `userId_userPhone_unique` (`user_id`,`user_phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户收件地址';

```

## Model 生成
```shell

# 在目录 /project/user/ 中执行
goctl model mysql datasource -url="root:Mysql_321@tcp(192.168.56.103:3306)/go_user" -table="user" -c -dir ./model
```

## Rpc 生成
```shell

# 在目录 /project/user/cmd/rpc/ 中执行
goctl rpc protoc ./pb/user.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero
```

## Api 生成
```shell

# 在目录 /project/user/cmd/api/ 中执行
goctl api go -api ./desc/user.api -dir . --style=goZero
```