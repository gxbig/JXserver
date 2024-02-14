CREATE TABLE jxserver.servers (
                                  id INT auto_increment NOT NULL COMMENT '主键',
                                  address varchar(100) NOT NULL COMMENT '服务器地址',
                                  code INT NOT NULL COMMENT '服务器编码',
                                  status INT NOT NULL COMMENT '服务器状态',
                                  max INT NOT NULL COMMENT '最大用户数',
                                  recommend INT NOT NULL COMMENT '0否1是',
                                  server_name varchar(100) NOT NULL COMMENT '服务器名称',
                                  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  `activated_at` datetime DEFAULT NULL COMMENT '激活时间',
                                  `deleted_at` datetime DEFAULT NULL,
                                  PRIMARY KEY (`id`)
)
    ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO jxserver.servers
(address, code, status, max, server_name, recommend)
VALUES('http://localhost:9000', 1000, 1, 20000, '1区 国士无双',1);

INSERT INTO jxserver.servers
(address, code, status, max, server_name, recommend)
VALUES('http://localhost:9000', 1001, 1, 20000, '2区 风云际会',0);

ALTER TABLE jxserver.servers MODIFY COLUMN status int NOT NULL COMMENT '服务器状态 1:流畅 2：爆满 3：维护';

