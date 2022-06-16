create table user(
                     id int(10) unsigned not null auto_increment,
                     username varchar(255) not null unique,
                     password varchar(255) not null,
                     email varchar(255) not null,
                     avatar varchar(255) default null,
                     create_at timestamp not null,
                     age int(10) unsigned default null,
                     phone_number varchar(255) default null,
                     sign varchar(255) default null,
                     primary key (id),
                     unique key(username)
);

create table article(
                        id int(10) unsigned not null auto_increment,
                        create_at timestamp not null,
                        user_id int not null,
                        username varchar(255) not null,
                        title varchar(255) unique not null,
                        favoriteNum int(10) not null default 0,
                        commentNum int(10) not null default 0,
                        WatchNum int(10) not null default 0,
                        primary key (id),
                        unique key(title)
);



create table comment(
                        id int(10) unsigned not null auto_increment,
                        created_at timestamp not null,
                        favoriteNum int(10) not null default 0,
                        user_id int not null,
                        username varchar(255) not null,
                        article_id int(10) not null,
                        content text not null,
                        parent_id int(10) not null default 0 comment '为0时为顶级评论',
                        primary key (id)
);

create table tag
(
    id         int unsigned auto_increment primary key,
    article_id int unsigned not null,
    tag_name   varchar(20)  not null
);

create table divide(
                       id int(10) unsigned not null  auto_increment,
                       name varchar(255) not null,
                       nickname varchar(255) not null,
                       `describe` varchar(255) not null default '暂无描述',
                       parent_id int(10) unsigned not null default 0 comment '为0时为顶级标签',
                       primary key (id)
)

