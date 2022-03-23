/**
    入驻小区
**/
(function(vc){
    var DEFAULT_PAGE = 1;
    var DEFAULT_ROWS = 10;
    vc.extends({
        data:{
            wafManageInfo:{
                wafs:[],
                total:0,
                records:1,
                moreCondition:false,
                wafId:'',
                conditions:{
                    wafId:'',
wafName:'',
state:'',

                }
            }
        },
        _initMethod:function(){
            vc.component._listWafs(DEFAULT_PAGE, DEFAULT_ROWS);
        },
        _initEvent:function(){
            
            vc.on('wafManage','listWaf',function(_param){
                  vc.component._listWafs(DEFAULT_PAGE, DEFAULT_ROWS);
            });
             vc.on('pagination','page_event',function(_currentPage){
                vc.component._listWafs(_currentPage,DEFAULT_ROWS);
            });
        },
        methods:{
            _listWafs:function(_page, _rows){

                vc.component.wafManageInfo.conditions.page = _page;
                vc.component.wafManageInfo.conditions.row = _rows;
                var param = {
                    params:vc.component.wafManageInfo.conditions
               };

               //发送get请求
               vc.http.apiGet('waf.listWafs',
                             param,
                             function(json,res){
                                var _wafManageInfo=JSON.parse(json);
                                vc.component.wafManageInfo.total = _wafManageInfo.total;
                                vc.component.wafManageInfo.records = _wafManageInfo.records;
                                vc.component.wafManageInfo.wafs = _wafManageInfo.data;
                                vc.emit('pagination','init',{
                                     total:vc.component.wafManageInfo.records,
                                     currentPage:_page
                                 });
                             },function(errInfo,error){
                                console.log('请求失败处理');
                             }
                           );
            },
            _openAddWafModal:function(){
                vc.emit('addWaf','openAddWafModal',{});
            },
            _openEditWafModel:function(_waf){
                vc.emit('editWaf','openEditWafModal',_waf);
            },
            _openDeleteWafModel:function(_waf){
                vc.emit('deleteWaf','openDeleteWafModal',_waf);
            },
            _queryWafMethod:function(){
                vc.component._listWafs(DEFAULT_PAGE, DEFAULT_ROWS);

            },
            _moreCondition:function(){
                if(vc.component.wafManageInfo.moreCondition){
                    vc.component.wafManageInfo.moreCondition = false;
                }else{
                    vc.component.wafManageInfo.moreCondition = true;
                }
            }

             
        }
    });
})(window.vc);
