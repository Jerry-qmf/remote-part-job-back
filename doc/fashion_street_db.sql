create DATABASE remote_part_job_db;

CREATE TABLE `job_info_tab` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `job_title` varchar(32) DEFAULT '其它',
    `job_pay` varchar(32) NOT NULL,
    `job_label` varchar(32) NOT NULL,
    `job_describe` varchar(1024) DEFAULT '-',
    `job_carousel_list` varchar(512) DEFAULT '-',
    `wechat_url` varchar(256) DEFAULT '-',
    `wechat_num` varchar(32) NOT NULL,
    `expires` int default '0',
    `is_top` tinyint(1) DEFAULT '0',
    `deleted` tinyint(1) DEFAULT '0',
    `created_at` int unsigned not NULL DEFAULT '0',
    `updated_at` int unsigned not NULL DEFAULT '0',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

insert into job_info_tab (id,job_title,job_pay,job_label,job_describe,job_carousel_list,wechat_url,wechat_num,created_at,updated_at) VALUES (1,"兼职写手","500-800/天","远程 高薪","<h1>Hello, World!</h1>","https://www.something.wang/api/v1/image/download?image=carousel-1.jpg","https://www.something.wang/api/v1/image/download?image=1-wechat.jpg","123456789",unix_timestamp(now()),unix_timestamp(now()));

CREATE TABLE `carousel_info_tab` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `carousel_url` varchar(256) DEFAULT '',
    `expires` int default '0',
    `is_top` tinyint(1) DEFAULT '0',
    `deleted` tinyint(1) DEFAULT '0',
    `created_at` int unsigned not NULL DEFAULT '0',
    `updated_at` int unsigned not NULL DEFAULT '0',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

insert into carousel_info_tab (carousel_url,created_at,updated_at) VALUES ("https://www.something.wang/remote_part_job/api/v1/image/download?image=1688275870228465410-CdMum.jpg",unix_timestamp(now()),unix_timestamp(now()));
insert into carousel_info_tab (carousel_url,created_at,updated_at) VALUES ("https://www.something.wang/remote_part_job/api/v1/image/download?image=1688275904019592387-Zkrnu.jpg",unix_timestamp(now()),unix_timestamp(now()));
insert into carousel_info_tab (carousel_url,created_at,updated_at) VALUES ("https://www.something.wang/remote_part_job/api/v1/image/download?image=1688275704976009709-ZpQhi.jpg",unix_timestamp(now()),unix_timestamp(now()));
