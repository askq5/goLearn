create table if not exists `gorm_learn` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `user_name` varchar(20) ,
    `age` tinyint(8) unsigned NOT NULL DEFAULT 0,
    `level` tinyint(8) unsigned ,
    `department` varchar(20) NOT NULL,
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key (`id`),
    unique key `uniq_user_name` (user_name),
    key `idx_level` (`level`) using BTREE
    ) ENGINE=InnoDB  Auto_increment=123456 default charset = utf8mb4;