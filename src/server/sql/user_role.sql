-- jxserver.user_role definition

CREATE TABLE jxserver.`user_role` (
                             `id` int NOT NULL AUTO_INCREMENT COMMENT '用户角色主键',
                             `user_id` int NOT NULL COMMENT '用户user表id',
                             `game_role_id` int NOT NULL COMMENT 'game_role表id',
                             `grade` int NOT NULL DEFAULT '1' COMMENT '等级',
                             `attack` int NOT NULL DEFAULT '0' COMMENT '攻击力',
                             `hp` int DEFAULT '0' COMMENT '生命值',
                             `mp` int NOT NULL DEFAULT '0' COMMENT '法力值',
                             `penetrate` int NOT NULL DEFAULT '0' COMMENT '穿透',
                             `avoid_injury` int NOT NULL DEFAULT '0' COMMENT '免伤',
                             `spell_damage` int NOT NULL DEFAULT '0' COMMENT '法术伤害',
                             `spell_defense` int NOT NULL DEFAULT '0' COMMENT '法术防御',
                             `physical_defense` int NOT NULL DEFAULT '0' COMMENT '物理防御',
                             `activated_at` datetime DEFAULT NULL COMMENT '激活时间',
                             `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                             `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `deleted_at` datetime DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             KEY `user_role_user_id_IDX` (`user_id`) USING BTREE,
                             KEY `user_role_FK_1` (`game_role_id`),
                             CONSTRAINT `user_role_FK` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
                             CONSTRAINT `user_role_FK_1` FOREIGN KEY (`game_role_id`) REFERENCES `game_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;