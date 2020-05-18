time = require("time")
process = require("process")
cmd = require("cmd")
util = require("util")

while (true) do
    time.sleep(time.minute * 10)

    remote_version,download_url = util.get_remote_version_info()
    if remote_version > util.get_current_version() then
        err = util.download_ytfs_node(remote_version,download_url)
        if err == nil then
            process.killSelf()
            print("更新完成")
        else
            print("更新失败",err)
        end
    end
end

