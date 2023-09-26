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

 Date: 26/09/2023 17:59:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

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
INSERT INTO `gen_columns` VALUES (163, 14, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:05.987', '2023-09-26 13:44:05.987', NULL);
INSERT INTO `gen_columns` VALUES (164, 14, 'username', '用户名', 'varchar(32)', 'string', 'Username', 'username', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:05.992', '2023-09-26 13:44:05.992', NULL);
INSERT INTO `gen_columns` VALUES (165, 14, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:05.996', '2023-09-26 13:44:05.996', NULL);
INSERT INTO `gen_columns` VALUES (166, 14, 'email', '邮箱', 'varchar(128)', 'string', 'Email', 'email', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:05.999', '2023-09-26 13:44:05.999', NULL);
INSERT INTO `gen_columns` VALUES (167, 14, 'password', '密码', 'varchar(128)', 'string', 'Password', 'password', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.002', '2023-09-26 13:44:06.002', NULL);
INSERT INTO `gen_columns` VALUES (168, 14, 'nickname', '昵称', 'varchar(128)', 'string', 'Nickname', 'nickname', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.005', '2023-09-26 13:44:06.005', NULL);
INSERT INTO `gen_columns` VALUES (169, 14, 'name', '姓名', 'varchar(64)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.008', '2023-09-26 13:44:06.008', NULL);
INSERT INTO `gen_columns` VALUES (170, 14, 'avatar', '头像', 'varchar(255)', 'string', 'Avatar', 'avatar', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.011', '2023-09-26 13:44:06.011', NULL);
INSERT INTO `gen_columns` VALUES (171, 14, 'bio', '签名', 'varchar(255)', 'string', 'Bio', 'bio', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.015', '2023-09-26 13:44:06.015', NULL);
INSERT INTO `gen_columns` VALUES (172, 14, 'birthday', '生日 格式 yyyy-MM-dd', 'date', 'string', 'Birthday', 'birthday', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.018', '2023-09-26 13:44:06.018', NULL);
INSERT INTO `gen_columns` VALUES (173, 14, 'gender', '性别 1男 2女 3未知', 'char(1)', 'string', 'Gender', 'gender', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.021', '2023-09-26 13:44:06.021', NULL);
INSERT INTO `gen_columns` VALUES (174, 14, 'role_id', '角色ID', 'int unsigned', 'int', 'RoleId', 'roleId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.023', '2023-09-26 13:44:06.023', NULL);
INSERT INTO `gen_columns` VALUES (175, 14, 'post', '岗位', 'varchar(32)', 'string', 'Post', 'post', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.027', '2023-09-26 13:44:06.027', NULL);
INSERT INTO `gen_columns` VALUES (176, 14, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.031', '2023-09-26 13:44:06.031', NULL);
INSERT INTO `gen_columns` VALUES (177, 14, 'status', '状态 1冻结 2正常 3默认', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.034', '2023-09-26 13:44:06.034', NULL);
INSERT INTO `gen_columns` VALUES (178, 14, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.037', '2023-09-26 13:44:06.037', NULL);
INSERT INTO `gen_columns` VALUES (179, 14, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.039', '2023-09-26 13:44:06.039', NULL);
INSERT INTO `gen_columns` VALUES (180, 14, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.042', '2023-09-26 13:44:06.042', NULL);
INSERT INTO `gen_columns` VALUES (181, 14, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 19, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.045', '2023-09-26 13:44:06.045', NULL);
INSERT INTO `gen_columns` VALUES (182, 14, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 20, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:06.049', '2023-09-26 13:44:06.049', NULL);
INSERT INTO `gen_columns` VALUES (183, 15, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.539', '2023-09-26 13:44:51.539', NULL);
INSERT INTO `gen_columns` VALUES (184, 15, 'menu_name', '菜单名', 'varchar(128)', 'string', 'MenuName', 'menuName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.545', '2023-09-26 13:44:51.545', NULL);
INSERT INTO `gen_columns` VALUES (185, 15, 'title', '显示名称', 'varchar(128)', 'string', 'Title', 'title', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.549', '2023-09-26 13:44:51.549', NULL);
INSERT INTO `gen_columns` VALUES (186, 15, 'icon', '图标', 'varchar(128)', 'string', 'Icon', 'icon', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.564', '2023-09-26 13:44:51.564', NULL);
INSERT INTO `gen_columns` VALUES (187, 15, 'path', '路径', 'varchar(128)', 'string', 'Path', 'path', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.570', '2023-09-26 13:44:51.570', NULL);
INSERT INTO `gen_columns` VALUES (188, 15, 'paths', '路径ids/分割', 'varchar(128)', 'string', 'Paths', 'paths', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.574', '2023-09-26 13:44:51.574', NULL);
INSERT INTO `gen_columns` VALUES (189, 15, 'menu_type', '菜单类型 1 分类 2菜单 3方法按钮', 'tinyint', 'int', 'MenuType', 'menuType', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.578', '2023-09-26 13:44:51.578', NULL);
INSERT INTO `gen_columns` VALUES (190, 15, 'permission', '权限', 'varchar(255)', 'string', 'Permission', 'permission', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.581', '2023-09-26 13:44:51.581', NULL);
INSERT INTO `gen_columns` VALUES (191, 15, 'parent_id', '菜单父id', 'int unsigned', 'int', 'ParentId', 'parentId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.583', '2023-09-26 13:44:51.583', NULL);
INSERT INTO `gen_columns` VALUES (192, 15, 'no_cache', '是否缓存', 'tinyint(1)', 'int', 'NoCache', 'noCache', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.585', '2023-09-26 13:44:51.585', NULL);
INSERT INTO `gen_columns` VALUES (193, 15, 'component', '前端组件路径', 'varchar(255)', 'string', 'Component', 'component', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.588', '2023-09-26 13:44:51.588', NULL);
INSERT INTO `gen_columns` VALUES (194, 15, 'sort', '排序倒叙', 'tinyint', 'int', 'Sort', 'sort', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.590', '2023-09-26 13:44:51.590', NULL);
INSERT INTO `gen_columns` VALUES (195, 15, 'hidden', '是否隐藏', 'tinyint(1)', 'int', 'Hidden', 'hidden', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.595', '2023-09-26 13:44:51.595', NULL);
INSERT INTO `gen_columns` VALUES (196, 15, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.598', '2023-09-26 13:44:51.598', NULL);
INSERT INTO `gen_columns` VALUES (197, 15, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.600', '2023-09-26 13:44:51.600', NULL);
INSERT INTO `gen_columns` VALUES (198, 15, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.603', '2023-09-26 13:44:51.603', NULL);
INSERT INTO `gen_columns` VALUES (199, 15, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.605', '2023-09-26 13:44:51.605', NULL);
INSERT INTO `gen_columns` VALUES (200, 15, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:44:51.608', '2023-09-26 13:44:51.608', NULL);
INSERT INTO `gen_columns` VALUES (201, 16, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.901', '2023-09-26 13:45:11.901', NULL);
INSERT INTO `gen_columns` VALUES (202, 16, 'name', '角色名称', 'varchar(128)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.904', '2023-09-26 13:45:11.904', NULL);
INSERT INTO `gen_columns` VALUES (203, 16, 'status', '状态', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.910', '2023-09-26 13:45:11.910', NULL);
INSERT INTO `gen_columns` VALUES (204, 16, 'role_key', '角色代码', 'varchar(128)', 'string', 'RoleKey', 'roleKey', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.915', '2023-09-26 13:45:11.915', NULL);
INSERT INTO `gen_columns` VALUES (205, 16, 'role_sort', '排序', 'int unsigned', 'int', 'RoleSort', 'roleSort', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.926', '2023-09-26 13:45:11.926', NULL);
INSERT INTO `gen_columns` VALUES (206, 16, 'flag', 'flag', 'varchar(128)', 'string', 'Flag', 'flag', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.932', '2023-09-26 13:45:11.932', NULL);
INSERT INTO `gen_columns` VALUES (207, 16, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.936', '2023-09-26 13:45:11.936', NULL);
INSERT INTO `gen_columns` VALUES (208, 16, 'admin', '管理员', 'tinyint(1)', 'int', 'Admin', 'admin', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.938', '2023-09-26 13:45:11.938', NULL);
INSERT INTO `gen_columns` VALUES (209, 16, 'data_scope', '数据权限', 'varchar(128)', 'string', 'DataScope', 'dataScope', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.942', '2023-09-26 13:45:11.942', NULL);
INSERT INTO `gen_columns` VALUES (210, 16, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.946', '2023-09-26 13:45:11.946', NULL);
INSERT INTO `gen_columns` VALUES (211, 16, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.949', '2023-09-26 13:45:11.949', NULL);
INSERT INTO `gen_columns` VALUES (212, 16, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.951', '2023-09-26 13:45:11.951', NULL);
INSERT INTO `gen_columns` VALUES (213, 16, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.953', '2023-09-26 13:45:11.953', NULL);
INSERT INTO `gen_columns` VALUES (214, 16, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:45:11.955', '2023-09-26 13:45:11.955', NULL);
INSERT INTO `gen_columns` VALUES (215, 17, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.212', '2023-09-26 13:46:06.212', NULL);
INSERT INTO `gen_columns` VALUES (216, 17, 'parent_id', '父id', 'int unsigned', 'int', 'ParentId', 'parentId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.216', '2023-09-26 13:46:06.216', NULL);
INSERT INTO `gen_columns` VALUES (217, 17, 'dept_path', '部门路径', 'varchar(255)', 'string', 'DeptPath', 'deptPath', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.221', '2023-09-26 13:46:06.221', NULL);
INSERT INTO `gen_columns` VALUES (218, 17, 'name', '部门名', 'varchar(128)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.227', '2023-09-26 13:46:06.227', NULL);
INSERT INTO `gen_columns` VALUES (219, 17, 'type', '类型', 'tinyint', 'int', 'Type', 'type', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.231', '2023-09-26 13:46:06.231', NULL);
INSERT INTO `gen_columns` VALUES (220, 17, 'principal', '部门领导', 'varchar(128)', 'string', 'Principal', 'principal', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.233', '2023-09-26 13:46:06.233', NULL);
INSERT INTO `gen_columns` VALUES (221, 17, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.235', '2023-09-26 13:46:06.235', NULL);
INSERT INTO `gen_columns` VALUES (222, 17, 'email', '邮箱', 'varchar(128)', 'string', 'Email', 'email', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.237', '2023-09-26 13:46:06.237', NULL);
INSERT INTO `gen_columns` VALUES (223, 17, 'sort', '排序', 'tinyint', 'int', 'Sort', 'sort', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.239', '2023-09-26 13:46:06.239', NULL);
INSERT INTO `gen_columns` VALUES (224, 17, 'status', '状态', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.243', '2023-09-26 13:46:06.243', NULL);
INSERT INTO `gen_columns` VALUES (225, 17, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.245', '2023-09-26 13:46:06.245', NULL);
INSERT INTO `gen_columns` VALUES (226, 17, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.248', '2023-09-26 13:46:06.248', NULL);
INSERT INTO `gen_columns` VALUES (227, 17, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.252', '2023-09-26 13:46:06.252', NULL);
INSERT INTO `gen_columns` VALUES (228, 17, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.254', '2023-09-26 13:46:06.254', NULL);
INSERT INTO `gen_columns` VALUES (229, 17, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.256', '2023-09-26 13:46:06.256', NULL);
INSERT INTO `gen_columns` VALUES (230, 17, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 13:46:06.260', '2023-09-26 13:46:06.260', NULL);
INSERT INTO `gen_columns` VALUES (241, 19, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.213', '2023-09-26 14:07:44.213', NULL);
INSERT INTO `gen_columns` VALUES (242, 19, 'dept_id', '部门id', 'int unsigned', 'int', 'DeptId', 'deptId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.216', '2023-09-26 14:07:44.216', NULL);
INSERT INTO `gen_columns` VALUES (243, 19, 'user_id', '用户id', 'int unsigned', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.223', '2023-09-26 14:07:44.223', NULL);
INSERT INTO `gen_columns` VALUES (244, 19, 'nickname', '昵称', 'varchar(128)', 'string', 'Nickname', 'nickname', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.226', '2023-09-26 14:07:44.226', NULL);
INSERT INTO `gen_columns` VALUES (245, 19, 'name', '姓名', 'varchar(64)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.239', '2023-09-26 14:07:44.239', NULL);
INSERT INTO `gen_columns` VALUES (246, 19, 'phone', '电话', 'varchar(11)', 'string', 'Phone', 'phone', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.243', '2023-09-26 14:07:44.243', NULL);
INSERT INTO `gen_columns` VALUES (247, 19, 'dept_path', '部门路径', 'varchar(255)', 'string', 'DeptPath', 'deptPath', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.249', '2023-09-26 14:07:44.249', NULL);
INSERT INTO `gen_columns` VALUES (248, 19, 'post_tag', '职位标签 1主管 2副主管 3员工', 'tinyint unsigned', 'int', 'PostTag', 'postTag', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.252', '2023-09-26 14:07:44.252', NULL);
INSERT INTO `gen_columns` VALUES (249, 19, 'status', '状态 1正常 ', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.254', '2023-09-26 14:07:44.254', NULL);
INSERT INTO `gen_columns` VALUES (250, 19, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.255', '2023-09-26 14:07:44.255', NULL);
INSERT INTO `gen_columns` VALUES (251, 19, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.257', '2023-09-26 14:07:44.257', NULL);
INSERT INTO `gen_columns` VALUES (252, 19, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.259', '2023-09-26 14:07:44.259', NULL);
INSERT INTO `gen_columns` VALUES (253, 19, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.263', '2023-09-26 14:07:44.263', NULL);
INSERT INTO `gen_columns` VALUES (254, 19, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-26 14:07:44.266', '2023-09-26 14:07:44.266', NULL);

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
INSERT INTO `gen_tables` VALUES (14, 'dilu-db', 'sys_user', '用户', 'SysUser', 'crud', 'sys', 'sys-user', '', 'sysUser', '用户', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-26 13:44:05.983', '2023-09-26 13:44:05.983', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (15, 'dilu-db', 'sys_menu', '菜单', 'SysMenu', 'crud', 'sys', 'sys-menu', '', 'sysMenu', '菜单', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-26 13:44:51.533', '2023-09-26 13:44:51.533', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (16, 'dilu-db', 'sys_role', '角色', 'SysRole', 'crud', 'sys', 'sys-role', '', 'sysRole', '角色', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-26 13:45:11.897', '2023-09-26 13:45:11.897', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (17, 'dilu-db', 'sys_dept', '部门', 'SysDept', 'crud', 'sys', 'sys-dept', '', 'sysDept', '部门', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-26 13:46:06.206', '2023-09-26 13:46:06.206', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (19, 'dilu-db', 'sys_member', '成员', 'SysMember', 'crud', 'sys', 'sys-member', '', 'sysMember', '成员', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-26 14:07:44.207', '2023-09-26 14:07:44.207', NULL, 0, 0);

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '接口' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_api
-- ----------------------------
INSERT INTO `sys_api` VALUES (1, '分页获取用户', 'POST', '/api/v1/sys/sys-user/page', '', 't', 3, 0, '2023-09-26 13:46:59.488');
INSERT INTO `sys_api` VALUES (2, '根据id获取用户', 'POST', '/api/v1/sys/sys-user/get', '', 't', 3, 0, '2023-09-26 13:46:59.518');
INSERT INTO `sys_api` VALUES (3, '创建用户', 'POST', '/api/v1/sys/sys-user/create', '', 't', 3, 0, '2023-09-26 13:46:59.532');
INSERT INTO `sys_api` VALUES (4, '修改用户', 'POST', '/api/v1/sys/sys-user/update', '', 't', 3, 0, '2023-09-26 13:46:59.539');
INSERT INTO `sys_api` VALUES (5, '删除用户', 'POST', '/api/v1/sys/sys-user/del', '', 't', 3, 0, '2023-09-26 13:46:59.550');
INSERT INTO `sys_api` VALUES (6, '分页获取菜单', 'POST', '/api/v1/sys/sys-menu/page', '', 't', 3, 0, '2023-09-26 13:47:37.020');
INSERT INTO `sys_api` VALUES (7, '根据id获取菜单', 'POST', '/api/v1/sys/sys-menu/get', '', 't', 3, 0, '2023-09-26 13:47:37.038');
INSERT INTO `sys_api` VALUES (8, '创建菜单', 'POST', '/api/v1/sys/sys-menu/create', '', 't', 3, 0, '2023-09-26 13:47:37.064');
INSERT INTO `sys_api` VALUES (9, '修改菜单', 'POST', '/api/v1/sys/sys-menu/update', '', 't', 3, 0, '2023-09-26 13:47:37.073');
INSERT INTO `sys_api` VALUES (10, '删除菜单', 'POST', '/api/v1/sys/sys-menu/del', '', 't', 3, 0, '2023-09-26 13:47:37.082');
INSERT INTO `sys_api` VALUES (11, '分页获取角色', 'POST', '/api/v1/sys/sys-role/page', '', 't', 3, 0, '2023-09-26 13:47:39.685');
INSERT INTO `sys_api` VALUES (12, '根据id获取角色', 'POST', '/api/v1/sys/sys-role/get', '', 't', 3, 0, '2023-09-26 13:47:39.701');
INSERT INTO `sys_api` VALUES (13, '创建角色', 'POST', '/api/v1/sys/sys-role/create', '', 't', 3, 0, '2023-09-26 13:47:39.734');
INSERT INTO `sys_api` VALUES (14, '修改角色', 'POST', '/api/v1/sys/sys-role/update', '', 't', 3, 0, '2023-09-26 13:47:39.762');
INSERT INTO `sys_api` VALUES (15, '删除角色', 'POST', '/api/v1/sys/sys-role/del', '', 't', 3, 0, '2023-09-26 13:47:39.773');
INSERT INTO `sys_api` VALUES (16, '分页获取部门', 'POST', '/api/v1/sys/sys-dept/page', '', 't', 3, 0, '2023-09-26 13:47:42.136');
INSERT INTO `sys_api` VALUES (17, '根据id获取部门', 'POST', '/api/v1/sys/sys-dept/get', '', 't', 3, 0, '2023-09-26 13:47:42.152');
INSERT INTO `sys_api` VALUES (18, '创建部门', 'POST', '/api/v1/sys/sys-dept/create', '', 't', 3, 0, '2023-09-26 13:47:42.172');
INSERT INTO `sys_api` VALUES (19, '修改部门', 'POST', '/api/v1/sys/sys-dept/update', '', 't', 3, 0, '2023-09-26 13:47:42.186');
INSERT INTO `sys_api` VALUES (20, '删除部门', 'POST', '/api/v1/sys/sys-dept/del', '', 't', 3, 0, '2023-09-26 13:47:42.197');
INSERT INTO `sys_api` VALUES (21, '分页获取用户部门', 'POST', '/api/v1/sys/sys-user-dept/page', '', 't', 3, 0, '2023-09-26 13:47:45.447');
INSERT INTO `sys_api` VALUES (22, '根据id获取用户部门', 'POST', '/api/v1/sys/sys-user-dept/get', '', 't', 3, 0, '2023-09-26 13:47:45.459');
INSERT INTO `sys_api` VALUES (23, '创建用户部门', 'POST', '/api/v1/sys/sys-user-dept/create', '', 't', 3, 0, '2023-09-26 13:47:45.483');
INSERT INTO `sys_api` VALUES (24, '修改用户部门', 'POST', '/api/v1/sys/sys-user-dept/update', '', 't', 3, 0, '2023-09-26 13:47:45.495');
INSERT INTO `sys_api` VALUES (25, '删除用户部门', 'POST', '/api/v1/sys/sys-user-dept/del', '', 't', 3, 0, '2023-09-26 13:47:45.503');
INSERT INTO `sys_api` VALUES (26, '分页获取成员', 'POST', '/api/v1/sys/sys-member/page', '', 't', 3, 0, '2023-09-26 14:09:40.443');
INSERT INTO `sys_api` VALUES (27, '根据id获取成员', 'POST', '/api/v1/sys/sys-member/get', '', 't', 3, 0, '2023-09-26 14:09:40.459');
INSERT INTO `sys_api` VALUES (28, '创建成员', 'POST', '/api/v1/sys/sys-member/create', '', 't', 3, 0, '2023-09-26 14:09:40.483');
INSERT INTO `sys_api` VALUES (29, '修改成员', 'POST', '/api/v1/sys/sys-member/update', '', 't', 3, 0, '2023-09-26 14:09:40.493');
INSERT INTO `sys_api` VALUES (30, '删除成员', 'POST', '/api/v1/sys/sys-member/del', '', 't', 3, 0, '2023-09-26 14:09:40.502');

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '配置' ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '部门' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (1, 0, '/0/1', 'xxx', 1, '的卢', '13800138000', NULL, 1, 1, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dept` VALUES (2, 1, '/0/1/2', 'xxx-xx', 1, '', '', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dept` VALUES (3, 1, '/0/1/3', 'xxx-xx3', 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dept` VALUES (4, 0, '/0/4', 'yyy', 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `sys_dept` VALUES (5, 4, '/0/4/5', 'yyy-5', 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '邮件' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_email
-- ----------------------------

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `job_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '名称',
  `job_group` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '组',
  `job_type` tinyint(0) NULL DEFAULT NULL COMMENT '类型',
  `cron_expression` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '表达式',
  `invoke_target` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '调用目标',
  `args` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '参数',
  `misfire_policy` bigint(0) NULL DEFAULT NULL COMMENT '策略',
  `concurrent` tinyint(0) NULL DEFAULT NULL COMMENT '并发',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `entry_id` smallint(0) NULL DEFAULT NULL COMMENT '任务id',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_job_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '定时任务' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job
-- ----------------------------

-- ----------------------------
-- Table structure for sys_member
-- ----------------------------
DROP TABLE IF EXISTS `sys_member`;
CREATE TABLE `sys_member`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `dept_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '部门id',
  `user_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '昵称',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '姓名',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '电话',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `post_tag` tinyint(0) UNSIGNED NULL DEFAULT NULL COMMENT '职位标签 1主管 2副主管 3员工',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态 1正常 ',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_did_uid`(`dept_id`, `user_id`) USING BTREE,
  INDEX `idx_sys_member_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '成员' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_member
-- ----------------------------

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
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
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_menu_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_menu_create_by`(`create_by`) USING BTREE,
  INDEX `idx_sys_menu_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, '', '系统管理', 'pass', '/sys', '/0/1', 1, '', 0, 0, 'Layout', 0, 0, 1, 0, '2023-09-26 13:46:59.480', '2023-09-26 13:46:59.481', NULL);
INSERT INTO `sys_menu` VALUES (2, 'SysUserManage', '用户管理', 'pass', '/sys/sys-user', '/0/1/2', 2, 'sys:sysUser:list', 1, 0, '/sys/user/index', 0, 0, 1, 1, '2023-09-26 13:46:59.493', '2023-09-26 13:46:59.503', NULL);
INSERT INTO `sys_menu` VALUES (3, '', '用户详情', '', 'sys_user_detail', '/0/1/2/3', 3, 'sys:sysUser:query', 2, 0, '', 0, 0, 1, 1, '2023-09-26 13:46:59.522', '2023-09-26 13:46:59.526', NULL);
INSERT INTO `sys_menu` VALUES (4, '', '用户创建', '', 'sys_user_create', '/0/1/2/4', 3, 'sys:sysUser:add', 2, 0, '', 0, 0, 1, 1, '2023-09-26 13:46:59.534', '2023-09-26 13:46:59.536', NULL);
INSERT INTO `sys_menu` VALUES (5, '', '用户修改', '', 'sys_user_update', '/0/1/2/5', 3, 'sys:sysUser:edit', 2, 0, '', 0, 0, 1, 1, '2023-09-26 13:46:59.543', '2023-09-26 13:46:59.546', NULL);
INSERT INTO `sys_menu` VALUES (6, '', '用户删除', '', 'sys_user_del', '/0/1/2/6', 3, 'sys:sysUser:remove', 2, 0, '', 0, 0, 1, 1, '2023-09-26 13:46:59.553', '2023-09-26 13:46:59.555', NULL);
INSERT INTO `sys_menu` VALUES (7, 'SysMenuManage', '菜单管理', 'pass', '/sys/sys-menu', '/0/1/7', 2, 'sys:sysMenu:list', 1, 0, '/sys/menu/index', 0, 0, 1, 1, '2023-09-26 13:47:37.025', '2023-09-26 13:47:37.029', NULL);
INSERT INTO `sys_menu` VALUES (8, '', '菜单详情', '', 'sys_menu_detail', '/0/1/7/8', 3, 'sys:sysMenu:query', 7, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:37.042', '2023-09-26 13:47:37.055', NULL);
INSERT INTO `sys_menu` VALUES (9, '', '菜单创建', '', 'sys_menu_create', '/0/1/7/9', 3, 'sys:sysMenu:add', 7, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:37.067', '2023-09-26 13:47:37.069', NULL);
INSERT INTO `sys_menu` VALUES (10, '', '菜单修改', '', 'sys_menu_update', '/0/1/7/10', 3, 'sys:sysMenu:edit', 7, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:37.076', '2023-09-26 13:47:37.079', NULL);
INSERT INTO `sys_menu` VALUES (11, '', '菜单删除', '', 'sys_menu_del', '/0/1/7/11', 3, 'sys:sysMenu:remove', 7, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:37.084', '2023-09-26 13:47:37.086', NULL);
INSERT INTO `sys_menu` VALUES (12, 'SysRoleManage', '角色管理', 'pass', '/sys/sys-role', '/0/1/12', 2, 'sys:sysRole:list', 1, 0, '/sys/role/index', 0, 0, 1, 1, '2023-09-26 13:47:39.688', '2023-09-26 13:47:39.694', NULL);
INSERT INTO `sys_menu` VALUES (13, '', '角色详情', '', 'sys_role_detail', '/0/1/12/13', 3, 'sys:sysRole:query', 12, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:39.714', '2023-09-26 13:47:39.717', NULL);
INSERT INTO `sys_menu` VALUES (14, '', '角色创建', '', 'sys_role_create', '/0/1/12/14', 3, 'sys:sysRole:add', 12, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:39.747', '2023-09-26 13:47:39.752', NULL);
INSERT INTO `sys_menu` VALUES (15, '', '角色修改', '', 'sys_role_update', '/0/1/12/15', 3, 'sys:sysRole:edit', 12, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:39.765', '2023-09-26 13:47:39.769', NULL);
INSERT INTO `sys_menu` VALUES (16, '', '角色删除', '', 'sys_role_del', '/0/1/12/16', 3, 'sys:sysRole:remove', 12, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:39.776', '2023-09-26 13:47:39.778', NULL);
INSERT INTO `sys_menu` VALUES (17, 'SysDeptManage', '部门管理', 'pass', '/sys/sys-dept', '/0/1/17', 2, 'sys:sysDept:list', 1, 0, '/sys/dept/index', 0, 0, 1, 1, '2023-09-26 13:47:42.140', '2023-09-26 13:47:42.144', NULL);
INSERT INTO `sys_menu` VALUES (18, '', '部门详情', '', 'sys_dept_detail', '/0/1/17/18', 3, 'sys:sysDept:query', 17, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:42.155', '2023-09-26 13:47:42.167', NULL);
INSERT INTO `sys_menu` VALUES (19, '', '部门创建', '', 'sys_dept_create', '/0/1/17/19', 3, 'sys:sysDept:add', 17, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:42.180', '2023-09-26 13:47:42.182', NULL);
INSERT INTO `sys_menu` VALUES (20, '', '部门修改', '', 'sys_dept_update', '/0/1/17/20', 3, 'sys:sysDept:edit', 17, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:42.188', '2023-09-26 13:47:42.190', NULL);
INSERT INTO `sys_menu` VALUES (21, '', '部门删除', '', 'sys_dept_del', '/0/1/17/21', 3, 'sys:sysDept:remove', 17, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:42.200', '2023-09-26 13:47:42.202', NULL);
INSERT INTO `sys_menu` VALUES (27, 'SysMemberManage', '成员管理', 'pass', '/sys/sys-member', '/0/1/27', 2, 'sys:sysMember:list', 1, 0, '/sys/sys-member/index', 0, 0, 1, 1, '2023-09-26 14:09:40.447', '2023-09-26 14:09:40.450', NULL);
INSERT INTO `sys_menu` VALUES (28, '', '成员详情', '', 'sys_member_detail', '/0/1/27/28', 3, 'sys:sysMember:query', 27, 0, '', 0, 0, 1, 1, '2023-09-26 14:09:40.462', '2023-09-26 14:09:40.465', NULL);
INSERT INTO `sys_menu` VALUES (29, '', '成员创建', '', 'sys_member_create', '/0/1/27/29', 3, 'sys:sysMember:add', 27, 0, '', 0, 0, 1, 1, '2023-09-26 14:09:40.487', '2023-09-26 14:09:40.489', NULL);
INSERT INTO `sys_menu` VALUES (30, '', '成员修改', '', 'sys_member_update', '/0/1/27/30', 3, 'sys:sysMember:edit', 27, 0, '', 0, 0, 1, 1, '2023-09-26 14:09:40.496', '2023-09-26 14:09:40.498', NULL);
INSERT INTO `sys_menu` VALUES (31, '', '成员删除', '', 'sys_member_del', '/0/1/27/31', 3, 'sys:sysMember:remove', 27, 0, '', 0, 0, 1, 1, '2023-09-26 14:09:40.504', '2023-09-26 14:09:40.506', NULL);

-- ----------------------------
-- Table structure for sys_menu_api_rule
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_api_rule`;
CREATE TABLE `sys_menu_api_rule`  (
  `sys_menu_id` int(0) UNSIGNED NOT NULL COMMENT '主键',
  `sys_api_id` int(0) UNSIGNED NOT NULL COMMENT '主键编码',
  PRIMARY KEY (`sys_menu_id`, `sys_api_id`) USING BTREE,
  INDEX `fk_sys_menu_api_rule_sys_api`(`sys_api_id`) USING BTREE,
  CONSTRAINT `fk_sys_menu_api_rule_sys_api` FOREIGN KEY (`sys_api_id`) REFERENCES `sys_api` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_menu_api_rule_sys_menu` FOREIGN KEY (`sys_menu_id`) REFERENCES `sys_menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu_api_rule
-- ----------------------------
INSERT INTO `sys_menu_api_rule` VALUES (2, 1);
INSERT INTO `sys_menu_api_rule` VALUES (3, 2);
INSERT INTO `sys_menu_api_rule` VALUES (4, 3);
INSERT INTO `sys_menu_api_rule` VALUES (5, 4);
INSERT INTO `sys_menu_api_rule` VALUES (6, 5);
INSERT INTO `sys_menu_api_rule` VALUES (7, 6);
INSERT INTO `sys_menu_api_rule` VALUES (8, 7);
INSERT INTO `sys_menu_api_rule` VALUES (9, 8);
INSERT INTO `sys_menu_api_rule` VALUES (10, 9);
INSERT INTO `sys_menu_api_rule` VALUES (11, 10);
INSERT INTO `sys_menu_api_rule` VALUES (12, 11);
INSERT INTO `sys_menu_api_rule` VALUES (13, 12);
INSERT INTO `sys_menu_api_rule` VALUES (14, 13);
INSERT INTO `sys_menu_api_rule` VALUES (15, 14);
INSERT INTO `sys_menu_api_rule` VALUES (16, 15);
INSERT INTO `sys_menu_api_rule` VALUES (17, 16);
INSERT INTO `sys_menu_api_rule` VALUES (18, 17);
INSERT INTO `sys_menu_api_rule` VALUES (19, 18);
INSERT INTO `sys_menu_api_rule` VALUES (20, 19);
INSERT INTO `sys_menu_api_rule` VALUES (21, 20);
INSERT INTO `sys_menu_api_rule` VALUES (27, 26);
INSERT INTO `sys_menu_api_rule` VALUES (28, 27);
INSERT INTO `sys_menu_api_rule` VALUES (29, 28);
INSERT INTO `sys_menu_api_rule` VALUES (30, 29);
INSERT INTO `sys_menu_api_rule` VALUES (31, 30);

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
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '操作日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_opera_log
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '角色名称',
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
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_role_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_menu`;
CREATE TABLE `sys_role_menu`  (
  `role_id` int(0) UNSIGNED NOT NULL COMMENT '主键',
  `menu_id` int(0) UNSIGNED NOT NULL COMMENT '主键',
  PRIMARY KEY (`role_id`, `menu_id`) USING BTREE,
  INDEX `fk_sys_role_menu_sys_menu`(`menu_id`) USING BTREE,
  CONSTRAINT `fk_sys_role_menu_sys_menu` FOREIGN KEY (`menu_id`) REFERENCES `sys_menu` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_sys_role_menu_sys_role` FOREIGN KEY (`role_id`) REFERENCES `sys_role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '短信' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_sms
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
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
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_user_create_by`(`create_by`) USING BTREE,
  INDEX `idx_sys_user_update_by`(`update_by`) USING BTREE,
  INDEX `idx_sys_user_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'dilu', '', NULL, '$2a$10$2OxaPJviu7NMSKMk5c2mPOvvb41Xg5ZiQB0153QpB77THK4sIXF1a', 'dilu', 'dilu', NULL, NULL, NULL, '2', 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '三方登录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of third_login
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
