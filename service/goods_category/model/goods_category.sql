CREATE TABLE `tb_newbee_mall_goods_category`
(
    `category_id`    bigint(20)  NOT NULL AUTO_INCREMENT COMMENT '分类id',
    `category_level` tinyint(4)  NOT NULL DEFAULT '0' COMMENT '分类级别(1-一级分类 2-二级分类 3-三级分类)',
    `parent_id`      bigint(20)  NOT NULL DEFAULT '0' COMMENT '父分类id',
    `category_name`  varchar(50) NOT NULL DEFAULT '' COMMENT '分类名称',
    `category_rank`  int(11)     NOT NULL DEFAULT '0' COMMENT '排序值(字段越大越靠前)',
    `is_deleted`     tinyint(4)  NOT NULL DEFAULT '0' COMMENT '删除标识字段(0-未删除 1-已删除)',
    `create_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_user`    int(11)     NOT NULL DEFAULT '0' COMMENT '创建者id',
    `update_time`    datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '修改时间',
    `update_user`    int(11)              DEFAULT '0' COMMENT '修改者id',
    PRIMARY KEY (`category_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  ROW_FORMAT = DYNAMIC;