/*
 Navicat MySQL Data Transfer

 Source Server         : 区块浏览器
 Source Server Type    : MySQL
 Source Server Version : 50730
 Source Host           : 47.107.147.245:13306
 Source Schema         : scan

 Target Server Type    : MySQL
 Target Server Version : 50730
 File Encoding         : 65001

 Date: 01/11/2021 09:55:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account_count_tbl
-- ----------------------------
DROP TABLE IF EXISTS `account_count_tbl`;
CREATE TABLE `account_count_tbl`  (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '自增 ID',
  `account_count` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '账户数量',
  `date` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '时间戳',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_idx`(`id`) USING BTREE COMMENT '自增 ID 索引',
  UNIQUE INDEX `date_idx`(`date`) USING BTREE COMMENT '日期索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_tbl
-- ----------------------------
DROP TABLE IF EXISTS `account_tbl`;
CREATE TABLE `account_tbl`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 ID',
  `addr` varbinary(40) NOT NULL COMMENT '地址',
  `balance` bigint(20) UNSIGNED NOT NULL COMMENT '余额',
  `vote` bigint(20) UNSIGNED NOT NULL COMMENT '票数',
  `create_time` int(10) UNSIGNED NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_idx`(`id`) USING BTREE COMMENT '自增ID索引',
  UNIQUE INDEX `addr_idx`(`addr`) USING BTREE COMMENT '地址索引',
  INDEX `balance_idx`(`balance`) USING BTREE COMMENT '余额索引',
  INDEX `vote_idx`(`vote`) USING BTREE COMMENT '票数索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for account_tx_tbl
-- ----------------------------
DROP TABLE IF EXISTS `account_tx_tbl`;
CREATE TABLE `account_tx_tbl`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 ID',
  `addr` varbinary(40) NULL DEFAULT NULL COMMENT '地址',
  `txid` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '交易哈希',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_idx`(`id`) USING BTREE COMMENT '自增ID索引',
  INDEX `addr_idx`(`addr`) USING BTREE COMMENT '地址索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for asset_tbl
-- ----------------------------
DROP TABLE IF EXISTS `asset_tbl`;
CREATE TABLE `asset_tbl`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 ID',
  `asset_symbol` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '资产符号',
  `asset_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '资产名称',
  `owner_addr` varbinary(40) NULL DEFAULT NULL COMMENT '所有者地址',
  `total_supply` bigint(20) NULL DEFAULT NULL COMMENT '发行总量',
  `mintable` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '是否可增发（0：不可增发；1：可增发）',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_idx`(`id`) USING BTREE COMMENT '自增ID索引',
  UNIQUE INDEX `asset_symbol_idx`(`asset_symbol`) USING BTREE COMMENT '资产符号索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for block_tbl
-- ----------------------------
DROP TABLE IF EXISTS `block_tbl`;
CREATE TABLE `block_tbl`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `height` int(10) UNSIGNED NOT NULL COMMENT '区块高度',
  `block_hash` varchar(64) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL COMMENT '区块哈希',
  `miner_uid` varbinary(66) NOT NULL COMMENT '矿工 RegID',
  `miner_address` varbinary(40) NOT NULL COMMENT '矿工地址',
  `size` int(10) UNSIGNED NOT NULL COMMENT '区块大小',
  `version` tinyint(3) UNSIGNED NOT NULL COMMENT '版本',
  `merkle_root` varchar(64) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL COMMENT '区块梅克尔根',
  `time` int(10) UNSIGNED NOT NULL COMMENT '区块时间戳',
  `previous_block_hash` varchar(64) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL COMMENT '上一个区块哈希',
  `next_block_hash` varchar(64) CHARACTER SET ascii COLLATE ascii_general_ci NULL DEFAULT NULL COMMENT '下一个区块哈希',
  `tx_count` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '交易数量',
  `tx` mediumtext CHARACTER SET ascii COLLATE ascii_general_ci NULL COMMENT '交易哈希集合',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_idx`(`id`) USING BTREE COMMENT '自增ID索引',
  UNIQUE INDEX `height_idx`(`height`) USING BTREE COMMENT '区块高度索引',
  UNIQUE INDEX `block_hash_idx`(`block_hash`) USING BTREE COMMENT '区块哈希索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tx_count_tbl
-- ----------------------------
DROP TABLE IF EXISTS `tx_count_tbl`;
CREATE TABLE `tx_count_tbl`  (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '自增 ID',
  `tx_count` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '交易数量',
  `date` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '时间戳',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_idx`(`id`) USING BTREE COMMENT '自增 ID 索引',
  UNIQUE INDEX `date_idx`(`date`) USING BTREE COMMENT '日期索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tx_tbl
-- ----------------------------
DROP TABLE IF EXISTS `tx_tbl`;
CREATE TABLE `tx_tbl`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 ID',
  `txid` varchar(64) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL COMMENT '交易哈希',
  `tx_type` varchar(30) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL COMMENT '交易类型',
  `version` tinyint(3) UNSIGNED NOT NULL COMMENT '版本',
  `valid_height` int(5) UNSIGNED NOT NULL COMMENT '有效高度',
  `confirmed_height` int(10) UNSIGNED NOT NULL COMMENT '确认高度',
  `confirmed_time` int(10) UNSIGNED NOT NULL COMMENT '确认时间',
  `block_hash` varchar(64) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL COMMENT '区块哈希',
  `receipts` mediumtext CHARACTER SET ascii COLLATE ascii_general_ci NULL COMMENT '交易收据',
  `rawtx` mediumtext CHARACTER SET ascii COLLATE ascii_general_ci NULL COMMENT '原始交易',
  `tx_uid` varbinary(66) NULL DEFAULT NULL COMMENT '发送方 UID',
  `from_addr` varbinary(40) NULL DEFAULT NULL COMMENT '发送方地址',
  `to_uid` varbinary(66) NULL DEFAULT NULL COMMENT '接收方 UID',
  `to_addr` varbinary(40) NULL DEFAULT NULL COMMENT '接收方地址',
  `fee_symbol` varchar(30) CHARACTER SET ascii COLLATE ascii_general_ci NULL DEFAULT NULL COMMENT '小费类型',
  `fees` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '小费',
  `pubkey` varbinary(66) NULL DEFAULT NULL COMMENT '公钥',
  `miner_pubkey` varbinary(66) NULL DEFAULT NULL COMMENT '挖矿公钥',
  `signature` varchar(200) CHARACTER SET ascii COLLATE ascii_general_ci NULL DEFAULT NULL COMMENT '交易签名',
  `candidate_votes` mediumtext CHARACTER SET ascii COLLATE ascii_general_ci NULL COMMENT '投票信息',
  `transfers` mediumtext CHARACTER SET ascii COLLATE ascii_general_ci NULL COMMENT '转账信息',
  `memo` varchar(200) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '备注',
  `required_sigs` tinyint(3) UNSIGNED NULL DEFAULT NULL COMMENT '需要的签名数量',
  `signatures` mediumtext CHARACTER SET ascii COLLATE ascii_general_ci NULL COMMENT '签名信息',
  `stake_type` varchar(30) CHARACTER SET ascii COLLATE ascii_general_ci NULL DEFAULT NULL COMMENT '抵押类型',
  `coin_symbol` varchar(30) CHARACTER SET ascii COLLATE ascii_general_ci NULL DEFAULT NULL COMMENT '币种符号',
  `coin_amount` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '币种数量',
  `vm_type` tinyint(3) UNSIGNED NULL DEFAULT NULL COMMENT '虚拟机类型',
  `upgradable` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '可升级',
  `code` mediumtext CHARACTER SET ascii COLLATE ascii_general_ci NULL COMMENT '合约代码',
  `abi` mediumtext CHARACTER SET ascii COLLATE ascii_general_ci NULL COMMENT '合约 ABI',
  `arguments` mediumtext CHARACTER SET ascii COLLATE ascii_general_ci NULL COMMENT '合约参数',
  `asset_name` varchar(60) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '资产名称',
  `asset_symbol` varchar(30) CHARACTER SET ascii COLLATE ascii_general_ci NULL DEFAULT NULL COMMENT '资产符号',
  `owner_uid` varbinary(66) NULL DEFAULT NULL COMMENT 'Owner UID',
  `owner_addr` varbinary(40) NULL DEFAULT NULL COMMENT 'Owner 地址',
  `total_supply` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '发行总量',
  `mintable` tinyint(1) UNSIGNED NULL DEFAULT NULL COMMENT '可增发',
  `update_type` varchar(30) CHARACTER SET ascii COLLATE ascii_general_ci NULL DEFAULT NULL COMMENT '更新类型',
  `update_value` varchar(30) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '更新值',
  `domain` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '存证 域',
  `key` varchar(100) CHARACTER SET utf8 COLLATE utf8_bin NULL DEFAULT NULL COMMENT '存证 Key',
  `value` text CHARACTER SET utf8 COLLATE utf8_bin NULL COMMENT '存证 Value',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_idx`(`id`) USING BTREE COMMENT '自增 ID 索引',
  UNIQUE INDEX `txid_idx`(`txid`) USING BTREE COMMENT '交易哈希索引',
  INDEX `from_addr_idx`(`from_addr`) USING BTREE COMMENT '发送方地址索引',
  INDEX `to_addr_idx`(`to_addr`) USING BTREE COMMENT '接收方地址索引',
  INDEX `block_hash_idx`(`block_hash`) USING BTREE COMMENT '区块哈希索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for tx_type_tbl
-- ----------------------------
DROP TABLE IF EXISTS `tx_type_tbl`;
CREATE TABLE `tx_type_tbl`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增 ID',
  `enum` tinyint(3) UNSIGNED NULL DEFAULT NULL COMMENT '交易类型枚举值',
  `tx_type` varchar(50) CHARACTER SET ascii COLLATE ascii_general_ci NULL DEFAULT NULL COMMENT '交易类型',
  `cn_description` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '中文描述',
  `en_description` varchar(50) CHARACTER SET ascii COLLATE ascii_general_ci NULL DEFAULT NULL COMMENT '英文描述',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `id_idx`(`id`) USING BTREE COMMENT '自增 ID 索引'
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tx_type_tbl
-- ----------------------------
INSERT INTO `tx_type_tbl` VALUES (1, 1, 'BLOCK_REWARD_TX', '区块奖励', 'Block Reward');
INSERT INTO `tx_type_tbl` VALUES (2, 2, 'ACCOUNT_REGISTER_TX', '账户激活', 'Account Register');
INSERT INTO `tx_type_tbl` VALUES (3, 3, 'DELEGATE_VOTE_TX', '投票', 'Vote');
INSERT INTO `tx_type_tbl` VALUES (4, 11, 'COIN_TRANSFER_TX', '转账', 'Transfer');
INSERT INTO `tx_type_tbl` VALUES (5, 12, 'COIN_TRANSFER_MTX', '多签转账', 'Mulsig Transfer');
INSERT INTO `tx_type_tbl` VALUES (6, 13, 'COIN_STAKE_TX', '抵押', 'Stake');
INSERT INTO `tx_type_tbl` VALUES (7, 21, 'CONTRACT_DEPLOY_TX', '合约部署', 'Contract Deploy');
INSERT INTO `tx_type_tbl` VALUES (8, 22, 'CONTRACT_INVOKE_TX', '合约调用', 'Contract Invoke');
INSERT INTO `tx_type_tbl` VALUES (9, 31, 'ASSET_ISSUE_TX', '资产发布', 'Asset Issue');
INSERT INTO `tx_type_tbl` VALUES (10, 32, 'ASSET_UPDATE_TX', '资产更新', 'Asset Update');
INSERT INTO `tx_type_tbl` VALUES (11, 41, 'RECORD_TX', '数字存证', 'Data Record');

SET FOREIGN_KEY_CHECKS = 1;
