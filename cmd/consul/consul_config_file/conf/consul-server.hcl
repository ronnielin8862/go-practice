# consul-server.hcl

server = true
bootstrap_expect = 1
ui = true
datacenter = "ithome"
data_dir = "/Users/ronnie/Library/Mobile Documents/com~apple~CloudDocs/Documents/coding/code/consul/data"
disable_update_check = true
enable_local_script_checks = true

node_name = "consul-server-1"
client_addr = "127.0.0.1"
bind_addr = "127.0.0.1"

connect {
  enabled = true
}

log_level = "DEBUG"
log_file = "/Users/ronnie/Library/Mobile Documents/com~apple~CloudDocs/Documents/coding/code/consul/logs/"
log_rotate_duration = "24h"
log_rotate_max_files = 0