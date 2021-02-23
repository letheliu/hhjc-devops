/**
    入驻小区
**/
(function (vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            monitorHostGroupManageInfo: {
                monitorHostGroups: [],
                total: 0,
                records: 1,
                moreCondition: false,
                mhgId: '',
                conditions: {
                    name: '',
                    state: '',
                    noticeType: '',

                }
            }
        },
        _initMethod: function () {
            vc.component._listMonitorHostGroups(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function () {

            vc.on('monitorHostGroupManage', 'listMonitorHostGroup', function (_param) {
                vc.component._listMonitorHostGroups(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function (_currentPage) {
                vc.component._listMonitorHostGroups(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listMonitorHostGroups: function (_page, _rows) {

                vc.component.monitorHostGroupManageInfo.conditions.page = _page;
                vc.component.monitorHostGroupManageInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.monitorHostGroupManageInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/monitor/getMonitorHosts',
                    param,
                    function (json, res) {
                        var _monitorHostGroupManageInfo = JSON.parse(json);
                        vc.component.monitorHostGroupManageInfo.total = _monitorHostGroupManageInfo.total;
                        vc.component.monitorHostGroupManageInfo.records = _monitorHostGroupManageInfo.records;
                        vc.component.monitorHostGroupManageInfo.monitorHostGroups = _monitorHostGroupManageInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.monitorHostGroupManageInfo.records,
                            currentPage: _page
                        });
                    }, function (errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openAddMonitorHostGroupModal: function () {
                vc.emit('addMonitorHostGroup', 'openAddMonitorHostGroupModal', {});
            },
            _openEditMonitorHostGroupModel: function (_monitorHostGroup) {
                vc.emit('editMonitorHostGroup', 'openEditMonitorHostGroupModal', _monitorHostGroup);
            },
            _openDeleteMonitorHostGroupModel: function (_monitorHostGroup) {
                vc.emit('deleteMonitorHostGroup', 'openDeleteMonitorHostGroupModal', _monitorHostGroup);
            },
            _queryMonitorHostGroupMethod: function () {
                vc.component._listMonitorHostGroups(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition: function () {
                if (vc.component.monitorHostGroupManageInfo.moreCondition) {
                    vc.component.monitorHostGroupManageInfo.moreCondition = false;
                } else {
                    vc.component.monitorHostGroupManageInfo.moreCondition = true;
                }
            }


        }
    });
})(window.vc);
