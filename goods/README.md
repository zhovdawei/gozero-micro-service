# Goods服务模块

## 数据模型
```sql

CREATE DATABASE go_goods DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

CREATE TABLE `goods` (
     `goods_id` bigint NOT NULL AUTO_INCREMENT COMMENT '商品id主键',
     `goods_name` varchar(64) NOT NULL COMMENT '商品名称',
     `title` varchar(64) NOT NULL COMMENT '商品标题',
     `company` varchar(32) NOT NULL COMMENT '公司名称',
     `category` ENUM('手机','笔记本电脑','平板电脑','耳机','裤子','上衣','鞋子','水果','蔬菜','零食','饮料','健身') NOT NULL COMMENT '商品类型',
     `logo` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '商品LOGO' ,
     `content` TEXT NOT NULL COMMENT '商品详情',
     `sale_price` int NOT NULL DEFAULT 0 COMMENT '零售价(单位分)',
     `min_price` int NOT NULL DEFAULT 0 COMMENT '最低价(单位分)',
     `max_price` int NOT NULL DEFAULT 0 COMMENT '最高价(单位分)',
     `stock` int NOT NULL DEFAULT 0 COMMENT '库存',
     `sale` int NOT NULL DEFAULT 0 COMMENT '销量',
     `is_attr` tinyint NOT NULL DEFAULT 0 COMMENT '是否含有规格(0-没有，1-含有)',
     `status` tinyint NOT NULL DEFAULT 0 COMMENT '状态(1-未审核，2-已审核，3-已上架)',
     `deleted` tinyint NOT NULL DEFAULT 0 COMMENT '删除(0-正常，1-已删除)',
     `create_at` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
     `update_at` bigint NOT NULL DEFAULT 0 COMMENT '更新时间',
     PRIMARY KEY (`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品表';

CREATE TABLE `goods_attrs` (
    `goods_attr_id` bigint NOT NULL AUTO_INCREMENT COMMENT '商品id',
    `goods_id` bigint NOT NULL COMMENT '商品品类ID',
    `attr_name` varchar(16) NOT NULL COMMENT '属性名',
    `attr_vals` json NOT NULL COMMENT '属性值',
    `deleted` tinyint NOT NULL DEFAULT 0 COMMENT '删除(0-正常，1-已删除)',
    `create_at` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_at` bigint NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`goods_attr_id`),
    INDEX goodsIdKey(`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品属性表';

CREATE TABLE `goods_spec` (
    `goods_spec_id` bigint NOT NULL AUTO_INCREMENT COMMENT '商品id',
    `goods_id` bigint NOT NULL COMMENT '商品品类ID',
    `spec` json NOT NULL COMMENT '规格组合值',
    `sale_price` int NOT NULL DEFAULT 0 COMMENT '零售价(单位分)',
    `stock` int NOT NULL DEFAULT 0 COMMENT '库存',
    `sale` int NOT NULL DEFAULT 0 COMMENT '销量',
    `preview` varchar(256) NOT NULL DEFAULT '' COMMENT '预览图',
    `deleted` tinyint NOT NULL DEFAULT 0 COMMENT '删除(0-正常，1-已删除)',
    `create_at` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_at` bigint NOT NULL DEFAULT 0 COMMENT '更新时间',
    PRIMARY KEY (`goods_spec_id`),
    INDEX goodsIdKey(`goods_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品规格表';

```

## Model 生成
```shell

# 在目录 /project/goods/ 中执行
goctl model mysql datasource -url="root:Mysql_321@tcp(192.168.56.103:3306)/go_goods" -table="goods" -c -dir ./model
goctl model mysql datasource -url="root:Mysql_321@tcp(192.168.56.103:3306)/go_goods" -table="goods_attrs" -c -dir ./model
goctl model mysql datasource -url="root:Mysql_321@tcp(192.168.56.103:3306)/go_goods" -table="goods_spec" -c -dir ./model
```

## Rpc 生成
```shell

# 在目录 /project/goods/cmd/rpc/ 中执行
goctl rpc protoc ./pb/goods.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style=goZero
```

## Api 生成
```shell

# 在目录 /project/goods/cmd/api/ 中执行
goctl api go -api ./desc/goods.api -dir . --style=goZero
```