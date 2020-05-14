time = require("time")
process = require("process")
cmd = require("cmd")

while (true) do
    time.sleep(time.minute * 10)
    res,_ = cmd.exec('./ytfs-node update')
    if string.match(res,"更新完成") == "更新完成" then
        process.killSelf()
    end
end

