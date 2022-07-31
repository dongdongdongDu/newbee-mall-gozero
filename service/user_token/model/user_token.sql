CREATE TABLE `tb_newbee_mall_user_token`
(
    `user_id`     bigint(20)  NOT NULL COMMENT '用户主键id',
    `token`       varchar(32) NOT NULL COMMENT 'token值(32位字符串)',
    `update_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    `expire_time` datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'token过期时间',
    PRIMARY KEY (`user_id`),
    UNIQUE KEY `uq_token` (`token`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;