config = require("config")
http = require("http")

module = {}

local current_version = 0

function module.download_ytfs_node(remote_version, download_url)
    err = http.download(download_url,"ytfs-node")
    if err == nil then
        current_version = remote_version
    end

    return err
end

function module.get_remote_version_info ()
    config.reset()
    config.setType("yaml")
    cfg_str = http.get('http://dnapi.yottachain.net/update-config/update.yaml')
    config.read(cfg_str)

    remote_version = config.getInt(string.format('%s.%s.remote_version',L_ARCH,L_OS))
    download_url = config.getString(string.format('%s.%s.download_url',L_ARCH,L_OS))

    return remote_version,download_url
end

function module.get_current_version()
    return current_version
end

return module