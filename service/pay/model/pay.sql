CREATE TABLE `pay` (
	`id` bigint NOT NULL AUTO_INCREMENT,
	`uid` bigint NOT NULL DEFAULT '0' COMMENT '用户ID',
	`oid` bigint NOT NULL DEFAULT '0' COMMENT '订单ID',
	`amount` int(10) NOT NULL DEFAULT '0'  COMMENT '产品金额',
	`source` tinyint(3) NOT NULL DEFAULT '0' COMMENT '支付方式',
	`status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '支付状态',
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `idx_uid` (`uid`),
	KEY `idx_oid` (`oid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;