# Order服务模块

## 数据模型
```sql

CREATE DATABASE go_order DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE `goods_order` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT '订单主键',
    `order_no` bigint NOT NULL COMMENT '订单主键',
    `pay_no` bigint NOT NULL DEFAULT 0 COMMENT '支付单号',
    `pack_no` bigint NOT NULL DEFAULT 0 COMMENT '打包单号',
    `post_no` bigint NOT NULL DEFAULT 0 COMMENT '快递单号',
    `user_id` bigint NOT NULL COMMENT '用户id',
    `post_city_code` int NOT NULL COMMENT '收货所属城市',
    `post_id` bigint NOT NULL COMMENT '收货地址id',
    `shop_id` bigint NOT NULL COMMENT '商铺',
    `sale_amount` int NOT NULL COMMENT '售价总金额(单位分)',
    `pay_amount` int NOT NULL COMMENT '实付总金额(单位分)',
    `goods_amount` int NOT NULL COMMENT '商品总数量',
    `status` tinyint NOT NULL COMMENT '状态(1:未支付,2:已支付,3:仓库打包,4:运输货物,5:完成交易,6:取消)',
    `deleted` tinyint NOT NULL DEFAULT 0 COMMENT '删除(0-正常，1-已删除)',
    `create_at` bigint NOT NULL COMMENT '创建时间',
    `pay_at` bigint NOT NULL DEFAULT 0 COMMENT '支付时间',
    `pack_at` bigint NOT NULL DEFAULT 0 COMMENT '打包时间',
    `post_at` bigint NOT NULL DEFAULT 0 COMMENT '发货时间',
    `update_at` bigint NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    INDEX orderNo(`order_no`),
    INDEX userId(`user_id`),
    INDEX payNo(`pay_no`),
    INDEX packNo(`pack_no`),
    INDEX postNo(`post_no`),
    INDEX shopId(`shop_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单表';


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

## Debug 设置
```
Run kind:  File
Files:     main方法路径，比如user-api(/home/zdw/workPlace/go/gozero-micro-service/user/cmd/api/user.go)
Working directory: main方法所在的目录，比如user-api(/home/zdw/workPlace/go/gozero-micro-service/user/cmd/api)


```