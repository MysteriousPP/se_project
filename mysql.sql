/*
 CareBridge 数据库完整建表脚本 (已增加 SOS 经纬度字段)
 包含：核心用户、家庭关系、健康监测、SOS(含经纬度)、家庭社交、智能提醒、紧急联系人、社区服务
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- 1. 用户表 (user)
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `userID` INT(10) NOT NULL AUTO_INCREMENT COMMENT '用户唯一ID',
  `phone` VARCHAR(20) NOT NULL UNIQUE COMMENT '登录账号（手机号）',
  `password` VARCHAR(128) NOT NULL COMMENT 'SHA256加盐加密存储',
  `role` TINYINT(4) NOT NULL COMMENT '1=银发端, 2=子女端',
  `nickname` VARCHAR(30) NOT NULL COMMENT '用户昵称',
  PRIMARY KEY (`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- 2. 绑定关系表 (bind_relation)
-- ----------------------------
DROP TABLE IF EXISTS `bind_relation`;
CREATE TABLE `bind_relation` (
  `bindID` INT(11) NOT NULL AUTO_INCREMENT COMMENT '绑定记录ID',
  `childUserID` INT(11) NOT NULL COMMENT '子女端用户ID',
  `elderUserId` INT(11) NOT NULL COMMENT '银发端用户ID',
  `bindTime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '绑定时间',
  `status` TINYINT(4) NOT NULL DEFAULT 1 COMMENT '1=有效, 0=已解绑',
  PRIMARY KEY (`bindID`),
  INDEX `idx_child` (`childUserID`),
  INDEX `idx_elder` (`elderUserId`),
  CONSTRAINT `fk_bind_child` FOREIGN KEY (`childUserID`) REFERENCES `user` (`userID`) ON DELETE CASCADE,
  CONSTRAINT `fk_bind_elder` FOREIGN KEY (`elderUserId`) REFERENCES `user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- 3. 健康数据表 (health_data)
-- ----------------------------
DROP TABLE IF EXISTS `health_data`;
CREATE TABLE `health_data` (
  `dataID` INT(11) NOT NULL AUTO_INCREMENT COMMENT '数据唯一ID',
  `userID` INT(11) NOT NULL COMMENT '银发端用户ID',
  `highBp` INT(11) NOT NULL COMMENT '高压值(mmHg)',
  `lowBp` INT(11) NOT NULL COMMENT '低压值(mmHg)',
  `measureTime` DATETIME NOT NULL COMMENT '测量时间',
  `uploadType` TINYINT(4) NOT NULL COMMENT '1=手动, 2=设备上传',
  PRIMARY KEY (`dataID`),
  INDEX `idx_health_user` (`userID`),
  CONSTRAINT `fk_health_user` FOREIGN KEY (`userID`) REFERENCES `user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- 4. 紧急求助表 (sos_alert) - 已更新经纬度字段
-- ----------------------------
DROP TABLE IF EXISTS `sos_alert`;
CREATE TABLE `sos_alert` (
  `alertId` INT(11) NOT NULL AUTO_INCREMENT COMMENT '求助记录ID',
  `elderUserId` INT(11) NOT NULL COMMENT '触发求助的老人ID',
  `triggerTime` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '触发时间',
  `location` VARCHAR(255) DEFAULT NULL COMMENT '求助时的详细地址描述',
  `longitude` VARCHAR(50) DEFAULT NULL COMMENT '经度',
  `latitude` VARCHAR(50) DEFAULT NULL COMMENT '纬度',
  `status` TINYINT(4) DEFAULT 1 COMMENT '1=待处理, 2=已接听, 3=已处理',
  `handledBy` INT(11) DEFAULT NULL COMMENT '处理人ID',
  `handledTime` DATETIME DEFAULT NULL COMMENT '处理时间',
  `notes` TEXT DEFAULT NULL COMMENT '处理备注',
  PRIMARY KEY (`alertId`),
  CONSTRAINT `fk_sos_elder` FOREIGN KEY (`elderUserId`) REFERENCES `user` (`userID`) ON DELETE CASCADE,
  CONSTRAINT `fk_sos_handler` FOREIGN KEY (`handledBy`) REFERENCES `user` (`userID`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- 5. 家庭动态表 (family_feed)
-- ----------------------------
DROP TABLE IF EXISTS `family_feed`;
CREATE TABLE `family_feed` (
  `feedId` INT(11) NOT NULL AUTO_INCREMENT COMMENT '动态ID',
  `authorId` INT(11) NOT NULL COMMENT '发布者ID',
  `content` TEXT NOT NULL COMMENT '动态内容',
  `images` JSON DEFAULT NULL COMMENT '图片URL数组',
  `createTime` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `isDeleted` TINYINT(4) DEFAULT 0 COMMENT '0=正常, 1=已删除',
  PRIMARY KEY (`feedId`),
  CONSTRAINT `fk_feed_author` FOREIGN KEY (`authorId`) REFERENCES `user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- 6. 智能提醒表 (reminder)
-- ----------------------------
DROP TABLE IF EXISTS `reminder`;
CREATE TABLE `reminder` (
  `reminderId` INT(11) NOT NULL AUTO_INCREMENT COMMENT '提醒ID',
  `createdBy` INT(11) NOT NULL COMMENT '创建人ID',
  `elderUserId` INT(11) NOT NULL COMMENT '提醒对象ID',
  `title` VARCHAR(100) NOT NULL COMMENT '提醒标题',
  `content` TEXT DEFAULT NULL COMMENT '提醒内容',
  `remindTime` DATETIME NOT NULL COMMENT '提醒时间',
  `repeatType` TINYINT(4) DEFAULT 0 COMMENT '0=不重复, 1=每天, 2=每周, 3=每月',
  `status` TINYINT(4) DEFAULT 1 COMMENT '1=待处理, 2=已完成',
  PRIMARY KEY (`reminderId`),
  CONSTRAINT `fk_rem_creator` FOREIGN KEY (`createdBy`) REFERENCES `user` (`userID`) ON DELETE CASCADE,
  CONSTRAINT `fk_rem_elder` FOREIGN KEY (`elderUserId`) REFERENCES `user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- 7. 紧急联系人表 (emergency_contact)
-- ----------------------------
DROP TABLE IF EXISTS `emergency_contact`;
CREATE TABLE `emergency_contact` (
  `contactId` INT(11) NOT NULL AUTO_INCREMENT COMMENT '联系人ID',
  `userId` INT(11) NOT NULL COMMENT '所属用户ID',
  `name` VARCHAR(30) NOT NULL COMMENT '联系人姓名',
  `phone` VARCHAR(20) NOT NULL COMMENT '联系人电话',
  `relationship` VARCHAR(20) DEFAULT NULL COMMENT '关系描述',
  `isCommunity` TINYINT(4) DEFAULT 0 COMMENT '0=个人, 1=社区联系人',
  PRIMARY KEY (`contactId`),
  CONSTRAINT `fk_contact_user` FOREIGN KEY (`userId`) REFERENCES `user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- 8. 社区服务表 (community_service)
-- ----------------------------
DROP TABLE IF EXISTS `community_service`;
CREATE TABLE `community_service` (
  `serviceId` INT(11) NOT NULL AUTO_INCREMENT COMMENT '服务ID',
  `serviceName` VARCHAR(50) NOT NULL COMMENT '服务名称',
  `price` DECIMAL(8,2) NOT NULL COMMENT '服务价格',
  `duration` INT(11) NOT NULL COMMENT '服务时长(分钟)',
  `provider` VARCHAR(50) NOT NULL COMMENT '服务提供方',
  `serviceStatus` TINYINT(4) DEFAULT 1 COMMENT '1=可预约, 0=已下架',
  PRIMARY KEY (`serviceId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- 9. 服务预约表 (service_booking)
-- ----------------------------
DROP TABLE IF EXISTS `service_booking`;
CREATE TABLE `service_booking` (
  `bookingId` INT(11) NOT NULL AUTO_INCREMENT COMMENT '预约ID',
  `serviceId` INT(11) NOT NULL COMMENT '服务ID',
  `elderUserId` INT(11) NOT NULL COMMENT '预约老人ID',
  `bookedBy` INT(11) NOT NULL COMMENT '预约人ID',
  `bookingTime` DATETIME NOT NULL COMMENT '预约服务时间',
  `status` TINYINT(4) DEFAULT 1 COMMENT '1=待确认, 2=已确认, 3=已完成',
  PRIMARY KEY (`bookingId`),
  CONSTRAINT `fk_book_service` FOREIGN KEY (`serviceId`) REFERENCES `community_service` (`serviceId`) ON DELETE CASCADE,
  CONSTRAINT `fk_book_elder` FOREIGN KEY (`elderUserId`) REFERENCES `user` (`userID`) ON DELETE CASCADE,
  CONSTRAINT `fk_book_by` FOREIGN KEY (`bookedBy`) REFERENCES `user` (`userID`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- 触发器：血压范围逻辑验证
-- ----------------------------
DELIMITER //
DROP TRIGGER IF EXISTS `before_health_data_insert` //
CREATE TRIGGER `before_health_data_insert`
BEFORE INSERT ON `health_data`
FOR EACH ROW
BEGIN
    IF NEW.highBp < 40 OR NEW.highBp > 250 THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '高压值超出合理范围(40-250)';
    END IF;
    IF NEW.lowBp < 30 OR NEW.lowBp > 150 THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '低压值超出合理范围(30-150)';
    END IF;
    IF NEW.highBp <= NEW.lowBp THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = '高压必须大于低压';
    END IF;
END //
DELIMITER ;

SET FOREIGN_KEY_CHECKS = 1;