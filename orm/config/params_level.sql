/*
 Navicat Premium Data Transfer

 Source Server         : 172.16.3.148
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : 172.16.3.148
 Source Database       : test

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : utf-8

 Date: 05/29/2019 16:18:45 PM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `params_level`
-- ----------------------------
DROP TABLE IF EXISTS `params_level`;
CREATE TABLE `params_level` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `level` int(11) DEFAULT NULL,
  `name` varchar(55) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `tree_path` varchar(55) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
