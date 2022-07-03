# user服务

## SQL语句
```mysql

CREATE DATABASE go_user DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;


// 
CREATE TABLE user (
  id bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  name varchar(128) NULL DEFAULT '' comment '姓名',
  nick_name varchar(64) NULL DEFAULT '' comment '昵称',
  birthday date NULL DEFAULT '1907-01-01' comment '出生年月日',
  phone varchar(11) NULL DEFAULT '' comment '手机号码',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT '用户基础信息表';

```
