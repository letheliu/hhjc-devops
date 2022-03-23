-- 2022-02-07

INSERT INTO `menu` VALUES ('700201904027', '发布者', '800201904008', '/index.html#/pages/admin/appPublisherManage', '2', '发布者', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904022', '发布者', '发布者', '2019-04-01 02:24:53', '0', '/pages/admin/appPublisherManage', '700201904027');
INSERT INTO `privilege_rel` VALUES ('40', '500201904022', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('41', '500201904022', '600201904002', '2019-04-01 08:18:29', '0');

-- 2022-02-09
create table install_app
(
    app_id         varchar(64) not null
        primary key,
    app_name       varchar(64) not null,
    version        varchar(64) not null,
    ext_app_id     varchar(64) not null,
    create_user_id varchar(64) not null,
    tenant_id varchar(64) not null ,
    create_time    timestamp  default CURRENT_TIMESTAMP not null,
    status_cd      varchar(2) default '0' not null
);

create table app_publisher
(
    publisher_id varchar(64)  not null
        primary key,
    username     varchar(256) not null,
    email        varchar(64)  not null,
    token        varchar(128) not null,
    phone        varchar(11)  not null,
    state        varchar(12) default '001' not null,
    create_time  timestamp   default CURRENT_TIMESTAMP not null,
    status_cd    varchar(2)  default '0' not null,
    tenant_id    varchar(64)  not null,
    ext_publisher_id varchar(64) not null
);

create table business_images_ext
(
    id             varchar(64)  not null
        primary key,
    images_id           varchar(64)  not null,
    app_id        varchar(64)  not null,
    app_name    varchar(128)  not null,
    ext_images_id       varchar(64) not null,
    ext_publisher_id    varchar(64)  not null,
    create_time    timestamp  default CURRENT_TIMESTAMP not null,
    status_cd      varchar(2) default '0' not null,
    tenant_id      varchar(64)  not null
);

-- 2022-02-13

INSERT INTO `menu` VALUES ('700201904028', 'ftp', '800201904003', '/index.html#/pages/admin/ftpManage', '3', 'ftp', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904023', 'ftp', 'ftp', '2019-04-01 02:24:53', '0', '/pages/admin/ftpManage', '700201904028');
INSERT INTO `privilege_rel` VALUES ('42', '500201904023', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('43', '500201904023', '600201904002', '2019-04-01 08:18:29', '0');

INSERT INTO `menu` VALUES ('700201904029', 'oss', '800201904003', '/index.html#/pages/admin/ossManage', '3', 'oss', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904024', 'oss', 'oss', '2019-04-01 02:24:53', '0', '/pages/admin/ossManage', '700201904029');
INSERT INTO `privilege_rel` VALUES ('44', '500201904024', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('45', '500201904024', '600201904002', '2019-04-01 08:18:29', '0');

INSERT INTO `menu` VALUES ('700201904030', '数据库', '800201904003', '/index.html#/pages/admin/dbManage', '3', '数据库', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904025', '数据库', '数据库', '2019-04-01 02:24:53', '0', '/pages/admin/dbManage', '700201904030');
INSERT INTO `privilege_rel` VALUES ('46', '500201904025', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('47', '500201904025', '600201904002', '2019-04-01 08:18:29', '0');

INSERT INTO `menu` VALUES ('700201904031', '资源备份', '800201904003', '/index.html#/pages/admin/backupManage', '3', '资源备份', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904026', '资源备份', '资源备份', '2019-04-01 02:24:53', '0', '/pages/admin/backupManage', '700201904031');
INSERT INTO `privilege_rel` VALUES ('48', '500201904026', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('49', '500201904026', '600201904002', '2019-04-01 08:18:29', '0');


create table resources_ftp
(
    ftp_id        varchar(64)  primary key  not null,
    name           varchar(64)    not null,
    ip             varchar(128)   not null,
    port             varchar(64)        not null,
    username       varchar(64)    not null,
    passwd         varchar(64)    not null,
    tenant_id      varchar(64)    not null,
    path      varchar(128)    not null,
    create_time    timestamp   default CURRENT_TIMESTAMP not null,
    status_cd      varchar(2)  default '0' not null
);


create table resources_oss
(
    oss_id        varchar(64)  primary key  not null,
    name           varchar(64)    not null,
    oss_type       varchar(12) not null,
    bucket             varchar(128)   not null,
    access_key_secret             varchar(128)        not null,
    access_key_id       varchar(128)    not null,
    endpoint         varchar(128)    not null,
    tenant_id      varchar(64)    not null,
    path      varchar(128)    not null,
    create_time    timestamp   default CURRENT_TIMESTAMP not null,
    status_cd      varchar(2)  default '0' not null
);

create table resources_db
(
    db_id             varchar(64)  not null primary key,
    name           varchar(64)    not null,
    ip             varchar(128) not null,
    port           varchar(12)  not null,
    username       varchar(64)  not null,
    password       varchar(128) not null,
    db_name        varchar(64)  not null,
    tenant_id      varchar(64)    not null,
    create_time    timestamp   default CURRENT_TIMESTAMP not null,
    status_cd      varchar(2)  default '0' not null
);

create table resources_backup
(
    id          varchar(64) not null
        primary key,
    name        varchar(64) not null,
    exec_time   varchar(64) not null,
    type_cd     varchar(12) not null,
    src_id      varchar(64) not null,
    src_object  longtext    not null,
    target_type_cd   varchar(12) not null,
    target_id   varchar(64) not null,
    tenant_id   varchar(64) not null,
    state       varchar(12) not null,
    back_time timestamp  not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

-- 2022-02-16 加入服务追踪

INSERT INTO `menu_group` VALUES ('800201904013', '服务追踪', 'fa fa-globe', '', '10', '服务追踪', '2019-04-01 07:55:51', '0', 'P_WEB');

INSERT INTO `menu` VALUES ('700201904032', '调用链', '800201904013', '/index.html#/pages/admin/logTrace', '1', '调用链', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904027', '调用链', '调用链', '2019-04-01 02:24:53', '0', '/pages/admin/logTrace', '700201904032');
INSERT INTO `privilege_rel` VALUES ('50', '500201904027', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('51', '500201904027', '600201904002', '2019-04-01 08:18:29', '0');

INSERT INTO `menu` VALUES ('700201904033', '调用sql', '800201904013', '/index.html#/pages/admin/logTraceDb', '2', '调用sql', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904028', '调用sql', '调用sql', '2019-04-01 02:24:53', '0', '/pages/admin/logTraceDb', '700201904033');
INSERT INTO `privilege_rel` VALUES ('52', '500201904028', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('53', '500201904028', '600201904002', '2019-04-01 08:18:29', '0');




-- 2022-02-19 修改 log_trace_param 表
drop table log_trace_param;
create table log_trace_param
(
    id          varchar(64) not null
        primary key,
    span_id     varchar(64) not null,
    req_header  logtext,
    req_param   logtext,
    res_header   logtext,
    res_param   logtext,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table log_trace_db
(
    id          varchar(64) not null
        primary key,
    span_id     varchar(64) not null,
    db_sql  logtext,
    param   logtext,
    duration   varchar(64),
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table waf
(
    waf_id           varchar(64) not null
        primary key,
    waf_name varchar(64) not null,
    port         varchar(64) not null,
    state varchar(12) not null,
    create_time  timestamp  default CURRENT_TIMESTAMP not null,
    status_cd    varchar(2) default '0' not null
);

create table waf_hosts
(
    waf_host_id           varchar(64) not null
        primary key,
    waf_id varchar(64) not null,
    host_id         varchar(64) not null,
    create_time  timestamp  default CURRENT_TIMESTAMP not null,
    status_cd    varchar(2) default '0' not null
);

create table waf_route
(
    route_id           varchar(64) not null
        primary key,
    waf_id varchar(64) not null,
    hostname varchar(128) not null,
    ip varchar(128) not null,
    port         varchar(64) not null,
    create_time  timestamp  default CURRENT_TIMESTAMP not null,
    status_cd    varchar(2) default '0' not null
);

create table waf_access_log
(
    request_id           varchar(64) not null
        primary key,
    waf_ip varchar(64) not null,
    host_id           varchar(64) not null,
    x_real_ip         varchar(64) not null,
    scheme        varchar(64) not null,
    response_code    varchar(64) not null,
    method    varchar(64) not null,
    http_host    varchar(64) not null,
    upstream_addr varchar(64) not null ,
    url    varchar(512) not null,
    request_length varchar(64) not null,
    response_length varchar(64) not null,
    state varchar(12) not null,
    message varchar(512) not null,
    create_time  timestamp  default CURRENT_TIMESTAMP not null,
    status_cd    varchar(2) default '0' not null
);