package models


type Data interface {}			//用于存储模板解析时的action

//以下是用到的两张数据表，使用的是这个github.com/go-sql-driver/mysql数据库驱动

//CREATE TABLE `categroy` (
//`id` bigint(20) NOT NULL AUTO_INCREMENT,
//`title` varchar(255) NOT NULL DEFAULT '',
//`created` datetime DEFAULT NULL,
//`count` int(3) DEFAULT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8


//CREATE TABLE `topic` (
//`id` bigint(20) NOT NULL AUTO_INCREMENT,
//`title` varchar(255) NOT NULL DEFAULT '',
//`content` text,
//`created` datetime NOT NULL,
//`updated` datetime NOT NULL,
//`views` bigint(20) NOT NULL DEFAULT '0',
//`reply_count` bigint(20) NOT NULL DEFAULT '0',
//PRIMARY KEY (`id`),
//KEY `created` (`created`),
//KEY `updated` (`updated`),
//KEY `views` (`views`)
//) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8