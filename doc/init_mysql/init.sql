CREATE DATABASE `demo` /*!40100 COLLATE 'utf8mb4_general_ci' */

CREATE TABLE `uploadfile` (
                              `num` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '序号',
                              `filename` CHAR(50) NOT NULL COMMENT '文件名称' COLLATE 'utf8mb4_general_ci',
                              `fullpath` VARCHAR(200) NOT NULL COMMENT '文件存放路径' COLLATE 'utf8mb4_general_ci',
                              `md5sum` CHAR(200) NOT NULL COMMENT 'md5值' COLLATE 'utf8mb4_general_ci',
                              `extract` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '是否已解压',
                              `delete` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '是否已删除',
                              PRIMARY KEY (`num`) USING BTREE
)
    COMMENT='保存上传文件名称、状态、md5、路径等信息'
    COLLATE='utf8mb4_general_ci'
    ENGINE=InnoDB
;

CREATE TABLE `extractfile` (
                               `num` INT(11) NOT NULL AUTO_INCREMENT COMMENT '序号',
                               `filename` CHAR(200) NOT NULL COMMENT '解压后的文件名称' COLLATE 'utf8mb4_general_ci',
                               `fromfile` CHAR(200) NOT NULL COMMENT '来自的哪个压缩文件' COLLATE 'utf8mb4_general_ci',
                               `job_exec` TINYINT(4) NOT NULL DEFAULT '0' COMMENT '是否已经执行job，0未执行、1已执行',
                               `job_status` TINYINT(4) NOT NULL DEFAULT '0' COMMENT 'job执行是否成功，0成功、1失败',
                               INDEX `num` (`num`) USING BTREE
)
    COMMENT='保存解压后文件数据\r\n'
    COLLATE='utf8mb4_general_ci'
    ENGINE=InnoDB
;


