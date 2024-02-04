-- jxserver.game_role definition

CREATE TABLE jxserver.`game_role` (
                                      `id` int NOT NULL AUTO_INCREMENT COMMENT '游戏角色主键',
                                      `role_name` varchar(100) DEFAULT NULL COMMENT '角色名称',
                                      `role_type` varchar(100) NOT NULL COMMENT '1、战士,2、法师, 3、射手',
                                      `hp` int NOT NULL COMMENT '生命值',
                                      `mp` int NOT NULL COMMENT '法力值',
                                      `attack` int NOT NULL COMMENT '攻击力',
                                      `penetrate` int NOT NULL DEFAULT '0' COMMENT '穿透',
                                      `avoid_injury` int NOT NULL DEFAULT '0' COMMENT '免伤',
                                      `hp_grow_up` int NOT NULL DEFAULT '0' COMMENT 'hp1~20每级增长',
                                      `hp_grow_m_up` int NOT NULL DEFAULT '0' COMMENT 'hp21~60每级增长',
                                      `hp_grow_l_up` int NOT NULL DEFAULT '0' COMMENT 'hp60~每级增长',
                                      `mp_grow_up` int NOT NULL DEFAULT '0' COMMENT '法力1~20每级增长',
                                      `mp_grow_m_up` int NOT NULL DEFAULT '0' COMMENT '法力21~60每级增长',
                                      `mp_grow_l_up` int NOT NULL DEFAULT '0' COMMENT '法力60~每级增长',
                                      `spell_damage` int NOT NULL DEFAULT '0' COMMENT '法术伤害',
                                      `attack_grow_up` int NOT NULL DEFAULT '0' COMMENT '攻击力1~20每级增长',
                                      `attack_grow_m_up` int NOT NULL DEFAULT '0' COMMENT '攻击力21~60每级增长',
                                      `attack_grow_l_up` int NOT NULL DEFAULT '0' COMMENT '攻击力60~每级增长',
                                      `spell_damage_grow_up` int NOT NULL DEFAULT '0' COMMENT '法术伤害1~20每级增长',
                                      `spell_damage_grow_m_up` int NOT NULL DEFAULT '0' COMMENT '法术伤害21~60每级增长',
                                      `spell_damage_grow_l_up` int NOT NULL DEFAULT '0' COMMENT '法术伤害60~每级增长',
                                      `spell_defense` int NOT NULL DEFAULT '0' COMMENT '法术防御',
                                      `physical_defense` int NOT NULL DEFAULT '0' COMMENT '物理防御',
                                      `spell_defense_grow_up` int NOT NULL DEFAULT '0' COMMENT '法术防御1~20增长',
                                      `spell_defense_grow_m_up` int NOT NULL DEFAULT '0' COMMENT '法术防御21~60增长',
                                      `spell_defense_grow_l_up` int NOT NULL DEFAULT '0' COMMENT '法术防御60~增长',
                                      `physical_defense_grow_up` int NOT NULL DEFAULT '0' COMMENT '物理防御1~20增长',
                                      `physical_defense_grow_m_up` int NOT NULL DEFAULT '0' COMMENT '物理防御21~60增长',
                                      `physical_defense_grow_l_up` int NOT NULL DEFAULT '0' COMMENT '物理防御60~增长',
                                      `grade` int NOT NULL DEFAULT '1' COMMENT '角色默认等级',
                                      `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                      `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                      `activated_at` datetime DEFAULT NULL COMMENT '激活时间',
                                      `deleted_at` datetime DEFAULT NULL,
                                      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
ALTER TABLE jxserver.game_role ADD speed INT NOT NULL COMMENT '速度';
ALTER TABLE jxserver.game_role ADD attack_speed INT NOT NULL COMMENT '攻速';


insert into jxserver.`game_role` (role_name, role_type, hp, mp, attack, penetrate, avoid_injury, hp_grow_up,
                                  hp_grow_m_up, hp_grow_l_up, mp_grow_up, mp_grow_m_up, mp_grow_l_up, spell_damage,
                                  attack_grow_up, attack_grow_m_up, attack_grow_l_up, spell_damage_grow_up, spell_damage_grow_m_up,
                                  spell_damage_grow_l_up, spell_defense, physical_defense, spell_defense_grow_up,
                                  spell_defense_grow_m_up, spell_defense_grow_l_up, physical_defense_grow_up,
                                  physical_defense_grow_m_up, physical_defense_grow_l_up, grade, speed, attack_speed)
values ('战士','1',200,100,20,0,0,180,270,360,100,100,100,0,20,30,40,0,0,0,10,10,10,15,20,10,15,20,1,100,1)

insert into jxserver.`game_role` (role_name, role_type, hp, mp, attack, penetrate, avoid_injury, hp_grow_up,
                                      hp_grow_m_up, hp_grow_l_up, mp_grow_up, mp_grow_m_up, mp_grow_l_up, spell_damage,
                                      attack_grow_up, attack_grow_m_up, attack_grow_l_up, spell_damage_grow_up, spell_damage_grow_m_up,
                                      spell_damage_grow_l_up, spell_defense, physical_defense, spell_defense_grow_up,
                                      spell_defense_grow_m_up, spell_defense_grow_l_up, physical_defense_grow_up,
                                      physical_defense_grow_m_up, physical_defense_grow_l_up, grade, speed, attack_speed)
values ('法师','2',200,200,10,0,0,150,230,320,200,200,200,30,10,10,10,30,45,60,5,5,5,7,10,5,7,10,1,400,1)


insert into jxserver.`game_role` (role_name, role_type, hp, mp, attack, penetrate, avoid_injury, hp_grow_up,
                                  hp_grow_m_up, hp_grow_l_up, mp_grow_up, mp_grow_m_up, mp_grow_l_up, spell_damage,
                                  attack_grow_up, attack_grow_m_up, attack_grow_l_up, spell_damage_grow_up, spell_damage_grow_m_up,
                                  spell_damage_grow_l_up, spell_defense, physical_defense, spell_defense_grow_up,
                                  spell_defense_grow_m_up, spell_defense_grow_l_up, physical_defense_grow_up,
                                  physical_defense_grow_m_up, physical_defense_grow_l_up, grade, speed, attack_speed)
values ('射手','3',200,100,30,0,0,150,230,320,100,100,100,0,30,45,60,0,0,0,4,4,4,6,8,4,6,8,1,400,2)



