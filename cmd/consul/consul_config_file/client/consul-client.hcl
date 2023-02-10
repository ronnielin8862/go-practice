datacenter = "ithome"
data_dir = "/Users/ronnie/Library/Mobile Documents/com~apple~CloudDocs/Documents/coding/code/consul/client/data"
disable_update_check = true

node_name = "consul-client-1"
# client_addr = "0.0.0.0"
bind_addr = "127.0.0.1"

connect {
  enabled = true
}

retry_join  = ["127.0.0.1"]
retry_interval = "20s"



log_level = "DEBUG"
log_file = "/Users/ronnie/Library/Mobile Documents/com~apple~CloudDocs/Documents/coding/code/consul/client/logs/"
log_rotate_duration = "24h"
log_rotate_max_files = 0

performance {
  raft_multiplier = 1
}