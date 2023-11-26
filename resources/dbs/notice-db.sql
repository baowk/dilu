/*
 Navicat Premium Data Transfer

 Source Server         : local_u
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : notice-db

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 26/11/2023 19:56:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ai_msg
-- ----------------------------
DROP TABLE IF EXISTS `ai_msg`;
CREATE TABLE `ai_msg`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `platform` tinyint NULL DEFAULT NULL COMMENT '平台',
  `model` tinyint NULL DEFAULT NULL COMMENT '模型',
  `user_id` int NULL DEFAULT NULL COMMENT '用户id',
  `topic_id` int NULL DEFAULT NULL COMMENT '话题',
  `is_user` tinyint NULL DEFAULT NULL COMMENT '用户',
  `text` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '文本',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ai_msg
-- ----------------------------

-- ----------------------------
-- Table structure for ai_topic
-- ----------------------------
DROP TABLE IF EXISTS `ai_topic`;
CREATE TABLE `ai_topic`  (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
  `platform` tinyint NULL DEFAULT NULL COMMENT '平台',
  `model` tinyint NULL DEFAULT NULL COMMENT '模型',
  `topic` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '话题',
  `user_id` int NULL DEFAULT NULL COMMENT '用户id',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ai_topic
-- ----------------------------

-- ----------------------------
-- Table structure for pub_notice
-- ----------------------------
DROP TABLE IF EXISTS `pub_notice`;
CREATE TABLE `pub_notice`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int NULL DEFAULT 0 COMMENT '针对组消息',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标题',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '内容',
  `notice_type` tinyint NULL DEFAULT NULL COMMENT '消息类型',
  `op` tinyint NULL DEFAULT NULL COMMENT '操作类型',
  `op_id` int NULL DEFAULT NULL COMMENT '操作id',
  `status` tinyint NULL DEFAULT NULL COMMENT '状态',
  `create_by` int NULL DEFAULT NULL COMMENT '创建人',
  `update_by` int NULL DEFAULT NULL COMMENT '更新人',
  `expired` datetime NULL DEFAULT NULL COMMENT '到期时间',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '公用通知' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of pub_notice
-- ----------------------------
INSERT INTO `pub_notice` VALUES (1, 0, 'test', 'testcontent', NULL, NULL, NULL, 1, 1, 1, '2023-12-07 14:40:06', '2023-11-24 14:40:13', '2023-11-24 14:40:17');
INSERT INTO `pub_notice` VALUES (2, 1, 'test2', 'content', NULL, NULL, NULL, 1, 1, 1, '2023-11-25 14:53:05', '2023-11-24 14:53:14', '2023-11-24 14:53:18');
INSERT INTO `pub_notice` VALUES (3, 0, '地方大师傅', '的发射点', 0, 0, 0, 1, 0, 0, '2023-11-28 05:02:02', '2023-11-26 09:12:03', '2023-11-26 13:01:20');
INSERT INTO `pub_notice` VALUES (4, 0, '活动啊', '活动通知一下', 0, 0, 0, 1, 0, 0, '2023-11-30 00:00:00', '2023-11-26 13:01:11', '2023-11-26 13:01:11');

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int NULL DEFAULT NULL COMMENT '用户id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '任务标题',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '任务内容',
  `task_type` tinyint NULL DEFAULT NULL COMMENT '任务类型',
  `op` int NULL DEFAULT NULL COMMENT '操作类型',
  `op_id` int NULL DEFAULT NULL COMMENT '操作id',
  `begin_at` datetime NULL DEFAULT NULL COMMENT '开始时间',
  `end_at` datetime NULL DEFAULT NULL COMMENT '结束时间',
  `reminder_time` datetime NULL DEFAULT NULL COMMENT '提醒时间',
  `status` tinyint NULL DEFAULT NULL COMMENT '状态1开启2关闭',
  `reminder_status` tinyint NULL DEFAULT NULL COMMENT '提醒状态 1开启 2关闭',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of task
-- ----------------------------
INSERT INTO `task` VALUES (1, 1, 2, '嘻嘻嘻', '顶顶顶顶顶', 0, 0, 0, '2023-11-26 15:04:02', '2023-11-27 00:00:00', '2023-11-26 13:02:03', 1, 1, '2023-11-26 13:02:08', '2023-11-26 13:02:08', NULL);
INSERT INTO `task` VALUES (2, -1, 1, '这是一个任务', '任务开始了', 0, 0, 0, '2023-11-26 19:52:36', '2023-11-26 19:52:38', '2023-11-26 19:52:40', 1, 1, '2023-11-26 19:52:47', '2023-11-26 19:52:47', NULL);

-- ----------------------------
-- Table structure for user_notice
-- ----------------------------
DROP TABLE IF EXISTS `user_notice`;
CREATE TABLE `user_notice`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int NULL DEFAULT NULL COMMENT '用户id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标题',
  `content` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '内容',
  `notice_type` tinyint NULL DEFAULT NULL COMMENT '消息类型',
  `op` tinyint NULL DEFAULT NULL COMMENT '操作类型',
  `op_id` int NULL DEFAULT NULL COMMENT '操作对象id',
  `status` tinyint NULL DEFAULT NULL COMMENT '状态 1未读 2已读 -1回收站',
  `pub_id` int NULL DEFAULT NULL COMMENT '公共id',
  `create_by` int NULL DEFAULT NULL COMMENT '创建人',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户通知' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_notice
-- ----------------------------
INSERT INTO `user_notice` VALUES (2, 1, 2, 'test2', 'content', 0, 0, 0, 1, 2, 1, '2023-11-24 14:53:14', '2023-11-24 14:53:18', NULL);
INSERT INTO `user_notice` VALUES (3, 1, 2, 'test', 'testcontent', 0, 0, 0, 1, 1, 1, '2023-11-24 14:40:13', '2023-11-24 14:40:17', NULL);
INSERT INTO `user_notice` VALUES (4, 1, 2, 'test3', '33333', 0, 0, 0, 1, NULL, 1, '2023-11-24 14:55:59', '2023-11-24 14:56:03', NULL);
INSERT INTO `user_notice` VALUES (5, -1, 1, '活动啊', '活动通知一下', 0, 0, 0, 1, 4, 0, '2023-11-26 13:01:11', '2023-11-26 13:01:11', NULL);
INSERT INTO `user_notice` VALUES (6, -1, 1, '地方大师傅', '的发射点', 0, 0, 0, 1, 3, 0, '2023-11-26 09:12:03', '2023-11-26 13:01:20', NULL);
INSERT INTO `user_notice` VALUES (7, -1, 1, 'test', 'testcontent', 0, 0, 0, 1, 1, 1, '2023-11-24 14:40:13', '2023-11-24 14:40:17', NULL);

SET FOREIGN_KEY_CHECKS = 1;
