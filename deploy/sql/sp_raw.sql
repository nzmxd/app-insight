CREATE DATABASE IF NOT EXISTS sp_raw;
       
-- sp_raw.scrapy_applovin_core_params definition
CREATE TABLE IF NOT EXISTS `sp_raw`.`scrapy_applovin_core_params`
(
    `id`              int(10) unsigned NOT NULL AUTO_INCREMENT,
    `params_id`       int(10) unsigned DEFAULT '0' COMMENT '对应泛化参数表中的主键id',
    `params_dict`     varchar(4096) COLLATE utf8_unicode_ci     DEFAULT '' COMMENT '参数字典',
    `params_dict_md5` char(32) COLLATE utf8_unicode_ci NOT NULL COMMENT '参数字典的md5值，主要用于去重',
    `is_available`    tinyint(3) unsigned DEFAULT '1' COMMENT '0：不可用，1: 可用',
    `os`              tinyint(3) unsigned DEFAULT '0' COMMENT '设备类型1:ios,2:android 3:pc',
    `source_app`      varchar(256) COLLATE utf8_unicode_ci      DEFAULT '' COMMENT '来源的app名称',
    `created_at`      timestamp                        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`      timestamp                        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `bk_int`          int(10) unsigned DEFAULT '0' COMMENT 'int 类型备用字段',
    `bk_string`       varchar(128) COLLATE utf8_unicode_ci      DEFAULT '' COMMENT 'string类型备用字段',
    `geo`             char(3) COLLATE utf8_unicode_ci           DEFAULT '' COMMENT '相关国家编码',
    `lang`            char(10) COLLATE utf8_unicode_ci          DEFAULT '' COMMENT '国家语言代码',
    PRIMARY KEY (`id`),
    KEY               `idx_params_id` (`params_id`),
    KEY               `idx_params_dict_md5` (`params_dict_md5`),
    KEY               `idx_is_avaliable` (`is_available`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;