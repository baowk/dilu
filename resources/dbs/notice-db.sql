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

 Date: 24/11/2023 15:55:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for pub_notice
-- ----------------------------
DROP TABLE IF EXISTS `pub_notice`;
CREATE TABLE `pub_notice`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int(0) NULL DEFAULT 0 COMMENT '针对组消息',
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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '公用通知' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of pub_notice
-- ----------------------------
INSERT INTO `pub_notice` VALUES (1, 0, 'test', 'testcontent', NULL, NULL, NULL, 1, 1, 1, '2023-12-07 14:40:06', '2023-11-24 14:40:13', '2023-11-24 14:40:17');
INSERT INTO `pub_notice` VALUES (2, 1, 'test2', 'content', NULL, NULL, NULL, 1, 1, 1, '2023-11-25 14:53:05', '2023-11-24 14:53:14', '2023-11-24 14:53:18');

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task`  (
  `id` bigint(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int(0) NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int(0) NULL DEFAULT NULL COMMENT '用户id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '任务标题',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '任务内容',
  `task_type` tinyint(0) NULL DEFAULT NULL COMMENT '任务类型',
  `op` int(0) NULL DEFAULT NULL COMMENT '操作类型',
  `op_id` int(0) NULL DEFAULT NULL COMMENT '操作id',
  `begin_at` datetime(0) NULL DEFAULT NULL COMMENT '开始时间',
  `end_at` datetime(0) NULL DEFAULT NULL COMMENT '结束时间',
  `reminder_time` datetime(0) NULL DEFAULT NULL COMMENT '提醒时间',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态1开启2关闭',
  `reminder_status` tinyint(0) NULL DEFAULT NULL COMMENT '提醒状态 1开启 2关闭',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of task
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
  `pub_id` int(0) NULL DEFAULT NULL COMMENT '公共id',
  `create_by` int(0) NULL DEFAULT NULL COMMENT '创建人',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime(0) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户通知' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_notice
-- ----------------------------
INSERT INTO `user_notice` VALUES (2, 1, 2, 'test2', 'content', 0, 0, 0, 1, 2, 1, '2023-11-24 14:53:14', '2023-11-24 14:53:18', NULL);
INSERT INTO `user_notice` VALUES (3, 1, 2, 'test', 'testcontent', 0, 0, 0, 1, 1, 1, '2023-11-24 14:40:13', '2023-11-24 14:40:17', NULL);
INSERT INTO `user_notice` VALUES (4, 1, 2, 'test3', '33333', 0, 0, 0, 1, NULL, 1, '2023-11-24 14:55:59', '2023-11-24 14:56:03', NULL);

SET FOREIGN_KEY_CHECKS = 1;
