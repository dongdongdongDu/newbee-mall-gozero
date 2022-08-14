CREATE TABLE `tb_newbee_mall_order_item`
(
    `order_item_id`   bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '订单关联购物项主键id',
    `order_id`        bigint(20)   NOT NULL DEFAULT '0' COMMENT '订单主键id',
    `goods_id`        bigint(20)   NOT NULL DEFAULT '0' COMMENT '关联商品id',
    `goods_name`      varchar(200) NOT NULL DEFAULT '' COMMENT '下单时商品的名称(订单快照)',
    `goods_cover_img` varchar(200) NOT NULL DEFAULT '' COMMENT '下单时商品的主图(订单快照)',
    `selling_price`   int(11)      NOT NULL DEFAULT '1' COMMENT '下单时商品的价格(订单快照)',
    `goods_count`     int(11)      NOT NULL DEFAULT '1' COMMENT '数量(订单快照)',
    `create_time`     datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`order_item_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;