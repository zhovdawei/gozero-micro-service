# user服务

## SQL语句
```mysql

CREATE DATABASE go_user DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;


// 
CREATE TABLE `user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '' COMMENT '用户名称',
  `phone` varchar(11) DEFAULT '' COMMENT '手机号',
  `password` varchar(255) DEFAULT '' COMMENT '用户密码',
  `gender` tinyint(4) DEFAULT '0' COMMENT '男-1｜女-2｜未公开-0',
  `birthday` date DEFAULT '1907-01-01' COMMENT '出生年月日',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone_unique` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户注册表';

```

## Model 生成
```shell



```