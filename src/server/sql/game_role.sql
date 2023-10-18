-- jxserver.game_role definition

CREATE TABLE jxserver.`game_role` (
                                      `id` int NOT NULL AUTO_INCREMENT COMMENT '游戏角色主键',
                                      `role_name` varchar(100) DEFAULT NULL COMMENT '角色名称',
                                      `type` enum('soldier','master','shooter') NOT NULL COMMENT '角色类型 1、战士： ''soldier''2、 法师：''master''  3、射手：''shooter''',
                                      `hp` int NOT NULL COMMENT '生命值',
                                      `mp` int NOT NULL COMMENT '法力值',
                                      `attack` int NOT NULL COMMENT '攻击力',
                                      `penetrate` int NOT NULL DEFAULT '0' COMMENT '穿透',
                                      `avoid_injury` int NOT NULL DEFAULT '0' COMMENT '免伤',
                                      `hp_grow_up` float NOT NULL DEFAULT '0' COMMENT 'hp每级增长率',
                                      `mp_grow_up` float NOT NULL DEFAULT '0' COMMENT '法力每级增长率',
                                      `spell_damage` int NOT NULL DEFAULT '0' COMMENT '法术伤害',
                                      `attack_grow_up` float NOT NULL DEFAULT '0' COMMENT '攻击力每级增长率',
                                      `spell damage__grow_up` float NOT NULL DEFAULT '0' COMMENT '法术伤害每级增长率',
                                      `spell_defense` int NOT NULL DEFAULT '0' COMMENT '法术防御',
                                      `physical_defense` int NOT NULL DEFAULT '0' COMMENT '物理防御',
                                      `spell_defense_group` float NOT NULL DEFAULT '0' COMMENT '法术防御增长率',
                                      `physical_defense_group` float NOT NULL DEFAULT '0' COMMENT '物理防御增长率',
                                      `grade` int NOT NULL DEFAULT '1' COMMENT '角色默认等级',
                                      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      `activated_at` datetime DEFAULT NULL COMMENT '激活时间',
                                      `deleted_at` datetime DEFAULT NULL,
                                      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;