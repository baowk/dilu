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

 Date: 31/10/2023 11:50:37
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
) ENGINE = InnoDB AUTO_INCREMENT = 474 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_columns
-- ----------------------------
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
INSERT INTO `gen_columns` VALUES (255, 20, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:06.365', '2023-09-29 08:33:06.365', NULL);
INSERT INTO `gen_columns` VALUES (256, 20, 'name', '团队名', 'varchar(32)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:06.373', '2023-09-29 08:33:06.373', NULL);
INSERT INTO `gen_columns` VALUES (257, 20, 'owner', '团队拥有者', 'int unsigned', 'int', 'Owner', 'owner', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:06.377', '2023-09-29 08:33:06.377', NULL);
INSERT INTO `gen_columns` VALUES (258, 20, 'status', '状态', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:06.382', '2023-09-29 08:33:06.382', NULL);
INSERT INTO `gen_columns` VALUES (259, 20, 'created_at', '创建时间', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:06.386', '2023-09-29 08:33:06.386', NULL);
INSERT INTO `gen_columns` VALUES (260, 20, 'updated_at', '更新时间', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:06.389', '2023-09-29 08:33:06.389', NULL);
INSERT INTO `gen_columns` VALUES (261, 21, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.155', '2023-09-29 08:33:20.155', NULL);
INSERT INTO `gen_columns` VALUES (262, 21, 'team_id', '团队id', 'int unsigned', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.182', '2023-09-29 08:33:20.182', NULL);
INSERT INTO `gen_columns` VALUES (263, 21, 'user_id', '用户id', 'int unsigned', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.187', '2023-09-29 08:33:20.187', NULL);
INSERT INTO `gen_columns` VALUES (264, 21, 'nickname', '昵称', 'varchar(128)', 'string', 'Nickname', 'nickname', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.191', '2023-09-29 08:33:20.191', NULL);
INSERT INTO `gen_columns` VALUES (265, 21, 'name', '姓名', 'varchar(64)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.194', '2023-09-29 08:33:20.194', NULL);
INSERT INTO `gen_columns` VALUES (266, 21, 'phone', '电话', 'varchar(11)', 'string', 'Phone', 'phone', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.197', '2023-09-29 08:33:20.197', NULL);
INSERT INTO `gen_columns` VALUES (267, 21, 'dept_path', '部门路径', 'varchar(255)', 'string', 'DeptPath', 'deptPath', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.201', '2023-09-29 08:33:20.201', NULL);
INSERT INTO `gen_columns` VALUES (268, 21, 'dept_id', '部门id', 'int unsigned', 'int', 'DeptId', 'deptId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.204', '2023-09-29 08:33:20.204', NULL);
INSERT INTO `gen_columns` VALUES (269, 21, 'post_tag', '职位标签 1主管 2副主管 3员工', 'tinyint unsigned', 'int', 'PostTag', 'postTag', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.208', '2023-09-29 08:33:20.208', NULL);
INSERT INTO `gen_columns` VALUES (270, 21, 'status', '状态 1正常 ', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.211', '2023-09-29 08:33:20.211', NULL);
INSERT INTO `gen_columns` VALUES (271, 21, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.215', '2023-09-29 08:33:20.215', NULL);
INSERT INTO `gen_columns` VALUES (272, 21, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.218', '2023-09-29 08:33:20.218', NULL);
INSERT INTO `gen_columns` VALUES (273, 21, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.221', '2023-09-29 08:33:20.221', NULL);
INSERT INTO `gen_columns` VALUES (274, 21, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.225', '2023-09-29 08:33:20.225', NULL);
INSERT INTO `gen_columns` VALUES (275, 21, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-29 08:33:20.229', '2023-09-29 08:33:20.229', NULL);
INSERT INTO `gen_columns` VALUES (276, 22, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.877', '2023-09-30 14:05:14.877', NULL);
INSERT INTO `gen_columns` VALUES (277, 22, 'no', '订单号', 'varchar(20)', 'string', 'No', 'no', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.887', '2023-09-30 14:05:14.887', NULL);
INSERT INTO `gen_columns` VALUES (278, 22, 'customer_id', '顾客', 'int unsigned', 'int', 'CustomerId', 'customerId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.892', '2023-09-30 14:05:14.892', NULL);
INSERT INTO `gen_columns` VALUES (279, 22, 'user_id', '用户id', 'int unsigned', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.898', '2023-09-30 14:05:14.898', NULL);
INSERT INTO `gen_columns` VALUES (280, 22, 'team_id', '团队id', 'int unsigned', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.902', '2023-09-30 14:05:14.902', NULL);
INSERT INTO `gen_columns` VALUES (281, 22, 'dept_path', '部门路径', 'varchar(255)', 'string', 'DeptPath', 'deptPath', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.906', '2023-09-30 14:05:14.906', NULL);
INSERT INTO `gen_columns` VALUES (282, 22, 'total', '金额', 'decimal(10,2)', 'string', 'Total', 'total', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.910', '2023-09-30 14:05:14.910', NULL);
INSERT INTO `gen_columns` VALUES (283, 22, 'real_total', '折后金额', 'decimal(10,2)', 'string', 'RealTotal', 'realTotal', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.913', '2023-09-30 14:05:14.913', NULL);
INSERT INTO `gen_columns` VALUES (284, 22, 'paid_total', '已支付金额', 'decimal(10,2)', 'string', 'PaidTotal', 'paidTotal', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.917', '2023-09-30 14:05:14.917', NULL);
INSERT INTO `gen_columns` VALUES (285, 22, 'link_id', '关联订单', 'int unsigned', 'int', 'LinkId', 'linkId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.921', '2023-09-30 14:05:14.921', NULL);
INSERT INTO `gen_columns` VALUES (286, 22, 'trade_at', '交易日期', 'datetime', 'time.Time', 'TradeAt', 'tradeAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.924', '2023-09-30 14:05:14.924', NULL);
INSERT INTO `gen_columns` VALUES (287, 22, 'trade_status', '交易类型1 成交 2补尾款  3补上月欠款 10退款', 'tinyint', 'int', 'TradeStatus', 'tradeStatus', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.927', '2023-09-30 14:05:14.927', NULL);
INSERT INTO `gen_columns` VALUES (288, 22, 'dental_count', '颗数', 'tinyint', 'int', 'DentalCount', 'dentalCount', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.931', '2023-09-30 14:05:14.931', NULL);
INSERT INTO `gen_columns` VALUES (289, 22, 'brand', '品牌', 'tinyint', 'int', 'Brand', 'brand', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.934', '2023-09-30 14:05:14.934', NULL);
INSERT INTO `gen_columns` VALUES (290, 22, 'implanted_count', '已种颗数', 'tinyint', 'int', 'ImplantedCount', 'implantedCount', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.939', '2023-09-30 14:05:14.939', NULL);
INSERT INTO `gen_columns` VALUES (291, 22, 'implant', '是否已种', 'tinyint', 'int', 'Implant', 'implant', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.944', '2023-09-30 14:05:14.944', NULL);
INSERT INTO `gen_columns` VALUES (292, 22, 'implant_date', '植入日期', 'datetime', 'time.Time', 'ImplantDate', 'implantDate', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.947', '2023-09-30 14:05:14.947', NULL);
INSERT INTO `gen_columns` VALUES (293, 22, 'doctor', '医生', 'varchar(32)', 'string', 'Doctor', 'doctor', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.950', '2023-09-30 14:05:14.950', NULL);
INSERT INTO `gen_columns` VALUES (294, 22, 'pack', '1 普通 2 半口 3 全口', 'tinyint', 'int', 'Pack', 'pack', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 19, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.955', '2023-09-30 14:05:14.955', NULL);
INSERT INTO `gen_columns` VALUES (295, 22, 'payback_date', '预定回款日期', 'datetime', 'time.Time', 'PaybackDate', 'paybackDate', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 20, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.958', '2023-09-30 14:05:14.958', NULL);
INSERT INTO `gen_columns` VALUES (296, 22, 'tags', '标签', 'varchar(255)', 'string', 'Tags', 'tags', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 21, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.961', '2023-09-30 14:05:14.961', NULL);
INSERT INTO `gen_columns` VALUES (297, 22, 'prj_name', '项目', 'varchar(255)', 'string', 'PrjName', 'prjName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 22, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.965', '2023-09-30 14:05:14.965', NULL);
INSERT INTO `gen_columns` VALUES (298, 22, 'other_prj', '其他项目', 'varchar(255)', 'string', 'OtherPrj', 'otherPrj', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 23, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.969', '2023-09-30 14:05:14.969', NULL);
INSERT INTO `gen_columns` VALUES (299, 22, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 24, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.972', '2023-09-30 14:05:14.972', NULL);
INSERT INTO `gen_columns` VALUES (300, 22, 'created_at', '创建时间', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 25, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.977', '2023-09-30 14:05:14.977', NULL);
INSERT INTO `gen_columns` VALUES (301, 22, 'updated_at', '更新时间', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 26, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:14.980', '2023-09-30 14:05:14.980', NULL);
INSERT INTO `gen_columns` VALUES (302, 23, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.624', '2023-09-30 14:05:22.624', NULL);
INSERT INTO `gen_columns` VALUES (303, 23, 'name', '姓名', 'varchar(32)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.631', '2023-09-30 14:05:22.631', NULL);
INSERT INTO `gen_columns` VALUES (304, 23, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.635', '2023-09-30 14:05:22.635', NULL);
INSERT INTO `gen_columns` VALUES (305, 23, 'wechat', '微信号', 'varchar(64)', 'string', 'Wechat', 'wechat', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.640', '2023-09-30 14:05:22.640', NULL);
INSERT INTO `gen_columns` VALUES (306, 23, 'gender', '性别', 'tinyint', 'int', 'Gender', 'gender', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.644', '2023-09-30 14:05:22.644', NULL);
INSERT INTO `gen_columns` VALUES (307, 23, 'age', '年龄', 'tinyint unsigned', 'int', 'Age', 'age', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.647', '2023-09-30 14:05:22.647', NULL);
INSERT INTO `gen_columns` VALUES (308, 23, 'birthday', '生日', 'bigint', 'int', 'Birthday', 'birthday', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.650', '2023-09-30 14:05:22.650', NULL);
INSERT INTO `gen_columns` VALUES (309, 23, 'source', '来源', 'varchar(64)', 'string', 'Source', 'source', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.654', '2023-09-30 14:05:22.654', NULL);
INSERT INTO `gen_columns` VALUES (310, 23, 'address', '地址', 'varchar(255)', 'string', 'Address', 'address', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.657', '2023-09-30 14:05:22.657', NULL);
INSERT INTO `gen_columns` VALUES (311, 23, 'remark', '描述', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.660', '2023-09-30 14:05:22.660', NULL);
INSERT INTO `gen_columns` VALUES (312, 23, 'user_id', '用户id', 'int unsigned', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.663', '2023-09-30 14:05:22.663', NULL);
INSERT INTO `gen_columns` VALUES (313, 23, 'team_id', '团队id', 'int unsigned', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.667', '2023-09-30 14:05:22.667', NULL);
INSERT INTO `gen_columns` VALUES (314, 23, 'dept_path', '部门路径', 'varchar(255)', 'string', 'DeptPath', 'deptPath', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.670', '2023-09-30 14:05:22.670', NULL);
INSERT INTO `gen_columns` VALUES (315, 23, 'inviter', '邀请人', 'int unsigned', 'int', 'Inviter', 'inviter', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.674', '2023-09-30 14:05:22.674', NULL);
INSERT INTO `gen_columns` VALUES (316, 23, 'inviter_name', '邀请人名', 'varchar(32)', 'string', 'InviterName', 'inviterName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.677', '2023-09-30 14:05:22.677', NULL);
INSERT INTO `gen_columns` VALUES (317, 23, 'created_at', '创建时间', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.680', '2023-09-30 14:05:22.680', NULL);
INSERT INTO `gen_columns` VALUES (318, 23, 'updated_at', '更新时间', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:22.683', '2023-09-30 14:05:22.683', NULL);
INSERT INTO `gen_columns` VALUES (332, 25, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:37.224', '2023-09-30 14:05:37.224', NULL);
INSERT INTO `gen_columns` VALUES (333, 25, 'day', '天', 'int unsigned', 'int', 'Day', 'day', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:37.250', '2023-09-30 14:05:37.250', NULL);
INSERT INTO `gen_columns` VALUES (334, 25, 'team_id', '团队id', 'int unsigned', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:37.257', '2023-09-30 14:05:37.257', NULL);
INSERT INTO `gen_columns` VALUES (335, 25, 'user_id', '用户id', 'int unsigned', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:37.261', '2023-09-30 14:05:37.261', NULL);
INSERT INTO `gen_columns` VALUES (336, 25, 'dept_path', '部门路径', 'varchar(255)', 'string', 'DeptPath', 'deptPath', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:37.265', '2023-09-30 14:05:37.265', NULL);
INSERT INTO `gen_columns` VALUES (337, 25, 'summary', '今日总结', 'text', 'string', 'Summary', 'summary', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:37.269', '2023-09-30 14:05:37.269', NULL);
INSERT INTO `gen_columns` VALUES (338, 25, 'plan', '明日计划', 'text', 'string', 'Plan', 'plan', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:37.272', '2023-09-30 14:05:37.272', NULL);
INSERT INTO `gen_columns` VALUES (339, 25, 'created_at', '创建时间', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:37.276', '2023-09-30 14:05:37.276', NULL);
INSERT INTO `gen_columns` VALUES (340, 25, 'updated_at', '更新时间', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-09-30 14:05:37.279', '2023-09-30 14:05:37.279', NULL);
INSERT INTO `gen_columns` VALUES (349, 27, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.384', '2023-10-03 16:12:23.384', NULL);
INSERT INTO `gen_columns` VALUES (350, 27, 'parent_id', '父id', 'int unsigned', 'int', 'ParentId', 'parentId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.390', '2023-10-03 16:12:23.390', NULL);
INSERT INTO `gen_columns` VALUES (351, 27, 'dept_path', '部门路径', 'varchar(255)', 'string', 'DeptPath', 'deptPath', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.399', '2023-10-03 16:12:23.399', NULL);
INSERT INTO `gen_columns` VALUES (352, 27, 'name', '部门名', 'varchar(128)', 'string', 'Name', 'name', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.403', '2023-10-03 16:12:23.403', NULL);
INSERT INTO `gen_columns` VALUES (353, 27, 'type', '类型', 'tinyint', 'int', 'Type', 'type', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.407', '2023-10-03 16:12:23.407', NULL);
INSERT INTO `gen_columns` VALUES (354, 27, 'principal', '部门领导', 'varchar(128)', 'string', 'Principal', 'principal', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.411', '2023-10-03 16:12:23.411', NULL);
INSERT INTO `gen_columns` VALUES (355, 27, 'phone', '手机号', 'varchar(11)', 'string', 'Phone', 'phone', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.414', '2023-10-03 16:12:23.414', NULL);
INSERT INTO `gen_columns` VALUES (356, 27, 'email', '邮箱', 'varchar(128)', 'string', 'Email', 'email', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.418', '2023-10-03 16:12:23.418', NULL);
INSERT INTO `gen_columns` VALUES (357, 27, 'sort', '排序', 'tinyint', 'int', 'Sort', 'sort', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.422', '2023-10-03 16:12:23.422', NULL);
INSERT INTO `gen_columns` VALUES (358, 27, 'status', '状态', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.425', '2023-10-03 16:12:23.425', NULL);
INSERT INTO `gen_columns` VALUES (359, 27, 'remark', '备注', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.429', '2023-10-03 16:12:23.429', NULL);
INSERT INTO `gen_columns` VALUES (360, 27, 'team_id', '团队id', 'int', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.433', '2023-10-03 16:12:23.433', NULL);
INSERT INTO `gen_columns` VALUES (361, 27, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.436', '2023-10-03 16:12:23.436', NULL);
INSERT INTO `gen_columns` VALUES (362, 27, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.440', '2023-10-03 16:12:23.440', NULL);
INSERT INTO `gen_columns` VALUES (363, 27, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.443', '2023-10-03 16:12:23.443', NULL);
INSERT INTO `gen_columns` VALUES (364, 27, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.448', '2023-10-03 16:12:23.448', NULL);
INSERT INTO `gen_columns` VALUES (365, 27, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 16:12:23.452', '2023-10-03 16:12:23.452', NULL);
INSERT INTO `gen_columns` VALUES (366, 28, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.076', '2023-10-03 17:15:17.076', NULL);
INSERT INTO `gen_columns` VALUES (367, 28, 'day_type', '时间类型:月 30,周 7', 'tinyint unsigned', 'int', 'DayType', 'dayType', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.083', '2023-10-03 17:15:17.083', NULL);
INSERT INTO `gen_columns` VALUES (368, 28, 'day', '时间:202310', 'int unsigned', 'int', 'Day', 'day', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.088', '2023-10-03 17:15:17.088', NULL);
INSERT INTO `gen_columns` VALUES (369, 28, 'team_id', '团队id', 'int unsigned', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.092', '2023-10-03 17:15:17.092', NULL);
INSERT INTO `gen_columns` VALUES (370, 28, 'user_id', '用户id', 'int unsigned', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.097', '2023-10-03 17:15:17.097', NULL);
INSERT INTO `gen_columns` VALUES (371, 28, 'dept_path', '部门路径', 'varchar(255)', 'string', 'DeptPath', 'deptPath', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.100', '2023-10-03 17:15:17.100', NULL);
INSERT INTO `gen_columns` VALUES (372, 28, 'new_customer_cnt', '留存任务', 'int unsigned', 'int', 'NewCustomerCnt', 'newCustomerCnt', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.104', '2023-10-03 17:15:17.104', NULL);
INSERT INTO `gen_columns` VALUES (373, 28, 'first_diagnosis', '导诊任务', 'int unsigned', 'int', 'FirstDiagnosis', 'firstDiagnosis', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.108', '2023-10-03 17:15:17.108', NULL);
INSERT INTO `gen_columns` VALUES (374, 28, 'deal', '成交任务', 'int unsigned', 'int', 'Deal', 'deal', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.112', '2023-10-03 17:15:17.112', NULL);
INSERT INTO `gen_columns` VALUES (375, 28, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.117', '2023-10-03 17:15:17.117', NULL);
INSERT INTO `gen_columns` VALUES (376, 28, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.120', '2023-10-03 17:15:17.120', NULL);
INSERT INTO `gen_columns` VALUES (377, 28, 'created_at', '创建时间', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.124', '2023-10-03 17:15:17.124', NULL);
INSERT INTO `gen_columns` VALUES (378, 28, 'updated_at', '更新时间', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-03 17:15:17.128', '2023-10-03 17:15:17.128', NULL);
INSERT INTO `gen_columns` VALUES (379, 29, 'id', '主键', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.670', '2023-10-06 20:17:08.670', NULL);
INSERT INTO `gen_columns` VALUES (380, 29, 'day', '时间', 'int unsigned', 'int', 'Day', 'day', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.676', '2023-10-06 20:17:08.676', NULL);
INSERT INTO `gen_columns` VALUES (381, 29, 'team_id', '团队id', 'int unsigned', 'int', 'TeamId', 'teamId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.681', '2023-10-06 20:17:08.681', NULL);
INSERT INTO `gen_columns` VALUES (382, 29, 'user_id', '用户id', 'int unsigned', 'int', 'UserId', 'userId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.685', '2023-10-06 20:17:08.685', NULL);
INSERT INTO `gen_columns` VALUES (383, 29, 'dept_path', '部门路径', 'varchar(255)', 'string', 'DeptPath', 'deptPath', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.689', '2023-10-06 20:17:08.689', NULL);
INSERT INTO `gen_columns` VALUES (384, 29, 'new_customer_cnt', '留存', 'int unsigned', 'int', 'NewCustomerCnt', 'newCustomerCnt', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.693', '2023-10-06 20:17:08.693', NULL);
INSERT INTO `gen_columns` VALUES (385, 29, 'first_diagnosis', '初诊', 'int unsigned', 'int', 'FirstDiagnosis', 'firstDiagnosis', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.697', '2023-10-06 20:17:08.697', NULL);
INSERT INTO `gen_columns` VALUES (386, 29, 'further_diagnosis', '复诊', 'int unsigned', 'int', 'FurtherDiagnosis', 'furtherDiagnosis', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.700', '2023-10-06 20:17:08.700', NULL);
INSERT INTO `gen_columns` VALUES (387, 29, 'deal', '成交', 'int unsigned', 'int', 'Deal', 'deal', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.705', '2023-10-06 20:17:08.705', NULL);
INSERT INTO `gen_columns` VALUES (388, 29, 'invitation', '明日邀约', 'int unsigned', 'int', 'Invitation', 'invitation', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.708', '2023-10-06 20:17:08.708', NULL);
INSERT INTO `gen_columns` VALUES (389, 29, 'rest', '休息', 'tinyint', 'int', 'Rest', 'rest', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.711', '2023-10-06 20:17:08.711', NULL);
INSERT INTO `gen_columns` VALUES (390, 29, 'created_at', '创建时间', 'datetime', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.716', '2023-10-06 20:17:08.716', NULL);
INSERT INTO `gen_columns` VALUES (391, 29, 'updated_at', '更新时间', 'datetime', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.720', '2023-10-06 20:17:08.720', NULL);
INSERT INTO `gen_columns` VALUES (392, 29, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.724', '2023-10-06 20:17:08.724', NULL);
INSERT INTO `gen_columns` VALUES (393, 29, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-06 20:17:08.728', '2023-10-06 20:17:08.728', NULL);
INSERT INTO `gen_columns` VALUES (394, 30, 'table_id', '', 'bigint', 'int', 'TableId', 'tableId', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.943', '2023-10-28 16:27:31.943', NULL);
INSERT INTO `gen_columns` VALUES (395, 30, 'db_name', '', 'varchar(64)', 'string', 'DbName', 'dbName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.947', '2023-10-28 16:27:31.947', NULL);
INSERT INTO `gen_columns` VALUES (396, 30, 'table_name', '', 'varchar(128)', 'string', 'TableName', 'tableName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.951', '2023-10-28 16:27:31.951', NULL);
INSERT INTO `gen_columns` VALUES (397, 30, 'table_comment', '', 'varchar(128)', 'string', 'TableComment', 'tableComment', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.955', '2023-10-28 16:27:31.955', NULL);
INSERT INTO `gen_columns` VALUES (398, 30, 'class_name', '', 'varchar(128)', 'string', 'ClassName', 'className', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.958', '2023-10-28 16:27:31.958', NULL);
INSERT INTO `gen_columns` VALUES (399, 30, 'tpl_category', '', 'varchar(128)', 'string', 'TplCategory', 'tplCategory', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.962', '2023-10-28 16:27:31.962', NULL);
INSERT INTO `gen_columns` VALUES (400, 30, 'package_name', '', 'varchar(128)', 'string', 'PackageName', 'packageName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.965', '2023-10-28 16:27:31.965', NULL);
INSERT INTO `gen_columns` VALUES (401, 30, 'module_name', '', 'varchar(128)', 'string', 'ModuleName', 'moduleName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.969', '2023-10-28 16:27:31.969', NULL);
INSERT INTO `gen_columns` VALUES (402, 30, 'module_front_name', '前端文件名', 'varchar(255)', 'string', 'ModuleFrontName', 'moduleFrontName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.971', '2023-10-28 16:27:31.971', NULL);
INSERT INTO `gen_columns` VALUES (403, 30, 'business_name', '', 'varchar(255)', 'string', 'BusinessName', 'businessName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.974', '2023-10-28 16:27:31.974', NULL);
INSERT INTO `gen_columns` VALUES (404, 30, 'function_name', '', 'varchar(255)', 'string', 'FunctionName', 'functionName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.977', '2023-10-28 16:27:31.977', NULL);
INSERT INTO `gen_columns` VALUES (405, 30, 'function_author', '', 'varchar(255)', 'string', 'FunctionAuthor', 'functionAuthor', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.981', '2023-10-28 16:27:31.981', NULL);
INSERT INTO `gen_columns` VALUES (406, 30, 'pk_column', '', 'varchar(255)', 'string', 'PkColumn', 'pkColumn', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.985', '2023-10-28 16:27:31.985', NULL);
INSERT INTO `gen_columns` VALUES (407, 30, 'pk_go_field', '', 'varchar(255)', 'string', 'PkGoField', 'pkGoField', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.988', '2023-10-28 16:27:31.988', NULL);
INSERT INTO `gen_columns` VALUES (408, 30, 'pk_json_field', '', 'varchar(255)', 'string', 'PkJsonField', 'pkJsonField', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.990', '2023-10-28 16:27:31.990', NULL);
INSERT INTO `gen_columns` VALUES (409, 30, 'options', '', 'varchar(255)', 'string', 'Options', 'options', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.994', '2023-10-28 16:27:31.994', NULL);
INSERT INTO `gen_columns` VALUES (410, 30, 'tree_code', '', 'varchar(255)', 'string', 'TreeCode', 'treeCode', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:31.997', '2023-10-28 16:27:31.997', NULL);
INSERT INTO `gen_columns` VALUES (411, 30, 'tree_parent_code', '', 'varchar(255)', 'string', 'TreeParentCode', 'treeParentCode', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.000', '2023-10-28 16:27:32.000', NULL);
INSERT INTO `gen_columns` VALUES (412, 30, 'tree_name', '', 'varchar(255)', 'string', 'TreeName', 'treeName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 19, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.003', '2023-10-28 16:27:32.003', NULL);
INSERT INTO `gen_columns` VALUES (413, 30, 'tree', '', 'tinyint(1)', 'int', 'Tree', 'tree', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 20, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.005', '2023-10-28 16:27:32.005', NULL);
INSERT INTO `gen_columns` VALUES (414, 30, 'crud', '', 'tinyint(1)', 'int', 'Crud', 'crud', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 21, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.008', '2023-10-28 16:27:32.008', NULL);
INSERT INTO `gen_columns` VALUES (415, 30, 'remark', '', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 22, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.012', '2023-10-28 16:27:32.012', NULL);
INSERT INTO `gen_columns` VALUES (416, 30, 'is_data_scope', '', 'tinyint', 'int', 'IsDataScope', 'isDataScope', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 23, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.015', '2023-10-28 16:27:32.015', NULL);
INSERT INTO `gen_columns` VALUES (417, 30, 'is_actions', '', 'tinyint', 'int', 'IsActions', 'isActions', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 24, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.018', '2023-10-28 16:27:32.018', NULL);
INSERT INTO `gen_columns` VALUES (418, 30, 'is_auth', '', 'tinyint', 'int', 'IsAuth', 'isAuth', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 25, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.021', '2023-10-28 16:27:32.021', NULL);
INSERT INTO `gen_columns` VALUES (419, 30, 'is_logical_delete', '', 'varchar(1)', 'string', 'IsLogicalDelete', 'isLogicalDelete', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 26, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.024', '2023-10-28 16:27:32.024', NULL);
INSERT INTO `gen_columns` VALUES (420, 30, 'logical_delete', '', 'tinyint(1)', 'int', 'LogicalDelete', 'logicalDelete', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 27, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.027', '2023-10-28 16:27:32.027', NULL);
INSERT INTO `gen_columns` VALUES (421, 30, 'logical_delete_column', '', 'varchar(128)', 'string', 'LogicalDeleteColumn', 'logicalDeleteColumn', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 28, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.031', '2023-10-28 16:27:32.031', NULL);
INSERT INTO `gen_columns` VALUES (422, 30, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 29, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.034', '2023-10-28 16:27:32.034', NULL);
INSERT INTO `gen_columns` VALUES (423, 30, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 30, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.036', '2023-10-28 16:27:32.036', NULL);
INSERT INTO `gen_columns` VALUES (424, 30, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 31, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.039', '2023-10-28 16:27:32.039', NULL);
INSERT INTO `gen_columns` VALUES (425, 30, 'create_by', '创建者', 'int unsigned', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 32, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.042', '2023-10-28 16:27:32.042', NULL);
INSERT INTO `gen_columns` VALUES (426, 30, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 33, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:27:32.045', '2023-10-28 16:27:32.045', NULL);
INSERT INTO `gen_columns` VALUES (427, 31, 'column_id', '', 'bigint', 'int', 'ColumnId', 'columnId', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.931', '2023-10-28 16:28:02.931', NULL);
INSERT INTO `gen_columns` VALUES (428, 31, 'table_id', '', 'bigint', 'int', 'TableId', 'tableId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.935', '2023-10-28 16:28:02.935', NULL);
INSERT INTO `gen_columns` VALUES (429, 31, 'column_name', '', 'varchar(128)', 'string', 'ColumnName', 'columnName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.938', '2023-10-28 16:28:02.938', NULL);
INSERT INTO `gen_columns` VALUES (430, 31, 'column_comment', '', 'varchar(128)', 'string', 'ColumnComment', 'columnComment', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.941', '2023-10-28 16:28:02.941', NULL);
INSERT INTO `gen_columns` VALUES (431, 31, 'column_type', '', 'varchar(128)', 'string', 'ColumnType', 'columnType', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.946', '2023-10-28 16:28:02.946', NULL);
INSERT INTO `gen_columns` VALUES (432, 31, 'go_type', '', 'varchar(128)', 'string', 'GoType', 'goType', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.949', '2023-10-28 16:28:02.949', NULL);
INSERT INTO `gen_columns` VALUES (433, 31, 'go_field', '', 'varchar(128)', 'string', 'GoField', 'goField', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.952', '2023-10-28 16:28:02.952', NULL);
INSERT INTO `gen_columns` VALUES (434, 31, 'json_field', '', 'varchar(128)', 'string', 'JsonField', 'jsonField', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.954', '2023-10-28 16:28:02.954', NULL);
INSERT INTO `gen_columns` VALUES (435, 31, 'is_pk', '', 'varchar(4)', 'string', 'IsPk', 'isPk', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 9, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.957', '2023-10-28 16:28:02.957', NULL);
INSERT INTO `gen_columns` VALUES (436, 31, 'is_increment', '', 'varchar(4)', 'string', 'IsIncrement', 'isIncrement', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 10, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.961', '2023-10-28 16:28:02.961', NULL);
INSERT INTO `gen_columns` VALUES (437, 31, 'is_required', '', 'varchar(4)', 'string', 'IsRequired', 'isRequired', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 11, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.964', '2023-10-28 16:28:02.964', NULL);
INSERT INTO `gen_columns` VALUES (438, 31, 'is_insert', '', 'varchar(4)', 'string', 'IsInsert', 'isInsert', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 12, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.967', '2023-10-28 16:28:02.967', NULL);
INSERT INTO `gen_columns` VALUES (439, 31, 'is_edit', '', 'varchar(4)', 'string', 'IsEdit', 'isEdit', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 13, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.969', '2023-10-28 16:28:02.969', NULL);
INSERT INTO `gen_columns` VALUES (440, 31, 'is_list', '', 'varchar(4)', 'string', 'IsList', 'isList', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 14, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.972', '2023-10-28 16:28:02.972', NULL);
INSERT INTO `gen_columns` VALUES (441, 31, 'is_query', '', 'varchar(4)', 'string', 'IsQuery', 'isQuery', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 15, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.975', '2023-10-28 16:28:02.975', NULL);
INSERT INTO `gen_columns` VALUES (442, 31, 'query_type', '', 'varchar(128)', 'string', 'QueryType', 'queryType', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 16, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.978', '2023-10-28 16:28:02.978', NULL);
INSERT INTO `gen_columns` VALUES (443, 31, 'html_type', '', 'varchar(128)', 'string', 'HtmlType', 'htmlType', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 17, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.981', '2023-10-28 16:28:02.981', NULL);
INSERT INTO `gen_columns` VALUES (444, 31, 'dict_type', '', 'varchar(128)', 'string', 'DictType', 'dictType', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 18, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.985', '2023-10-28 16:28:02.985', NULL);
INSERT INTO `gen_columns` VALUES (445, 31, 'sort', '', 'bigint', 'int', 'Sort', 'sort', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 19, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.987', '2023-10-28 16:28:02.987', NULL);
INSERT INTO `gen_columns` VALUES (446, 31, 'list', '', 'varchar(1)', 'string', 'List', 'list', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 20, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.990', '2023-10-28 16:28:02.990', NULL);
INSERT INTO `gen_columns` VALUES (447, 31, 'pk', '', 'tinyint(1)', 'int', 'Pk', 'pk', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 21, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.993', '2023-10-28 16:28:02.993', NULL);
INSERT INTO `gen_columns` VALUES (448, 31, 'required', '', 'tinyint(1)', 'int', 'Required', 'required', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 22, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.996', '2023-10-28 16:28:02.996', NULL);
INSERT INTO `gen_columns` VALUES (449, 31, 'super_column', '', 'tinyint(1)', 'int', 'SuperColumn', 'superColumn', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 23, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:02.999', '2023-10-28 16:28:02.999', NULL);
INSERT INTO `gen_columns` VALUES (450, 31, 'usable_column', '', 'tinyint(1)', 'int', 'UsableColumn', 'usableColumn', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 24, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.002', '2023-10-28 16:28:03.002', NULL);
INSERT INTO `gen_columns` VALUES (451, 31, 'increment', '', 'tinyint(1)', 'int', 'Increment', 'increment', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 25, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.004', '2023-10-28 16:28:03.004', NULL);
INSERT INTO `gen_columns` VALUES (452, 31, 'insert', '', 'tinyint(1)', 'int', 'Insert', 'insert', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 26, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.007', '2023-10-28 16:28:03.007', NULL);
INSERT INTO `gen_columns` VALUES (453, 31, 'edit', '', 'tinyint(1)', 'int', 'Edit', 'edit', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 27, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.011', '2023-10-28 16:28:03.011', NULL);
INSERT INTO `gen_columns` VALUES (454, 31, 'query', '', 'tinyint(1)', 'int', 'Query', 'query', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 28, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.014', '2023-10-28 16:28:03.014', NULL);
INSERT INTO `gen_columns` VALUES (455, 31, 'remark', '', 'varchar(255)', 'string', 'Remark', 'remark', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 29, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.017', '2023-10-28 16:28:03.017', NULL);
INSERT INTO `gen_columns` VALUES (456, 31, 'fk_table_name', '', 'longtext', 'string', 'FkTableName', 'fkTableName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 30, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.019', '2023-10-28 16:28:03.019', NULL);
INSERT INTO `gen_columns` VALUES (457, 31, 'fk_table_name_class', '', 'longtext', 'string', 'FkTableNameClass', 'fkTableNameClass', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 31, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.022', '2023-10-28 16:28:03.022', NULL);
INSERT INTO `gen_columns` VALUES (458, 31, 'fk_table_name_package', '', 'longtext', 'string', 'FkTableNamePackage', 'fkTableNamePackage', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 32, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.026', '2023-10-28 16:28:03.026', NULL);
INSERT INTO `gen_columns` VALUES (459, 31, 'fk_label_id', '', 'longtext', 'string', 'FkLabelId', 'fkLabelId', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 33, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.029', '2023-10-28 16:28:03.029', NULL);
INSERT INTO `gen_columns` VALUES (460, 31, 'fk_label_name', '', 'varchar(255)', 'string', 'FkLabelName', 'fkLabelName', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 34, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.032', '2023-10-28 16:28:03.032', NULL);
INSERT INTO `gen_columns` VALUES (461, 31, 'create_by', '', 'mediumint', 'int', 'CreateBy', 'createBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 35, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.035', '2023-10-28 16:28:03.035', NULL);
INSERT INTO `gen_columns` VALUES (462, 31, 'update_By', '', 'mediumint', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 36, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.037', '2023-10-28 16:28:03.037', NULL);
INSERT INTO `gen_columns` VALUES (463, 31, 'created_at', '创建时间', 'datetime(3)', 'time.Time', 'CreatedAt', 'createdAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 37, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.040', '2023-10-28 16:28:03.040', NULL);
INSERT INTO `gen_columns` VALUES (464, 31, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 38, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.044', '2023-10-28 16:28:03.044', NULL);
INSERT INTO `gen_columns` VALUES (465, 31, 'deleted_at', '删除时间', 'datetime(3)', 'time.Time', 'DeletedAt', 'deletedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 39, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:28:03.047', '2023-10-28 16:28:03.047', NULL);
INSERT INTO `gen_columns` VALUES (466, 32, 'id', '主键编码', 'int unsigned', 'int', 'Id', 'id', '1', '', '1', '1', '1', '1', '', 'EQ', 'input', '', 1, '', 1, 1, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:51:28.750', '2023-10-28 16:51:28.750', NULL);
INSERT INTO `gen_columns` VALUES (467, 32, 'title', '标题', 'varchar(128)', 'string', 'Title', 'title', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 2, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:51:28.754', '2023-10-28 16:51:28.754', NULL);
INSERT INTO `gen_columns` VALUES (468, 32, 'method', '请求类型', 'varchar(16)', 'string', 'Method', 'method', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 3, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:51:28.756', '2023-10-28 16:51:28.756', NULL);
INSERT INTO `gen_columns` VALUES (469, 32, 'path', '请求地址', 'varchar(128)', 'string', 'Path', 'path', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 4, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:51:28.758', '2023-10-28 16:51:28.758', NULL);
INSERT INTO `gen_columns` VALUES (470, 32, 'perm_type', '权限类型（1：无需认证 2:须token 3：须鉴权）', 'bigint', 'int', 'PermType', 'permType', '0', '', '0', '1', '1', '1', '', 'EQ', 'input', '', 5, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:51:28.761', '2023-10-28 16:51:28.761', NULL);
INSERT INTO `gen_columns` VALUES (471, 32, 'status', '状态 3 DEF 2 OK 1 del', 'tinyint', 'int', 'Status', 'status', '0', '', '0', '1', '1', '1', '1', 'EQ', 'input', '', 6, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:51:28.765', '2023-10-28 16:51:28.765', NULL);
INSERT INTO `gen_columns` VALUES (472, 32, 'update_by', '更新者', 'int unsigned', 'int', 'UpdateBy', 'updateBy', '0', '', '0', '1', '', '1', '', 'EQ', 'input', '', 7, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:51:28.768', '2023-10-28 16:51:28.768', NULL);
INSERT INTO `gen_columns` VALUES (473, 32, 'updated_at', '最后更新时间', 'datetime(3)', 'time.Time', 'UpdatedAt', 'updatedAt', '0', '', '0', '1', '', '1', '', 'EQ', 'datetime', '', 8, '', 0, 0, 0, 0, 0, 1, 0, 0, '', '', '', '', '', '', 0, 0, '2023-10-28 16:51:28.770', '2023-10-28 16:51:28.770', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of gen_tables
-- ----------------------------
INSERT INTO `gen_tables` VALUES (14, 'dilu-db', 'sys_user', '用户', 'SysUser', 'crud', 'sys', 'sys-user', '', 'sysUser', '用户', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-26 13:44:05.983', '2023-09-26 13:44:05.983', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (15, 'dilu-db', 'sys_menu', '菜单', 'SysMenu', 'crud', 'sys', 'sys-menu', '', 'sysMenu', '菜单', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-26 13:44:51.533', '2023-09-26 13:44:51.533', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (16, 'dilu-db', 'sys_role', '角色', 'SysRole', 'crud', 'sys', 'sys-role', '', 'sysRole', '角色', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-26 13:45:11.897', '2023-09-26 13:45:11.897', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (20, 'dilu-db', 'sys_team', '团队', 'SysTeam', 'crud', 'sys', 'sys-team', '', 'sysTeam', 'SysTeam', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-29 08:33:06.345', '2023-09-29 08:33:06.345', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (21, 'dilu-db', 'sys_member', '会员', 'SysMember', 'crud', 'sys', 'sys-member', '', 'sysMember', 'SysMember', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-29 08:33:20.152', '2023-09-29 08:33:20.152', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (22, 'dental-db', 'bill', '账单', 'Bill', 'crud', 'dental', 'bill', '', 'bill', 'Bill', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-30 14:05:14.853', '2023-09-30 14:05:14.853', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (23, 'dental-db', 'customer', '客户', 'Customer', 'crud', 'dental', 'customer', '', 'customer', 'Customer', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-30 14:05:22.598', '2023-09-30 14:05:22.598', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (25, 'dental-db', 'summary_plan_day', '总结与计划', 'SummaryPlanDay', 'crud', 'dental', 'summary-plan-day', '', 'summaryPlanDay', 'SummaryPlanDay', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-09-30 14:05:37.221', '2023-09-30 14:05:37.221', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (27, 'dilu-db', 'sys_dept', '部门', 'SysDept', 'crud', 'sys', 'sys-dept', '', 'sysDept', '部门', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-10-03 16:12:23.378', '2023-10-03 16:12:23.378', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (28, 'dental-db', 'target_task', 'TargetTask', 'TargetTask', 'crud', 'dental', 'target-task', '', 'targetTask', 'TargetTask', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-10-03 17:15:17.051', '2023-10-03 17:15:17.051', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (29, 'dental-db', 'event_day_st', 'EventDaySt', 'EventDaySt', 'crud', 'dental', 'event-day-st', '', 'eventDaySt', 'EventDaySt', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-10-06 20:17:08.646', '2023-10-06 20:17:08.646', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (30, 'dilu-db', 'gen_tables', 'GenTables', 'GenTables', 'crud', 'sys', 'gen-tables', '', 'genTables', 'GenTables', 'baowk', 'table_id', 'TableId', 'tableId', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-10-28 16:27:31.937', '2023-10-28 16:27:31.937', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (31, 'dilu-db', 'gen_columns', 'GenColumns', 'GenColumns', 'crud', 'sys', 'gen-columns', '', 'genColumns', 'GenColumns', 'baowk', 'column_id', 'ColumnId', 'columnId', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-10-28 16:28:02.925', '2023-10-28 16:28:02.925', NULL, 0, 0);
INSERT INTO `gen_tables` VALUES (32, 'dilu-db', 'sys_api', '接口', 'SysApi', 'crud', 'sys', 'sys-api', '', 'sysApi', '接口', 'baowk', 'id', 'Id', 'id', '', '', '', '', 0, 1, '', 1, 2, 1, '1', 1, 'is_del', '2023-10-28 16:51:28.742', '2023-10-28 16:51:28.742', NULL, 0, 0);

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `title` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '标题',
  `method` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求类型',
  `path` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '请求地址',
  `perm_type` bigint(0) NULL DEFAULT NULL COMMENT '权限类型（1：无需认证 2:须token 3：须鉴权）',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态 3 DEF 2 OK 1 del',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_method_path`(`method`, `path`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 80 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '接口' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_api
-- ----------------------------
INSERT INTO `sys_api` VALUES (1, '分页获取用户', 'POST', '/api/v1/sys/sys-user/page', 3, 3, 0, '2023-09-26 13:46:59.488');
INSERT INTO `sys_api` VALUES (2, '根据id获取用户', 'POST', '/api/v1/sys/sys-user/get', 3, 3, 0, '2023-09-26 13:46:59.518');
INSERT INTO `sys_api` VALUES (3, '创建用户', 'POST', '/api/v1/sys/sys-user/create', 3, 3, 0, '2023-09-26 13:46:59.532');
INSERT INTO `sys_api` VALUES (4, '修改用户', 'POST', '/api/v1/sys/sys-user/update', 3, 3, 0, '2023-09-26 13:46:59.539');
INSERT INTO `sys_api` VALUES (5, '删除用户', 'POST', '/api/v1/sys/sys-user/del', 3, 3, 0, '2023-09-26 13:46:59.550');
INSERT INTO `sys_api` VALUES (6, '分页获取菜单', 'POST', '/api/v1/sys/sys-menu/page', 3, 3, 0, '2023-09-26 13:47:37.020');
INSERT INTO `sys_api` VALUES (7, '根据id获取菜单', 'POST', '/api/v1/sys/sys-menu/get', 3, 3, 0, '2023-09-26 13:47:37.038');
INSERT INTO `sys_api` VALUES (8, '创建菜单', 'POST', '/api/v1/sys/sys-menu/create', 3, 3, 0, '2023-09-26 13:47:37.064');
INSERT INTO `sys_api` VALUES (9, '修改菜单', 'POST', '/api/v1/sys/sys-menu/update', 3, 3, 0, '2023-09-26 13:47:37.073');
INSERT INTO `sys_api` VALUES (10, '删除菜单', 'POST', '/api/v1/sys/sys-menu/del', 3, 3, 0, '2023-09-26 13:47:37.082');
INSERT INTO `sys_api` VALUES (11, '分页获取角色', 'POST', '/api/v1/sys/sys-role/page', 3, 3, 0, '2023-09-26 13:47:39.685');
INSERT INTO `sys_api` VALUES (12, '根据id获取角色', 'POST', '/api/v1/sys/sys-role/get', 3, 3, 0, '2023-09-26 13:47:39.701');
INSERT INTO `sys_api` VALUES (13, '创建角色', 'POST', '/api/v1/sys/sys-role/create', 3, 3, 0, '2023-09-26 13:47:39.734');
INSERT INTO `sys_api` VALUES (14, '修改角色', 'POST', '/api/v1/sys/sys-role/update', 3, 3, 0, '2023-09-26 13:47:39.762');
INSERT INTO `sys_api` VALUES (15, '删除角色', 'POST', '/api/v1/sys/sys-role/del', 3, 3, 0, '2023-09-26 13:47:39.773');
INSERT INTO `sys_api` VALUES (16, '分页获取部门', 'POST', '/api/v1/sys/sys-dept/all', 3, 3, 0, '2023-09-26 13:47:42.136');
INSERT INTO `sys_api` VALUES (17, '根据id获取部门', 'POST', '/api/v1/sys/sys-dept/get', 3, 3, 0, '2023-09-26 13:47:42.152');
INSERT INTO `sys_api` VALUES (18, '创建部门', 'POST', '/api/v1/sys/sys-dept/create', 3, 3, 0, '2023-09-26 13:47:42.172');
INSERT INTO `sys_api` VALUES (19, '修改部门', 'POST', '/api/v1/sys/sys-dept/update', 3, 3, 0, '2023-09-26 13:47:42.186');
INSERT INTO `sys_api` VALUES (20, '删除部门', 'POST', '/api/v1/sys/sys-dept/del', 3, 3, 0, '2023-09-26 13:47:42.197');
INSERT INTO `sys_api` VALUES (21, '分页获取用户部门', 'POST', '/api/v1/sys/sys-user-dept/page', 3, 3, 0, '2023-09-26 13:47:45.447');
INSERT INTO `sys_api` VALUES (22, '根据id获取用户部门', 'POST', '/api/v1/sys/sys-user-dept/get', 3, 3, 0, '2023-09-26 13:47:45.459');
INSERT INTO `sys_api` VALUES (23, '创建用户部门', 'POST', '/api/v1/sys/sys-user-dept/create', 3, 3, 0, '2023-09-26 13:47:45.483');
INSERT INTO `sys_api` VALUES (24, '修改用户部门', 'POST', '/api/v1/sys/sys-user-dept/update', 3, 3, 0, '2023-09-26 13:47:45.495');
INSERT INTO `sys_api` VALUES (25, '删除用户部门', 'POST', '/api/v1/sys/sys-user-dept/del', 3, 3, 0, '2023-09-26 13:47:45.503');
INSERT INTO `sys_api` VALUES (26, '分页获取成员', 'POST', '/api/v1/sys/sys-member/page', 3, 3, 0, '2023-09-26 14:09:40.443');
INSERT INTO `sys_api` VALUES (27, '根据id获取成员', 'POST', '/api/v1/sys/sys-member/get', 3, 3, 0, '2023-09-26 14:09:40.459');
INSERT INTO `sys_api` VALUES (28, '创建成员', 'POST', '/api/v1/sys/sys-member/create', 3, 3, 0, '2023-09-26 14:09:40.483');
INSERT INTO `sys_api` VALUES (29, '修改成员', 'POST', '/api/v1/sys/sys-member/update', 3, 3, 0, '2023-09-26 14:09:40.493');
INSERT INTO `sys_api` VALUES (30, '删除成员', 'POST', '/api/v1/sys/sys-member/del', 3, 3, 0, '2023-09-26 14:09:40.502');
INSERT INTO `sys_api` VALUES (31, '分页获取团队', 'POST', '/api/v1/sys/sys-team/page', 3, 3, 0, '2023-09-29 08:40:00.840');
INSERT INTO `sys_api` VALUES (32, '根据id获取团队', 'POST', '/api/v1/sys/sys-team/get', 3, 3, 0, '2023-09-29 08:40:00.855');
INSERT INTO `sys_api` VALUES (33, '创建团队', 'POST', '/api/v1/sys/sys-team/create', 3, 3, 0, '2023-09-29 08:40:00.867');
INSERT INTO `sys_api` VALUES (34, '修改团队', 'POST', '/api/v1/sys/sys-team/update', 3, 3, 0, '2023-09-29 08:40:00.879');
INSERT INTO `sys_api` VALUES (35, '删除团队', 'POST', '/api/v1/sys/sys-team/del', 3, 3, 0, '2023-09-29 08:40:00.891');
INSERT INTO `sys_api` VALUES (36, '分页获取账单', 'POST', '/api/v1/dental/bill/page', 3, 3, 0, '2023-09-29 08:45:53.790');
INSERT INTO `sys_api` VALUES (37, '根据id获取账单', 'POST', '/api/v1/dental/bill/get', 3, 3, 0, '2023-09-29 08:45:53.800');
INSERT INTO `sys_api` VALUES (38, '创建账单', 'POST', '/api/v1/dental/bill/create', 3, 3, 0, '2023-09-29 08:45:53.814');
INSERT INTO `sys_api` VALUES (39, '修改账单', 'POST', '/api/v1/dental/bill/update', 3, 3, 0, '2023-09-29 08:45:53.827');
INSERT INTO `sys_api` VALUES (40, '删除账单', 'POST', '/api/v1/dental/bill/del', 3, 3, 0, '2023-09-29 08:45:53.838');
INSERT INTO `sys_api` VALUES (41, '分页获取客户', 'POST', '/api/v1/dental/customer/page', 3, 3, 0, '2023-09-29 08:46:25.136');
INSERT INTO `sys_api` VALUES (42, '根据id获取客户', 'POST', '/api/v1/dental/customer/get', 3, 3, 0, '2023-09-29 08:46:25.149');
INSERT INTO `sys_api` VALUES (43, '创建客户', 'POST', '/api/v1/dental/customer/create', 3, 3, 0, '2023-09-29 08:46:25.162');
INSERT INTO `sys_api` VALUES (44, '修改客户', 'POST', '/api/v1/dental/customer/update', 3, 3, 0, '2023-09-29 08:46:25.173');
INSERT INTO `sys_api` VALUES (45, '删除客户', 'POST', '/api/v1/dental/customer/del', 3, 3, 0, '2023-09-29 08:46:25.184');
INSERT INTO `sys_api` VALUES (46, '分页获取日统计', 'POST', '/api/v1/dental/event-day-st/page', 3, 3, 0, '2023-09-29 08:46:27.734');
INSERT INTO `sys_api` VALUES (47, '根据id获取日统计', 'POST', '/api/v1/dental/event-day-st/get', 3, 3, 0, '2023-09-29 08:46:27.768');
INSERT INTO `sys_api` VALUES (48, '创建日统计', 'POST', '/api/v1/dental/event-day-st/create', 3, 3, 0, '2023-09-29 08:46:27.778');
INSERT INTO `sys_api` VALUES (49, '修改日统计', 'POST', '/api/v1/dental/event-day-st/update', 3, 3, 0, '2023-09-29 08:46:27.787');
INSERT INTO `sys_api` VALUES (50, '删除日统计', 'POST', '/api/v1/dental/event-day-st/del', 3, 3, 0, '2023-09-29 08:46:27.798');
INSERT INTO `sys_api` VALUES (51, '分页获取日总结与计划', 'POST', '/api/v1/dental/summary-plan-day/page', 3, 3, 0, '2023-09-29 08:46:30.205');
INSERT INTO `sys_api` VALUES (52, '根据id获取日总结与计划', 'POST', '/api/v1/dental/summary-plan-day/get', 3, 3, 0, '2023-09-29 08:46:30.219');
INSERT INTO `sys_api` VALUES (53, '创建日总结与计划', 'POST', '/api/v1/dental/summary-plan-day/create', 3, 3, 0, '2023-09-29 08:46:30.230');
INSERT INTO `sys_api` VALUES (54, '修改日总结与计划', 'POST', '/api/v1/dental/summary-plan-day/update', 3, 3, 0, '2023-09-29 08:46:30.241');
INSERT INTO `sys_api` VALUES (55, '删除日总结与计划', 'POST', '/api/v1/dental/summary-plan-day/del', 3, 3, 0, '2023-09-29 08:46:30.252');
INSERT INTO `sys_api` VALUES (56, '分页获取任务目标', 'POST', '/api/v1/dental/target-task/page', 3, 3, 0, '2023-09-29 08:46:32.788');
INSERT INTO `sys_api` VALUES (57, '根据id获取任务目标', 'POST', '/api/v1/dental/target-task/get', 3, 3, 0, '2023-09-29 08:46:32.824');
INSERT INTO `sys_api` VALUES (58, '创建任务目标', 'POST', '/api/v1/dental/target-task/create', 3, 3, 0, '2023-09-29 08:46:32.834');
INSERT INTO `sys_api` VALUES (59, '修改任务目标', 'POST', '/api/v1/dental/target-task/update', 3, 3, 0, '2023-09-29 08:46:32.844');
INSERT INTO `sys_api` VALUES (60, '删除任务目标', 'POST', '/api/v1/dental/target-task/del', 3, 3, 0, '2023-09-29 08:46:32.854');
INSERT INTO `sys_api` VALUES (61, '账单智能识别', 'POST', '/api/v1/dental/bill/identify', 3, 3, 0, '2023-10-15 10:45:32.000');
INSERT INTO `sys_api` VALUES (62, '日统计', 'POST', '/api/v1/dental/st/day', 3, 3, 0, '2023-10-16 13:51:54.000');
INSERT INTO `sys_api` VALUES (63, '月统计', 'POST', '/api/v1/dental/st/month', 3, 3, 0, '2023-10-16 14:30:58.000');
INSERT INTO `sys_api` VALUES (64, '分页获取GenTables', 'POST', '/api/v1/sys/gen-tables/page', 3, 3, 0, '2023-10-28 16:31:47.687');
INSERT INTO `sys_api` VALUES (65, '根据id获取GenTables', 'POST', '/api/v1/sys/gen-tables/get', 3, 3, 0, '2023-10-28 16:31:47.701');
INSERT INTO `sys_api` VALUES (66, '创建GenTables', 'POST', '/api/v1/sys/gen-tables/create', 3, 3, 0, '2023-10-28 16:31:47.710');
INSERT INTO `sys_api` VALUES (67, '修改GenTables', 'POST', '/api/v1/sys/gen-tables/update', 3, 3, 0, '2023-10-28 16:31:47.720');
INSERT INTO `sys_api` VALUES (68, '删除GenTables', 'POST', '/api/v1/sys/gen-tables/del', 3, 3, 0, '2023-10-28 16:31:47.728');
INSERT INTO `sys_api` VALUES (69, '分页获取GenColumns', 'POST', '/api/v1/sys/gen-columns/page', 3, 3, 0, '2023-10-28 16:35:01.328');
INSERT INTO `sys_api` VALUES (70, '根据id获取GenColumns', 'POST', '/api/v1/sys/gen-columns/get', 3, 3, 0, '2023-10-28 16:35:01.350');
INSERT INTO `sys_api` VALUES (71, '创建GenColumns', 'POST', '/api/v1/sys/gen-columns/create', 3, 3, 0, '2023-10-28 16:35:01.360');
INSERT INTO `sys_api` VALUES (72, '修改GenColumns', 'POST', '/api/v1/sys/gen-columns/update', 3, 3, 0, '2023-10-28 16:35:01.368');
INSERT INTO `sys_api` VALUES (73, '删除GenColumns', 'POST', '/api/v1/sys/gen-columns/del', 3, 3, 0, '2023-10-28 16:35:01.388');
INSERT INTO `sys_api` VALUES (74, '分页获取接口', 'POST', '/api/v1/sys/sys-api/page', 3, 3, 0, '2023-10-28 16:51:44.921');
INSERT INTO `sys_api` VALUES (75, '根据id获取接口', 'POST', '/api/v1/sys/sys-api/get', 3, 3, 0, '2023-10-28 16:51:44.933');
INSERT INTO `sys_api` VALUES (76, '创建接口', 'POST', '/api/v1/sys/sys-api/create', 3, 3, 0, '2023-10-28 16:51:44.943');
INSERT INTO `sys_api` VALUES (77, '修改接口', 'POST', '/api/v1/sys/sys-api/update', 3, 3, 0, '2023-10-28 16:51:44.952');
INSERT INTO `sys_api` VALUES (78, '删除接口', 'POST', '/api/v1/sys/sys-api/del', 3, 3, 0, '2023-10-28 16:51:44.960');
INSERT INTO `sys_api` VALUES (79, '账单查询统计', 'POST', '/api/v1/dental/st/query', 3, 3, 0, '2023-10-29 18:56:28.000');

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '配置' ROW_FORMAT = Dynamic;

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
  `type` tinyint(0) NULL DEFAULT NULL COMMENT '类型 1分公司 2部门',
  `principal` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门领导',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '手机号',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '邮箱',
  `sort` tinyint(0) NULL DEFAULT NULL COMMENT '排序',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态 1正常 2关闭',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `team_id` int(0) NULL DEFAULT NULL COMMENT '团队id',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dept_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '部门' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
INSERT INTO `sys_dept` VALUES (1, 0, '/0/1/', '销售一部', 2, '糖糖', '13800138000', NULL, 1, 1, NULL, 1, NULL, 2, NULL, '2023-10-26 20:57:48.321', NULL);
INSERT INTO `sys_dept` VALUES (2, 1, '/0/1/2/', '糖糖组', 2, '', '', NULL, NULL, 1, NULL, 1, NULL, 2, NULL, '2023-10-26 21:03:59.644', NULL);
INSERT INTO `sys_dept` VALUES (9, 10, '/0/10/9/', '哈哈组', 2, '哈哈', '', '', 0, 1, '', 1, 2, 2, '2023-10-26 20:50:45.790', '2023-10-26 20:59:33.924', NULL);
INSERT INTO `sys_dept` VALUES (10, 0, '/0/10/', '销售二部', 0, '', '', '', 0, 1, '', 1, 2, 0, '2023-10-26 20:51:11.338', '2023-10-26 20:51:11.342', NULL);
INSERT INTO `sys_dept` VALUES (11, 0, '/0/11/', '销售三部', 1, '', '', '', 0, 1, '', 1, 2, 2, '2023-10-26 21:00:05.460', '2023-10-26 21:04:04.980', NULL);
INSERT INTO `sys_dept` VALUES (12, 11, '/0/11/12/', '呵呵组', 2, '', '', '', 0, 1, '', 1, 2, 2, '2023-10-26 21:00:22.635', '2023-10-26 21:00:45.968', NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '邮件' ROW_FORMAT = Dynamic;

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '定时任务' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_job
-- ----------------------------

-- ----------------------------
-- Table structure for sys_member
-- ----------------------------
DROP TABLE IF EXISTS `sys_member`;
CREATE TABLE `sys_member`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `team_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '团队id',
  `user_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  `nickname` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '昵称',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '姓名',
  `py` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '姓名拼音',
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '电话',
  `dept_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '部门路径',
  `dept_id` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '部门id',
  `roles` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '角色id',
  `post_id` tinyint(0) UNSIGNED NULL DEFAULT NULL COMMENT '职位 1系统超管 2 团队拥有者 4主管 8副主管 16员工',
  `entry_time` datetime(3) NULL DEFAULT NULL COMMENT '入职时间',
  `retire_time` datetime(3) NULL DEFAULT NULL COMMENT '离职时间',
  `gender` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '2' COMMENT '性别 1男 2女 3未知',
  `birthday` date NULL DEFAULT NULL COMMENT '生日 格式 yyyy-MM-dd',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态 1正常 ',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_member
-- ----------------------------
INSERT INTO `sys_member` VALUES (1, 1, 2, '糖糖', '唐敦霞', NULL, '', '/0/1/2/', 2, '-1', 8, '2023-02-20 00:00:00.000', NULL, '2', NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_member` VALUES (2, 1, 3, '藏春梅', '藏春梅', NULL, NULL, '/0/1/2/', 2, '1', 16, '2021-01-13 00:00:00.000', NULL, '2', NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_member` VALUES (3, 1, 4, '李艳雷', '李艳雷', NULL, NULL, '/0/1/2/', 2, NULL, 16, '2023-06-19 00:00:00.000', NULL, '1', NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_member` VALUES (4, 1, 5, '简小丽', '简小丽', NULL, NULL, '/0/1/2/', 2, NULL, 16, '2022-11-18 00:00:00.000', NULL, '2', NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_member` VALUES (5, 1, 6, '胡珊', '胡珊', NULL, NULL, '/0/1/2/', 2, NULL, 16, '2023-06-02 00:00:00.000', NULL, '2', NULL, 1, NULL, NULL, NULL, NULL);
INSERT INTO `sys_member` VALUES (6, 1, 7, '余鸿雁', '余鸿雁', NULL, NULL, '/0/1/2/', 2, NULL, 16, '2023-08-07 00:00:00.000', NULL, '2', NULL, 1, NULL, NULL, NULL, NULL);

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
  `platform_type` bigint(0) NULL DEFAULT NULL COMMENT '平台类型 1 平台管理 2团队管理',
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
) ENGINE = InnoDB AUTO_INCREMENT = 110 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '菜单' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, '', '系统管理', 'setting', '/sys', 2, 1, '', 0, 0, 'Layout', 0, 0, 1, 0, '2023-09-26 13:46:59.480', '2023-09-26 13:46:59.481', NULL);
INSERT INTO `sys_menu` VALUES (2, 'SysUserManage', '用户管理', 'flUser', '/sys/sys-user', 1, 2, 'sys:sysUser:list', 1, 0, '/sys/sys-user/index', 0, 0, 1, 1, '2023-09-26 13:46:59.493', '2023-09-26 13:46:59.503', NULL);
INSERT INTO `sys_menu` VALUES (3, '', '用户详情', '', 'sys_user_detail', 1, 3, 'sys:sysUser:query', 2, 0, '', 0, 0, 1, 1, '2023-09-26 13:46:59.522', '2023-09-26 13:46:59.526', NULL);
INSERT INTO `sys_menu` VALUES (4, '', '用户创建', '', 'sys_user_create', 1, 3, 'sys:sysUser:add', 2, 0, '', 0, 0, 1, 1, '2023-09-26 13:46:59.534', '2023-09-26 13:46:59.536', NULL);
INSERT INTO `sys_menu` VALUES (5, '', '用户修改', '', 'sys_user_update', 1, 3, 'sys:sysUser:edit', 2, 0, '', 0, 0, 1, 1, '2023-09-26 13:46:59.543', '2023-09-26 13:46:59.546', NULL);
INSERT INTO `sys_menu` VALUES (6, '', '用户删除', '', 'sys_user_del', 1, 3, 'sys:sysUser:remove', 2, 0, '', 0, 0, 1, 1, '2023-09-26 13:46:59.553', '2023-09-26 13:46:59.555', NULL);
INSERT INTO `sys_menu` VALUES (7, 'SysMenuManage', '菜单管理', 'menu', '/sys/sys-menu', 1, 2, 'sys:sysMenu:list', 1, 0, '/sys/sys-menu/index', 0, 0, 1, 1, '2023-09-26 13:47:37.025', '2023-09-26 13:47:37.029', NULL);
INSERT INTO `sys_menu` VALUES (8, '', '菜单详情', '', 'sys_menu_detail', 1, 3, 'sys:sysMenu:query', 7, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:37.042', '2023-09-26 13:47:37.055', NULL);
INSERT INTO `sys_menu` VALUES (9, '', '菜单创建', '', 'sys_menu_create', 1, 3, 'sys:sysMenu:add', 7, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:37.067', '2023-09-26 13:47:37.069', NULL);
INSERT INTO `sys_menu` VALUES (10, '', '菜单修改', '', 'sys_menu_update', 1, 3, 'sys:sysMenu:edit', 7, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:37.076', '2023-09-26 13:47:37.079', NULL);
INSERT INTO `sys_menu` VALUES (11, '', '菜单删除', '', 'sys_menu_del', 1, 3, 'sys:sysMenu:remove', 7, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:37.084', '2023-09-26 13:47:37.086', NULL);
INSERT INTO `sys_menu` VALUES (12, 'SysRoleManage', '角色管理', 'role', '/sys/sys-role', 2, 2, 'sys:sysRole:list', 1, 0, '/sys/sys-role/index', 0, 0, 1, 1, '2023-09-26 13:47:39.688', '2023-09-26 13:47:39.694', NULL);
INSERT INTO `sys_menu` VALUES (13, '', '角色详情', '', 'sys_role_detail', 2, 3, 'sys:sysRole:query', 12, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:39.714', '2023-09-26 13:47:39.717', NULL);
INSERT INTO `sys_menu` VALUES (14, '', '角色创建', '', 'sys_role_create', 2, 3, 'sys:sysRole:add', 12, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:39.747', '2023-09-26 13:47:39.752', NULL);
INSERT INTO `sys_menu` VALUES (15, '', '角色修改', '', 'sys_role_update', 2, 3, 'sys:sysRole:edit', 12, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:39.765', '2023-09-26 13:47:39.769', NULL);
INSERT INTO `sys_menu` VALUES (16, '', '角色删除', '', 'sys_role_del', 2, 3, 'sys:sysRole:remove', 12, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:39.776', '2023-09-26 13:47:39.778', NULL);
INSERT INTO `sys_menu` VALUES (17, 'SysDeptManage', '部门管理', 'dept', '/sys/sys-dept', 2, 2, 'sys:sysDept:list', 1, 0, '/sys/sys-dept/index', 0, 0, 1, 1, '2023-09-26 13:47:42.140', '2023-09-26 13:47:42.144', NULL);
INSERT INTO `sys_menu` VALUES (18, '', '部门详情', '', 'sys_dept_detail', 2, 3, 'sys:sysDept:query', 17, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:42.155', '2023-09-26 13:47:42.167', NULL);
INSERT INTO `sys_menu` VALUES (19, '', '部门创建', '', 'sys_dept_create', 2, 3, 'sys:sysDept:add', 17, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:42.180', '2023-09-26 13:47:42.182', NULL);
INSERT INTO `sys_menu` VALUES (20, '', '部门修改', '', 'sys_dept_update', 2, 3, 'sys:sysDept:edit', 17, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:42.188', '2023-09-26 13:47:42.190', NULL);
INSERT INTO `sys_menu` VALUES (21, '', '部门删除', '', 'sys_dept_del', 2, 3, 'sys:sysDept:remove', 17, 0, '', 0, 0, 1, 1, '2023-09-26 13:47:42.200', '2023-09-26 13:47:42.202', NULL);
INSERT INTO `sys_menu` VALUES (52, 'SysTeamManage', '团队管理', 'team', '/sys/sys-team', 1, 2, 'sys:sysTeam:list', 1, 0, '/sys/sys-team/index', 0, 0, 1, 1, '2023-09-29 08:44:07.932', '2023-09-29 08:44:07.935', NULL);
INSERT INTO `sys_menu` VALUES (53, '', '团队详情', '', 'sys_team_detail', 1, 3, 'sys:sysTeam:query', 52, 0, '', 0, 0, 1, 1, '2023-09-29 08:44:07.941', '2023-09-29 08:44:07.945', NULL);
INSERT INTO `sys_menu` VALUES (54, '', '团队创建', '', 'sys_team_create', 1, 3, 'sys:sysTeam:add', 52, 0, '', 0, 0, 1, 1, '2023-09-29 08:44:07.951', '2023-09-29 08:44:07.954', NULL);
INSERT INTO `sys_menu` VALUES (55, '', '团队修改', '', 'sys_team_update', 1, 3, 'sys:sysTeam:edit', 52, 0, '', 0, 0, 1, 1, '2023-09-29 08:44:07.958', '2023-09-29 08:44:07.961', NULL);
INSERT INTO `sys_menu` VALUES (56, '', '团队删除', '', 'sys_team_del', 1, 3, 'sys:sysTeam:remove', 52, 0, '', 0, 0, 1, 1, '2023-09-29 08:44:07.966', '2023-09-29 08:44:07.971', NULL);
INSERT INTO `sys_menu` VALUES (57, 'SysMemberManage', '成员管理', 'openArm', '/sys/sys-member', 2, 2, 'sys:sysMember:list', 1, 0, '/sys/sys-member/index', 0, 0, 1, 1, '2023-09-29 08:44:33.484', '2023-09-29 08:44:33.508', NULL);
INSERT INTO `sys_menu` VALUES (58, '', '成员详情', '', 'sys_member_detail', 2, 3, 'sys:sysMember:query', 57, 0, '', 0, 0, 1, 1, '2023-09-29 08:44:33.517', '2023-09-29 08:44:33.520', NULL);
INSERT INTO `sys_menu` VALUES (59, '', '成员创建', '', 'sys_member_create', 2, 3, 'sys:sysMember:add', 57, 0, '', 0, 0, 1, 1, '2023-09-29 08:44:33.526', '2023-09-29 08:44:33.529', NULL);
INSERT INTO `sys_menu` VALUES (60, '', '成员修改', '', 'sys_member_update', 2, 3, 'sys:sysMember:edit', 57, 0, '', 0, 0, 1, 1, '2023-09-29 08:44:33.534', '2023-09-29 08:44:33.538', NULL);
INSERT INTO `sys_menu` VALUES (61, '', '会员删除', '', 'sys_member_del', 2, 3, 'sys:sysMember:remove', 57, 0, '', 0, 0, 1, 1, '2023-09-29 08:44:33.543', '2023-09-29 08:44:33.546', NULL);
INSERT INTO `sys_menu` VALUES (62, '', '营销管理', 'bill', '/dental', 2, 1, '', 0, 0, 'Layout', 0, 0, 1, 0, '2023-09-29 08:45:53.780', '2023-09-29 08:45:53.783', NULL);
INSERT INTO `sys_menu` VALUES (63, 'BillManage', '账单管理', 'bill', '/dental/bill', 2, 2, 'dental:bill:list', 62, 0, '/dental/bill/index', 0, 0, 1, 1, '2023-09-29 08:45:53.794', '2023-09-29 08:45:53.797', NULL);
INSERT INTO `sys_menu` VALUES (64, '', '账单详情', '', 'bill_detail', 2, 3, 'dental:bill:query', 63, 0, '', 0, 0, 1, 1, '2023-09-29 08:45:53.803', '2023-09-29 08:45:53.806', NULL);
INSERT INTO `sys_menu` VALUES (65, '', '账单创建', '', 'bill_create', 2, 3, 'dental:bill:add', 63, 0, '', 0, 0, 1, 1, '2023-09-29 08:45:53.818', '2023-09-29 08:45:53.821', NULL);
INSERT INTO `sys_menu` VALUES (66, '', '账单修改', '', 'bill_update', 2, 3, 'dental:bill:edit', 63, 0, '', 0, 0, 1, 1, '2023-09-29 08:45:53.829', '2023-09-29 08:45:53.833', NULL);
INSERT INTO `sys_menu` VALUES (67, '', '账单删除', '', 'bill_del', 2, 3, 'dental:bill:remove', 63, 0, '', 0, 0, 1, 1, '2023-09-29 08:45:53.840', '2023-09-29 08:45:53.843', NULL);
INSERT INTO `sys_menu` VALUES (68, 'CustomerManage', '客户管理', 'customer', '/dental/customer', 2, 2, 'dental:customer:list', 62, 0, '/dental/customer/index', 0, 0, 1, 1, '2023-09-29 08:46:25.139', '2023-09-29 08:46:25.143', NULL);
INSERT INTO `sys_menu` VALUES (69, '', '客户详情', '', 'customer_detail', 2, 3, 'dental:customer:query', 68, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:25.153', '2023-09-29 08:46:25.156', NULL);
INSERT INTO `sys_menu` VALUES (70, '', '客户创建', '', 'customer_create', 2, 3, 'dental:customer:add', 68, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:25.165', '2023-09-29 08:46:25.168', NULL);
INSERT INTO `sys_menu` VALUES (71, '', '客户修改', '', 'customer_update', 2, 3, 'dental:customer:edit', 68, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:25.176', '2023-09-29 08:46:25.179', NULL);
INSERT INTO `sys_menu` VALUES (72, '', '客户删除', '', 'customer_del', 2, 3, 'dental:customer:remove', 68, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:25.187', '2023-09-29 08:46:25.190', NULL);
INSERT INTO `sys_menu` VALUES (73, 'EventDayStManage', '日统计管理', 'calendarCheck', '/dental/event-day-st', 2, 2, 'dental:eventDaySt:list', 62, 0, '/dental/event-day-st/index', 0, 0, 1, 1, '2023-09-29 08:46:27.736', '2023-09-29 08:46:27.762', NULL);
INSERT INTO `sys_menu` VALUES (74, '', '日统计详情', '', 'event_day_st_detail', 2, 3, 'dental:eventDaySt:query', 73, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:27.771', '2023-09-29 08:46:27.774', NULL);
INSERT INTO `sys_menu` VALUES (75, '', '日统计创建', '', 'event_day_st_create', 2, 3, 'dental:eventDaySt:add', 73, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:27.780', '2023-09-29 08:46:27.783', NULL);
INSERT INTO `sys_menu` VALUES (76, '', '日统计修改', '', 'event_day_st_update', 2, 3, 'dental:eventDaySt:edit', 73, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:27.790', '2023-09-29 08:46:27.794', NULL);
INSERT INTO `sys_menu` VALUES (77, '', '日统计删除', '', 'event_day_st_del', 2, 3, 'dental:eventDaySt:remove', 73, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:27.801', '2023-09-29 08:46:27.806', NULL);
INSERT INTO `sys_menu` VALUES (78, 'SummaryPlanDayManage', '日总结与计划管理', 'planet', '/dental/summary-plan-day', 2, 2, 'dental:summaryPlanDay:list', 62, 0, '/dental/summary-plan-day/index', 0, 0, 1, 1, '2023-09-29 08:46:30.209', '2023-09-29 08:46:30.215', NULL);
INSERT INTO `sys_menu` VALUES (79, '', '日总结与计划详情', '', 'summary_plan_day_detail', 2, 3, 'dental:summaryPlanDay:query', 78, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:30.222', '2023-09-29 08:46:30.225', NULL);
INSERT INTO `sys_menu` VALUES (80, '', '日总结与计划创建', '', 'summary_plan_day_create', 2, 3, 'dental:summaryPlanDay:add', 78, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:30.233', '2023-09-29 08:46:30.236', NULL);
INSERT INTO `sys_menu` VALUES (81, '', '日总结与计划修改', '', 'summary_plan_day_update', 2, 3, 'dental:summaryPlanDay:edit', 78, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:30.244', '2023-09-29 08:46:30.247', NULL);
INSERT INTO `sys_menu` VALUES (82, '', '日总结与计划删除', '', 'summary_plan_day_del', 2, 3, 'dental:summaryPlanDay:remove', 78, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:30.254', '2023-09-29 08:46:30.258', NULL);
INSERT INTO `sys_menu` VALUES (83, 'TargetTaskManage', '任务目标管理', 'task', '/dental/target-task', 2, 2, 'dental:targetTask:list', 62, 0, '/dental/target-task/index', 0, 0, 1, 1, '2023-09-29 08:46:32.790', '2023-09-29 08:46:32.815', NULL);
INSERT INTO `sys_menu` VALUES (84, '', '任务目标详情', '', 'target_task_detail', 2, 3, 'dental:targetTask:query', 83, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:32.828', '2023-09-29 08:46:32.830', NULL);
INSERT INTO `sys_menu` VALUES (85, '', '任务目标创建', '', 'target_task_create', 2, 3, 'dental:targetTask:add', 83, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:32.837', '2023-09-29 08:46:32.840', NULL);
INSERT INTO `sys_menu` VALUES (86, '', '任务目标修改', '', 'target_task_update', 2, 3, 'dental:targetTask:edit', 83, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:32.847', '2023-09-29 08:46:32.850', NULL);
INSERT INTO `sys_menu` VALUES (87, '', '任务目标删除', '', 'target_task_del', 2, 3, 'dental:targetTask:remove', 83, 0, '', 0, 0, 1, 1, '2023-09-29 08:46:32.857', '2023-09-29 08:46:32.861', NULL);
INSERT INTO `sys_menu` VALUES (88, NULL, '账单智能识别', NULL, 'bill_identify', 2, 3, 'dental:bill:identify', 63, 0, '', 0, 0, 1, 1, '2023-10-15 00:00:00.000', '2023-10-15 00:00:00.000', NULL);
INSERT INTO `sys_menu` VALUES (89, NULL, '日统计', NULL, 'st_day', 2, 3, 'dental:st:day', 108, 0, NULL, 0, 0, 1, 1, '2023-10-16 13:53:40.000', '2023-10-16 13:53:44.000', NULL);
INSERT INTO `sys_menu` VALUES (90, NULL, '月统计', NULL, 'st_month', 2, 3, 'dental:st:month', 108, 0, NULL, 0, 0, 1, 1, '2023-10-16 14:32:02.000', '2023-10-16 14:32:06.000', NULL);
INSERT INTO `sys_menu` VALUES (91, '', '工具管理', 'setUp', '/tools', 1, 1, NULL, 0, 0, 'Layout', 0, 0, 1, 1, '2023-10-28 09:41:26.000', '2023-10-28 09:41:29.000', NULL);
INSERT INTO `sys_menu` VALUES (92, 'ToolManage', '代码生成', 'setUp', '/tools/gen', 1, 2, 'tools:gen:list', 91, 0, '/tool/gen/index', 0, 0, 1, 1, '2023-10-28 09:44:15.000', '2023-10-28 09:44:18.000', NULL);
INSERT INTO `sys_menu` VALUES (93, 'GenTablesManage', 'GenTables管理', 'table', '/sys/gen-tables', 1, 2, 'sys:genTables:list', 91, 0, '/sys/gen-tables/index', 0, 0, 1, 1, '2023-10-28 16:31:47.693', '2023-10-28 16:31:47.693', NULL);
INSERT INTO `sys_menu` VALUES (94, '', 'GenTables详情', '', 'gen_tables_detail', 1, 3, 'sys:genTables:query', 93, 0, '', 0, 0, 1, 1, '2023-10-28 16:31:47.704', '2023-10-28 16:31:47.704', NULL);
INSERT INTO `sys_menu` VALUES (95, '', 'GenTables创建', '', 'gen_tables_create', 1, 3, 'sys:genTables:add', 93, 0, '', 0, 0, 1, 1, '2023-10-28 16:31:47.714', '2023-10-28 16:31:47.714', NULL);
INSERT INTO `sys_menu` VALUES (96, '', 'GenTables修改', '', 'gen_tables_update', 1, 3, 'sys:genTables:edit', 93, 0, '', 0, 0, 1, 1, '2023-10-28 16:31:47.723', '2023-10-28 16:31:47.723', NULL);
INSERT INTO `sys_menu` VALUES (97, '', 'GenTables删除', '', 'gen_tables_del', 1, 3, 'sys:genTables:remove', 93, 0, '', 0, 0, 1, 1, '2023-10-28 16:31:47.731', '2023-10-28 16:31:47.731', NULL);
INSERT INTO `sys_menu` VALUES (98, 'GenColumnsManage', 'GenColumns管理', 'generate', '/sys/gen-columns', 1, 2, 'sys:genColumns:list', 91, 0, '/sys/gen-columns/index', 0, 0, 1, 1, '2023-10-28 16:35:01.334', '2023-10-28 16:35:01.334', NULL);
INSERT INTO `sys_menu` VALUES (99, '', 'GenColumns详情', '', 'gen_columns_detail', 1, 3, 'sys:genColumns:query', 98, 0, '', 0, 0, 1, 1, '2023-10-28 16:35:01.354', '2023-10-28 16:35:01.354', NULL);
INSERT INTO `sys_menu` VALUES (100, '', 'GenColumns创建', '', 'gen_columns_create', 1, 3, 'sys:genColumns:add', 98, 0, '', 0, 0, 1, 1, '2023-10-28 16:35:01.363', '2023-10-28 16:35:01.363', NULL);
INSERT INTO `sys_menu` VALUES (101, '', 'GenColumns修改', '', 'gen_columns_update', 1, 3, 'sys:genColumns:edit', 98, 0, '', 0, 0, 1, 1, '2023-10-28 16:35:01.371', '2023-10-28 16:35:01.371', NULL);
INSERT INTO `sys_menu` VALUES (102, '', 'GenColumns删除', '', 'gen_columns_del', 1, 3, 'sys:genColumns:remove', 98, 0, '', 0, 0, 1, 1, '2023-10-28 16:35:01.392', '2023-10-28 16:35:01.392', NULL);
INSERT INTO `sys_menu` VALUES (103, 'SysApiManage', '接口管理', 'pass', '/sys/sys-api', 1, 2, 'sys:sysApi:list', 91, 0, '/sys/sys-api/index', 0, 0, 1, 1, '2023-10-28 16:51:44.924', '2023-10-28 16:51:44.924', NULL);
INSERT INTO `sys_menu` VALUES (104, '', '接口详情', '', 'sys_api_detail', 1, 3, 'sys:sysApi:query', 103, 0, '', 0, 0, 1, 1, '2023-10-28 16:51:44.938', '2023-10-28 16:51:44.938', NULL);
INSERT INTO `sys_menu` VALUES (105, '', '接口创建', '', 'sys_api_create', 1, 3, 'sys:sysApi:add', 103, 0, '', 0, 0, 1, 1, '2023-10-28 16:51:44.947', '2023-10-28 16:51:44.947', NULL);
INSERT INTO `sys_menu` VALUES (106, '', '接口修改', '', 'sys_api_update', 1, 3, 'sys:sysApi:edit', 103, 0, '', 0, 0, 1, 1, '2023-10-28 16:51:44.955', '2023-10-28 16:51:44.955', NULL);
INSERT INTO `sys_menu` VALUES (107, '', '接口删除', '', 'sys_api_del', 1, 3, 'sys:sysApi:remove', 103, 0, '', 0, 0, 1, 1, '2023-10-28 16:51:44.964', '2023-10-28 16:51:44.964', NULL);
INSERT INTO `sys_menu` VALUES (108, 'BillStQuery', '统计查询', 'histogram', '/dental/st', 2, 2, 'dental:st:list', 62, 0, '/dental/billst/index', 0, 0, 1, 1, '2023-10-29 14:34:14.000', '2023-10-29 14:34:23.000', NULL);
INSERT INTO `sys_menu` VALUES (109, NULL, '查询统计', NULL, 'st_bill_query', 2, 3, 'dental:st:query', 108, 0, NULL, 0, 0, 1, 1, '2023-10-29 18:47:57.000', '2023-10-29 18:48:02.000', NULL);

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
INSERT INTO `sys_menu_api_rule` VALUES (57, 26);
INSERT INTO `sys_menu_api_rule` VALUES (58, 27);
INSERT INTO `sys_menu_api_rule` VALUES (59, 28);
INSERT INTO `sys_menu_api_rule` VALUES (60, 29);
INSERT INTO `sys_menu_api_rule` VALUES (61, 30);
INSERT INTO `sys_menu_api_rule` VALUES (52, 31);
INSERT INTO `sys_menu_api_rule` VALUES (53, 32);
INSERT INTO `sys_menu_api_rule` VALUES (54, 33);
INSERT INTO `sys_menu_api_rule` VALUES (55, 34);
INSERT INTO `sys_menu_api_rule` VALUES (56, 35);
INSERT INTO `sys_menu_api_rule` VALUES (63, 36);
INSERT INTO `sys_menu_api_rule` VALUES (64, 37);
INSERT INTO `sys_menu_api_rule` VALUES (65, 38);
INSERT INTO `sys_menu_api_rule` VALUES (66, 39);
INSERT INTO `sys_menu_api_rule` VALUES (67, 40);
INSERT INTO `sys_menu_api_rule` VALUES (68, 41);
INSERT INTO `sys_menu_api_rule` VALUES (69, 42);
INSERT INTO `sys_menu_api_rule` VALUES (70, 43);
INSERT INTO `sys_menu_api_rule` VALUES (71, 44);
INSERT INTO `sys_menu_api_rule` VALUES (72, 45);
INSERT INTO `sys_menu_api_rule` VALUES (73, 46);
INSERT INTO `sys_menu_api_rule` VALUES (74, 47);
INSERT INTO `sys_menu_api_rule` VALUES (75, 48);
INSERT INTO `sys_menu_api_rule` VALUES (76, 49);
INSERT INTO `sys_menu_api_rule` VALUES (77, 50);
INSERT INTO `sys_menu_api_rule` VALUES (78, 51);
INSERT INTO `sys_menu_api_rule` VALUES (79, 52);
INSERT INTO `sys_menu_api_rule` VALUES (80, 53);
INSERT INTO `sys_menu_api_rule` VALUES (81, 54);
INSERT INTO `sys_menu_api_rule` VALUES (82, 55);
INSERT INTO `sys_menu_api_rule` VALUES (83, 56);
INSERT INTO `sys_menu_api_rule` VALUES (84, 57);
INSERT INTO `sys_menu_api_rule` VALUES (85, 58);
INSERT INTO `sys_menu_api_rule` VALUES (86, 59);
INSERT INTO `sys_menu_api_rule` VALUES (87, 60);
INSERT INTO `sys_menu_api_rule` VALUES (88, 61);
INSERT INTO `sys_menu_api_rule` VALUES (89, 62);
INSERT INTO `sys_menu_api_rule` VALUES (90, 63);
INSERT INTO `sys_menu_api_rule` VALUES (93, 64);
INSERT INTO `sys_menu_api_rule` VALUES (94, 65);
INSERT INTO `sys_menu_api_rule` VALUES (95, 66);
INSERT INTO `sys_menu_api_rule` VALUES (96, 67);
INSERT INTO `sys_menu_api_rule` VALUES (97, 68);
INSERT INTO `sys_menu_api_rule` VALUES (98, 69);
INSERT INTO `sys_menu_api_rule` VALUES (99, 70);
INSERT INTO `sys_menu_api_rule` VALUES (100, 71);
INSERT INTO `sys_menu_api_rule` VALUES (101, 72);
INSERT INTO `sys_menu_api_rule` VALUES (102, 73);
INSERT INTO `sys_menu_api_rule` VALUES (103, 74);
INSERT INTO `sys_menu_api_rule` VALUES (104, 75);
INSERT INTO `sys_menu_api_rule` VALUES (105, 76);
INSERT INTO `sys_menu_api_rule` VALUES (106, 77);
INSERT INTO `sys_menu_api_rule` VALUES (107, 78);
INSERT INTO `sys_menu_api_rule` VALUES (109, 79);

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '操作日志' ROW_FORMAT = Dynamic;

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
  `role_key` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '角色代码',
  `role_sort` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '排序',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `team_id` tinyint(1) NULL DEFAULT NULL COMMENT '团队',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `create_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_role_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, 'test', 'test', NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

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
INSERT INTO `sys_role_menu` VALUES (1, 1);
INSERT INTO `sys_role_menu` VALUES (1, 2);
INSERT INTO `sys_role_menu` VALUES (1, 3);
INSERT INTO `sys_role_menu` VALUES (1, 4);
INSERT INTO `sys_role_menu` VALUES (1, 5);
INSERT INTO `sys_role_menu` VALUES (1, 6);
INSERT INTO `sys_role_menu` VALUES (1, 7);
INSERT INTO `sys_role_menu` VALUES (1, 8);

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '短信' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_sms
-- ----------------------------

-- ----------------------------
-- Table structure for sys_team
-- ----------------------------
DROP TABLE IF EXISTS `sys_team`;
CREATE TABLE `sys_team`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '团队名',
  `owner` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '团队拥有者',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态',
  `created_at` datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_team
-- ----------------------------
INSERT INTO `sys_team` VALUES (1, '好牙（杭州）', 0, 2, '2023-09-30 00:22:43', '2023-09-30 14:18:31');

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
  `platform_role_id` mediumint(0) NULL DEFAULT NULL COMMENT '平台角色ID 大于0为平台账户,0为团队账户',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '备注',
  `lock_time` datetime(3) NULL DEFAULT NULL COMMENT '锁定结束时间',
  `status` tinyint(0) NULL DEFAULT NULL COMMENT '状态 1正常 ',
  `update_by` int(0) UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `created_at` datetime(3) NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) NULL DEFAULT NULL COMMENT '最后更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_user_update_by`(`update_by`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '用户' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'dilu', '', NULL, '$2a$10$2OxaPJviu7NMSKMk5c2mPOvvb41Xg5ZiQB0153QpB77THK4sIXF1a', 'dilu', 'dilu', NULL, NULL, NULL, '2', -1, NULL, NULL, 1, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (2, 'tangtang', '13800138001', NULL, '$2a$10$2OxaPJviu7NMSKMk5c2mPOvvb41Xg5ZiQB0153QpB77THK4sIXF1a', '糖糖', '唐敦霞', NULL, NULL, NULL, '2', NULL, NULL, NULL, 1, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (3, 'zcm', NULL, NULL, '$2a$10$2OxaPJviu7NMSKMk5c2mPOvvb41Xg5ZiQB0153QpB77THK4sIXF1a', '藏春梅', '藏春梅', NULL, NULL, NULL, '2', NULL, NULL, NULL, 1, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (4, 'liyanlei', NULL, NULL, NULL, '李艳雷', '李艳雷', NULL, NULL, NULL, '1', NULL, NULL, NULL, 1, NULL, NULL, '2023-10-29 14:01:31.999');
INSERT INTO `sys_user` VALUES (5, NULL, NULL, NULL, NULL, '简小丽', '简小丽', NULL, NULL, NULL, '2', NULL, NULL, NULL, 1, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (6, NULL, NULL, NULL, NULL, '胡珊', '胡珊', NULL, NULL, NULL, '2', NULL, NULL, NULL, 1, NULL, NULL, NULL);
INSERT INTO `sys_user` VALUES (7, NULL, NULL, NULL, NULL, '余鸿雁', '余鸿雁', NULL, NULL, NULL, '2', NULL, NULL, NULL, 1, NULL, NULL, NULL);

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
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '三方登录' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of third_login
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
