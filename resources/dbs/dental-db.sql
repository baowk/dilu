/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : localhost:3306
 Source Schema         : dental-db

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 11/10/2023 18:02:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bill
-- ----------------------------
DROP TABLE IF EXISTS `bill`;
CREATE TABLE `bill`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `no` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '订单号',
  `customer_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '顾客',
  `user_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `team_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `total` decimal(10, 2) NULL DEFAULT NULL COMMENT '金额',
  `real_total` decimal(10, 2) NULL DEFAULT NULL COMMENT '折后金额',
  `paid_total` decimal(10, 2) NULL DEFAULT NULL COMMENT '已支付金额',
  `link_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '关联订单',
  `trade_at` datetime(0) NULL DEFAULT NULL COMMENT '交易日期',
  `trade_status` tinyint(0) NULL DEFAULT NULL COMMENT '交易类型1 成交 2补尾款  3补上月欠款 10退款',
  `dental_count` tinyint(0) NULL DEFAULT NULL COMMENT '颗数',
  `brand` tinyint(0) NULL DEFAULT NULL COMMENT '品牌',
  `implanted_count` tinyint(0) NULL DEFAULT NULL COMMENT '已种颗数',
  `implant` tinyint(0) NULL DEFAULT NULL COMMENT '是否已种',
  `implant_date` datetime(0) NULL DEFAULT NULL COMMENT '植入日期',
  `doctor` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '医生',
  `pack` tinyint(0) NULL DEFAULT NULL COMMENT '1 普通 2 半口 3 全口',
  `payback_date` datetime(0) NULL DEFAULT NULL COMMENT '预定回款日期',
  `tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标签',
  `prj_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '项目',
  `other_prj` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '其他项目',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_bill_create_by`(`create_by`) USING BTREE,
  INDEX `idx_bill_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of bill
-- ----------------------------

-- ----------------------------
-- Table structure for customer
-- ----------------------------
DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '姓名',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `wechat` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '微信号',
  `gender` tinyint(0) NULL DEFAULT NULL COMMENT '性别',
  `age` tinyint(0) UNSIGNED NULL DEFAULT NULL COMMENT '年龄',
  `birthday` bigint(0) NULL DEFAULT NULL COMMENT '生日',
  `source` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '来源',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '地址',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '描述',
  `user_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `team_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `inviter` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '邀请人',
  `inviter_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邀请人名',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_customer_user_id`(`user_id`) USING BTREE,
  INDEX `idx_customer_team_id`(`team_id`) USING BTREE,
  INDEX `idx_customer_create_by`(`create_by`) USING BTREE,
  INDEX `idx_customer_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customer
-- ----------------------------

-- ----------------------------
-- Table structure for event_day_st
-- ----------------------------
DROP TABLE IF EXISTS `event_day_st`;
CREATE TABLE `event_day_st`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `day` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '时间',
  `team_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `new_customer_cnt` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '留存',
  `first_diagnosis` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '初诊',
  `further_diagnosis` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '复诊',
  `deal` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '成交',
  `invitation` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '明日邀约',
  `rest` tinyint(0) NULL DEFAULT NULL COMMENT '休息',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_event_day_st_create_by`(`create_by`) USING BTREE,
  INDEX `idx_event_day_st_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of event_day_st
-- ----------------------------
INSERT INTO `event_day_st` VALUES (1, 20231001, 1, 2, '', 0, 0, 0, 0, 0, 1, '2023-10-03 18:32:55', '2023-10-03 18:32:55', 0, 0);

-- ----------------------------
-- Table structure for summary_plan_day
-- ----------------------------
DROP TABLE IF EXISTS `summary_plan_day`;
CREATE TABLE `summary_plan_day`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `day` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '天',
  `team_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `summary` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '今日总结',
  `plan` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '明日计划',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_summary_plan_day_create_by`(`create_by`) USING BTREE,
  INDEX `idx_summary_plan_day_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of summary_plan_day
-- ----------------------------
INSERT INTO `summary_plan_day` VALUES (1, 20231001, 1, 2, '', '今日休息，督促组员邀约初诊', '明日有事休息，组内有两个初诊，已安排同事帮忙接待，督促组员邀约初诊以及提高留存数量', '2023-10-03 17:19:57', '2023-10-03 17:19:57', 0, 0);
INSERT INTO `summary_plan_day` VALUES (2, 20231002, 1, 2, '', '今日休息，督促组员邀约初诊', '上午去场地留存，下午参加市场大会，晚上集体打邀约初诊', '2023-10-03 17:21:29', '2023-10-03 17:21:29', 0, 0);

-- ----------------------------
-- Table structure for target_task
-- ----------------------------
DROP TABLE IF EXISTS `target_task`;
CREATE TABLE `target_task`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `day_type` tinyint(0) UNSIGNED NULL DEFAULT NULL COMMENT '时间类型:月 30,周 7',
  `day` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '时间:202310',
  `team_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `new_customer_cnt` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '留存任务',
  `first_diagnosis` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '导诊任务',
  `deal` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '成交任务',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_target_task_create_by`(`create_by`) USING BTREE,
  INDEX `idx_target_task_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of target_task
-- ----------------------------
INSERT INTO `target_task` VALUES (1, 30, 202310, 1, 2, '/0/1/2', 72, 16, 80000, NULL, NULL, NULL, NULL);
INSERT INTO `target_task` VALUES (2, 30, 202310, 1, 3, '/0/1/2', 72, 16, 80000, NULL, NULL, NULL, NULL);
INSERT INTO `target_task` VALUES (3, 30, 202310, 1, 4, '/0/1/2', 72, 16, 80000, NULL, NULL, NULL, NULL);
INSERT INTO `target_task` VALUES (4, 30, 202310, 1, 5, '/0/1/2', 72, 16, 80000, NULL, NULL, NULL, NULL);
INSERT INTO `target_task` VALUES (5, 30, 202310, 1, 6, '/0/1/2', 72, 16, 80000, NULL, NULL, NULL, NULL);
INSERT INTO `target_task` VALUES (6, 30, 202310, 1, 7, '/0/1/2', 72, 16, 80000, NULL, NULL, NULL, NULL);

SET FOREIGN_KEY_CHECKS = 1;
