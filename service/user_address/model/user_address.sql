CREATE TABLE `tb_newbee_mall_user_address`
(
    `address_id`     bigint(20)  NOT NULL AUTO_INCREMENT,
    `user_id`        bigint(20)  NOT NULL DEFAULT '0' COMMENT '用户主键id',
    `user_name`      varchar(30) NOT NULL DEFAULT '' COMMENT '收货人姓名',
    `user_phone`     varchar(11) NOT NULL DEFAULT '' COMMENT '收货人手机号',
    `default_flag`   tinyint(4)  NOT NULL DEFAULT '0' COMMENT '是否为默认 0-非默认 1-是默认',
    `province_name`  varchar(32) NOT NULL DEFAULT '' COMMENT '省',
    `city_name`      varchar(32) NOT NULL DEFAULT '' COMMENT '城',
    `region_name`    varchar(32) NOT NULL DEFAULT '' COMMENT '区',
    `detail_address` varchar(64) NOT NULL DEFAULT '' COMMENT '收件详细地址(街道/楼宇/单元)',
    `is_deleted`     tinyint(4)  NOT NULL DEFAULT '0' COMMENT '删除标识字段(0-未删除 1-已删除)',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`address_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='收货地址表';