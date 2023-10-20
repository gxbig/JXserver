-- jxserver.`user` definition

CREATE TABLE jxserver.`user` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `age` int DEFAULT NULL COMMENT '年龄',
                        `gender` enum('0','1') CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '性别',
                        `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '姓名',
                        `identity` varchar(50) DEFAULT NULL COMMENT '身份证号码',
                        `date_of_birth` date DEFAULT NULL COMMENT '出生日期',
                        `phone` int DEFAULT NULL COMMENT '手机号',
                        `address` varchar(100) DEFAULT NULL COMMENT '地址',
                        `server` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '服务器区',
                        `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱地址',
                        `postal-code` varchar(30) DEFAULT NULL COMMENT '邮政编码',
                        `qq` varchar(30) DEFAULT NULL COMMENT 'qq号',
                        `wx` varchar(100) DEFAULT NULL COMMENT '微信号',
                        `zfb` varchar(100) DEFAULT NULL COMMENT '支付宝账号',
                        `account` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '账户',
                        `pw` varchar(30) NOT NULL COMMENT '密码',
                        `deleted_at` datetime DEFAULT NULL,
                        `activated_at` datetime DEFAULT NULL COMMENT '激活时间',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `user_account_IDX` (`account`) USING BTREE,
                        UNIQUE KEY `user_un` (`identity`),
                        UNIQUE KEY `user_phone` (`phone`),
                        UNIQUE KEY `user_un_qq` (`qq`),
                        UNIQUE KEY `user_un_wx` (`wx`),
                        UNIQUE KEY `user_un_zfb` (`zfb`),
                        UNIQUE KEY `user_un_email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

ALTER TABLE jxserver.`user` CHANGE `postal-code` postal_code varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '邮政编码';
ALTER TABLE jxserver.`user` MODIFY COLUMN name varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '姓名';
ALTER TABLE jxserver.`user` MODIFY COLUMN email varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '邮箱地址';

ALTER TABLE jxserver.`user` MODIFY COLUMN gender enum('0','1',"") CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '性别';
ALTER TABLE jxserver.`user` MODIFY COLUMN date_of_birth varchar(20) NULL COMMENT '出生日期';
ALTER TABLE jxserver.`user` MODIFY COLUMN phone varchar(20) NULL COMMENT '手机号';
CREATE INDEX idx_deleted_at USING BTREE ON jxserver.`user` (deleted_at);

ALTER TABLE jxserver.`user` DROP KEY user_un;
ALTER TABLE jxserver.`user` DROP KEY user_phone;
ALTER TABLE jxserver.`user` DROP KEY user_un_qq;
ALTER TABLE jxserver.`user` DROP KEY user_un_wx;
ALTER TABLE jxserver.`user` DROP KEY user_un_zfb;
