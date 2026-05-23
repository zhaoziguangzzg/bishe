后端使用golang + gin
使用golang渲染html页面
前端使用Vue3 + ElementPlus

web\views存放html模板 
每个page使用独立的html模板

前台页面除了登录和注册页，其他页面都使用layout/layout.html模板，参考path /page/circle/detail，handler注意调用RenderIndexPage，注意不要出现index_layout_extra_scripts部分渲染两次的问题，使用layout中的.container样式，不要自己定义.container样式  
后台页面都使用admin/_layout.html模板  

