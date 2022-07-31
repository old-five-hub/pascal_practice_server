CREATE TABLE `language_tag` (
                                `id` int unsigned NOT NULL AUTO_INCREMENT,
                                `name` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
                                `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
                                `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                `hot` int DEFAULT '0' COMMENT '1 true 0 false',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;