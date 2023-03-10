(function(vc, vm) {

    vc.extends({
        data: {
            editTenantSettingInfo: {
                settingId: '',
                specCd: '',
                value: '',
                specCds: []
            }
        },
        _initMethod: function() {

            vc.getDict('tenant_setting', 'spec_cd', function(data) {

                $that.editTenantSettingInfo.specCds = data.data;
            })

        },
        _initEvent: function() {
            vc.on('editTenantSetting', 'openEditTenantSettingModal', function(_params) {
                vc.component.refreshEditTenantSettingInfo();
                $('#editTenantSettingModel').modal('show');
                vc.copyObject(_params, vc.component.editTenantSettingInfo);
            });
        },
        methods: {
            editTenantSettingValidate: function() {
                return vc.validate.validate({
                    editTenantSettingInfo: vc.component.editTenantSettingInfo
                }, {
                    'editTenantSettingInfo.specCd': [{
                            limit: "required",
                            param: "",
                            errInfo: "规格不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "64",
                            errInfo: "规格错误"
                        },
                    ],
                    'editTenantSettingInfo.value': [{
                            limit: "required",
                            param: "",
                            errInfo: "值不能为空"
                        },
                        {
                            limit: "maxLength",
                            param: "1024",
                            errInfo: "值太长"
                        },
                    ],
                    'editTenantSettingInfo.settingId': [{
                        limit: "required",
                        param: "",
                        errInfo: "ID不能为空"
                    }]

                });
            },
            editTenantSetting: function() {
                if (!vc.component.editTenantSettingValidate()) {
                    vc.toast(vc.validate.errInfo);
                    return;
                }

                vc.http.apiPost(
                    '/tenant/updateTenantSetting',
                    JSON.stringify(vc.component.editTenantSettingInfo), {
                        emulateJSON: true
                    },
                    function(json, res) {
                        //vm.menus = vm.refreshMenuActive(JSON.parse(json),0);
                        let _json = JSON.parse(json);
                        if (_json.code == 0) {
                            //关闭model
                            $('#editTenantSettingModel').modal('hide');
                            vc.emit('tenantSettingManage', 'listTenantSetting', {});
                            return;
                        }
                        vc.toast(_json.msg);
                    },
                    function(errInfo, error) {
                        console.log('请求失败处理');

                        vc.toast(errInfo);
                    });
            },
            refreshEditTenantSettingInfo: function() {
                let _specCds = $that.addTenantSettingInfo.specCds
                vc.component.editTenantSettingInfo = {
                    settingId: '',
                    specCd: '',
                    value: '',
                    specCds: _specCds
                }
            }
        }
    });

})(window.vc, window.vc.component);