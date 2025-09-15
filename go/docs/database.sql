CREATE TABLE `passport` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) NOT NULL DEFAULT '',
  `user_password` varchar(100) NOT NULL DEFAULT '',
  `user_email` varchar(100) NOT NULL DEFAULT '',
  `user_info` varchar(1024) NOT NULL DEFAULT '',
  `head_img` varchar(1024) NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL DEFAULT 1 COMMENT "账号状态：1正常，0拉黑",
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_name` (`user_name`),
  UNIQUE KEY `uk_user_email` (`user_email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;