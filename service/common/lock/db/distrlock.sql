/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1_3306
Source Server Version : 50724
Source Host           : 127.0.0.1:3306
Source Database       : distrlock

Target Server Type    : MYSQL
Target Server Version : 50724
File Encoding         : 65001

Date: 2019-11-28 12:00:08
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for method_lock
-- ----------------------------
DROP TABLE IF EXISTS `method_lock`;
CREATE TABLE `method_lock` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `method_name` varchar(64) NOT NULL COMMENT '锁定的方法名',
  `desc` varchar(255) NOT NULL COMMENT '备注信息',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uidx_method_name` (`method_name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23741 DEFAULT CHARSET=utf8 COMMENT='锁定中的方法';

-- ----------------------------
-- Event structure for task_clear_timeout_lock
-- ----------------------------
DROP EVENT IF EXISTS `task_clear_timeout_lock`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` EVENT `task_clear_timeout_lock` ON SCHEDULE EVERY 3 SECOND STARTS '2019-11-28 11:59:37' ON COMPLETION NOT PRESERVE ENABLE DO DELETE FROM method_lock WHERE TIMESTAMPDIFF(SECOND,update_time,NOW()) > 5
;;
DELIMITER ;
