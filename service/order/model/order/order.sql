CREATE TABLE `tb_newbee_mall_order`
(
    `order_id`     bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '订单表主键id',
    `order_no`     varchar(20)  NOT NULL DEFAULT '' COMMENT '订单号',
    `user_id`      bigint(20)   NOT NULL DEFAULT '0' COMMENT '用户主键id',
    `total_price`  int(11)      NOT NULL DEFAULT '1' COMMENT '订单总价',
    `pay_status`   tinyint(4)   NOT NULL DEFAULT '0' COMMENT '支付状态:0.未支付,1.支付成功,-1:支付失败',
    `pay_type`     tinyint(4)   NOT NULL DEFAULT '0' COMMENT '0.无 1.支付宝支付 2.微信支付',
    `pay_time`     datetime              DEFAULT NULL COMMENT '支付时间',
    `order_status` tinyint(4)   NOT NULL DEFAULT '0' COMMENT '订单状态:0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭',
    `extra_info`   varchar(100) NOT NULL DEFAULT '' COMMENT '订单body',
    `is_deleted`   tinyint(4)   NOT NULL DEFAULT '0' COMMENT '删除标识字段(0-未删除 1-已删除)',
    `create_time`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最新修改时间',
    PRIMARY KEY (`order_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;