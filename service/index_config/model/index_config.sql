CREATE TABLE `tb_newbee_mall_index_config`
(
    `config_id`    bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '首页配置项主键id',
    `config_name`  varchar(50)  NOT NULL DEFAULT '' COMMENT '显示字符(配置搜索时不可为空，其他可为空)',
    `config_type`  tinyint(4)   NOT NULL DEFAULT '0' COMMENT '1-搜索框热搜 2-搜索下拉框热搜 3-(首页)热销商品 4-(首页)新品上线 5-(首页)为你推荐',
    `goods_id`     bigint(20)   NOT NULL DEFAULT '0' COMMENT '商品id 默认为0',
    `redirect_url` varchar(100) NOT NULL DEFAULT '##' COMMENT '点击后的跳转地址(默认不跳转)',
    `config_rank`  int(11)      NOT NULL DEFAULT '0' COMMENT '排序值(字段越大越靠前)',
    `is_deleted`   tinyint(4)   NOT NULL DEFAULT '0' COMMENT '删除标识字段(0-未删除 1-已删除)',
    `create_time`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `create_user`  int(11)      NOT NULL DEFAULT '0' COMMENT '创建者id',
    `update_time`  datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最新修改时间',
    `update_user`  int(11)               DEFAULT '0' COMMENT '修改者id',
    PRIMARY KEY (`config_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;