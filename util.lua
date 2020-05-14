config = require("config")
http = require("http")

module = {}

function module.download_ytfs_node()
    config.reset()
    config.setType("yaml")
    cfg_str = http.get('http://dnapi.yottachain.net/update-config/update.yaml')
    config.read(cfg_str)

    remote_version = config.getInt(string.format('%s.%s.remote_version',L_ARCH,L_OS))
    download_url = config.getString(string.format('%s.%s.download_url',L_ARCH,L_OS))

    http.download(download_url,"ytfs-node")
end

return module