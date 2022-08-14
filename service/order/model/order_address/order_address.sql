CREATE TABLE `tb_newbee_mall_order_address`
(
    `order_id`       bigint(20)  NOT NULL,
    `user_name`      varchar(30) NOT NULL DEFAULT '' COMMENT '收货人姓名',
    `user_phone`     varchar(11) NOT NULL DEFAULT '' COMMENT '收货人手机号',
    `province_name`  varchar(32) NOT NULL DEFAULT '' COMMENT '省',
    `city_name`      varchar(32) NOT NULL DEFAULT '' COMMENT '城',
    `region_name`    varchar(32) NOT NULL DEFAULT '' COMMENT '区',
    `detail_address` varchar(64) NOT NULL DEFAULT '' COMMENT '收件详细地址(街道/楼宇/单元)',
    PRIMARY KEY (`order_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8 COMMENT ='订单收货地址关联表';