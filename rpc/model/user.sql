CREATE TABLE `user`
(
    `id` int NOT NULL AUTO_INCREMENT COMMENT 'id',
    `username` varchar(255) NOT NULL DEFAULT '' UNIQUE COMMENT 'username',
    `password` varchar(255) NOT NULL DEFAULT '' COMMENT 'password',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='user';
