#MHW-Server



#sql

create database mhw default charset = 'utf8';

CREATE USER mhw_admin IDENTIFIED BY 'mhw_admin_1803';

grant all on *.* to 'mhw_admin'@'localhost' identified by 'mhw_admin_1803'



drop table articles;
create table articles(
  `id` bigint(20) unsigned primary key  NOT NULL AUTO_INCREMENT COMMENT '主键',

  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_modified` datetime NOT NULL COMMENT '修改时间',

  `title` varchar(256) not null comment '标题',
  `content` longtext not null comment 'content',
  `source` varchar(256) not null comment 'source',
  `source_type` varchar(45) not null default 'news',
  `image` varchar(256)
);

alter table articles add index source_type(source_type);

insert into articles(`gmt_create`, `gmt_modified`, `title`, `content`, `source`, `source_type`)
values(now(), now(), '索尼推《怪物猎人世界》限定版联动周边 又来骗钱？',
'《怪物猎人世界（Monster Hunter World）》将于2018年1月26日发售，登陆PS4平台。为了配合该游戏的推出，SONY将与CAPCOM合作推出一系列《怪物猎人世界》周边，包含音乐播放器、头戴式蓝牙耳机、便携蓝牙音箱等',
'NGA', 'news');

select * from articles \G


create table weixin_users(
  `id` bigint(20) unsigned primary key  NOT NULL AUTO_INCREMENT COMMENT '主键',

  `gmt_create` datetime NOT NULL COMMENT '创建时间',
  `gmt_modified` datetime NOT NULL COMMENT '修改时间',

  `tokens` varchar(128) not null comment 'token',
  `head_img` varchar(256) not null comment 'head imgs',
  `name` varchar(128) not null comment 'name',
  `gender` int default 2 comment '0: girl, 1: man, 2: unknown',
  `vip_date` datetime not null default now()
);

alter table weixin_users add index tokens(tokens);





