--cmd = require('cmd')
--http = require('http')
--time = require('time')
--
--print("开始下载矿机程序")
--err = http.download('https://gengwenjuan.oss-cn-beijing.aliyuncs.com/ytfs-node', 'ytfs-node')
--if err == nil then
--    print("矿机程序下载完成")
--else
--    print("矿机程序下载失败",err)
--end
--
--
--time.sleep(1000)
--print("启动矿机程序")
--err = cmd.command('./ytfs-node daemon').run()
--print(err)
--

time = require("time")
print("111")
time.sleep(5000)
print("222")