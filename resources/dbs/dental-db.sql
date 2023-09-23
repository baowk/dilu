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

 Date: 23/09/2023 10:04:50
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
  `customer_id` bigint(0) NULL DEFAULT NULL COMMENT '顾客',
  `user_id` bigint(0) NULL DEFAULT NULL COMMENT '用户id',
  `team_id` bigint(0) NULL DEFAULT NULL COMMENT '团队id',
  `total` decimal(10, 2) NULL DEFAULT NULL COMMENT '金额',
  `real_total` decimal(10, 2) NULL DEFAULT NULL COMMENT '折后金额',
  `paid_total` decimal(10, 2) NULL DEFAULT NULL COMMENT '已支付金额',
  `link_id` bigint(0) NULL DEFAULT NULL COMMENT '关联订单',
  `trade_at` datetime(0) NULL DEFAULT NULL COMMENT '交易日期',
  `trade_status` tinyint(0) NULL DEFAULT NULL COMMENT '交易类型 1 成交 2补尾款 3退款',
  `dental_count` tinyint(0) NULL DEFAULT NULL COMMENT '颗数',
  `brand` tinyint(0) NULL DEFAULT NULL COMMENT '品牌',
  `implanted_count` tinyint(0) NULL DEFAULT NULL COMMENT '已种颗数',
  `implant` tinyint(0) NULL DEFAULT NULL COMMENT '是否已种',
  `implant_date` datetime(0) NULL DEFAULT NULL COMMENT '植入日期',
  `doctor` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '医生',
  `pack` tinyint(0) NULL DEFAULT NULL COMMENT '1 普通 2 半口 3 全口',
  `payback_date` datetime(0) NULL DEFAULT NULL COMMENT '预定回款日期',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

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
  `birthday` date NULL DEFAULT NULL COMMENT '生日',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `wechat` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '微信号',
  `gender` tinyint(0) NULL DEFAULT NULL COMMENT '性别',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '地址',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '描述',
  `sales_id` bigint(0) NULL DEFAULT NULL COMMENT '销售人员',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customer
-- ----------------------------

-- ----------------------------
-- Table structure for event_day_st
-- ----------------------------
DROP TABLE IF EXISTS `event_day_st`;
CREATE TABLE `event_day_st`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `day` datetime(0) NULL DEFAULT NULL COMMENT '时间',
  `team_id` bigint(0) NULL DEFAULT NULL COMMENT '团队id',
  `user_id` bigint(0) NULL DEFAULT NULL COMMENT '用户id',
  `new_customer_cnt` bigint(0) NULL DEFAULT NULL COMMENT '留存',
  `first_diagnosis` bigint(0) NULL DEFAULT NULL COMMENT '初诊',
  `further_diagnosis` bigint(0) NULL DEFAULT NULL COMMENT '复诊',
  `deal` bigint(0) NULL DEFAULT NULL COMMENT '成交',
  `rest` tinyint(0) NULL DEFAULT NULL COMMENT '休息',
  `created_at` bigint(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of event_day_st
-- ----------------------------

-- ----------------------------
-- Table structure for summary_plan_day
-- ----------------------------
DROP TABLE IF EXISTS `summary_plan_day`;
CREATE TABLE `summary_plan_day`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `day` bigint(0) NULL DEFAULT NULL COMMENT '天',
  `team_id` bigint(0) NULL DEFAULT NULL COMMENT '团队id',
  `user_id` bigint(0) NULL DEFAULT NULL COMMENT '用户id',
  `summary` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '今日总结',
  `plan` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '明日计划',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of summary_plan_day
-- ----------------------------

-- ----------------------------
-- Table structure for target_task
-- ----------------------------
DROP TABLE IF EXISTS `target_task`;
CREATE TABLE `target_task`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `month` bigint(0) NULL DEFAULT NULL COMMENT '月',
  `team_id` bigint(0) NULL DEFAULT NULL COMMENT '团队id',
  `user_id` bigint(0) NULL DEFAULT NULL COMMENT '用户id',
  `new_customer_cnt` bigint(0) NULL DEFAULT NULL COMMENT '留存任务',
  `first_diagnosis` bigint(0) NULL DEFAULT NULL COMMENT '导诊任务',
  `deal` bigint(0) NULL DEFAULT NULL COMMENT '成交任务',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of target_task
-- ----------------------------

-- ----------------------------
-- Table structure for team
-- ----------------------------
DROP TABLE IF EXISTS `team`;
CREATE TABLE `team`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` bigint(0) NULL DEFAULT NULL COMMENT '上级团队',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '团队路径',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '团队名',
  `owner` bigint(0) NULL DEFAULT NULL COMMENT '团队拥有者',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of team
-- ----------------------------

-- ----------------------------
-- Table structure for team_member
-- ----------------------------
DROP TABLE IF EXISTS `team_member`;
CREATE TABLE `team_member`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` bigint(0) NULL DEFAULT NULL COMMENT '用户id',
  `team_id` bigint(0) NULL DEFAULT NULL COMMENT '团队id',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '姓名',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '电话',
  `gender` tinyint(0) NULL DEFAULT NULL COMMENT '性别',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `role` tinyint(0) NULL DEFAULT NULL COMMENT '角色 1主管 2副主管 4普通',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of team_member
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
