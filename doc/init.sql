USE beego;
DROP TABLE IF EXISTS `beego_user`;
CREATE TABLE `beego_user` (
    `id` INT(11) NOT NULL AUTO_INCREMENT COMMENT '主键 ID',
    `uid` INT(11) NOT NULL DEFAULT 0 COMMENT '用户 uid',
    `name` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户姓名',
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户';

INSERT INTO `beego_user` (uid, name) VALUES (1, '张三'), (2, '李四');