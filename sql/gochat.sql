/*
 Navicat Premium Data Transfer

 Source Server         : 本机mysql
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : 127.0.0.1:3306
 Source Schema         : gochat

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 08/05/2024 12:09:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for communities
-- ----------------------------
DROP TABLE IF EXISTS `communities`;
CREATE TABLE `communities`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `owner_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `img` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_communities_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of communities
-- ----------------------------
INSERT INTO `communities` VALUES (19, '2024-05-03 21:49:21.416', '2024-05-03 21:49:21.416', NULL, '大雪菜', 1, '', '123');
INSERT INTO `communities` VALUES (20, '2024-05-03 21:59:25.394', '2024-05-03 21:59:25.394', NULL, '群21', 13, '', '324');
INSERT INTO `communities` VALUES (21, '2024-05-04 13:47:47.739', '2024-05-04 13:47:47.739', NULL, '计算机社团', 1, '', '学习群');

-- ----------------------------
-- Table structure for contacts
-- ----------------------------
DROP TABLE IF EXISTS `contacts`;
CREATE TABLE `contacts`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `owner_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `target_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `type` bigint(20) NULL DEFAULT NULL,
  `desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_contacts_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 53 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of contacts
-- ----------------------------
INSERT INTO `contacts` VALUES (40, '2024-05-03 21:49:21.429', '2024-05-03 21:49:21.429', NULL, 1, 19, 2, '');
INSERT INTO `contacts` VALUES (41, '2024-05-03 21:51:08.531', '2024-05-03 21:51:08.531', NULL, 1, 13, 1, '');
INSERT INTO `contacts` VALUES (42, '2024-05-03 21:51:08.531', '2024-05-03 21:51:08.531', NULL, 13, 1, 1, '');
INSERT INTO `contacts` VALUES (43, '2024-05-03 21:51:12.495', '2024-05-03 21:51:12.495', NULL, 1, 14, 1, '');
INSERT INTO `contacts` VALUES (44, '2024-05-03 21:51:12.496', '2024-05-03 21:51:12.496', NULL, 14, 1, 1, '');
INSERT INTO `contacts` VALUES (45, '2024-05-03 21:51:15.368', '2024-05-03 21:51:15.368', NULL, 1, 15, 1, '');
INSERT INTO `contacts` VALUES (46, '2024-05-03 21:51:15.368', '2024-05-03 21:51:15.368', NULL, 15, 1, 1, '');
INSERT INTO `contacts` VALUES (47, '2024-05-03 21:51:35.009', '2024-05-03 21:51:35.009', NULL, 13, 19, 2, '');
INSERT INTO `contacts` VALUES (48, '2024-05-03 21:59:25.394', '2024-05-03 21:59:25.394', NULL, 13, 20, 2, '');
INSERT INTO `contacts` VALUES (49, '2024-05-03 21:59:44.570', '2024-05-03 21:59:44.570', NULL, 1, 20, 2, '');
INSERT INTO `contacts` VALUES (50, '2024-05-04 13:47:47.739', '2024-05-04 13:47:47.739', NULL, 1, 21, 2, '');
INSERT INTO `contacts` VALUES (51, '2024-05-04 13:48:00.138', '2024-05-04 13:48:00.138', NULL, 13, 21, 2, '');
INSERT INTO `contacts` VALUES (52, '2024-05-04 13:48:05.481', '2024-05-04 13:48:05.481', NULL, 14, 21, 2, '');

-- ----------------------------
-- Table structure for group_basics
-- ----------------------------
DROP TABLE IF EXISTS `group_basics`;
CREATE TABLE `group_basics`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `owner_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `icon` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `type` bigint(20) NULL DEFAULT NULL,
  `desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_group_basics_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of group_basics
-- ----------------------------

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `target_id` bigint(20) NULL DEFAULT NULL,
  `type` bigint(20) NULL DEFAULT NULL,
  `media` bigint(20) NULL DEFAULT NULL,
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `create_time` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `read_time` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `pic` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `url` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `desc` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `amount` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_message_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of message
-- ----------------------------

-- ----------------------------
-- Table structure for user_basics
-- ----------------------------
DROP TABLE IF EXISTS `user_basics`;
CREATE TABLE `user_basics`  (
  `name` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `phone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `email` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `avatar` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `identity` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `client_ip` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `client_port` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `login_time` datetime(3) NULL DEFAULT NULL,
  `heartbeat_time` datetime(3) NULL DEFAULT NULL,
  `login_out_time` datetime(3) NULL DEFAULT NULL,
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `salt` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_basics_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_basics
-- ----------------------------
INSERT INTO `user_basics` VALUES ('zs', '6f51e660866139a84e84f8b98f786bd4', '1234566', '91231@qq.com', '/asset/images/avatar1.jpg', 'a597acb023e0bead24e74933020d77b4', '', '', '2024-04-29 12:28:38.127', '2024-04-29 12:28:38.127', '2024-04-29 12:28:38.127', 1, '2024-04-29 12:28:38.127', '2024-05-06 19:46:01.914', NULL, '1497517903');
INSERT INTO `user_basics` VALUES ('ls', '3b67487c81b3afcc9a268cb4be6fea80', '12313143', '12313132', NULL, '32c39fd283b0f6e824b53d950c43ddb0', '', '', '2024-04-29 12:30:48.143', '2024-04-29 12:30:48.143', '2024-04-29 12:30:48.143', 13, '2024-04-29 12:30:48.143', '2024-05-06 19:49:43.522', NULL, '906694880');
INSERT INTO `user_basics` VALUES ('wu', '582e3fefbe63a2368092027fc00f7540', '231313', '123123123', NULL, 'd08fd121438ef2c598b3aeb84e4fa185', '', '', '2024-04-29 12:30:54.511', '2024-04-29 12:30:54.511', '2024-04-29 12:30:54.511', 14, '2024-04-29 12:30:54.511', '2024-05-06 19:49:53.321', NULL, '1331477368');
INSERT INTO `user_basics` VALUES ('zl', '700b1ccb007f8b5a9117bad7a89e388b', '1231313', '12312313', NULL, '59f9d9bacd67ffb97cfd60b2aafa1912', '', '', '2024-04-29 12:30:57.197', '2024-04-29 12:30:57.197', '2024-04-29 12:30:57.197', 15, '2024-04-29 12:30:57.197', '2024-05-03 16:12:58.365', NULL, '1723865637');
INSERT INTO `user_basics` VALUES ('a1', '0c5d19da1c9ad94fe114cc7c70b65b73', '', '', NULL, '', '', '', '2024-04-29 12:34:27.603', '2024-04-29 12:34:27.603', '2024-04-29 12:34:27.603', 16, '2024-04-29 12:34:27.603', '2024-04-29 12:34:27.603', NULL, '1197460819');
INSERT INTO `user_basics` VALUES ('a2', 'c51fca93616a861eb57345f3ba1e6de4', '', '', NULL, '', '', '', '2024-04-29 12:34:29.892', '2024-04-29 12:34:29.892', '2024-04-29 12:34:29.892', 17, '2024-04-29 12:34:29.892', '2024-04-29 12:34:29.892', NULL, '1371209940');
INSERT INTO `user_basics` VALUES ('a3', '8979632b8a032692c4f9e956a026af60', '', '', NULL, '', '', '', '2024-04-29 12:34:31.538', '2024-04-29 12:34:31.538', '2024-04-29 12:34:31.538', 18, '2024-04-29 12:34:31.538', '2024-04-29 12:34:31.538', NULL, '963407454');
INSERT INTO `user_basics` VALUES ('a4', 'f083914335e18870638bd4d5bf9d2418', '', '', NULL, '', '', '', '2024-04-29 12:34:33.774', '2024-04-29 12:34:33.774', '2024-04-29 12:34:33.774', 19, '2024-04-29 12:34:33.774', '2024-04-29 12:34:33.774', NULL, '1511233307');
INSERT INTO `user_basics` VALUES ('a5', 'f79c93cb3f18b80122dbe4023cd72ed1', '', '', NULL, '', '', '', '2024-04-29 12:34:35.801', '2024-04-29 12:34:35.801', '2024-04-29 12:34:35.801', 20, '2024-04-29 12:34:35.801', '2024-04-29 12:34:35.801', NULL, '1580232736');

SET FOREIGN_KEY_CHECKS = 1;
