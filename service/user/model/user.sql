CREATE TABLE `user`
(
    `user_id`           bigint unsigned     NOT NULL    AUTO_INCREMENT,
    `nike_name`         varchar(50)         NOT NULL    DEFAULT ''                 COMMENT '用户昵称',
    `login_name`        varchar(11)         NOT NULL    DEFAULT '0'                COMMENT '登陆名称(默认为手机号)',
    `password_md5`      varchar(32)         NOT NULL    DEFAULT ''                 COMMENT 'MD5加密后的密码',
    `introduce_sign`    varchar(100)        NOT NULL    DEFAULT ''                 COMMENT '个性签名',
    `is_deleted`        tinyint(3) unsigned NULL        DEFAULT 0                  COMMENT '注销标识字段(0-正常 1-已注销)',
    `locked_flag`       tinyint(3) unsigned NULL        DEFAULT 0                  COMMENT '锁定标识字段(0-未锁定 1-已锁定)',
    `create_time`       timestamp           NULL        DEFAULT CURRENT_TIMESTAMP  COMMENT '注册时间',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `idx_login_name_unique` (`login_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;