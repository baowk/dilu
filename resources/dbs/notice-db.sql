/*
 Navicat Premium Data Transfer

 Source Server         : wsl
 Source Server Type    : MySQL
 Source Server Version : 80035
 Source Host           : 172.29.173.47:3306
 Source Schema         : notice-db

 Target Server Type    : MySQL
 Target Server Version : 80035
 File Encoding         : 65001

 Date: 21/11/2023 18:00:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pub_notice
-- ----------------------------
DROP TABLE IF EXISTS `pub_notice`;
CREATE TABLE `pub_notice`  (
  `id` int(0) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int(0) NULL DEFAULT NULL COMMENT '针对组消息',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标题',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '内容',
  `notice_type` tinyint(0) NULL DEFAULT NULL COMMENT '消息类型',
  `op` tinyint(0) NULL DEFAULT NULL COMMENT '操作类型',
  `op_id` int(0) NULL DEFAULT NULL COMMENT '操作id',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `create_by` int(0) NULL DEFAULT NULL COMMENT '创建人',
  `update_by` int(0) NULL DEFAULT NULL COMMENT '更新人',
  `expired` datetime(0) NULL DEFAULT NULL COMMENT '到期时间',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '公用通知' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pub_notice
-- ----------------------------

-- ----------------------------
-- Table structure for user_notice
-- ----------------------------
DROP TABLE IF EXISTS `user_notice`;
CREATE TABLE `user_notice`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int(0) NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int(0) NULL DEFAULT NULL COMMENT '用户id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标题',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '内容',
  `notice_type` tinyint(0) NULL DEFAULT NULL COMMENT '消息类型',
  `op` tinyint(0) NULL DEFAULT NULL COMMENT '操作类型',
  `op_id` int(0) NULL DEFAULT NULL COMMENT '操作对象id',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态 1未读 2已读 -1回收站',
  `create_by` int(0) NULL DEFAULT NULL COMMENT '创建人',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `delete_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户通知' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_notice
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
