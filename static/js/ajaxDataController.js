/*$.extend({
	addLoading:function(){
   	 //插入loading
       var html = "";
           html += '<div class="js_loading centered loading" style="display:none;">';
           html +=     '<i class="fa fa-spinner fa-pulse"></i>';
           html += '</div>';
       $("body").append(html);
   }
});
$.addLoading();*/


var myAjax = function () {

    //打印ajax错误日志
    function printLog(result, url, params, response) {
        console.error('AJAX 请求异常 - %s\n错误信息：\n%c%s\n%c请求链接：%s\n%c请求参数：%c%s\n%c返回数据：%c%s',
            'color:red;',
            result,
            'color:#333;',
            'color:blue',
            url,
            'color:#333;',
            'color:green',
            JSON.stringify(params),
            'color:#333;',
            'color:#643A3A',
            response)
    }

    function dataHandle(url, params, callback, async, method) {

        if (!method) {
            throw 'method 参数未设置'
        }

        if ((typeof params) === 'function') {
            callback = params
            params = null
        }

        params = params || {};
        params = $.extend({ date: new Date().getTime().toString() }, params);
        async = async == null ? true : async;

        var ERROR_PROCESS_MODE = 0;

        if (typeof (params.ERROR_PROCESS_MODE) != "undefined") {
            if (params.ERROR_PROCESS_MODE==1) {
                ERROR_PROCESS_MODE = 1;
                try {
                    delete params.ERROR_PROCESS_MODE;
                } catch (e) {

                }
               
            }
        } 



        $.ajax({
            async: async,
            url: url,
            dataType: 'json',
            data: params,
            type: method,
            beforeSend: function(){  //开始loading
               // $(".js_loading").show();
               // common.mui.showLoading();
            },
            success: function (result, textStatus, xhr) {
                if (result.success === false) {
                    printLog(result, url, params, xhr.responseText)
                }
                callback(result);

            },
            error: function (xhr, textStatus, error) {
            	var loginHtml =  xhr.responseText;
            	if(loginHtml!=undefined){
            		if(loginHtml.indexOf('window.location.href = "/login.tpl"')>0){
                        //common.mui.alert("登录验证失败，请重新登录");
                	}
            	}
            },
            complete: function(){   //结束loading
            	setTimeout(function () { 
            		 //$(".js_loading").hide();
                    //common.mui.hideLoading();
            	}, 300);  
            }
        });
    }

    return {
        post: function (url, params, callback, async) {
            dataHandle(url, params, callback, async, 'post');
        },
        get: function (url, params, callback, async) {
            dataHandle(url, params, callback, async, 'get');
        }
    };
}