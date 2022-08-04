CREATE TABLE `tb_newbee_mall_goods_info`
(
    `goods_id`             bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品表主键id',
    `goods_name`           varchar(200)        NOT NULL DEFAULT '' COMMENT '商品名',
    `goods_intro`          varchar(200)        NOT NULL DEFAULT '' COMMENT '商品简介',
    `goods_category_id`    bigint(20)          NOT NULL DEFAULT '0' COMMENT '关联分类id',
    `goods_cover_img`      varchar(200)        NOT NULL DEFAULT '/admin/dist/img/no-img.png' COMMENT '商品主图',
    `goods_carousel`       varchar(500)        NOT NULL DEFAULT '/admin/dist/img/no-img.png' COMMENT '商品轮播图',
    `goods_detail_content` text                NOT NULL COMMENT '商品详情',
    `original_price`       int(11)             NOT NULL DEFAULT '1' COMMENT '商品价格',
    `selling_price`        int(11)             NOT NULL DEFAULT '1' COMMENT '商品实际售价',
    `stock_num`            int(11) unsigned    NOT NULL DEFAULT '0' COMMENT '商品库存数量',
    `tag`                  varchar(20)         NOT NULL DEFAULT '' COMMENT '商品标签',
    `goods_sell_status`    tinyint(4)          NOT NULL DEFAULT '0' COMMENT '商品上架状态 1-下架 0-上架',
    `create_user`          int(11)             NOT NULL DEFAULT '0' COMMENT '添加者主键id',
    `create_time`          datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '商品添加时间',
    `update_user`          int(11)             NOT NULL DEFAULT '0' COMMENT '修改者主键id',
    `update_time`          datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '商品修改时间',
    PRIMARY KEY (`goods_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  ROW_FORMAT = DYNAMIC;