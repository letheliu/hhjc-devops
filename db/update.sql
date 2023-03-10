-- 2022-02-07
alter table app_service rename to app_service_0331 ;

create table app_service
(
    as_id          varchar(64)  not null,
    as_name        varchar(128) not null,
    as_type        varchar(12)  not null,
    tenant_id      varchar(64)  not null,
    as_desc        varchar(512) default '''''''NULL''''''',
    state          varchar(12)  default '''10012''' not null,
    as_count       int(11)      default '''1''' not null,
    create_time    timestamp    default CURRENT_TIMESTAMP not null,
    status_cd      varchar(2)   default 0 not null,
    as_group_id    VARCHAR(64),
    as_deploy_type VARCHAR(64),
    as_deploy_id   VARCHAR(64),
    images_id      VARCHAR(64),
    ver_id         varchar(64)
);

insert into app_service(as_id, as_name, as_type, tenant_id, as_desc, state, as_count, create_time, status_cd, as_group_id, as_deploy_type, as_deploy_id, images_id, ver_id)
select as_id, as_name, as_type, tenant_id, as_desc, state, as_count, create_time, status_cd, as_group_id, as_deploy_type, as_deploy_id, images_id,'1' from app_service_0331

-- insert into business_images_ver(id, images_id, version, type_url, tenant_id)
-- select t.id,t.id,t.version,t.type_url,t.tenant_id from business_images t

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
    https_port varchar(64) not null,
    state varchar(12) not null,
    create_time  timestamp  default CURRENT_TIMESTAMP not null,
    status_cd    varchar(2) default '0' not null
);
insert into waf values (1,'waf',80,443,1001,'2020-01-01','0');
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
    route_id    varchar(64)  not null
        primary key,
    waf_id      varchar(64)  not null,
    scheme varchar(64) not null,
    hostname    varchar(128) not null,
    ip          varchar(128) not null,
    port        varchar(64)  not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
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

create table waf_hostname_cert
(
    cert_id    varchar(64)  not null primary key,
    hostname    varchar(128) not null,
    cert_content      longtext  not null,
    priv_key_content     longtext  not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

-- 2022-03-23 加入waf

INSERT INTO `menu` VALUES ('700201904034', '网站防火墙（waf）', '800201904006', '/index.html#/pages/admin/wafManage', '1', '网站防火墙（waf）', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904031', '网站防火墙（waf）', '网站防火墙（waf）', '2019-04-01 02:24:53', '0', '/pages/admin/wafManage', '700201904034');
INSERT INTO `privilege_rel` VALUES ('58', '500201904031', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('59', '500201904031', '600201904002', '2019-04-01 08:18:29', '0');


INSERT INTO `menu` VALUES ('700201904035', 'waf访问记录', '800201904006', '/index.html#/pages/admin/wafAccessLog', '1', 'waf访问', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904032', 'waf访问', 'waf访问', '2019-04-01 02:24:53', '0', '/pages/admin/wafAccessLog', '700201904035');
INSERT INTO `privilege_rel` VALUES ('60', '500201904032', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('61', '500201904032', '600201904002', '2019-04-01 08:18:29', '0');

INSERT INTO `menu` VALUES ('700201904036', 'waf路由', '800201904006', '/index.html#/pages/admin/wafRouteManage', '1', 'waf路由', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904033', 'waf访问', 'waf访问', '2019-04-01 02:24:53', '0', '/pages/admin/wafRouteManage', '700201904036');
INSERT INTO `privilege_rel` VALUES ('62', '500201904033', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('63', '500201904033', '600201904002', '2019-04-01 08:18:29', '0');


INSERT INTO `menu` VALUES ('700201904037', 'waf监控', '800201904006', '/waf/index.html', '4', 'waf监控', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904034', 'waf监控', 'waf访问', '2019-04-01 02:24:53', '0', '/waf/waf.html', '700201904037');
INSERT INTO `privilege_rel` VALUES ('64', '500201904034', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('65', '500201904034', '600201904002', '2019-04-01 08:18:29', '0');


INSERT INTO `t_dict` VALUES ('35', 'sqli', 'SQL注入', 'SQL注入', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('36', 'xss', '跨站脚本', '跨站脚本', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('20', 'webshell', 'WebShell攻击', 'WebShell攻击', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('21', 'robot', '恶意爬虫', '恶意爬虫', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('22', 'cmdi', '命令注入', '命令注入', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('23', 'rfi', '远程文件', '远程文件', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('24', 'lfi', '本地文件包含', '本地文件包含', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('25', 'illegal', '非法请求', '非法请求', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('26', 'vuln', '漏洞攻击', '漏洞攻击', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('27', 'cc', '命中CC防护', '命中CC防护', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('28', 'custom_custom', '命中精准防护', '命中精准防护', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('29', 'custom_whiteip', '命中IP黑白名单', '命中IP黑白名单', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('30', 'custom_geoip', '命中地理位置控制', '命中地理位置控制', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('31', 'antitamper', '命中网页防篡改', '命中网页防篡改', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('32', 'anticrawler', '命中JS挑战反爬虫', '命中JS挑战反爬虫', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('33', 'leakage', '命中敏感信息泄露', '命中敏感信息泄露', '2021-02-24 22:49:33', 'waf_access_log', 'state');
INSERT INTO `t_dict` VALUES ('34', 'followed_action', '攻击惩罚', '攻击惩罚', '2021-02-24 22:49:33', 'waf_access_log', 'state');

create table waf_rule_group
(
    group_id      varchar(64) not null
        primary key,
    group_name    varchar(64) not null,
    state varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table waf_rule
(
    rule_id     varchar(64) not null
        primary key,
    group_id varchar(64) not null,
    rule_name   varchar(64) not null,
    scope       varchar(64) not null,
    obj_id      varchar(64) not null,
    obj_type    varchar(64) not null,
    seq         int         not null,
    state       varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);


INSERT INTO `menu` VALUES ('700201904038', '安全组', '800201904006', '/index.html#/pages/admin/wafRuleGroupManage', '5', '安全组', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904035', '安全组', '安全组', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/wafRuleGroupManage', '700201904038');
INSERT INTO `privilege_rel` VALUES ('66', '500201904035', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('67', '500201904035', '600201904002', '2019-04-01 08:18:29', '0');

create table waf_ip_black_white
(
    id     varchar(64) not null
        primary key,
    type_cd    varchar(64) not null,
    ip   varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);


INSERT INTO `menu` VALUES ('700201904039', '黑白名单', '800201904006', '/index.html#/pages/admin/wafIpBlackWhiteManage', '5', '黑白名单', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904036', '黑白名单', '黑白名单', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/wafIpBlackWhiteManage', '700201904039');
INSERT INTO `privilege_rel` VALUES ('68', '500201904036', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('69', '500201904036', '600201904002', '2019-04-01 08:18:29', '0');

create table waf_area
(
    id          varchar(64) not null
        primary key,
    type_cd     varchar(64) not null,
    area_name          longtext not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);


INSERT INTO `menu` VALUES ('700201904040', '地理位置', '800201904006', '/index.html#/pages/admin/wafAreaManage', '6', '地理位置', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904040', '地理位置', '地理位置', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/wafAreaManage', '700201904040');
INSERT INTO `privilege_rel` VALUES ('70', '500201904040', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('71', '500201904040', '600201904002', '2019-04-01 08:18:29', '0');

create table waf_cc
(
    id          varchar(64) not null
        primary key,
    path     varchar(512) not null,
    visit_count   varchar(12)    not null,
    visit_sec   varchar(12)    not null,
    block_sec   varchar(12)    not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);


INSERT INTO `menu` VALUES ('700201904041', 'CC防护', '800201904006', '/index.html#/pages/admin/wafCCManage', '7', 'CC防护', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904041', '地理位置', '地理位置', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/wafCCManage', '700201904041');
INSERT INTO `privilege_rel` VALUES ('72', '500201904041', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('73', '500201904041', '600201904002', '2019-04-01 08:18:29', '0');

create table waf_accurate
(
    id          varchar(64)  not null
        primary key,
    action varchar(12) not null ,
    type_cd        varchar(12) not null,
    include varchar(512)  not null,
    include_value   varchar(512)  not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

INSERT INTO `menu` VALUES ('700201904042', '精准防护', '800201904006', '/index.html#/pages/admin/wafAccurateManage', '7', 'CC防护', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904042', '精准防护', '精准防护', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/wafAccurateManage', '700201904042');
INSERT INTO `privilege_rel` VALUES ('74', '500201904042', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('75', '500201904042', '600201904002', '2019-04-01 08:18:29', '0');

drop table vpn;
create table inner_net
(
    inner_net_id      varchar(64) not null
        primary key,
    inner_net_port    varchar(64) not null,
    tun        varchar(64) not null,
    tun_name  varchar(64) not null,
    dns  varchar(64) not null,
    protocol  varchar(64) not null,
    state       varchar(12) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

insert into inner_net(inner_net_id, inner_net_port, tun, tun_name, dns, protocol, state) VALUES ('1','5555','192.168.1.0/24','virName','8.8.8.8','tcp','2002');
create table inner_net_hosts
(
    inner_net_host_id varchar(64) not null
        primary key,
    inner_net_id      varchar(64) not null,
    host_id     varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table inner_net_users
(
    user_id varchar(64) not null
        primary key,
    username      varchar(64) not null,
    password     varchar(64) not null,
    tel varchar(64) not null,
    ip     varchar(64) not null,
    login_time datetime not null ,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);


INSERT INTO `menu_group` VALUES ('800201904014', '虚拟组网', 'fa fa-globe', '', '8', '虚拟组网', '2019-04-01 07:55:51', '0', 'P_WEB');


INSERT INTO `menu` VALUES ('700201904043', '中心节点', '800201904014', '/index.html#/pages/admin/innerNetManage', '1', '中心节点', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904043', '中心节点', '中心节点', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/innerNetManage', '700201904043');
INSERT INTO `privilege_rel` VALUES ('76', '500201904043', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('77', '500201904043', '600201904002', '2019-04-01 08:18:29', '0');


INSERT INTO `menu` VALUES ('700201904044', '组网电脑', '800201904014', '/index.html#/pages/admin/innerNetUserManage', '1', '组网电脑', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904044', '组网电脑', '组网电脑', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/innerNetUserManage', '700201904044');
INSERT INTO `privilege_rel` VALUES ('78', '500201904044', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('79', '500201904044', '600201904002', '2019-04-01 08:18:29', '0');

create table dns
(
    dns_id   varchar(64) not null
        primary key,
    dns_port varchar(64) not null,
    dns_ip            varchar(64) not null,
    state          varchar(12) not null,
    create_time    timestamp  default CURRENT_TIMESTAMP not null,
    status_cd      varchar(2) default '0' not null
);

create table dns_hosts
(
    dns_host_id varchar(64) not null
        primary key,
    dns_id      varchar(64) not null,
    host_id     varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table dns_map
(
    dns_map_id     varchar(64) not null
        primary key,
    host    varchar(64) not null,
    type    varchar(64) not null,
    value         varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

insert into dns(dns_id, dns_ip, dns_port, state) VALUES ('1','8.8.8.8','53','2002');

INSERT INTO `menu` VALUES ('700201904045', 'dns', '800201904014', '/index.html#/pages/admin/dnsManage', '4', 'dns', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904045', 'dns', 'dns', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/dnsManage', '700201904045');
INSERT INTO `privilege_rel` VALUES ('80', '500201904045', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('81', '500201904045', '600201904002', '2019-04-01 08:18:29', '0');

INSERT INTO `menu` VALUES ('700201904046', 'hosts', '800201904014', '/index.html#/pages/admin/dnsMapManage', '5', 'hosts', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904046', 'dns', 'dns', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/dnsMapManage', '700201904046');
INSERT INTO `privilege_rel` VALUES ('82', '500201904046', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('83', '500201904046', '600201904002', '2019-04-01 08:18:29', '0');

create table inner_net_privilege
(
    p_id     varchar(64) not null
        primary key,
    src_user_id    varchar(64) not null,
    target_user_id   varchar(64) not null,
    target_port varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);


INSERT INTO `menu` VALUES ('700201904047', '组网访问', '800201904014', '/index.html#/pages/admin/innerNetPrivilegeManage', '3', 'hosts', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904047', 'dns', 'dns', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/innerNetPrivilegeManage', '700201904047');
INSERT INTO `privilege_rel` VALUES ('84', '500201904047', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('85', '500201904047', '600201904002', '2019-04-01 08:18:29', '0');

create table inner_net_log
(
    log_id   varchar(64) not null
        primary key,
    username varchar(64) not null,
    ip            varchar(64) not null,
    src_ip       varchar(64) not null,
    state varchar(64) not null,
    create_time    timestamp  default CURRENT_TIMESTAMP not null,
    status_cd      varchar(2) default '0' not null
);


INSERT INTO `menu` VALUES ('700201904048', '组网日志', '800201904014', '/index.html#/pages/admin/innerNetLogManage', '4', 'hosts', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904048', 'dns', 'dns', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/innerNetLogManage', '700201904048');
INSERT INTO `privilege_rel` VALUES ('86', '500201904048', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('87', '500201904048', '600201904002', '2019-04-01 08:18:29', '0');

create table host_attr
(
    attr_id     varchar(64) not null  primary key,
    host_id     varchar(64) not null ,
    spec_cd     varchar(12) not null,
    value       varchar(50) not null,
    create_time timestamp   default CURRENT_TIMESTAMP not null,
    status_cd   varchar(12) default '0' not null
);

create table firewall_rule_group
(
    group_id    varchar(64) not null
        primary key,
    group_name  varchar(64) not null,
    state       varchar(64) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table host_firewall_group
(
    hfg_id    varchar(64) not null
        primary key,
    group_id        varchar(64) not null,
    host_id varchar(64) not null,
    create_time timestamp    default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2)   default '0' not null
);

create table firewall_rule
(
    rule_id     varchar(64) not null
        primary key,
    group_id    varchar(64) not null,
    in_out   varchar(64) not null,
    allow_limit       varchar(64) not null,
    seq      int not null,
    protocol    varchar(64) not null,
    src_obj         varchar(64)         not null,
    dst_obj       varchar(64) not null,
    remark varchar(512) ,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);



INSERT INTO `menu` VALUES ('700201904049', '开放端口', '800201904006', '/index.html#/pages/admin/firewallRuleGroupManage', '8', 'hosts', '2019-04-09 14:50:56', '0', 'Y');
INSERT INTO `privilege` VALUES ('500201904049', '开放端口', 'dns', '2019-04-01 02:24:53', '0', '/index.html#/pages/admin/firewallRuleGroupManage', '700201904049');
INSERT INTO `privilege_rel` VALUES ('88', '500201904049', '600201904000', '2019-04-01 08:18:29', '0');
INSERT INTO `privilege_rel` VALUES ('89', '500201904049', '600201904002', '2019-04-01 08:18:29', '0');


create table workflow_step
(
    step_id       varchar(64)  not null
        primary key,
    step     varchar(64) not null,
    name      varchar(64) not null,
    create_time  timestamp   default CURRENT_TIMESTAMP not null,
    status_cd    varchar(2)  default '0' not null
);

create table workflow_step_param
(
    param_id varchar(64) not null primary key,
    step_id     varchar(64) not null,
    param_name  varchar(64) not null,
    param_spec  varchar(64) not null,
    param_value varchar(128),
    seq varchar(11) not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);

create table workflow
(
    workflow_id    varchar(64) not null
        primary key,
    name     varchar(64) not null,
    yaml  longtext,
    tenant_id  varchar(64) not null,
    state varchar(12),
    job_time     datetime     not null,
    create_time timestamp  default CURRENT_TIMESTAMP not null,
    status_cd   varchar(2) default '0' not null
);