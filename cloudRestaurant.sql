/*
Navicat MySQL Data Transfer

Source Server         : 43.139.78.177
Source Server Version : 50740
Source Host           : 43.139.78.177:33061
Source Database       : cloudRestaurant

Target Server Type    : MYSQL
Target Server Version : 50740
File Encoding         : 65001

Date: 2023-06-20 15:37:09
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for food_category
-- ----------------------------
DROP TABLE IF EXISTS `food_category`;
CREATE TABLE `food_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(20) DEFAULT NULL,
  `description` varchar(30) DEFAULT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  `link_url` varchar(255) DEFAULT NULL,
  `is_in_serving` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of food_category
-- ----------------------------
INSERT INTO `food_category` VALUES ('1', '预定早餐', '预定早餐', 'https://img1.baidu.com/it/u=1770839340,1650524274&fm=253&fmt=auto&app=138&f=JPEG?w=499&h=500', '', '1');
INSERT INTO `food_category` VALUES ('2', '果蔬生鲜', '果蔬生鲜', 'https://img0.baidu.com/it/u=3561394708,981723676&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=375', '', '1');
INSERT INTO `food_category` VALUES ('3', '鲜花蛋糕', '鲜花蛋糕', 'https://img0.baidu.com/it/u=3863240341,4151205672&fm=253&fmt=auto&app=138&f=PNG?w=500&h=500', '', '1');
INSERT INTO `food_category` VALUES ('4', '商超便利', '商超便利', 'https://img1.baidu.com/it/u=1835296012,1920502616&fm=253&fmt=auto&app=120&f=JPEG?w=800&h=1025', '', '0');
INSERT INTO `food_category` VALUES ('5', '美食', '美食', 'https://img0.baidu.com/it/u=1950708306,3755148504&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500', '', '0');
INSERT INTO `food_category` VALUES ('6', '甜品饮品', '甜品饮品', 'https://img2.baidu.com/it/u=1456269164,211471403&fm=253&fmt=auto&app=138&f=PNG?w=500&h=500', '', '0');
INSERT INTO `food_category` VALUES ('7', '土豪推荐', '土豪推荐', 'https://img2.baidu.com/it/u=1526784363,3592891821&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500', '', '0');
INSERT INTO `food_category` VALUES ('8', '准时达', '准时达', 'https://img1.baidu.com/it/u=3720391912,2487976670&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=788', '', '0');
INSERT INTO `food_category` VALUES ('9', '简餐', '简餐', 'https://img0.baidu.com/it/u=2033050375,3976304783&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500', '', '0');
INSERT INTO `food_category` VALUES ('10', '汉堡薯条', '汉堡薯条', 'https://img2.baidu.com/it/u=1176031919,3396668158&fm=253&fmt=auto&app=138&f=PNG?w=511&h=500', '', '0');
INSERT INTO `food_category` VALUES ('11', '日韩料理', '日韩料理', 'https://img0.baidu.com/it/u=1456118985,1698535492&fm=253&fmt=auto&app=138&f=JPEG?w=610&h=449', '', '0');
INSERT INTO `food_category` VALUES ('12', '麻辣烫', '麻辣烫', 'https://img0.baidu.com/it/u=1365112546,2237439435&fm=253&fmt=auto&app=138&f=JPEG?w=640&h=452', '', '0');
INSERT INTO `food_category` VALUES ('13', '披萨意面', '披萨意面', 'https://img0.baidu.com/it/u=1565538418,4150868021&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=560', '', '0');
INSERT INTO `food_category` VALUES ('14', '川湘菜', '川湘菜', 'https://img2.baidu.com/it/u=3947210249,1593226444&fm=253&fmt=auto&app=138&f=JPEG?w=780&h=500', '', '0');
INSERT INTO `food_category` VALUES ('15', '包子粥铺', '包子粥铺', 'https://img0.baidu.com/it/u=3573340081,3518034932&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500', '', '0');
INSERT INTO `food_category` VALUES ('16', '新店特惠', '新店特惠', 'https://img2.baidu.com/it/u=391822331,1420647278&fm=253&fmt=auto&app=138&f=PNG?w=500&h=500', '', '0');

-- ----------------------------
-- Table structure for tb_goods
-- ----------------------------
DROP TABLE IF EXISTS `tb_goods`;
CREATE TABLE `tb_goods` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `sell_count` bigint(20) DEFAULT NULL,
  `price` double DEFAULT NULL,
  `old_price` double DEFAULT NULL,
  `shop_id` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_goods
-- ----------------------------
INSERT INTO `tb_goods` VALUES ('1', '小小鲜肉包', '滑蛋牛肉粥(1份)+小小鲜肉包(4只)', '', '14', '25', '29', '1');
INSERT INTO `tb_goods` VALUES ('2', '滑蛋牛肉粥+小小鲜肉包', '滑蛋牛肉粥(1份)+小小鲜肉包(3只)', '', '6', '35', '41', '1');
INSERT INTO `tb_goods` VALUES ('3', '滑蛋牛肉粥+绿甘蓝馅饼', '滑蛋牛肉粥(1份)+绿甘蓝馅饼(1张)', '', '2', '25', '30', '1');
INSERT INTO `tb_goods` VALUES ('4', '茶香卤味蛋', '咸鸡蛋', '', '688', '2.5', '3', '1');
INSERT INTO `tb_goods` VALUES ('5', '韭菜鸡蛋馅饼(2张)', '韭菜鸡蛋馅饼', '', '381', '10', '12', '1');
INSERT INTO `tb_goods` VALUES ('6', '小小鲜肉包+豆浆套餐', '小小鲜肉包(8只)装+豆浆(1杯)', '', '335', '9.899999618530273', '11.899999618530273', '479');
INSERT INTO `tb_goods` VALUES ('7', '翠香炒素饼', '咸鲜翠香素炒饼', '', '260', '17.899999618530273', '20.899999618530273', '485');
INSERT INTO `tb_goods` VALUES ('8', '香煎鲜肉包', '咸鲜猪肉鲜肉包', '', '173', '10.899999618530273', '12.899999618530273', '486');

-- ----------------------------
-- Table structure for tb_member
-- ----------------------------
DROP TABLE IF EXISTS `tb_member`;
CREATE TABLE `tb_member` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) DEFAULT NULL,
  `mobile` varchar(11) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `register_time` bigint(255) DEFAULT NULL,
  `avatar` varchar(255) DEFAULT NULL,
  `balance` double(255,0) DEFAULT NULL,
  `is_active` tinyint(255) DEFAULT NULL,
  `city` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_member
-- ----------------------------
INSERT INTO `tb_member` VALUES ('1', '17343200634', '17343200634', '', '1685719641', 'group1/M00/00/00/CgAMEWSNw4mANpzFADlpRcWUC-I421.jpg', '0', '0', null);
INSERT INTO `tb_member` VALUES ('2', 'admin', '', '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918', '1686236625', 'group1/M00/00/00/CgAMEWSRUCWAfaw3AADfRohuYQY548.png', '0', '0', null);
INSERT INTO `tb_member` VALUES ('3', '19372436051', '19372436051', '', '1686586966', '', '0', '0', null);
INSERT INTO `tb_member` VALUES ('4', 'zzzz', '', '2d6ccd34ad7af363159ed4bbe18c0e43c681f606877d9ffc96b62200720d7291', '1686671970', '', '0', '0', null);
INSERT INTO `tb_member` VALUES ('5', '1808208820', '', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', '1686839273', 'group1/M00/00/00/CgAMEWSNwkiAa_U0AACRgo5vphY138.png', '0', '0', '');
INSERT INTO `tb_member` VALUES ('6', '1808208820', '', '8bb0cf6eb9b17d0f7d22b456f121257dc1254e1f01665370476383ea776df414', '1687012112', '', '0', '0', '');
INSERT INTO `tb_member` VALUES ('7', '1808208820', '', 'ef797c8118f02dfb649607dd5d3f8c7623048c9c063d532cc95c5ed7a898a64f', '1687012203', '', '0', '0', '');
INSERT INTO `tb_member` VALUES ('8', '1575197604@', '', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', '1687183044', '', '0', '0', '');
INSERT INTO `tb_member` VALUES ('9', '1575197604@', '', '2558a34d4d20964ca1d272ab26ccce9511d880579593cd4c9e01ab91ed00f325', '1687183608', '', '0', '0', '');
INSERT INTO `tb_member` VALUES ('10', '15171618793', '', '8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92', '1687188598', '', '0', '0', '');

-- ----------------------------
-- Table structure for tb_service
-- ----------------------------
DROP TABLE IF EXISTS `tb_service`;
CREATE TABLE `tb_service` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `icon_name` varchar(255) DEFAULT NULL,
  `icon_color` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_service
-- ----------------------------
INSERT INTO `tb_service` VALUES ('4', '开发票', '该商家支持开发票，请在下单时填写好发票抬头', '票', '999999');
INSERT INTO `tb_service` VALUES ('7', '外卖保', '已加入“外卖保”计划，食品安全有保障', '保', '999999');
INSERT INTO `tb_service` VALUES ('9', '准时达', '准时必达，超时秒赔', '准', '57A9FF');

-- ----------------------------
-- Table structure for tb_shop
-- ----------------------------
DROP TABLE IF EXISTS `tb_shop`;
CREATE TABLE `tb_shop` (
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `promotion_info` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `longitude` double DEFAULT NULL,
  `latitude` double DEFAULT NULL,
  `image_path` varchar(255) DEFAULT NULL,
  `is_new` tinyint(1) DEFAULT NULL,
  `is_premium` tinyint(1) DEFAULT NULL,
  `rating` double DEFAULT NULL,
  `rating_count` bigint(20) DEFAULT NULL,
  `recent_order_num` bigint(20) DEFAULT NULL,
  `minimum_order_amount` int(11) DEFAULT NULL,
  `delivery_fee` int(11) DEFAULT NULL,
  `opening_hours` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=500 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_shop
-- ----------------------------
INSERT INTO `tb_shop` VALUES ('1', '嘉和一品（温都水城）', '欢迎光临，用餐高峰请提前下单，谢谢', '北京市昌平区宏福苑温都水城F1', '13437850035', '1', '116.36868', '40.10039', 'https://img0.baidu.com/it/u=776607826,3851276909&fm=253&fmt=auto&app=138&f=JPEG?w=400&h=473', '1', '1', '4.699999809265137', '961', '106', '20', '5', '8:30/20:30');
INSERT INTO `tb_shop` VALUES ('479', '杨国福麻辣烫', '欢迎光临，用餐高峰请提前下单，谢谢', '北京市市蜀山区南二环路天鹅湖万达广场8号楼1705室', '13167583411', '1', '116.22124', '40.81948', 'https://img0.baidu.com/it/u=4256118608,437326922&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=376', '1', '1', '4.199999809265137', '167', '755', '20', '5', '8:30/20:30');
INSERT INTO `tb_shop` VALUES ('485', '好适口', '欢迎光临，用餐高峰请提前下单，谢谢', '北京市海淀区西二旗大街58号', '12345678901', '1', '116.65355', '40.26578', 'https://img0.baidu.com/it/u=2395082049,4167035839&fm=253&fmt=auto&app=138&f=JPEG?w=623&h=500', '1', '1', '4.599999904632568', '576', '58', '20', '5', '8:30/20:30');
INSERT INTO `tb_shop` VALUES ('486', '东来顺旗舰店', '老北京正宗涮羊肉,非物质文化遗产', '北京市天河区东圃镇汇彩路38号1领汇创展商务中心401', '13544323775', '1', '116.41724', '40.1127', 'https://img1.baidu.com/it/u=2247943454,2919106826&fm=253&fmt=auto&app=138&f=JPEG?w=325&h=260', '1', '1', '4.199999809265137', '372', '542', '20', '5', '09:00/21:30');
INSERT INTO `tb_shop` VALUES ('487', '北京酒家', '北京第一家传承300年酒家', '北京市海淀区上下九商业步行街内', '13257482341', '0', '116.24826', '40.11488', 'https://img0.baidu.com/it/u=232497929,3516278409&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500', '1', '1', '4.199999809265137', '871', '923', '20', '5', '8:30/20:30');
INSERT INTO `tb_shop` VALUES ('488', '和平鸽饺子馆', '吃饺子就来和平鸽饺子馆', '北京市越秀区德政中路171', '17098764762', '1', '116.27521', '40.12092', 'https://img2.baidu.com/it/u=2528840159,3220488960&fm=253&fmt=auto&app=138&f=JPEG?w=667&h=500', '1', '1', '4.199999809265137', '273', '483', '20', '5', '8:30/20:30');
INSERT INTO `tb_shop` VALUES ('499', '黄焖鸡米饭', '一只鸡的传说,一道菜的餐厅。', '北京市海淀区上下九商业步行街内', '13812115655', '1', '116.12345', '40.32145', 'https://img2.baidu.com/it/u=1970706776,1468429064&fm=253&fmt=auto&app=138&f=JPEG?w=600&h=395', '1', '1', '4.199999809265137', '654', '465', '20', '5', '09:00/21:30');

-- ----------------------------
-- Table structure for tb_shop_service
-- ----------------------------
DROP TABLE IF EXISTS `tb_shop_service`;
CREATE TABLE `tb_shop_service` (
  `shop_id` int(20) NOT NULL,
  `service_id` int(20) NOT NULL,
  PRIMARY KEY (`shop_id`,`service_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_shop_service
-- ----------------------------
INSERT INTO `tb_shop_service` VALUES ('1', '4');
INSERT INTO `tb_shop_service` VALUES ('1', '7');
INSERT INTO `tb_shop_service` VALUES ('1', '9');
INSERT INTO `tb_shop_service` VALUES ('479', '4');
INSERT INTO `tb_shop_service` VALUES ('479', '7');
INSERT INTO `tb_shop_service` VALUES ('479', '9');

-- ----------------------------
-- Table structure for tb_smsCode
-- ----------------------------
DROP TABLE IF EXISTS `tb_smsCode`;
CREATE TABLE `tb_smsCode` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `phone` varchar(255) DEFAULT NULL,
  `biz_id` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `create_time` bigint(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of tb_smsCode
-- ----------------------------
INSERT INTO `tb_smsCode` VALUES ('1', '17343200634', '714117885631123135^0', '345585', '1685631121');
INSERT INTO `tb_smsCode` VALUES ('2', '17343200634', '350504585803336447^0', '804587', '1685803334');
INSERT INTO `tb_smsCode` VALUES ('3', '17343200634', '509007185803456104^0', '397981', '1685803454');
INSERT INTO `tb_smsCode` VALUES ('4', '17343200634', '271707186234367986^0', '672962', '1686234368');
INSERT INTO `tb_smsCode` VALUES ('5', '17343200634', '344815486234557675^0', '492553', '1686234557');
INSERT INTO `tb_smsCode` VALUES ('6', '19372436051', '464422086586491437^0', '637090', '1686586491');
INSERT INTO `tb_smsCode` VALUES ('7', '19372436051', '928402186586743483^0', '880520', '1686586743');
INSERT INTO `tb_smsCode` VALUES ('8', '19372436051', '439107286586896313^0', '772818', '1686586896');
INSERT INTO `tb_smsCode` VALUES ('9', '19372436051', '972407486586954646^0', '623492', '1686586954');
INSERT INTO `tb_smsCode` VALUES ('10', '17343200634', '103919686669762313^0', '035119', '1686669762');
INSERT INTO `tb_smsCode` VALUES ('11', '17343200634', '944712086670063180^0', '875263', '1686670063');
INSERT INTO `tb_smsCode` VALUES ('12', '17343200634', '308718586670587954^0', '487762', '1686670587');
INSERT INTO `tb_smsCode` VALUES ('13', '17343200634', '997218986670987430^0', '364147', '1686670987');
INSERT INTO `tb_smsCode` VALUES ('14', '17343200634', '581204786671068080^0', '800201', '1686671067');
INSERT INTO `tb_smsCode` VALUES ('15', '17343200634', '107719786703623544^0', '169870', '1686703621');
INSERT INTO `tb_smsCode` VALUES ('16', '17343200634', '841807687243892439^0', '984277', '1687243891');
