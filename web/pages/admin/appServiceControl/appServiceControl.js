/**
 入驻小区
 **/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    var TEMP_SEARCH = "simplifyAcceptanceSearch";
    vc.extends({
        data: {
            appServiceControlInfo: {
                _currentTab: 'appServiceControlPort',
                asId: '',
                asName: '',
                asCount: '',
                stateName: '',
                imagesName: '',
                imagesVersion: '',
                asType: '',
                avgName: '',
                hostGroupName: '',
                hostName: '',
            }
        },
        _initMethod: function() {
            $that.appServiceControlInfo.asId = vc.getParam('asId');
            $that._listAppServices();
        },
        _initEvent: function() {
            vc.on('simplifyAcceptance', 'chooseRoom', function(_room) {
                vc.copyObject(_room, $that.appServiceControlInfo);
                vc.emit('simplifyRoomFee', 'switch', $that.appServiceControlInfo)
            });
        },
        methods: {
            changeTab: function(_tab) {
                $that.appServiceControlInfo._currentTab = _tab;
                vc.emit(_tab, 'switch', {
                    asId: $that.appServiceControlInfo.asId
                })
            },
            _clearData: function() {
                $that.appServiceControlInfo = {
                    _currentTab: 'hostContainers',
                    asId: '',
                    asName: '',
                    asCount: '',
                    stateName: '',
                    imagesName: '',
                    imagesVersion: '',
                    asType: '',
                    avgName: '',
                    hostGroupName: '',
                    hostName: '',
                }
            },
            _listAppServices: function() {

                let param = {
                    params: {
                        page: 1,
                        row: 1,
                        asId: $that.appServiceControlInfo.asId
                    }
                };

                //发送get请求
                vc.http.apiGet('/appService/getAppService',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);

                        vc.copyObject(_hostManageInfo.data[0], $that.appServiceControlInfo)

                        $that.changeTab($that.appServiceControlInfo._currentTab);

                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _openSsh: function() {
                //获取主机访问token
                var param = {
                    params: {
                        hostId: $that.appServiceControlInfo.hostId
                    }
                };

                //发送get请求
                vc.http.apiGet('/host/getHostToken',
                    param,
                    function(json, res) {
                        let _hostManageInfo = JSON.parse(json);
                        let _zihaoToken = _hostManageInfo.data;
                        window.open("/webshell/console.html?hostId=" + $that.appServiceControlInfo.hostId + "&zihaoToken=" + _zihaoToken, '_blank')
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _goBack: function() {
                vc.goBack();
            },
            _getAsType: function(_asType) {
                if (_asType == '001') {
                    return '数据库';
                } else if (_asType == '002') {
                    return '缓存';
                }
                return '计算应用';

            }

        }
    });
})(window.vc);