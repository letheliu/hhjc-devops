(function(vc, vm) {

    vc.extends({
        data: {
            renameFileInfo: {
                hostId:'',
                fileName:'',
                newFileName:'',
                fileGroupName:'',
                curPath:''
            }
        },
        _initMethod: function() {

        },
        _initEvent: function() {
            vc.on('renameFile', 'openRenameFileModal', function(_params) {
                vc.component.refreshRenameFileInfo();
                $('#renameFileModel').modal('show');
                vc.copyObject(_params, vc.component.renameFileInfo);
            });
        },
        methods: {
            renameFile: function() {
                let _curPath = $that.renameFileInfo.curPath;
                let _newCurPath = $that.renameFileInfo.curPath;

                if (!_curPath.endsWith('/')) {
                    _curPath += ('/'+ $that.renameFileInfo.fileName);
                    _newCurPath += ('/'+ $that.renameFileInfo.newFileName);
                }else{
                    _curPath += ( $that.renameFileInfo.fileName);
                    _newCurPath += ( $that.renameFileInfo.newFileName);
                }
                let _data = {
                    hostId:$that.renameFileInfo.hostId,
                    fileName:_curPath,
                    newFileName: _newCurPath,
                    fileGroupName:$that.renameFileInfo.fileGroupName,
                }
                vc.http.apiPost(
                    '/host/renameFile',
                    JSON.stringify(_data), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#renameFileModel').modal('hide');
                            vc.emit('fileManager','listFiles', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshRenameFileInfo: function() {
                vc.component.renameFileInfo = {
                    hostId:'',
                    fileName:'',
                    newFileName:'',
                    fileGroupName:'',
                    curPath:'',
                }
            },
           
        }
    });

})(window.vc, window.vc.component);