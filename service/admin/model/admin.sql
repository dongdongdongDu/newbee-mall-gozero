CREATE TABLE `tb_newbee_mall_admin_user`
(
    `admin_user_id`   bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '管理员id',
    `login_user_name` varchar(50) NOT NULL COMMENT '管理员登陆名称',
    `login_password`  varchar(50) NOT NULL COMMENT '管理员登陆密码',
    `nick_name`       varchar(50) NOT NULL COMMENT '管理员显示昵称',
    `locked`          tinyint(4) DEFAULT '0' COMMENT '是否锁定 0未锁定 1已锁定无法登陆',
    PRIMARY KEY (`admin_user_id`) USING BTREE,
    UNIQUE KEY `idx_login_user_name_unique` (`login_user_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  ROW_FORMAT = DYNAMIC;
