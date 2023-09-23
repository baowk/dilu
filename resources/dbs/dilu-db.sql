/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : localhost:3306
 Source Schema         : dilu-db

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 23/09/2023 10:04:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for email_log
-- ----------------------------
DROP TABLE IF EXISTS `email_log`;
CREATE TABLE `email_log`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱地址',
  `code` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '验证码',
  `type` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '类型',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `use_status` tinyint(0) NULL DEFAULT NULL COMMENT '使用状态',
  `created_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of email_log
-- ----------------------------

-- ----------------------------
-- Table structure for gen_columns
-- ----------------------------
DROP TABLE IF EXISTS `gen_columns`;
CREATE TABLE `gen_columns`  (
  `column_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `table_id` bigint(0) NULL DEFAULT NULL,
  `column_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `column_comment` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `column_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `go_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `go_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `json_field` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_pk` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_increment` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_required` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_insert` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_edit` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_list` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_query` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `query_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `html_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `dict_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `sort` bigint(0) NULL DEFAULT NULL,
  `list` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `pk` tinyint(1) NULL DEFAULT NULL,
  `required` tinyint(1) NULL DEFAULT NULL,
  `super_column` tinyint(1) NULL DEFAULT NULL,
  `usable_column` tinyint(1) NULL DEFAULT NULL,
  `increment` tinyint(1) NULL DEFAULT NULL,
  `insert` tinyint(1) NULL DEFAULT NULL,
  `edit` tinyint(1) NULL DEFAULT NULL,
  `query` tinyint(1) NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `fk_table_name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `fk_table_name_class` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `fk_table_name_package` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `fk_label_id` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL,
  `fk_label_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `create_by` mediumint(0) NULL DEFAULT NULL,
  `update_By` mediumint(0) NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`column_id`) USING BTREE,
  INDEX `idx_gen_columns_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_columns
-- ----------------------------
INSERT INTO `gen_columns` VALUES (11, 2, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-14 09:02:59.966', '2023-09-14 09:02:59.966', NULL);
INSERT INTO `gen_columns` VALUES (12, 2, 'dept_id', '部门id', 'int unsigned', 'int', 'DeptId', 'deptId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-14 09:02:59.971', '2023-09-14 09:02:59.971', NULL);
INSERT INTO `gen_columns` VALUES (13, 2, 'user_id', '上级部门', 'int unsigned', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-14 09:02:59.974', '2023-09-14 09:02:59.974', NULL);
INSERT INTO `gen_columns` VALUES (14, 2, 'post_tag', '职位标签 1主管 2副主管 3员工', 'tinyint unsigned', 'int', 'PostTag', 'postTag', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-14 09:02:59.978', '2023-09-14 09:02:59.978', NULL);
INSERT INTO `gen_columns` VALUES (15, 2, 'status', '状态 1正常 ', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-14 09:02:59.981', '2023-09-14 09:02:59.981', NULL);
INSERT INTO `gen_columns` VALUES (16, 2, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-14 09:02:59.984', '2023-09-14 09:02:59.984', NULL);
INSERT INTO `gen_columns` VALUES (17, 2, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-14 09:02:59.987', '2023-09-14 09:02:59.987', NULL);
INSERT INTO `gen_columns` VALUES (18, 2, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-14 09:02:59.989', '2023-09-14 09:02:59.989', NULL);
INSERT INTO `gen_columns` VALUES (19, 2, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-14 09:02:59.992', '2023-09-14 09:02:59.992', NULL);
INSERT INTO `gen_columns` VALUES (20, 2, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-14 09:02:59.997', '2023-09-14 09:02:59.997', NULL);
INSERT INTO `gen_columns` VALUES (21, 3, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.465', '2023-09-23 09:34:14.465', NULL);
INSERT INTO `gen_columns` VALUES (22, 3, 'no', '订单号', 'varchar(20)', 'string', 'No', 'no', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.470', '2023-09-23 09:34:14.470', NULL);
INSERT INTO `gen_columns` VALUES (23, 3, 'customer_id', '顾客', 'bigint', 'int', 'CustomerId', 'customerId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.485', '2023-09-23 09:34:14.485', NULL);
INSERT INTO `gen_columns` VALUES (24, 3, 'user_id', '用户id', 'bigint', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.492', '2023-09-23 09:34:14.492', NULL);
INSERT INTO `gen_columns` VALUES (25, 3, 'team_id', '团队id', 'bigint', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.496', '2023-09-23 09:34:14.496', NULL);
INSERT INTO `gen_columns` VALUES (26, 3, 'total', '金额', 'decimal(10,2)', 'string', 'Total', 'total', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.499', '2023-09-23 09:34:14.499', NULL);
INSERT INTO `gen_columns` VALUES (27, 3, 'real_total', '折后金额', 'decimal(10,2)', 'string', 'RealTotal', 'realTotal', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.501', '2023-09-23 09:34:14.501', NULL);
INSERT INTO `gen_columns` VALUES (28, 3, 'paid_total', '已支付金额', 'decimal(10,2)', 'string', 'PaidTotal', 'paidTotal', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.504', '2023-09-23 09:34:14.504', NULL);
INSERT INTO `gen_columns` VALUES (29, 3, 'link_id', '关联订单', 'bigint', 'int', 'LinkId', 'linkId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.509', '2023-09-23 09:34:14.509', NULL);
INSERT INTO `gen_columns` VALUES (30, 3, 'trade_at', '交易日期', 'datetime', 'time.Time', 'TradeAt', 'tradeAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.511', '2023-09-23 09:34:14.511', NULL);
INSERT INTO `gen_columns` VALUES (31, 3, 'trade_status', '交易类型 1 成交 2补尾款 3退款', 'tinyint', 'int', 'TradeStatus', 'tradeStatus', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.514', '2023-09-23 09:34:14.514', NULL);
INSERT INTO `gen_columns` VALUES (32, 3, 'dental_count', '颗数', 'tinyint', 'int', 'DentalCount', 'dentalCount', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.515', '2023-09-23 09:34:14.515', NULL);
INSERT INTO `gen_columns` VALUES (33, 3, 'brand', '品牌', 'tinyint', 'int', 'Brand', 'brand', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.517', '2023-09-23 09:34:14.517', NULL);
INSERT INTO `gen_columns` VALUES (34, 3, 'implanted_count', '已种颗数', 'tinyint', 'int', 'ImplantedCount', 'implantedCount', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.519', '2023-09-23 09:34:14.519', NULL);
INSERT INTO `gen_columns` VALUES (35, 3, 'implant', '是否已种', 'tinyint', 'int', 'Implant', 'implant', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.523', '2023-09-23 09:34:14.523', NULL);
INSERT INTO `gen_columns` VALUES (36, 3, 'implant_date', '植入日期', 'datetime', 'time.Time', 'ImplantDate', 'implantDate', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.527', '2023-09-23 09:34:14.527', NULL);
INSERT INTO `gen_columns` VALUES (37, 3, 'doctor', '医生', 'varchar(32)', 'string', 'Doctor', 'doctor', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.529', '2023-09-23 09:34:14.529', NULL);
INSERT INTO `gen_columns` VALUES (38, 3, 'pack', '1 普通 2 半口 3 全口', 'tinyint', 'int', 'Pack', 'pack', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.530', '2023-09-23 09:34:14.530', NULL);
INSERT INTO `gen_columns` VALUES (39, 3, 'payback_date', '预定回款日期', 'datetime', 'time.Time', 'PaybackDate', 'paybackDate', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 19, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.532', '2023-09-23 09:34:14.532', NULL);
INSERT INTO `gen_columns` VALUES (40, 3, 'created_at', '创建时间', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 20, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.534', '2023-09-23 09:34:14.534', NULL);
INSERT INTO `gen_columns` VALUES (41, 3, 'updated_at', '更新时间', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 21, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:14.536', '2023-09-23 09:34:14.536', NULL);
INSERT INTO `gen_columns` VALUES (42, 4, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:25.062', '2023-09-23 09:34:25.062', NULL);
INSERT INTO `gen_columns` VALUES (43, 4, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:25.065', '2023-09-23 09:34:25.065', NULL);
INSERT INTO `gen_columns` VALUES (44, 4, 'birthday', '生日', 'date', 'string', 'Birthday', 'birthday', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:25.070', '2023-09-23 09:34:25.070', NULL);
INSERT INTO `gen_columns` VALUES (45, 4, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:25.086', '2023-09-23 09:34:25.086', NULL);
INSERT INTO `gen_columns` VALUES (46, 4, 'wechat', '微信号', 'varchar(64)', 'string', 'Wechat', 'wechat', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:25.092', '2023-09-23 09:34:25.092', NULL);
INSERT INTO `gen_columns` VALUES (47, 4, 'gender', '性别', 'tinyint', 'int', 'Gender', 'gender', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:25.098', '2023-09-23 09:34:25.098', NULL);
INSERT INTO `gen_columns` VALUES (48, 4, 'address', '地址', 'varchar(255)', 'string', 'Address', 'address', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:25.100', '2023-09-23 09:34:25.100', NULL);
INSERT INTO `gen_columns` VALUES (49, 4, 'remark', '描述', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:25.102', '2023-09-23 09:34:25.102', NULL);
INSERT INTO `gen_columns` VALUES (50, 4, 'sales_id', '销售人员', 'bigint', 'int', 'SalesId', 'salesId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:25.106', '2023-09-23 09:34:25.106', NULL);
INSERT INTO `gen_columns` VALUES (51, 5, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.208', '2023-09-23 09:34:34.208', NULL);
INSERT INTO `gen_columns` VALUES (52, 5, 'day', '时间', 'datetime', 'time.Time', 'Day', 'day', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.211', '2023-09-23 09:34:34.211', NULL);
INSERT INTO `gen_columns` VALUES (53, 5, 'team_id', '团队id', 'bigint', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.215', '2023-09-23 09:34:34.215', NULL);
INSERT INTO `gen_columns` VALUES (54, 5, 'user_id', '用户id', 'bigint', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.230', '2023-09-23 09:34:34.230', NULL);
INSERT INTO `gen_columns` VALUES (55, 5, 'new_customer_cnt', '留存', 'bigint', 'int', 'NewCustomerCnt', 'newCustomerCnt', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.234', '2023-09-23 09:34:34.234', NULL);
INSERT INTO `gen_columns` VALUES (56, 5, 'first_diagnosis', '初诊', 'bigint', 'int', 'FirstDiagnosis', 'firstDiagnosis', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.241', '2023-09-23 09:34:34.241', NULL);
INSERT INTO `gen_columns` VALUES (57, 5, 'further_diagnosis', '复诊', 'bigint', 'int', 'FurtherDiagnosis', 'furtherDiagnosis', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.245', '2023-09-23 09:34:34.245', NULL);
INSERT INTO `gen_columns` VALUES (58, 5, 'deal', '成交', 'bigint', 'int', 'Deal', 'deal', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.248', '2023-09-23 09:34:34.248', NULL);
INSERT INTO `gen_columns` VALUES (59, 5, 'rest', '休息', 'tinyint', 'int', 'Rest', 'rest', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.250', '2023-09-23 09:34:34.250', NULL);
INSERT INTO `gen_columns` VALUES (60, 5, 'created_at', '创建时间', 'bigint', 'int', 'CreatedAt', 'createdAt', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.252', '2023-09-23 09:34:34.252', NULL);
INSERT INTO `gen_columns` VALUES (61, 5, 'updated_at', '更新时间', 'bigint', 'int', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:34.255', '2023-09-23 09:34:34.255', NULL);
INSERT INTO `gen_columns` VALUES (62, 6, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:41.593', '2023-09-23 09:34:41.593', NULL);
INSERT INTO `gen_columns` VALUES (63, 6, 'day', '天', 'bigint', 'int', 'Day', 'day', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:41.597', '2023-09-23 09:34:41.597', NULL);
INSERT INTO `gen_columns` VALUES (64, 6, 'team_id', '团队id', 'bigint', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:41.601', '2023-09-23 09:34:41.601', NULL);
INSERT INTO `gen_columns` VALUES (65, 6, 'user_id', '用户id', 'bigint', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:41.617', '2023-09-23 09:34:41.617', NULL);
INSERT INTO `gen_columns` VALUES (66, 6, 'summary', '今日总结', 'text', 'string', 'Summary', 'summary', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:41.622', '2023-09-23 09:34:41.622', NULL);
INSERT INTO `gen_columns` VALUES (67, 6, 'plan', '明日计划', 'text', 'string', 'Plan', 'plan', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:41.627', '2023-09-23 09:34:41.627', NULL);
INSERT INTO `gen_columns` VALUES (68, 6, 'created_at', '创建时间', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:41.631', '2023-09-23 09:34:41.631', NULL);
INSERT INTO `gen_columns` VALUES (69, 6, 'updated_at', '更新时间', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:41.633', '2023-09-23 09:34:41.633', NULL);
INSERT INTO `gen_columns` VALUES (70, 7, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:49.022', '2023-09-23 09:34:49.022', NULL);
INSERT INTO `gen_columns` VALUES (71, 7, 'month', '月', 'bigint', 'int', 'Month', 'month', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:49.025', '2023-09-23 09:34:49.025', NULL);
INSERT INTO `gen_columns` VALUES (72, 7, 'team_id', '团队id', 'bigint', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:49.030', '2023-09-23 09:34:49.030', NULL);
INSERT INTO `gen_columns` VALUES (73, 7, 'user_id', '用户id', 'bigint', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:49.034', '2023-09-23 09:34:49.034', NULL);
INSERT INTO `gen_columns` VALUES (74, 7, 'new_customer_cnt', '留存任务', 'bigint', 'int', 'NewCustomerCnt', 'newCustomerCnt', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:49.046', '2023-09-23 09:34:49.046', NULL);
INSERT INTO `gen_columns` VALUES (75, 7, 'first_diagnosis', '导诊任务', 'bigint', 'int', 'FirstDiagnosis', 'firstDiagnosis', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:49.051', '2023-09-23 09:34:49.051', NULL);
INSERT INTO `gen_columns` VALUES (76, 7, 'deal', '成交任务', 'bigint', 'int', 'Deal', 'deal', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:49.058', '2023-09-23 09:34:49.058', NULL);
INSERT INTO `gen_columns` VALUES (77, 8, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:57.497', '2023-09-23 09:34:57.497', NULL);
INSERT INTO `gen_columns` VALUES (78, 8, 'parent_id', '上级团队', 'bigint', 'int', 'ParentId', 'parentId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:57.501', '2023-09-23 09:34:57.501', NULL);
INSERT INTO `gen_columns` VALUES (79, 8, 'path', '团队路径', 'varchar(255)', 'string', 'Path', 'path', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:57.507', '2023-09-23 09:34:57.507', NULL);
INSERT INTO `gen_columns` VALUES (80, 8, 'name', '团队名', 'varchar(32)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:57.513', '2023-09-23 09:34:57.513', NULL);
INSERT INTO `gen_columns` VALUES (81, 8, 'owner', '团队拥有者', 'bigint', 'int', 'Owner', 'owner', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:57.518', '2023-09-23 09:34:57.518', NULL);
INSERT INTO `gen_columns` VALUES (82, 8, 'status', '状态', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:57.522', '2023-09-23 09:34:57.522', NULL);
INSERT INTO `gen_columns` VALUES (83, 8, 'created_at', '创建时间', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:57.525', '2023-09-23 09:34:57.525', NULL);
INSERT INTO `gen_columns` VALUES (84, 8, 'updated_at', '更新时间', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:34:57.528', '2023-09-23 09:34:57.528', NULL);
INSERT INTO `gen_columns` VALUES (85, 9, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:35:05.441', '2023-09-23 09:35:05.441', NULL);
INSERT INTO `gen_columns` VALUES (86, 9, 'user_id', '用户id', 'bigint', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:35:05.444', '2023-09-23 09:35:05.444', NULL);
INSERT INTO `gen_columns` VALUES (87, 9, 'team_id', '团队id', 'bigint', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:35:05.450', '2023-09-23 09:35:05.450', NULL);
INSERT INTO `gen_columns` VALUES (88, 9, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:35:05.455', '2023-09-23 09:35:05.455', NULL);
INSERT INTO `gen_columns` VALUES (89, 9, 'phone', '电话', 'varchar(11)', 'string', 'Phone', 'phone', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:35:05.460', '2023-09-23 09:35:05.460', NULL);
INSERT INTO `gen_columns` VALUES (90, 9, 'gender', '性别', 'tinyint', 'int', 'Gender', 'gender', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:35:05.465', '2023-09-23 09:35:05.465', NULL);
INSERT INTO `gen_columns` VALUES (91, 9, 'status', '状态', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:35:05.467', '2023-09-23 09:35:05.467', NULL);
INSERT INTO `gen_columns` VALUES (92, 9, 'role', '角色 1主管 2副主管 4普通', 'tinyint', 'int', 'Role', 'role', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:35:05.469', '2023-09-23 09:35:05.469', NULL);
INSERT INTO `gen_columns` VALUES (93, 9, 'created_at', '创建时间', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:35:05.473', '2023-09-23 09:35:05.473', NULL);
INSERT INTO `gen_columns` VALUES (94, 9, 'updated_at', '更新时间', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-23 09:35:05.476', '2023-09-23 09:35:05.476', NULL);

-- ----------------------------
-- Table structure for gen_tables
-- ----------------------------
DROP TABLE IF EXISTS `gen_tables`;
CREATE TABLE `gen_tables`  (
  `table_id` bigint(0) NOT NULL AUTO_INCREMENT,
  `db_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `table_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `table_comment` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `class_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tpl_category` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `package_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `module_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `module_front_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '前端文件名',
  `business_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `function_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `function_author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `pk_column` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `pk_go_field` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `pk_json_field` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `options` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tree_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tree_parent_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tree_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tree` tinyint(1) NULL DEFAULT 0,
  `crud` tinyint(1) NULL DEFAULT 1,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `is_data_scope` tinyint(0) NULL DEFAULT NULL,
  `is_actions` tinyint(0) NULL DEFAULT NULL,
  `is_auth` tinyint(0) NULL DEFAULT NULL,
  `is_logical_delete` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `logical_delete` tinyint(1) NULL DEFAULT NULL,
  `logical_delete_column` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`table_id`) USING BTREE,
  INDEX `idx_gen_tables_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_gen_tables_create_by`(`create_by`) USING BTREE,
  INDEX `idx_gen_tables_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_tables
-- ----------------------------
INSERT INTO `gen_tables` VALUES (2, 'dilu-db', 'sys_user_dept', 'SysUserDept', 'SysUserDept', 'crud', 'sys', 'sys-user-dept', '', 'sysUserDept', 'SysUserDept', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-14 09:02:59.961', '2023-09-14 09:02:59.961', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (3, 'dental-db', 'bill', 'Bill', 'Bill', 'crud', 'dental', 'bill', '', 'bill', 'Bill', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:14.460', '2023-09-23 09:34:14.460', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (4, 'dental-db', 'customer', 'Customer', 'Customer', 'crud', 'dental', 'customer', '', 'customer', 'Customer', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:25.057', '2023-09-23 09:34:25.057', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (5, 'dental-db', 'event_day_st', 'EventDaySt', 'EventDaySt', 'crud', 'dental', 'event-day-st', '', 'eventDaySt', 'EventDaySt', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:34.203', '2023-09-23 09:34:34.203', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (6, 'dental-db', 'summary_plan_day', 'SummaryPlanDay', 'SummaryPlanDay', 'crud', 'dental', 'summary-plan-day', '', 'summaryPlanDay', 'SummaryPlanDay', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:41.589', '2023-09-23 09:34:41.589', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (7, 'dental-db', 'target_task', 'TargetTask', 'TargetTask', 'crud', 'dental', 'target-task', '', 'targetTask', 'TargetTask', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:49.019', '2023-09-23 09:34:49.019', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (8, 'dental-db', 'team', 'Team', 'Team', 'crud', 'dental', 'team', '', 'team', 'Team', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:57.492', '2023-09-23 09:34:57.492', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (9, 'dental-db', 'team_member', 'TeamMember', 'TeamMember', 'crud', 'dental', 'team-member', '', 'teamMember', 'TeamMember', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:35:05.438', '2023-09-23 09:35:05.438', NULL, 0, 0);

-- ----------------------------
-- Table structure for sms_log
-- ----------------------------
DROP TABLE IF EXISTS `sms_log`;
CREATE TABLE `sms_log`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `code` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '验证码',
  `type` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '类型',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `use_status` tinyint(0) NULL DEFAULT NULL COMMENT '使用状态',
  `created_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sms_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标题',
  `method` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求类型',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求地址',
  `type` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '接口类型',
  `perm_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '权限类型（n：无需任何认证 t:须token p：须权限）',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态 3 DEF 2 OK 1 del',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `action` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求类型',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_api
-- ----------------------------
INSERT INTO `sys_api` VALUES (1, '分页获取SysUserDept', 'POST', '/api/v1/sys/sys-user-dept/page', '', 't', 3, 0, '2023-09-21 09:56:12.005', NULL);
INSERT INTO `sys_api` VALUES (2, '根据id获取SysUserDept', 'POST', '/api/v1/sys/sys-user-dept/get', '', 't', 3, 0, '2023-09-21 09:56:12.013', NULL);
INSERT INTO `sys_api` VALUES (3, '创建SysUserDept', 'POST', '/api/v1/sys/sys-user-dept/create', '', 't', 3, 0, '2023-09-21 09:56:12.025', NULL);
INSERT INTO `sys_api` VALUES (4, '修改SysUserDept', 'POST', '/api/v1/sys/sys-user-dept/update', '', 't', 3, 0, '2023-09-21 09:56:12.031', NULL);
INSERT INTO `sys_api` VALUES (5, '删除SysUserDept', 'POST', '/api/v1/sys/sys-user-dept/del', '', 't', 3, 0, '2023-09-21 09:56:12.042', NULL);

-- ----------------------------
-- Table structure for sys_cfg
-- ----------------------------
DROP TABLE IF EXISTS `sys_cfg`;
CREATE TABLE `sys_cfg`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '名字',
  `key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'key',
  `value` json NULL COMMENT 'Value',
  `type` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'Type',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'Remark',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT 'Status',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_cfg
-- ----------------------------

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept`  (
  `dept_id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT 'ParentId',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'DeptPath',
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'DeptName',
  `sort` tinyint(0) NULL DEFAULT NULL COMMENT 'Sort',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT 'Status',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `leader` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`dept_id`) USING BTREE,
  INDEX `idx_sys_dept_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_dept_create_by`(`create_by`) USING BTREE,
  INDEX `idx_sys_dept_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------

-- ----------------------------
-- Table structure for sys_email
-- ----------------------------
DROP TABLE IF EXISTS `sys_email`;
CREATE TABLE `sys_email`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱地址',
  `code` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '验证码',
  `type` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '类型',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `use_status` tinyint(0) NULL DEFAULT NULL COMMENT '使用状态',
  `created_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_email
-- ----------------------------

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `job_id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `job_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'JobName',
  `job_group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'JobGroup',
  `job_type` tinyint(0) NULL DEFAULT NULL COMMENT 'JobType',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'CronExpression',
  `invoke_target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'InvokeTarget',
  `args` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'Args',
  `misfire_policy` bigint(0) NULL DEFAULT NULL COMMENT 'MisfirePolicy',
  `concurrent` tinyint(0) NULL DEFAULT NULL COMMENT 'Concurrent',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT 'Status',
  `entry_id` smallint(0) NULL DEFAULT NULL COMMENT 'EntryId',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`job_id`) USING BTREE,
  INDEX `idx_sys_job_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_job_create_by`(`create_by`) USING BTREE,
  INDEX `idx_sys_job_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `menu_id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `menu_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单名',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '显示名称',
  `icon` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '图标',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '路径',
  `paths` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '路径ids/分割',
  `menu_type` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '菜单类型 M 菜单 C 分类 F 方法 O 外链',
  `permission` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '权限',
  `parent_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '菜单父id',
  `no_cache` tinyint(1) NULL DEFAULT NULL COMMENT '是否缓存',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '前端组件路径',
  `sort` tinyint(0) NULL DEFAULT NULL COMMENT '排序倒叙',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '是否隐藏',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `action` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求方式',
  `breadcrumb` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '是否面包屑',
  `visible` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '是否显示',
  `is_frame` varchar(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '0' COMMENT '是否frame',
  PRIMARY KEY (`menu_id`) USING BTREE,
  INDEX `idx_sys_menu_create_by`(`create_by`) USING BTREE,
  INDEX `idx_sys_menu_update_by`(`update_by`) USING BTREE,
  INDEX `idx_sys_menu_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu_api_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_api_rule`;
CREATE TABLE `sys_menu_api_rule`  (
  `sys_menu_menu_id` int(0) UNSIGNED NOT NULL COMMENT '主键',
  `sys_api_id` int(0) UNSIGNED NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`sys_menu_menu_id`, `sys_api_id`) USING BTREE,
  INDEX `fk_sys_menu_api_rule_sys_api`(`sys_api_id`) USING BTREE,
  CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_menu_id`) REFERENCES `sys_menu` (`menu_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu_api_rule
-- ----------------------------

-- ----------------------------
-- Table structure for sys_opera_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_opera_log`;
CREATE TABLE `sys_opera_log`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作模块',
  `business_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作类型',
  `business_types` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'BusinessTypes',
  `method` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '函数',
  `request_method` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求方式 GET POST PUT DELETE',
  `operator_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作类型',
  `oper_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作者',
  `dept_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门名称',
  `oper_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '访问地址',
  `oper_ip` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '客户端ip',
  `oper_location` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '访问位置',
  `oper_param` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '请求参数',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '操作状态 1:成功 2:失败',
  `oper_time` datetime(3) NULL DEFAULT NULL COMMENT '操作时间',
  `json_result` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '返回数据',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `latency_time` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '耗时',
  `user_agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'ua',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_opera_log_create_by`(`create_by`) USING BTREE,
  INDEX `idx_sys_opera_log_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_opera_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `role_id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '角色名称',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '角色代码',
  `role_sort` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '排序',
  `flag` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT 'flag',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `admin` tinyint(1) NULL DEFAULT NULL COMMENT '管理员',
  `data_scope` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '数据权限',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`role_id`) USING BTREE,
  INDEX `idx_sys_role_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_role_create_by`(`create_by`) USING BTREE,
  INDEX `idx_sys_role_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 'super_admin', 1, 'admin', 1, '1', '超级管理员', 1, NULL, NULL, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `role_id` int(0) UNSIGNED NOT NULL COMMENT '主键',
  `menu_id` int(0) UNSIGNED NOT NULL COMMENT '主键',
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE,
  INDEX `fk_sys_role_menu_sys_menu`(`menu_id`) USING BTREE,
  CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`menu_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`role_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------

-- ----------------------------
-- Table structure for sys_sms
-- ----------------------------
DROP TABLE IF EXISTS `sys_sms`;
CREATE TABLE `sys_sms`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `code` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '验证码',
  `type` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '类型',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `use_status` tinyint(0) NULL DEFAULT NULL COMMENT '使用状态',
  `created_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_sms
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `user_id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户名',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '密码',
  `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '昵称',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '姓名',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '头像',
  `bio` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '签名',
  `birthday` date NULL DEFAULT NULL COMMENT '生日 格式 yyyy-MM-dd',
  `gender` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '2' COMMENT '性别 1男 2女 3未知',
  `role_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '角色ID',
  `post` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '岗位',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态 1冻结 2正常 3默认',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `nick_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '昵称',
  `dept_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '部门',
  PRIMARY KEY (`user_id`) USING BTREE,
  INDEX `idx_sys_user_create_by`(`create_by`) USING BTREE,
  INDEX `idx_sys_user_update_by`(`update_by`) USING BTREE,
  INDEX `idx_sys_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (2, 'dilu', '13800138000', '', '$2a$10$bMP2OrE4mG0noWl125X7peqUlz5VP2AfHhRemmKHvyNgZABd0H8Da', 'dilu', '', '', NULL, NULL, '2', 1, '', '', 0, 0, 0, '2023-09-19 11:49:13.139', '2023-09-19 11:49:13.139', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for sys_user_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_dept`;
CREATE TABLE `sys_user_dept`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dept_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '部门id',
  `user_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '上级部门',
  `post_tag` tinyint(0) UNSIGNED NULL DEFAULT NULL COMMENT '职位标签 1主管 2副主管 3员工',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态 1正常 ',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_dept_user`(`dept_id`, `user_id`) USING BTREE,
  INDEX `idx_sys_user_dept_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_dept
-- ----------------------------

-- ----------------------------
-- Table structure for third_login
-- ----------------------------
DROP TABLE IF EXISTS `third_login`;
CREATE TABLE `third_login`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `user_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `platform` tinyint(0) UNSIGNED NULL DEFAULT NULL COMMENT '平台 1 微信 2 钉钉',
  `open_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '第三方open_id',
  `union_id` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '第三方union_id',
  `third_data` json NULL COMMENT '第三方返回数据',
  `created_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of third_login
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
