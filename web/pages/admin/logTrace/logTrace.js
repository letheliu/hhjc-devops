/**
    入驻小区
**/
(function(vc) {
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data: {
            logTraceInfo: {
                traces: [],
                total: 0,
                records: 1,
                moreCondition: false,
                mhgId: '',
                conditions: {
                    name: '',
                    traceId: '',
                    id: '',
                    parentSpanId: ''

                }
            }
        },
        _initMethod: function() {
            vc.component._listLogTraces(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent: function() {

            vc.on('monitorHostGroupManage', 'listMonitorHostGroup', function(_param) {
                vc.component._listLogTraces(DEFAULT_PAGE, DEFAULT_ROWS);
            });
            vc.on('pagination', 'page_event', function(_currentPage) {
                vc.component._listLogTraces(_currentPage, DEFAULT_ROWS);
            });
        },
        methods: {
            _listLogTraces: function(_page, _rows) {

                vc.component.logTraceInfo.conditions.page = _page;
                vc.component.logTraceInfo.conditions.row = _rows;
                var param = {
                    params: vc.component.logTraceInfo.conditions
                };

                //发送get请求
                vc.http.apiGet('/monitor/getLogTrace',
                    param,
                    function(json, res) {
                        var _logTraceInfo = JSON.parse(json);
                        vc.component.logTraceInfo.total = _logTraceInfo.total;
                        vc.component.logTraceInfo.records = _logTraceInfo.records;
                        vc.component.logTraceInfo.traces = _logTraceInfo.data;
                        vc.emit('pagination', 'init', {
                            total: vc.component.logTraceInfo.records,
                            currentPage: _page
                        });
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');
                    }
                );
            },
            _queryMonitorHostGroupMethod: function() {
                vc.component._listLogTraces(DEFAULT_PAGE, DEFAULT_ROWS);
            },
            _moreCondition: function() {
                if (vc.component.logTraceInfo.moreCondition) {
                    vc.component.logTraceInfo.moreCondition = false;
                } else {
                    vc.component.logTraceInfo.moreCondition = true;
                }
            },
            _toTraceDetail: function(_trace) {
                vc.jumpToPage('/index.html#/pages/admin/logTraceDetail?traceId=' + _trace.traceId)
            }

        }
    });
})(window.vc);