
USE `test_db`;


DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`(
    `id`              bigint(20)    NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username`        varchar(64)   NOT NULL DEFAULT '' COMMENT 'username',
    `password`        varchar(64)   NOT NULL DEFAULT '' COMMENT 'password',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT = 'user';