## PVE Core
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/version`                                                  |:material-check-all:|

## Access
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/access/acl`                                               |:material-close:||:material-close:|
| `/access/password`                                          |||:material-close:|
| `/access/permissions`                                       |:material-close:|
| `/access/ticket`                                            |:material-close:|:material-close:|

### Users
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/access/users`                                             |:material-close:|:material-close:|
| `/access/users/:userId`                                     |:material-close:||:material-close:|:material-close:|
| `/access/users/:userId/tfa`                                 |:material-close:|
| `/access/users/:userId/unlock-tfa`                          |||:material-close:|
| `/access/users/:userId/token`                               |:material-close:|
| `/access/users/:userId/token/:tokenId`                      |:material-close:|:material-close:|:material-close:|:material-close:|

### TFA
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/access/tfa`                                               |:material-close:|
| `/access/tfa/:userId`                                       |:material-close:|:material-close:|
| `/access/tfa/:userId/:id`                                   |:material-close:||:material-close:|:material-close:|

### Roles
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/access/roles`                                             |:material-close:|:material-close:|
| `/access/roles/:roleId`                                     |:material-close:||:material-close:|:material-close:|

### OpenId
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/access/openid/auth-url`                                   ||:material-close:|
| `/access/openid/login`                                      ||:material-close:|


### Groups
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/access/groups`                                            |:material-close:|:material-close:|
| `/access/groups/:groupId`                                   |:material-close:||:material-close:|:material-close:|

### Domains
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/access/domains`                                           |:material-close:|:material-close:|
| `/access/domains/:realmId`                                  |:material-close:||:material-close:|:material-close:|
| `/access/domains/:realmId/sync`                             ||:material-close:|

## Cluster
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster`                                                  |:material-close:|
| `/cluster/log`                                              |:material-close:|
| `/cluster/nextid`                                           |:material-close:|
| `/cluster/options`                                          |:material-close:||:material-close:|
| `/cluster/resources`                                        |:material-close:|
| `/cluster/status`                                           |:material-close:|
| `/cluster/tasks`                                            |:material-close:|

### Acme
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/acme`                                             |:material-close:|
| `/cluster/acme/challenge-schema`                            |:material-close:|
| `/cluster/acme/directories`                                 |:material-close:|
| `/cluster/acme/meta`                                        |:material-close:|
| `/cluster/acme/tos`                                         |:material-close:|
| `/cluster/acme/account`                                     |:material-close:|:material-close:|
| `/cluster/acme/account/:name`                               |:material-close:||:material-close:|:material-close:|
| `/cluster/acme/plugins`                                     |:material-close:|:material-close:|
| `/cluster/acme/plugins/:id`                                 |:material-close:||:material-close:|:material-close:|


### Backup
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/backup`                                           |:material-close:|:material-close:|
| `/cluster/backup/:id`                                       |:material-close:||:material-close:|:material-close:|
| `/cluster/backup/:id/included_volumes`                      |:material-close:|

### Backup info
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/backup-info`                                      |:material-close:|
| `/cluster/backup-info/not-backed-up`                        |:material-close:|

### Ceph
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/ceph`                                             |:material-close:|
| `/cluster/ceph/metadata`                                    |:material-close:|
| `/cluster/ceph/status`                                      |:material-close:|
| `/cluster/ceph/flags`                                       |:material-close:||:material-close:|
| `/cluster/ceph/flags/:flag`                                 |:material-close:||:material-close:|

### Config
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/config`                                           |:material-close:|:material-close:|
| `/cluster/config/apiversion`                                |:material-close:|
| `/cluster/config/join`                                      |:material-close:|:material-close:|
| `/cluster/config/qdevice`                                   |:material-close:|:material-close:|
| `/cluster/config/totem`                                     |:material-close:|
| `/cluster/config/nodes`                                     |:material-close:|
| `/cluster/config/nodes/:node`                               ||:material-close:||:material-close:|

### Firewall
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/firewall`                                         |:material-close:|
| `/cluster/firewall/macros`                                  |:material-close:|
| `/cluster/firewall/options`                                 |:material-close:||:material-close:|
| `/cluster/firewall/refs`                                    |:material-close:|
| `/cluster/firewall/aliases`                                 |:material-close:|:material-close:|
| `/cluster/firewall/aliases/:name`                           |:material-close:||:material-close:|:material-close:|
| `/cluster/firewall/groups`                                  |:material-close:|:material-close:|
| `/cluster/firewall/groups/:group`                           |:material-close:|:material-close:||:material-close:|
| `/cluster/firewall/groups/:group/:pos`                      |:material-close:||:material-close:|:material-close:|
| `/cluster/firewall/ipset`                                   |:material-close:|:material-close:|
| `/cluster/firewall/ipset/:name`                             |:material-close:|:material-close:||:material-close:|
| `/cluster/firewall/ipset/:name/:cidr`                       |:material-close:||:material-close:|:material-close:|
| `/cluster/firewall/rules`                                   |:material-close:|:material-close:|
| `/cluster/firewall/rules/:pos`                              |:material-close:||:material-close:|:material-close:|

### High availability
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/ha`                                               |:material-close:|
| `/cluster/ha/groups`                                        |:material-close:|:material-close:|
| `/cluster/ha/groups/:group`                                 |:material-close:||:material-close:|:material-close:|
| `/cluster/ha/resources`                                     |:material-close:|:material-close:|
| `/cluster/ha/resources/:sid`                                |:material-close:||:material-close:|:material-close:|
| `/cluster/ha/resources/:sid/migrate`                        ||:material-close:|
| `/cluster/ha/resources/:sid/relocate`                       ||:material-close:|
| `/cluster/ha/status`                                        |:material-close:|
| `/cluster/ha/status/current`                                |:material-close:|
| `/cluster/ha/status/manager_status`                         |:material-close:|

### Jobs
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/jobs`                                             |:material-close:|
| `/cluster/jobs/schedule-analyze`                            |:material-close:|
| `/cluster/jobs/realm-sync`                                  |:material-close:|
| `/cluster/jobs/realm-sync/:id`                              |:material-close:|:material-close:|:material-close:|:material-close:|

### Mapping
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/mapping`                                          |:material-close:|
| `/cluster/mapping/pci`                                      |:material-close:|:material-close:|
| `/cluster/mapping/pci/:id`                                  |:material-close:||:material-close:|:material-close:|
| `/cluster/mapping/usb`                                      |:material-close:|:material-close:|
| `/cluster/mapping/usb/:id`                                  |:material-close:||:material-close:|:material-close:|

### Metrics
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/metrics`                                          |:material-close:|
| `/cluster/metrics/export`                                   |:material-close:|
| `/cluster/metrics/server`                                   |:material-close:|
| `/cluster/metrics/server/:id`                               |:material-close:|:material-close:|:material-close:|:material-close:|

### Notifications
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/notifications`                                    |:material-close:|
| `/cluster/notifications/matcher-field-values`               |:material-close:|
| `/cluster/notifications/matcher-fields`                     |:material-close:|
| `/cluster/notifications/endpoints`                          |:material-close:|
| `/cluster/notifications/endpoints/gotify`                   |:material-close:|:material-close:|
| `/cluster/notifications/endpoints/gotify/:name`             |:material-close:||:material-close:|:material-close:|
| `/cluster/notifications/endpoints/sendmail`                 |:material-close:|:material-close:|
| `/cluster/notifications/endpoints/sendmail/:name`           |:material-close:||:material-close:|:material-close:|
| `/cluster/notifications/endpoints/smpt`                     |:material-close:|:material-close:|
| `/cluster/notifications/endpoints/smpt/:name`               |:material-close:||:material-close:|:material-close:|
| `/cluster/notifications/endpoints/webhook`                  |:material-close:|:material-close:|
| `/cluster/notifications/endpoints/webhook/:name`            |:material-close:||:material-close:|:material-close:|
| `/cluster/notifications/matchers`                           |:material-close:|:material-close:|
| `/cluster/notifications/matchers/:name`                     |:material-close:||:material-close:|:material-close:|
| `/cluster/notifications/targets`                            |:material-close:|
| `/cluster/notifications/targets/:name`                      |:material-close:||:material-close:|:material-close:|

### Replication
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/replication`                                      |:material-close:|:material-close:|
| `/cluster/replication/:id`                                  |:material-close:||:material-close:|:material-close:|

### SDN
| path                                                        | get                | post           | put | delete |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/cluster/sdn`                                              |:material-close:||:material-close:|
| `/cluster/sdn/controllers`                                  |:material-close:|:material-close:|
| `/cluster/sdn/controllers/:controller`                      |:material-close:||:material-close:|:material-close:|
| `/cluster/sdn/dns`                                          |:material-close:|:material-close:|
| `/cluster/sdn/dns/:dns`                                     |:material-close:||:material-close:|:material-close:|
| `/cluster/sdn/ipams`                                        |:material-close:|:material-close:|
| `/cluster/sdn/ipams/:ipam`                                  |:material-close:||:material-close:|:material-close:|
| `/cluster/sdn/vnets`                                        |:material-close:|:material-close:|
| `/cluster/sdn/vnets/:vnet`                                  |:material-close:||:material-close:|:material-close:|
| `/cluster/sdn/vnets/:vnet/ips`                              ||:material-close:|:material-close:|:material-close:|
| `/cluster/sdn/vnets/:vnet/firewall`                         |:material-close:|
| `/cluster/sdn/vnets/:vnet/firewall/options`                 |:material-close:||:material-close:|
| `/cluster/sdn/vnets/:vnet/firewall/rules`                   |:material-close:|:material-close:|
| `/cluster/sdn/vnets/:vnet/firewall/rules/:pos`              |:material-close:||:material-close:|:material-close:|
| `/cluster/sdn/vnets/:vnet/subnets`                          |:material-close:|:material-close:|
| `/cluster/sdn/vnets/:vnet/subnets/:subnet`                  |:material-close:||:material-close:|:material-close:|
| `/cluster/sdn/zones`                                        |:material-close:|:material-close:|
| `/cluster/sdn/zones/:zone`                                  |:material-close:||:material-close:|:material-close:|

## Nodes
### LXC
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/lxc`                                          |:material-check-all:|:material-check:|
| `/nodes/:node/lxc/:vmid`                                    |:material-close:| | |:material-close:|
| `/nodes/:node/lxc/:vmid/firewall`                           |:material-close:|
| `/nodes/:node/lxc/:vmid/firewall/log`                       |:material-close:|
| `/nodes/:node/lxc/:vmid/firewall/option`                    |:material-close:||:material-close:|
| `/nodes/:node/lxc/:vmid/firewall/refs`                      |:material-close:|
| `/nodes/:node/lxc/:vmid/firewall/aliases`                   |:material-close:|:material-close:|
| `/nodes/:node/lxc/:vmid/firewall/aliases/:name`             |:material-close:||:material-close:|:material-close:|
| `/nodes/:node/lxc/:vmid/firewall/ipset`                     |:material-close:|:material-close:|
| `/nodes/:node/lxc/:vmid/firewall/ipset/:name`               |:material-close:|:material-close:||:material-close:|
| `/nodes/:node/lxc/:vmid/firewall/ipset/:name/:cidr`         |:material-close:||:material-close:|:material-close:|
| `/nodes/:node/lxc/:vmid/firewall/rules`                     |:material-close:|:material-close:|
| `/nodes/:node/lxc/:vmid/firewall/rules/:pos`                |:material-close:||:material-close:|:material-close:|
| `/nodes/:node/lxc/:vmid/snapshot`                           |:material-close:|:material-close:|
| `/nodes/:node/lxc/:vmid/snapshot/:name`                     |:material-close:|||:material-close:|
| `/nodes/:node/lxc/:vmid/snapshot/:name/config`              |:material-close:||:material-close:|
| `/nodes/:node/lxc/:vmid/snapshot/:name/rollback`            ||:material-close:|
| `/nodes/:node/lxc/:vmid/status`                             |:material-close:|
| `/nodes/:node/lxc/:vmid/status/current`                     |:material-close:|
| `/nodes/:node/lxc/:vmid/status/reboot`                      ||:material-close:|
| `/nodes/:node/lxc/:vmid/status/resume`                      ||:material-close:|
| `/nodes/:node/lxc/:vmid/status/shutdown`                    ||:material-close:|
| `/nodes/:node/lxc/:vmid/status/start`                       ||:material-close:|
| `/nodes/:node/lxc/:vmid/status/stop`                        ||:material-close:|
| `/nodes/:node/lxc/:vmid/status/suspend`                     ||:material-close:|
| `/nodes/:node/lxc/:vmid/clone`                              ||:material-close:|
| `/nodes/:node/lxc/:vmid/config`                             |:material-close:||:material-close:|
| `/nodes/:node/lxc/:vmid/feature`                            |:material-close:|
| `/nodes/:node/lxc/:vmid/interfaces`                         |:material-close:|
| `/nodes/:node/lxc/:vmid/migrate`                            ||:material-close:|
| `/nodes/:node/lxc/:vmid/move_volume`                        ||:material-close:|
| `/nodes/:node/lxc/:vmid/mtunnel`                            ||:material-close:|
| `/nodes/:node/lxc/:vmid/mtunnelwebsocket`                   |:material-close:|
| `/nodes/:node/lxc/:vmid/pending`                            |:material-close:|
| `/nodes/:node/lxc/:vmid/remote_migrate`                     ||:material-close:|
| `/nodes/:node/lxc/:vmid/rdd`                                |:material-close:|
| `/nodes/:node/lxc/:vmid/rdddata`                            |:material-close:|
| `/nodes/:node/lxc/:vmid/spiceproxy`                         ||:material-close:|
| `/nodes/:node/lxc/:vmid/template`                           ||:material-close:|
| `/nodes/:node/lxc/:vmid/termproxy`                          ||:material-close:|
| `/nodes/:node/lxc/:vmid/vncproxy`                           ||:material-close:|
| `/nodes/:node/lxc/:vmid/vncwebsocket`                       |:material-close:|

