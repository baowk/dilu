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

 Date: 26/09/2023 11:08:42
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
) ENGINE = InnoDB AUTO_INCREMENT = 163 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

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
INSERT INTO `gen_columns` VALUES (95, 10, 'user_id', '主键', 'int unsigned', 'int', 'UserId', 'userId', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.558', '2023-09-25 14:14:24.558', NULL);
INSERT INTO `gen_columns` VALUES (96, 10, 'username', '用户名', 'varchar(32)', 'string', 'Username', 'username', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.566', '2023-09-25 14:14:24.566', NULL);
INSERT INTO `gen_columns` VALUES (97, 10, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.572', '2023-09-25 14:14:24.572', NULL);
INSERT INTO `gen_columns` VALUES (98, 10, 'email', '邮箱', 'varchar(128)', 'string', 'Email', 'email', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.576', '2023-09-25 14:14:24.576', NULL);
INSERT INTO `gen_columns` VALUES (99, 10, 'password', '密码', 'varchar(128)', 'string', 'Password', 'password', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.580', '2023-09-25 14:14:24.580', NULL);
INSERT INTO `gen_columns` VALUES (100, 10, 'nickname', '昵称', 'varchar(128)', 'string', 'Nickname', 'nickname', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.584', '2023-09-25 14:14:24.584', NULL);
INSERT INTO `gen_columns` VALUES (101, 10, 'name', '姓名', 'varchar(64)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.588', '2023-09-25 14:14:24.588', NULL);
INSERT INTO `gen_columns` VALUES (102, 10, 'avatar', '头像', 'varchar(255)', 'string', 'Avatar', 'avatar', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.593', '2023-09-25 14:14:24.593', NULL);
INSERT INTO `gen_columns` VALUES (103, 10, 'bio', '签名', 'varchar(255)', 'string', 'Bio', 'bio', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.599', '2023-09-25 14:14:24.599', NULL);
INSERT INTO `gen_columns` VALUES (104, 10, 'birthday', '生日 格式 yyyy-MM-dd', 'date', 'string', 'Birthday', 'birthday', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.603', '2023-09-25 14:14:24.603', NULL);
INSERT INTO `gen_columns` VALUES (105, 10, 'gender', '性别 1男 2女 3未知', 'char(1)', 'string', 'Gender', 'gender', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.606', '2023-09-25 14:14:24.606', NULL);
INSERT INTO `gen_columns` VALUES (106, 10, 'role_id', '角色ID', 'int unsigned', 'int', 'RoleId', 'roleId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.608', '2023-09-25 14:14:24.608', NULL);
INSERT INTO `gen_columns` VALUES (107, 10, 'post', '岗位', 'varchar(32)', 'string', 'Post', 'post', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.611', '2023-09-25 14:14:24.611', NULL);
INSERT INTO `gen_columns` VALUES (108, 10, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.617', '2023-09-25 14:14:24.617', NULL);
INSERT INTO `gen_columns` VALUES (109, 10, 'status', '状态 1冻结 2正常 3默认', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.621', '2023-09-25 14:14:24.621', NULL);
INSERT INTO `gen_columns` VALUES (110, 10, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.625', '2023-09-25 14:14:24.625', NULL);
INSERT INTO `gen_columns` VALUES (111, 10, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.627', '2023-09-25 14:14:24.627', NULL);
INSERT INTO `gen_columns` VALUES (112, 10, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.634', '2023-09-25 14:14:24.634', NULL);
INSERT INTO `gen_columns` VALUES (113, 10, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 19, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.637', '2023-09-25 14:14:24.637', NULL);
INSERT INTO `gen_columns` VALUES (114, 10, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 20, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.639', '2023-09-25 14:14:24.639', NULL);
INSERT INTO `gen_columns` VALUES (115, 10, 'nick_name', '昵称', 'varchar(128)', 'string', 'NickName', 'nickName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 21, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.642', '2023-09-25 14:14:24.642', NULL);
INSERT INTO `gen_columns` VALUES (116, 10, 'dept_id', '部门', 'int unsigned', 'int', 'DeptId', 'deptId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 22, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 14:14:24.644', '2023-09-25 14:14:24.644', NULL);
INSERT INTO `gen_columns` VALUES (117, 11, 'menu_id', '主键', 'int unsigned', 'int', 'MenuId', 'menuId', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.672', '2023-09-25 17:36:31.672', NULL);
INSERT INTO `gen_columns` VALUES (118, 11, 'menu_name', '菜单名', 'varchar(128)', 'string', 'MenuName', 'menuName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.675', '2023-09-25 17:36:31.675', NULL);
INSERT INTO `gen_columns` VALUES (119, 11, 'title', '显示名称', 'varchar(128)', 'string', 'Title', 'title', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.679', '2023-09-25 17:36:31.679', NULL);
INSERT INTO `gen_columns` VALUES (120, 11, 'icon', '图标', 'varchar(128)', 'string', 'Icon', 'icon', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.686', '2023-09-25 17:36:31.686', NULL);
INSERT INTO `gen_columns` VALUES (121, 11, 'path', '路径', 'varchar(128)', 'string', 'Path', 'path', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.688', '2023-09-25 17:36:31.688', NULL);
INSERT INTO `gen_columns` VALUES (122, 11, 'paths', '路径ids/分割', 'varchar(128)', 'string', 'Paths', 'paths', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.691', '2023-09-25 17:36:31.691', NULL);
INSERT INTO `gen_columns` VALUES (123, 11, 'menu_type', '菜单类型 1 分类 2菜单 3方法按钮', 'tinyint', 'int', 'MenuType', 'menuType', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.692', '2023-09-25 17:36:31.692', NULL);
INSERT INTO `gen_columns` VALUES (124, 11, 'permission', '权限', 'varchar(255)', 'string', 'Permission', 'permission', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.695', '2023-09-25 17:36:31.695', NULL);
INSERT INTO `gen_columns` VALUES (125, 11, 'parent_id', '菜单父id', 'int unsigned', 'int', 'ParentId', 'parentId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.699', '2023-09-25 17:36:31.699', NULL);
INSERT INTO `gen_columns` VALUES (126, 11, 'no_cache', '是否缓存', 'tinyint(1)', 'int', 'NoCache', 'noCache', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.701', '2023-09-25 17:36:31.701', NULL);
INSERT INTO `gen_columns` VALUES (127, 11, 'component', '前端组件路径', 'varchar(255)', 'string', 'Component', 'component', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.704', '2023-09-25 17:36:31.704', NULL);
INSERT INTO `gen_columns` VALUES (128, 11, 'sort', '排序倒叙', 'tinyint', 'int', 'Sort', 'sort', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.707', '2023-09-25 17:36:31.707', NULL);
INSERT INTO `gen_columns` VALUES (129, 11, 'hidden', '是否隐藏', 'tinyint(1)', 'int', 'Hidden', 'hidden', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.709', '2023-09-25 17:36:31.709', NULL);
INSERT INTO `gen_columns` VALUES (130, 11, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.713', '2023-09-25 17:36:31.713', NULL);
INSERT INTO `gen_columns` VALUES (131, 11, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.718', '2023-09-25 17:36:31.718', NULL);
INSERT INTO `gen_columns` VALUES (132, 11, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.721', '2023-09-25 17:36:31.721', NULL);
INSERT INTO `gen_columns` VALUES (133, 11, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.723', '2023-09-25 17:36:31.723', NULL);
INSERT INTO `gen_columns` VALUES (134, 11, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-25 17:36:31.725', '2023-09-25 17:36:31.725', NULL);
INSERT INTO `gen_columns` VALUES (135, 12, 'dept_id', '主键', 'int unsigned', 'int', 'DeptId', 'deptId', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.176', '2023-09-26 09:21:19.176', NULL);
INSERT INTO `gen_columns` VALUES (136, 12, 'parent_id', 'ParentId', 'int unsigned', 'int', 'ParentId', 'parentId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.181', '2023-09-26 09:21:19.181', NULL);
INSERT INTO `gen_columns` VALUES (137, 12, 'dept_path', 'DeptPath', 'varchar(255)', 'string', 'DeptPath', 'deptPath', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.184', '2023-09-26 09:21:19.184', NULL);
INSERT INTO `gen_columns` VALUES (138, 12, 'dept_name', 'DeptName', 'varchar(128)', 'string', 'DeptName', 'deptName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.189', '2023-09-26 09:21:19.189', NULL);
INSERT INTO `gen_columns` VALUES (139, 12, 'sort', 'Sort', 'tinyint', 'int', 'Sort', 'sort', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.195', '2023-09-26 09:21:19.195', NULL);
INSERT INTO `gen_columns` VALUES (140, 12, 'status', 'Status', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.202', '2023-09-26 09:21:19.202', NULL);
INSERT INTO `gen_columns` VALUES (141, 12, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.208', '2023-09-26 09:21:19.208', NULL);
INSERT INTO `gen_columns` VALUES (142, 12, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.212', '2023-09-26 09:21:19.212', NULL);
INSERT INTO `gen_columns` VALUES (143, 12, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.214', '2023-09-26 09:21:19.214', NULL);
INSERT INTO `gen_columns` VALUES (144, 12, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.217', '2023-09-26 09:21:19.217', NULL);
INSERT INTO `gen_columns` VALUES (145, 12, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.221', '2023-09-26 09:21:19.221', NULL);
INSERT INTO `gen_columns` VALUES (146, 12, 'leader', '', 'varchar(128)', 'string', 'Leader', 'leader', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.226', '2023-09-26 09:21:19.226', NULL);
INSERT INTO `gen_columns` VALUES (147, 12, 'phone', '', 'varchar(11)', 'string', 'Phone', 'phone', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.229', '2023-09-26 09:21:19.229', NULL);
INSERT INTO `gen_columns` VALUES (148, 12, 'email', '', 'varchar(64)', 'string', 'Email', 'email', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:19.239', '2023-09-26 09:21:19.239', NULL);
INSERT INTO `gen_columns` VALUES (149, 13, 'role_id', '主键', 'int unsigned', 'int', 'RoleId', 'roleId', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.172', '2023-09-26 09:21:31.172', NULL);
INSERT INTO `gen_columns` VALUES (150, 13, 'role_name', '角色名称', 'varchar(128)', 'string', 'RoleName', 'roleName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.175', '2023-09-26 09:21:31.175', NULL);
INSERT INTO `gen_columns` VALUES (151, 13, 'status', '状态', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.180', '2023-09-26 09:21:31.180', NULL);
INSERT INTO `gen_columns` VALUES (152, 13, 'role_key', '角色代码', 'varchar(128)', 'string', 'RoleKey', 'roleKey', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.196', '2023-09-26 09:21:31.196', NULL);
INSERT INTO `gen_columns` VALUES (153, 13, 'role_sort', '排序', 'int unsigned', 'int', 'RoleSort', 'roleSort', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.201', '2023-09-26 09:21:31.201', NULL);
INSERT INTO `gen_columns` VALUES (154, 13, 'flag', 'flag', 'varchar(128)', 'string', 'Flag', 'flag', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.205', '2023-09-26 09:21:31.205', NULL);
INSERT INTO `gen_columns` VALUES (155, 13, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.208', '2023-09-26 09:21:31.208', NULL);
INSERT INTO `gen_columns` VALUES (156, 13, 'admin', '管理员', 'tinyint(1)', 'int', 'Admin', 'admin', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.213', '2023-09-26 09:21:31.213', NULL);
INSERT INTO `gen_columns` VALUES (157, 13, 'data_scope', '数据权限', 'varchar(128)', 'string', 'DataScope', 'dataScope', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.216', '2023-09-26 09:21:31.216', NULL);
INSERT INTO `gen_columns` VALUES (158, 13, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.219', '2023-09-26 09:21:31.219', NULL);
INSERT INTO `gen_columns` VALUES (159, 13, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.223', '2023-09-26 09:21:31.223', NULL);
INSERT INTO `gen_columns` VALUES (160, 13, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.227', '2023-09-26 09:21:31.227', NULL);
INSERT INTO `gen_columns` VALUES (161, 13, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.230', '2023-09-26 09:21:31.230', NULL);
INSERT INTO `gen_columns` VALUES (162, 13, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 09:21:31.232', '2023-09-26 09:21:31.232', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_tables
-- ----------------------------
INSERT INTO `gen_tables` VALUES (3, 'dental-db', 'bill', '账单', 'Bill', 'crud', 'dental', 'bill', '', 'bill', 'Bill', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:14.460', '2023-09-23 09:34:14.460', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (4, 'dental-db', 'customer', '客户', 'Customer', 'crud', 'dental', 'customer', '', 'customer', 'Customer', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:25.057', '2023-09-23 09:34:25.057', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (5, 'dental-db', 'event_day_st', '日统计', 'EventDaySt', 'crud', 'dental', 'event-day-st', '', 'eventDaySt', 'EventDaySt', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:34.203', '2023-09-23 09:34:34.203', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (6, 'dental-db', 'summary_plan_day', '日总结与计划', 'SummaryPlanDay', 'crud', 'dental', 'summary-plan-day', '', 'summaryPlanDay', 'SummaryPlanDay', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:41.589', '2023-09-23 09:34:41.589', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (7, 'dental-db', 'target_task', '任务目标', 'TargetTask', 'crud', 'dental', 'target-task', '', 'targetTask', 'TargetTask', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:49.019', '2023-09-23 09:34:49.019', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (8, 'dental-db', 'team', 'Team', 'Team', 'crud', 'dental', 'team', '', 'team', 'Team', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:34:57.492', '2023-09-23 09:34:57.492', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (9, 'dental-db', 'team_member', 'TeamMember', 'TeamMember', 'crud', 'dental', 'team-member', '', 'teamMember', 'TeamMember', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-23 09:35:05.438', '2023-09-23 09:35:05.438', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (10, 'dilu-db', 'sys_user', '用户', 'SysUser', 'crud', 'sys', 'sys-user', '', 'sysUser', 'SysUser', 'baowk', 'user_id', 'UserId', 'userId', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-25 14:14:24.554', '2023-09-25 14:14:24.554', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (11, 'dilu-db', 'sys_menu', '菜单', 'SysMenu', 'crud', 'sys', 'sys-menu', '', 'sysMenu', 'SysMenu', 'baowk', 'menu_id', 'MenuId', 'menuId', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-25 17:36:31.664', '2023-09-25 17:36:31.664', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (12, 'dilu-db', 'sys_dept', '部门', 'SysDept', 'crud', 'sys', 'sys-dept', '', 'sysDept', 'SysDept', 'baowk', 'dept_id', 'DeptId', 'deptId', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-26 09:21:19.168', '2023-09-26 09:21:19.168', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (13, 'dilu-db', 'sys_role', '角色', 'SysRole', 'crud', 'sys', 'sys-role', '', 'sysRole', 'SysRole', 'baowk', 'role_id', 'RoleId', 'roleId', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-26 09:21:31.168', '2023-09-26 09:21:31.168', NULL, 0, 0);

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
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_method_path`(`method`, `path`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 75 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_api
-- ----------------------------
INSERT INTO `sys_api` VALUES (55, '分页获取SysUser', 'POST', '/api/v1/sys/sys-user/page', '', 't', 3, 0, '2023-09-25 22:01:45.040');
INSERT INTO `sys_api` VALUES (56, '根据id获取SysUser', 'POST', '/api/v1/sys/sys-user/get', '', 't', 3, 0, '2023-09-25 22:01:45.062');
INSERT INTO `sys_api` VALUES (57, '创建SysUser', 'POST', '/api/v1/sys/sys-user/create', '', 't', 3, 0, '2023-09-25 22:01:45.083');
INSERT INTO `sys_api` VALUES (58, '修改SysUser', 'POST', '/api/v1/sys/sys-user/update', '', 't', 3, 0, '2023-09-25 22:01:45.106');
INSERT INTO `sys_api` VALUES (59, '删除SysUser', 'POST', '/api/v1/sys/sys-user/del', '', 't', 3, 0, '2023-09-25 22:01:45.161');
INSERT INTO `sys_api` VALUES (60, '分页获取SysMenu', 'POST', '/api/v1/sys/sys-menu/page', '', 't', 3, 0, '2023-09-25 22:02:23.409');
INSERT INTO `sys_api` VALUES (61, '根据id获取SysMenu', 'POST', '/api/v1/sys/sys-menu/get', '', 't', 3, 0, '2023-09-25 22:02:23.435');
INSERT INTO `sys_api` VALUES (62, '创建SysMenu', 'POST', '/api/v1/sys/sys-menu/create', '', 't', 3, 0, '2023-09-25 22:02:23.454');
INSERT INTO `sys_api` VALUES (63, '修改SysMenu', 'POST', '/api/v1/sys/sys-menu/update', '', 't', 3, 0, '2023-09-25 22:02:23.469');
INSERT INTO `sys_api` VALUES (64, '删除SysMenu', 'POST', '/api/v1/sys/sys-menu/del', '', 't', 3, 0, '2023-09-25 22:02:23.482');
INSERT INTO `sys_api` VALUES (65, '分页获取部门', 'POST', '/api/v1/sys/sys-dept/page', '', 't', 3, 0, '2023-09-26 09:28:25.235');
INSERT INTO `sys_api` VALUES (66, '根据id获取部门', 'POST', '/api/v1/sys/sys-dept/get', '', 't', 3, 0, '2023-09-26 09:28:25.269');
INSERT INTO `sys_api` VALUES (67, '创建部门', 'POST', '/api/v1/sys/sys-dept/create', '', 't', 3, 0, '2023-09-26 09:28:25.282');
INSERT INTO `sys_api` VALUES (68, '修改部门', 'POST', '/api/v1/sys/sys-dept/update', '', 't', 3, 0, '2023-09-26 09:28:25.294');
INSERT INTO `sys_api` VALUES (69, '删除部门', 'POST', '/api/v1/sys/sys-dept/del', '', 't', 3, 0, '2023-09-26 09:28:25.302');
INSERT INTO `sys_api` VALUES (70, '分页获取角色', 'POST', '/api/v1/sys/sys-role/page', '', 't', 3, 0, '2023-09-26 09:28:32.326');
INSERT INTO `sys_api` VALUES (71, '根据id获取角色', 'POST', '/api/v1/sys/sys-role/get', '', 't', 3, 0, '2023-09-26 09:28:32.345');
INSERT INTO `sys_api` VALUES (72, '创建角色', 'POST', '/api/v1/sys/sys-role/create', '', 't', 3, 0, '2023-09-26 09:28:32.371');
INSERT INTO `sys_api` VALUES (73, '修改角色', 'POST', '/api/v1/sys/sys-role/update', '', 't', 3, 0, '2023-09-26 09:28:32.382');
INSERT INTO `sys_api` VALUES (74, '删除角色', 'POST', '/api/v1/sys/sys-role/del', '', 't', 3, 0, '2023-09-26 09:28:32.393');

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
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '父id',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门名',
  `type` tinyint(0) NULL DEFAULT NULL COMMENT '类型',
  `principal` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门领导',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `sort` tinyint(0) NULL DEFAULT NULL COMMENT '排序',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dept_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

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
  `menu_type` tinyint(0) NULL DEFAULT NULL COMMENT '菜单类型 1 分类 2菜单 3方法按钮',
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
  PRIMARY KEY (`menu_id`) USING BTREE,
  INDEX `idx_sys_menu_create_by`(`create_by`) USING BTREE,
  INDEX `idx_sys_menu_update_by`(`update_by`) USING BTREE,
  INDEX `idx_sys_menu_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 83 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (72, '', '系统管理', 'pass', '/sys', '/0/72', 1, '', 0, 0, 'Layout', 0, 0, 1, 0, '2023-09-25 22:01:45.029', '2023-09-25 22:01:45.032', NULL);
INSERT INTO `sys_menu` VALUES (73, 'SysUserManage', '用户管理', 'pass', '/sys/user', '/0/72/73', 2, 'sys:sysUser:list', 72, 0, '/sys/user/index', 0, 0, 1, 1, '2023-09-25 22:01:45.046', '2023-09-25 22:01:45.054', NULL);
INSERT INTO `sys_menu` VALUES (74, '', '用户详情', '', 'sys_user', '/0/72/73/74', 3, 'sys:sysUser:query', 73, 0, '', 0, 0, 1, 1, '2023-09-25 22:01:45.066', '2023-09-25 22:01:45.075', NULL);
INSERT INTO `sys_menu` VALUES (75, '', '创建用户', '', 'sys_user', '/0/72/73/75', 3, 'sys:sysUser:add', 73, 0, '', 0, 0, 1, 1, '2023-09-25 22:01:45.089', '2023-09-25 22:01:45.096', NULL);
INSERT INTO `sys_menu` VALUES (76, '', '修改修改', '', 'sys_user', '/0/72/73/76', 3, 'sys:sysUser:edit', 73, 0, '', 0, 0, 1, 1, '2023-09-25 22:01:45.149', '2023-09-25 22:01:45.154', NULL);
INSERT INTO `sys_menu` VALUES (77, '', '删除用户', '', 'sys_user', '/0/72/73/77', 3, 'sys:sysUser:remove', 73, 0, '', 0, 0, 1, 1, '2023-09-25 22:01:45.165', '2023-09-25 22:01:45.170', NULL);
INSERT INTO `sys_menu` VALUES (78, 'SysMenuManage', '菜单管理', 'pass', '/sys/menu', '/0/72/78', 2, 'sys:sysMenu:list', 72, 0, '/sys/menu/index', 0, 0, 1, 1, '2023-09-25 22:02:23.418', '2023-09-25 22:02:23.426', NULL);
INSERT INTO `sys_menu` VALUES (79, '', '菜单详情', '', 'sys_menu', '/0/72/78/79', 3, 'sys:sysMenu:query', 78, 0, '', 0, 0, 1, 1, '2023-09-25 22:02:23.440', '2023-09-25 22:02:23.446', NULL);
INSERT INTO `sys_menu` VALUES (80, '', '创建菜单', '', 'sys_menu', '/0/72/78/80', 3, 'sys:sysMenu:add', 78, 0, '', 0, 0, 1, 1, '2023-09-25 22:02:23.459', '2023-09-25 22:02:23.463', NULL);
INSERT INTO `sys_menu` VALUES (81, '', '修改菜单', '', 'sys_menu', '/0/72/78/81', 3, 'sys:sysMenu:edit', 78, 0, '', 0, 0, 1, 1, '2023-09-25 22:02:23.472', '2023-09-25 22:02:23.477', NULL);
INSERT INTO `sys_menu` VALUES (82, '', '删除菜单', '', 'sys_menu', '/0/72/78/82', 3, 'sys:sysMenu:remove', 78, 0, '', 0, 0, 1, 1, '2023-09-25 22:02:23.485', '2023-09-25 22:02:23.489', NULL);
INSERT INTO `sys_menu` VALUES (83, 'SysDeptManage', '部门管理', 'pass', '/sys/dept', '/0/72/83', 2, 'sys:sysDept:list', 72, 0, '/sys/dept/index', 0, 0, 1, 1, '2023-09-26 09:28:25.241', '2023-09-26 09:28:25.258', NULL);
INSERT INTO `sys_menu` VALUES (84, '', '部门详情', '', 'sys_dept', '/0/72/83/84', 3, 'sys:sysDept:query', 83, 0, '', 0, 0, 1, 1, '2023-09-26 09:28:25.275', '2023-09-26 09:28:25.278', NULL);
INSERT INTO `sys_menu` VALUES (85, '', '创建部门', '', 'sys_dept', '/0/72/83/85', 3, 'sys:sysDept:add', 83, 0, '', 0, 0, 1, 1, '2023-09-26 09:28:25.285', '2023-09-26 09:28:25.287', NULL);
INSERT INTO `sys_menu` VALUES (86, '', '修改部门', '', 'sys_dept', '/0/72/83/86', 3, 'sys:sysDept:edit', 83, 0, '', 0, 0, 1, 1, '2023-09-26 09:28:25.297', '2023-09-26 09:28:25.299', NULL);
INSERT INTO `sys_menu` VALUES (87, '', '删除部门', '', 'sys_dept', '/0/72/83/87', 3, 'sys:sysDept:remove', 83, 0, '', 0, 0, 1, 1, '2023-09-26 09:28:25.306', '2023-09-26 09:28:25.310', NULL);
INSERT INTO `sys_menu` VALUES (88, 'SysRoleManage', '角色管理', 'pass', '/sys/role', '/0/72/88', 2, 'sys:sysRole:list', 72, 0, '/sys/role/index', 0, 0, 1, 1, '2023-09-26 09:28:32.331', '2023-09-26 09:28:32.335', NULL);
INSERT INTO `sys_menu` VALUES (89, '', '角色详情', '', 'sys_role', '/0/72/88/89', 3, 'sys:sysRole:query', 88, 0, '', 0, 0, 1, 1, '2023-09-26 09:28:32.361', '2023-09-26 09:28:32.365', NULL);
INSERT INTO `sys_menu` VALUES (90, '', '创建角色', '', 'sys_role', '/0/72/88/90', 3, 'sys:sysRole:add', 88, 0, '', 0, 0, 1, 1, '2023-09-26 09:28:32.375', '2023-09-26 09:28:32.378', NULL);
INSERT INTO `sys_menu` VALUES (91, '', '修改角色', '', 'sys_role', '/0/72/88/91', 3, 'sys:sysRole:edit', 88, 0, '', 0, 0, 1, 1, '2023-09-26 09:28:32.384', '2023-09-26 09:28:32.386', NULL);
INSERT INTO `sys_menu` VALUES (92, '', '删除角色', '', 'sys_role', '/0/72/88/92', 3, 'sys:sysRole:remove', 88, 0, '', 0, 0, 1, 1, '2023-09-26 09:28:32.396', '2023-09-26 09:28:32.398', NULL);

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
INSERT INTO `sys_menu_api_rule` VALUES (73, 55);
INSERT INTO `sys_menu_api_rule` VALUES (74, 56);
INSERT INTO `sys_menu_api_rule` VALUES (75, 57);
INSERT INTO `sys_menu_api_rule` VALUES (76, 58);
INSERT INTO `sys_menu_api_rule` VALUES (77, 59);
INSERT INTO `sys_menu_api_rule` VALUES (78, 60);
INSERT INTO `sys_menu_api_rule` VALUES (79, 61);
INSERT INTO `sys_menu_api_rule` VALUES (80, 62);
INSERT INTO `sys_menu_api_rule` VALUES (81, 63);
INSERT INTO `sys_menu_api_rule` VALUES (82, 64);
INSERT INTO `sys_menu_api_rule` VALUES (83, 65);
INSERT INTO `sys_menu_api_rule` VALUES (84, 66);
INSERT INTO `sys_menu_api_rule` VALUES (85, 67);
INSERT INTO `sys_menu_api_rule` VALUES (86, 68);
INSERT INTO `sys_menu_api_rule` VALUES (87, 69);
INSERT INTO `sys_menu_api_rule` VALUES (88, 70);
INSERT INTO `sys_menu_api_rule` VALUES (89, 71);
INSERT INTO `sys_menu_api_rule` VALUES (90, 72);
INSERT INTO `sys_menu_api_rule` VALUES (91, 73);
INSERT INTO `sys_menu_api_rule` VALUES (92, 74);

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
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色' ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (2, 'dilu', '13800138000', '', '$2a$10$2OxaPJviu7NMSKMk5c2mPOvvb41Xg5ZiQB0153QpB77THK4sIXF1a', 'dilu', '', '', NULL, NULL, '2', 1, '', '', 0, 0, 0, '2023-09-19 11:49:13.139', '2023-09-19 11:49:13.139', NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (3, 'tangtang', '13800138001', '', '$2a$10$2OxaPJviu7NMSKMk5c2mPOvvb41Xg5ZiQB0153QpB77THK4sIXF1a', '糖糖', '', '', NULL, NULL, '2', 0, '', '', 0, 0, 0, '2023-09-25 20:27:23.737', '2023-09-25 20:27:23.737', NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (4, NULL, NULL, NULL, NULL, '臧春梅', NULL, NULL, NULL, NULL, '2', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (5, NULL, NULL, NULL, NULL, '李艳雷', NULL, NULL, NULL, NULL, '1', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (6, NULL, NULL, NULL, NULL, '张华', NULL, NULL, NULL, NULL, '1', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (7, NULL, NULL, NULL, NULL, '简小丽', NULL, NULL, NULL, NULL, '2', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (8, NULL, NULL, NULL, NULL, '余鸿雁', NULL, NULL, NULL, NULL, '2', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (9, NULL, NULL, NULL, NULL, '胡珊', NULL, NULL, NULL, NULL, '2', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
