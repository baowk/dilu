/*
 Navicat Premium Data Transfer

 Source Server         : local_u
 Source Server Type    : MySQL
 Source Server Version : 80032 (8.0.32)
 Source Host           : localhost:3306
 Source Schema         : dental-db

 Target Server Type    : MySQL
 Target Server Version : 80032 (8.0.32)
 File Encoding         : 65001

 Date: 04/11/2023 21:10:01
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bill
-- ----------------------------
DROP TABLE IF EXISTS `bill`;
CREATE TABLE `bill`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `no` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '订单号',
  `customer_id` int UNSIGNED NULL DEFAULT NULL COMMENT '顾客',
  `user_id` int UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `team_id` int UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '金额',
  `real_amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '折后金额',
  `paid_amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '已支付金额',
  `debt_amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '回收上月欠款',
  `refund_amount` decimal(10, 2) NULL DEFAULT NULL COMMENT '退款',
  `link_id` int UNSIGNED NULL DEFAULT NULL COMMENT '关联订单',
  `trade_at` datetime NULL DEFAULT NULL COMMENT '交易日期',
  `trade_type` tinyint NULL DEFAULT NULL COMMENT '交易类型1 成交 2补尾款  3补上月欠款 10退款',
  `dental_count` tinyint NULL DEFAULT NULL COMMENT '颗数',
  `brand` tinyint NULL DEFAULT NULL COMMENT '品牌',
  `implanted_count` tinyint NULL DEFAULT NULL COMMENT '已种颗数',
  `implant` tinyint NULL DEFAULT NULL COMMENT '种植状态：1 未种 2部分 3已种',
  `implant_date` datetime NULL DEFAULT NULL COMMENT '植入日期',
  `doctor` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '医生',
  `pack` tinyint NULL DEFAULT NULL COMMENT '1 普通 2 半口 3 全口',
  `payback_date` datetime NULL DEFAULT NULL COMMENT '预定回款日期',
  `tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标签',
  `prj_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '项目',
  `other_prj` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '其他项目',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `create_by` int UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_bill_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_bill_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of bill
-- ----------------------------
INSERT INTO `bill` VALUES (15, '20231031205949384029', 10, 2, 1, '/0/1/2/', 0.00, 0.00, 0.00, 0.00, 0.00, 0, '2023-10-31 08:00:00', 1, 21, 0, 0, 1, NULL, '', 4, NULL, '', '', '根管加冠', '', '2023-10-31 20:59:49', '2023-10-31 20:59:49', 2, 0);
INSERT INTO `bill` VALUES (16, '20231101203741461501', 11, 2, 1, '/0/1/2/', 0.00, 72000.00, 10000.00, 0.00, 0.00, 0, '2023-11-01 08:00:00', 1, 12, 1, 12, 3, NULL, '', 3, NULL, '', '', '', '', '2023-11-01 20:37:41', '2023-11-01 20:39:21', 2, 2);
INSERT INTO `bill` VALUES (17, '20231101203901764452', 12, 2, 1, '/0/1/2/', 0.00, 23000.00, 15000.00, 0.00, 0.00, 0, '2023-11-01 08:00:00', 1, 5, 2, 5, 3, '2023-11-01 08:00:00', '', 1, NULL, '', '', '', '', '2023-11-01 20:39:02', '2023-11-01 20:39:02', 2, 0);
INSERT INTO `bill` VALUES (18, '20231102212406378265', 13, 3, 1, '/0/1/2/', 0.00, 380.00, 380.00, 0.00, 0.00, 0, '2023-11-02 08:00:00', 1, 0, 0, 0, 1, NULL, '', 4, NULL, '', '', '补牙', '', '2023-11-02 21:24:06', '2023-11-02 21:24:06', 2, 0);
INSERT INTO `bill` VALUES (19, '20231102212456325620', 11, 2, 1, '/0/1/2/', 0.00, 0.00, 62000.00, 0.00, 0.00, 16, '2023-11-02 21:24:56', 2, 0, 0, 0, 1, NULL, '', 0, NULL, '', '', '', '', '2023-11-02 21:24:56', '2023-11-02 21:24:56', 2, 0);
INSERT INTO `bill` VALUES (20, '20231104203621052718', 14, 2, 1, '/0/1/2/', 0.00, 240.00, 240.00, 0.00, 0.00, 0, '2023-11-04 20:36:21', 1, 0, 0, 0, 1, NULL, '', 4, NULL, '', '', '补牙', '', '2023-11-04 20:36:21', '2023-11-04 20:36:21', 2, 0);

-- ----------------------------
-- Table structure for customer
-- ----------------------------
DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '姓名',
  `py` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '姓名拼音',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `wechat` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '微信号',
  `gender` tinyint NULL DEFAULT NULL COMMENT '性别',
  `age` tinyint UNSIGNED NULL DEFAULT NULL COMMENT '年龄',
  `birthday` bigint NULL DEFAULT NULL COMMENT '生日',
  `source` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '来源',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '地址',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '描述',
  `user_id` int UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `team_id` int UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `inviter` int UNSIGNED NULL DEFAULT NULL COMMENT '邀请人',
  `inviter_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邀请人名',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `create_by` int UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_customer_user_id`(`user_id` ASC) USING BTREE,
  INDEX `idx_customer_team_id`(`team_id` ASC) USING BTREE,
  INDEX `idx_customer_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_customer_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of customer
-- ----------------------------
INSERT INTO `customer` VALUES (1, '李杏利', NULL, '', '', 0, 0, 0, '', '', '', 2, 1, '/0/1/2/', 0, '', '2023-10-25 11:32:57', '2023-10-25 11:32:57', 0, 0);
INSERT INTO `customer` VALUES (2, '傅见英', NULL, '', '', 0, 0, 0, '', '', '', 2, 1, '/0/1/2/', 0, '', '2023-10-25 15:06:24', '2023-10-25 15:06:24', 0, 0);
INSERT INTO `customer` VALUES (3, '鲁慧萍', NULL, '', '', 0, 0, 0, '', '', '', 5, 1, '/0/1/2/', 0, '', '2023-10-25 15:22:23', '2023-10-25 15:22:23', 0, 0);
INSERT INTO `customer` VALUES (4, '孔友祥', NULL, '', '', 0, 0, 0, '', '', '', 5, 1, '/0/1/2/', 0, '', '2023-10-25 15:30:54', '2023-10-25 15:30:54', 2, 0);
INSERT INTO `customer` VALUES (5, '诸红夏', 'zhu-hong-xia', '', '', 0, 0, 0, '', '', '', 2, 1, '/0/1/2/', 0, '', '2023-10-26 21:10:57', '2023-10-26 21:10:57', 2, 0);
INSERT INTO `customer` VALUES (6, '沈玉祥', 'shen-yu-xiang', '', '', 0, 0, 0, '', '', '', 2, 1, '/0/1/2/', 0, '', '2023-10-30 21:13:01', '2023-10-30 21:13:01', 2, 0);
INSERT INTO `customer` VALUES (7, '滕国俊', 'teng-guo-jun', '', '', 0, 0, 0, '', '', '', 2, 1, '/0/1/2/', 0, '', '2023-10-30 21:13:33', '2023-10-30 21:13:33', 2, 0);
INSERT INTO `customer` VALUES (8, '李明景', 'li-ming-jing', '', '', 0, 0, 0, '', '', '', 7, 1, '/0/1/2/', 0, '', '2023-10-30 21:13:57', '2023-10-30 21:13:57', 2, 0);
INSERT INTO `customer` VALUES (9, '黄振', 'huang-zhen', '', '', 0, 0, 0, '', '', '', 2, 1, '/0/1/2/', 0, '', '2023-10-31 20:59:14', '2023-10-31 20:59:14', 2, 0);
INSERT INTO `customer` VALUES (10, '赵光炎', 'zhao-guang-yan', '', '', 0, 0, 0, '', '', '', 2, 1, '/0/1/2/', 0, '', '2023-10-31 20:59:49', '2023-10-31 20:59:49', 2, 0);
INSERT INTO `customer` VALUES (11, '陈纪生', 'chen-ji-sheng', '', '', 0, 0, 0, '', '', '', 2, 1, '/0/1/2/', 0, '', '2023-11-01 20:37:41', '2023-11-01 20:37:41', 2, 0);
INSERT INTO `customer` VALUES (12, '韩爱琴', 'han-ai-qin', '', '', 0, 0, 0, '', '', '', 2, 1, '/0/1/2/', 0, '', '2023-11-01 20:39:02', '2023-11-01 20:39:02', 2, 0);
INSERT INTO `customer` VALUES (13, '凌曼卿', 'ling-man-qing', '', '', 0, 0, 0, '', '', '', 3, 1, '/0/1/2/', 0, '', '2023-11-02 21:24:06', '2023-11-02 21:24:06', 2, 0);
INSERT INTO `customer` VALUES (14, '宋国甄', 'song-guo-zhen', '', '', 0, 0, 0, '', '', '', 2, 1, '/0/1/2/', 0, '', '2023-11-04 20:36:21', '2023-11-04 20:36:21', 2, 0);

-- ----------------------------
-- Table structure for event_day_st
-- ----------------------------
DROP TABLE IF EXISTS `event_day_st`;
CREATE TABLE `event_day_st`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `day` date NULL DEFAULT NULL COMMENT '时间',
  `team_id` int UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `new_customer_cnt` int UNSIGNED NULL DEFAULT NULL COMMENT '留存',
  `first_diagnosis` int UNSIGNED NULL DEFAULT NULL COMMENT '初诊',
  `further_diagnosis` int UNSIGNED NULL DEFAULT NULL COMMENT '复诊',
  `deal` int UNSIGNED NULL DEFAULT NULL COMMENT '成交',
  `invitation` int UNSIGNED NULL DEFAULT NULL COMMENT '明日邀约',
  `rest` tinyint NULL DEFAULT NULL COMMENT ' 1上班 2休息',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `create_by` int UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_event_day_st_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_event_day_st_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of event_day_st
-- ----------------------------
INSERT INTO `event_day_st` VALUES (1, '2023-10-01', 1, 2, '/0/1/2/', 0, 0, 0, 0, 0, 1, '2023-10-03 18:32:55', '2023-10-03 18:32:55', 0, 0);
INSERT INTO `event_day_st` VALUES (2, '2023-10-27', 1, 2, '/0/1/2/', 1, 0, 0, 0, 0, 1, '2023-10-27 20:29:08', '2023-10-27 20:29:08', 2, 0);
INSERT INTO `event_day_st` VALUES (3, '2023-10-30', 1, 2, '/0/1/2/', 2, 1, 0, 2, 1, 1, '2023-10-30 21:10:36', '2023-10-30 21:10:36', 2, 0);
INSERT INTO `event_day_st` VALUES (4, '2023-10-30', 1, 3, '/0/1/2/', 4, 0, 0, 0, 0, 1, '2023-10-30 21:10:52', '2023-10-30 21:10:52', 2, 0);
INSERT INTO `event_day_st` VALUES (5, '2023-10-30', 1, 4, '/0/1/2/', 3, 0, 0, 0, 0, 1, '2023-10-30 21:11:06', '2023-10-30 21:11:06', 2, 0);
INSERT INTO `event_day_st` VALUES (6, '2023-10-30', 1, 5, '/0/1/2/', 3, 0, 0, 0, 0, 1, '2023-10-30 21:11:14', '2023-10-30 21:11:14', 2, 0);
INSERT INTO `event_day_st` VALUES (7, '2023-10-30', 1, 7, '/0/1/2/', 3, 0, 1, 1, 0, 1, '2023-10-30 21:11:34', '2023-10-30 21:11:34', 2, 0);
INSERT INTO `event_day_st` VALUES (8, '2023-10-30', 1, 6, '/0/1/2/', 1, 0, 0, 0, 0, 1, '2023-10-30 21:11:43', '2023-10-30 21:11:43', 2, 0);
INSERT INTO `event_day_st` VALUES (9, '2023-10-31', 1, 2, '/0/1/2/', 3, 2, 1, 1, 0, 1, '2023-10-31 21:14:06', '2023-10-31 21:14:06', 2, 0);
INSERT INTO `event_day_st` VALUES (10, '2023-10-31', 1, 3, '/0/1/2/', 1, 0, 0, 0, 0, 1, '2023-10-31 21:14:25', '2023-10-31 21:14:25', 2, 0);
INSERT INTO `event_day_st` VALUES (11, '2023-10-31', 1, 4, '/0/1/2/', 3, 0, 0, 0, 0, 1, '2023-10-31 21:14:41', '2023-10-31 21:14:41', 2, 0);
INSERT INTO `event_day_st` VALUES (12, '2023-10-31', 1, 5, '/0/1/2/', 2, 0, 0, 0, 0, 1, '2023-10-31 21:14:53', '2023-10-31 21:14:53', 2, 0);
INSERT INTO `event_day_st` VALUES (13, '2023-10-31', 1, 7, '/0/1/2/', 9, 1, 0, 0, 0, 1, '2023-10-31 21:15:05', '2023-10-31 21:15:05', 2, 0);
INSERT INTO `event_day_st` VALUES (14, '2023-10-31', 1, 6, '/0/1/2/', 0, 0, 0, 0, 0, 1, '2023-10-31 21:15:11', '2023-10-31 21:15:11', 2, 0);
INSERT INTO `event_day_st` VALUES (15, '2023-11-01', 1, 2, '/0/1/2/', 0, 1, 1, 2, 0, 1, '2023-11-01 20:50:15', '2023-11-01 20:50:15', 2, 0);
INSERT INTO `event_day_st` VALUES (16, '2023-11-01', 1, 5, '/0/1/2/', 2, 0, 0, 0, 0, 1, '2023-11-01 20:50:38', '2023-11-01 20:50:38', 2, 0);
INSERT INTO `event_day_st` VALUES (17, '2023-11-01', 1, 7, '/0/1/2/', 4, 0, 0, 0, 0, 1, '2023-11-01 20:50:59', '2023-11-01 20:50:59', 2, 0);
INSERT INTO `event_day_st` VALUES (18, '2023-11-01', 1, 6, '/0/1/2/', 3, 0, 0, 0, 0, 1, '2023-11-01 20:51:10', '2023-11-01 20:51:10', 2, 0);
INSERT INTO `event_day_st` VALUES (19, '2023-11-01', 1, 3, '/0/1/2/', 2, 0, 0, 0, 0, 1, '2023-11-01 20:51:31', '2023-11-01 20:51:31', 2, 0);
INSERT INTO `event_day_st` VALUES (20, '2023-11-01', 1, 4, '/0/1/2/', 2, 0, 0, 0, 0, 1, '2023-11-01 20:51:53', '2023-11-01 20:51:53', 2, 0);
INSERT INTO `event_day_st` VALUES (21, '2023-11-02', 1, 2, '/0/1/2/', 0, 0, 0, 0, 0, 2, '2023-11-02 21:25:40', '2023-11-02 21:25:40', 2, 0);
INSERT INTO `event_day_st` VALUES (22, '2023-11-02', 1, 3, '/0/1/2/', 4, 1, 0, 1, 0, 1, '2023-11-02 21:26:10', '2023-11-02 21:26:10', 2, 0);
INSERT INTO `event_day_st` VALUES (23, '2023-11-02', 1, 4, '/0/1/2/', 3, 0, 0, 0, 0, 1, '2023-11-02 21:26:39', '2023-11-02 21:26:39', 2, 0);
INSERT INTO `event_day_st` VALUES (24, '2023-11-02', 1, 5, '/0/1/2/', 4, 0, 0, 0, 0, 1, '2023-11-02 21:26:56', '2023-11-02 21:26:56', 2, 0);
INSERT INTO `event_day_st` VALUES (25, '2023-11-02', 1, 7, '/0/1/2/', 4, 0, 0, 0, 0, 1, '2023-11-02 21:27:19', '2023-11-02 21:27:19', 2, 0);
INSERT INTO `event_day_st` VALUES (26, '2023-11-02', 1, 6, '/0/1/2/', 0, 0, 0, 0, 0, 2, '2023-11-02 21:27:28', '2023-11-02 21:27:28', 2, 0);
INSERT INTO `event_day_st` VALUES (27, '2023-11-03', 1, 2, '/0/1/2/', 2, 0, 0, 0, 0, 1, '2023-11-03 20:54:39', '2023-11-03 20:54:39', 2, 0);
INSERT INTO `event_day_st` VALUES (28, '2023-11-03', 1, 3, '/0/1/2/', 0, 1, 0, 0, 0, 1, '2023-11-03 20:54:59', '2023-11-03 20:54:59', 2, 0);
INSERT INTO `event_day_st` VALUES (29, '2023-11-03', 1, 6, '/0/1/2/', 4, 0, 0, 0, 0, 1, '2023-11-03 20:55:13', '2023-11-03 20:55:13', 2, 0);
INSERT INTO `event_day_st` VALUES (30, '2023-11-03', 1, 5, '/0/1/2/', 2, 0, 0, 0, 0, 1, '2023-11-03 20:55:49', '2023-11-03 20:55:49', 2, 0);
INSERT INTO `event_day_st` VALUES (31, '2023-11-03', 1, 4, '/0/1/2/', 0, 0, 0, 0, 0, 2, '2023-11-03 20:56:05', '2023-11-03 20:56:05', 2, 0);
INSERT INTO `event_day_st` VALUES (32, '2023-11-03', 1, 7, '/0/1/2/', 3, 0, 0, 0, 0, 1, '2023-11-03 20:56:24', '2023-11-03 20:56:24', 2, 0);
INSERT INTO `event_day_st` VALUES (33, '2023-11-04', 1, 2, '/0/1/2/', 2, 1, 1, 1, 4, 1, '2023-11-04 20:37:07', '2023-11-04 20:37:07', 2, 0);
INSERT INTO `event_day_st` VALUES (34, '2023-11-04', 1, 3, '/0/1/2/', 3, 0, 0, 0, 0, 1, '2023-11-04 20:37:44', '2023-11-04 20:37:44', 2, 0);
INSERT INTO `event_day_st` VALUES (35, '2023-11-04', 1, 4, '/0/1/2/', 5, 0, 0, 0, 0, 1, '2023-11-04 20:38:03', '2023-11-04 20:38:03', 2, 0);
INSERT INTO `event_day_st` VALUES (36, '2023-11-04', 1, 5, '/0/1/2/', 4, 0, 0, 0, 0, 1, '2023-11-04 20:38:18', '2023-11-04 20:38:18', 2, 0);
INSERT INTO `event_day_st` VALUES (37, '2023-11-04', 1, 6, '/0/1/2/', 4, 0, 0, 0, 0, 1, '2023-11-04 20:38:32', '2023-11-04 20:38:32', 2, 0);
INSERT INTO `event_day_st` VALUES (38, '2023-11-04', 1, 7, '/0/1/2/', 5, 0, 0, 0, 0, 1, '2023-11-04 20:38:45', '2023-11-04 20:38:45', 2, 0);

-- ----------------------------
-- Table structure for summary_plan_day
-- ----------------------------
DROP TABLE IF EXISTS `summary_plan_day`;
CREATE TABLE `summary_plan_day`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `day` date NULL DEFAULT NULL COMMENT '天',
  `team_id` int UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `summary` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '今日总结',
  `plan` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '明日计划',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `create_by` int UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_summary_plan_day_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_summary_plan_day_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of summary_plan_day
-- ----------------------------
INSERT INTO `summary_plan_day` VALUES (6, '2023-11-01', 1, 2, '/0/1/2/', '上午一个初诊多颗一个复诊全口均成交，全程陪同，种完陪患者挂盐水，下午参加市场大会，晚上打邀约初诊，送全口叔叔回家', '今日身体不舒服，明日休息，组内有个初诊，督促组员抓留存', '2023-11-01 22:23:19', '2023-11-01 22:23:19', 0, 0);
INSERT INTO `summary_plan_day` VALUES (7, '2023-11-02', 1, 2, '/0/1/2/', '今日休息，收回昨天全口欠款', '明天上午有一个初诊促成交，另有几个复诊患者要接待，下午打邀约，督促组员留存数量和质量', '2023-11-02 23:17:51', '2023-11-02 23:17:51', 0, 0);
INSERT INTO `summary_plan_day` VALUES (8, '2023-11-03', 1, 2, '/0/1/2/', '上午接待三个复诊患者，下午邀约初诊患者，安抚组员的半口戴牙患者，因为咬模没咬好，下次还要来试戴，督促组员邀约', '明天上午三个初诊，促成交，另有6个复诊患者要接待，督促组员打邀约，自己打回访', '2023-11-03 21:05:35', '2023-11-03 21:05:35', 0, 0);
INSERT INTO `summary_plan_day` VALUES (9, '2023-11-04', 1, 2, '/0/1/2/', '上午接待3个复查患者，下午去一个小区看望几位种牙患者，开发另外的患者，然后回到医院带初诊和复诊，因为ITI的价格比别的医院贵500没成交，后面推了诺贝尔患者回去考虑，督促组员邀约明日初诊', '明日组内有初诊，协助促成交，另有几位复查患者要接待，督促组员月初发力月底才不累', '2023-11-04 20:44:54', '2023-11-04 21:08:01', 0, 0);

-- ----------------------------
-- Table structure for target_task
-- ----------------------------
DROP TABLE IF EXISTS `target_task`;
CREATE TABLE `target_task`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `day_type` tinyint UNSIGNED NULL DEFAULT NULL COMMENT '时间类型:月 30,周 7',
  `day` int UNSIGNED NULL DEFAULT NULL COMMENT '时间:202310',
  `team_id` int UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `task_type` int UNSIGNED NULL DEFAULT NULL COMMENT '任务类型 1正式 算人员数量',
  `new_customer_cnt` int UNSIGNED NULL DEFAULT NULL COMMENT '留存任务',
  `first_diagnosis` int UNSIGNED NULL DEFAULT NULL COMMENT '导诊任务',
  `deal` int UNSIGNED NULL DEFAULT NULL COMMENT '成交任务',
  `create_by` int UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_target_task_create_by`(`create_by` ASC) USING BTREE,
  INDEX `idx_target_task_update_by`(`update_by` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of target_task
-- ----------------------------
INSERT INTO `target_task` VALUES (1, 30, 202310, 1, 2, '/0/1/2/', NULL, 72, 16, 80000, NULL, 2, NULL, '2023-10-27 20:21:04');
INSERT INTO `target_task` VALUES (2, 30, 202310, 1, 3, '/0/1/2/', NULL, 72, 16, 80000, NULL, 2, NULL, '2023-10-27 20:22:14');
INSERT INTO `target_task` VALUES (3, 30, 202310, 1, 4, '/0/1/2/', NULL, 72, 16, 80000, NULL, 2, NULL, '2023-10-27 20:22:17');
INSERT INTO `target_task` VALUES (4, 30, 202310, 1, 5, '/0/1/2/', NULL, 72, 16, 80000, NULL, 2, NULL, '2023-10-27 20:22:20');
INSERT INTO `target_task` VALUES (5, 30, 202310, 1, 6, '/0/1/2/', NULL, 72, 16, 80000, NULL, 2, NULL, '2023-10-27 20:22:24');
INSERT INTO `target_task` VALUES (6, 30, 202310, 1, 7, '/0/1/2/', NULL, 72, 16, 80000, NULL, 2, NULL, '2023-10-27 20:22:27');
INSERT INTO `target_task` VALUES (9, 30, 202311, 1, 2, '/0/1/2/', 0, 72, 16, 80000, 2, 2, '2023-11-01 20:46:31', '2023-11-01 20:47:46');
INSERT INTO `target_task` VALUES (10, 30, 202311, 1, 3, '/0/1/2/', 0, 72, 16, 80000, 2, 0, '2023-11-01 20:46:57', '2023-11-01 20:46:57');
INSERT INTO `target_task` VALUES (11, 30, 202311, 1, 4, '/0/1/2/', 0, 72, 16, 80000, 2, 0, '2023-11-01 20:47:37', '2023-11-01 20:47:37');
INSERT INTO `target_task` VALUES (12, 30, 202311, 1, 5, '/0/1/2/', 0, 72, 16, 80000, 2, 0, '2023-11-01 20:48:08', '2023-11-01 20:48:08');
INSERT INTO `target_task` VALUES (13, 30, 202311, 1, 6, '/0/1/2/', 0, 72, 16, 80000, 2, 0, '2023-11-01 20:48:35', '2023-11-01 20:48:35');
INSERT INTO `target_task` VALUES (14, 30, 202311, 1, 7, '/0/1/2/', 0, 72, 16, 80000, 2, 0, '2023-11-01 20:49:23', '2023-11-01 20:49:23');

SET FOREIGN_KEY_CHECKS = 1;
