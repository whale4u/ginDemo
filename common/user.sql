/*建表语句*/
CREATE TABLE `users`
(
    `id`       bigint(20)                        NOT NULL AUTO_INCREMENT,
    `username` varchar(255) CHARACTER SET latin1 NOT NULL,
    `password` varchar(255) CHARACTER SET latin1 DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;