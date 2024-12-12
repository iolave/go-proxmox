This is a development-purposed list to keep track of the implemented proxmox [endpoints](https://pve.proxmox.com/pve-docs/api-viewer/) as a callable golang func but **not** necessarily as a CLI command.

| Symbol | Description |
|:------:|:-----------:|
|:material-close:|Not implemented|
|:material-check:|Partially implemented (notes will be added in the docs)|
|:material-check-all:|Fully implemented|

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
| `/cluster/nextid`                                           |:material-check-all:|
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
| `/cluster/firewall/aliases`                                 |:material-check-all:|:material-check-all:|
| `/cluster/firewall/aliases/:name`                           |:material-check-all:||:material-check-all:|:material-check-all:|
| `/cluster/firewall/groups`                                  |:material-close:|:material-close:|
| `/cluster/firewall/groups/:group`                           |:material-close:|:material-close:||:material-close:|
| `/cluster/firewall/groups/:group/:pos`                      |:material-close:||:material-close:|:material-close:|
| `/cluster/firewall/ipset`                                   |:material-check-all:|:material-close:|
| `/cluster/firewall/ipset/:name`                             |:material-close:|:material-close:||:material-close:|
| `/cluster/firewall/ipset/:name/:cidr`                       |:material-close:||:material-close:|:material-close:|
| `/cluster/firewall/rules`                                   |:material-check-all:|:material-close:|
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
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes`                                                    |:material-check-all:|
| `/nodes/:node`                                              |:material-check-all:|
| `/nodes/:node/aplinfo`                                      |:material-close:|:material-close:|
| `/nodes/:node/config`                                       |:material-close:||:material-close:|
| `/nodes/:node/dns`                                          |:material-close:||:material-close:|
| `/nodes/:node/execute`                                      ||:material-close:|
| `/nodes/:node/hosts`                                        |:material-close:|:material-close:|
| `/nodes/:node/journal`                                      |:material-close:|
| `/nodes/:node/migrateall`                                   ||:material-close:|
| `/nodes/:node/netstat`                                      |:material-close:|
| `/nodes/:node/query-url-metadata`                           |:material-close:|
| `/nodes/:node/report`                                       |:material-close:|
| `/nodes/:node/rrd`                                          |:material-close:|
| `/nodes/:node/rrddata`                                      |:material-close:|
| `/nodes/:node/spiceshell`                                   ||:material-close:|
| `/nodes/:node/startall`                                     ||:material-close:|
| `/nodes/:node/status`                                       |:material-close:|:material-close:|
| `/nodes/:node/stopall`                                      ||:material-close:|
| `/nodes/:node/subscription`                                 |:material-close:|:material-close:|:material-close:|:material-close:|
| `/nodes/:node/suspendall`                                   ||:material-close:|
| `/nodes/:node/syslog`                                       |:material-close:|
| `/nodes/:node/termproxy`                                    ||:material-close:|
| `/nodes/:node/time`                                         |:material-close:||:material-close:|
| `/nodes/:node/version`                                      |:material-close:|
| `/nodes/:node/vncshell`                                     ||:material-close:|
| `/nodes/:node/vncwebsocket`                                 |:material-close:|
| `/nodes/:node/wakeonlan`                                    ||:material-close:|

### Node: apt
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/apt`                                          |:material-close:|
| `/nodes/:node/apt/changelog`                                |:material-close:|
| `/nodes/:node/apt/repositories`                             |:material-close:|:material-close:|:material-close:|
| `/nodes/:node/apt/update`                                   |:material-close:|:material-close:|
| `/nodes/:node/apt/versions`                                 |:material-close:|

### Node: Capabilities
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/capabilities`                                 |:material-close:|
| `/nodes/:node/capabilities/qemu`                            |:material-close:|
| `/nodes/:node/capabilities/qemu/cpu`                        |:material-close:|
| `/nodes/:node/capabilities/qemu/machines`                   |:material-close:|

### Node: ceph
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/ceph`                                         |:material-close:|
| `/nodes/:node/ceph/cmd-safety`                              |:material-close:|
| `/nodes/:node/ceph/crush`                                   |:material-close:|
| `/nodes/:node/ceph/init`                                    ||:material-close:|
| `/nodes/:node/ceph/log`                                     |:material-close:|
| `/nodes/:node/ceph/restart`                                 ||:material-close:|
| `/nodes/:node/ceph/rules`                                   |:material-close:|
| `/nodes/:node/ceph/start`                                   ||:material-close:|
| `/nodes/:node/ceph/status`                                  |:material-close:|
| `/nodes/:node/ceph/stop`                                    ||:material-close:|
| `/nodes/:node/ceph/cfg`                                     |:material-close:|
| `/nodes/:node/ceph/cfg/db`                                  |:material-close:|
| `/nodes/:node/ceph/cfg/raw`                                 |:material-close:|
| `/nodes/:node/ceph/cfg/value`                               |:material-close:|
| `/nodes/:node/ceph/fs`                                      |:material-close:|
| `/nodes/:node/ceph/fs/:name`                                ||:material-close:|
| `/nodes/:node/ceph/mds`                                     |:material-close:|
| `/nodes/:node/ceph/mds/:name`                               ||:material-close:||:material-close:|
| `/nodes/:node/ceph/mgr`                                     |:material-close:|
| `/nodes/:node/ceph/mgr/:id`                                 ||:material-close:||:material-close:|
| `/nodes/:node/ceph/mon`                                     |:material-close:|
| `/nodes/:node/ceph/mon/:monid`                              ||:material-close:||:material-close:|
| `/nodes/:node/ceph/osd`                                     |:material-close:|:material-close:|
| `/nodes/:node/ceph/osd/:osdid`                              |:material-close:|||:material-close:|
| `/nodes/:node/ceph/osd/:osdid/in`                           ||:material-close:|
| `/nodes/:node/ceph/osd/:osdid/lv-info`                      |:material-close:|
| `/nodes/:node/ceph/osd/:osdid/metadata`                     |:material-close:|
| `/nodes/:node/ceph/osd/:osdid/out`                          ||:material-close:|
| `/nodes/:node/ceph/osd/:osdid/scrub`                        ||:material-close:|
| `/nodes/:node/ceph/pool`                                    |:material-close:|:material-close:|
| `/nodes/:node/ceph/pool/:name`                              |:material-close:||:material-close:|:material-close:|
| `/nodes/:node/ceph/pool/:name/status`                       |:material-close:|


### Node: Certificates
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/certificates`                                 |:material-close:|
| `/nodes/:node/certificates/custom`                          ||:material-close:||:material-close:|
| `/nodes/:node/certificates/info`                            |:material-close:|
| `/nodes/:node/certificates/acme`                            |:material-close:|
| `/nodes/:node/certificates/acme/certificate`                ||:material-close:|:material-close:|:material-close:|

### Node: Disks
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/disks`                                        |:material-close:|
| `/nodes/:node/disks/directory`                              |:material-close:|:material-close:|
| `/nodes/:node/disks/directory/:name`                        ||||:material-close:|
| `/nodes/:node/disks/lvm`                                    |:material-close:|:material-close:|
| `/nodes/:node/disks/lvm/:name`                              ||||:material-close:|
| `/nodes/:node/disks/lvmthin`                                |:material-close:|:material-close:|
| `/nodes/:node/disks/lvmthin/:name`                          ||||:material-close:|
| `/nodes/:node/disks/zfs`                                    |:material-close:|:material-close:| 
| `/nodes/:node/disks/zfs/:name`                              |:material-close:|||:material-close:| 
| `/nodes/:node/disks/initgpt`                                ||:material-close:|
| `/nodes/:node/disks/list`                                   |:material-close:|
| `/nodes/:node/disks/smart`                                  |:material-close:|
| `/nodes/:node/disks/wipedisk`                               |||:material-close:|

### Node: Firewall
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/firewall`                                     |:material-close:|
| `/nodes/:node/firewall/rules`                               |:material-check-all:|:material-close:|
| `/nodes/:node/firewall/rules/:pos`                          |:material-check-all:||:material-close:|:material-close:|
| `/nodes/:node/firewall/log`                                 |:material-check:|
| `/nodes/:node/firewall/options`                             |:material-close:||:material-close:|

### Node: Hardware
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/hardware`                                     |:material-close:|
| `/nodes/:node/hardware/pci/:pci-id-or-mapping`              |:material-close:|
| `/nodes/:node/hardware/pci/:pci-id-or-mapping/mdev`         |:material-close:|
| `/nodes/:node/hardware/usb`                                 |:material-close:|

### Node: lxc
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

### Node: Network
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/network`                                      |:material-close:|:material-close:|:material-close:|:material-close:|
| `/nodes/:node/network/:iface`                               |:material-close:||:material-close:|:material-close:|

### Node: qemu
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/qemu`                                         |:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid`                                   |:material-close:|||:material-close:|
| `/nodes/:node/qemu/:vmid/agent`                             |:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid/agent/exec`                        ||:material-close:|
| `/nodes/:node/qemu/:vmid/agent/exec-status`                 |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/file-read`                   |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/file-write`                  ||:material-close:|
| `/nodes/:node/qemu/:vmid/agent/fsfreeze-freeze`             ||:material-close:|
| `/nodes/:node/qemu/:vmid/agent/fsfreeze-status`             ||:material-close:|
| `/nodes/:node/qemu/:vmid/agent/fsfreeze-thaw`               ||:material-close:|
| `/nodes/:node/qemu/:vmid/agent/fstrim`                      ||:material-close:|
| `/nodes/:node/qemu/:vmid/agent/get-fsinfo`                  |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/get-host-name`               |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/get-memory-block-info`       |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/get-memory-blocks`           |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/get-osinfo`                  |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/get-time`                    |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/get-timezone`                |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/get-users`                   |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/get-vcpus`                   |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/info`                        |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/network-get-interfaces`      |:material-close:|
| `/nodes/:node/qemu/:vmid/agent/set-user-password`           ||:material-close:|
| `/nodes/:node/qemu/:vmid/agent/shutdown`                    ||:material-close:|
| `/nodes/:node/qemu/:vmid/agent/suspend-disk`                ||:material-close:|
| `/nodes/:node/qemu/:vmid/agent/suspend-hybrid`              ||:material-close:|
| `/nodes/:node/qemu/:vmid/agent/suspend-ram`                 ||:material-close:|
| `/nodes/:node/qemu/:vmid/cloudinit`                         |:material-close:||:material-close:|
| `/nodes/:node/qemu/:vmid/cloudinit/dump`                    |:material-close:|
| `/nodes/:node/qemu/:vmid/firewall`                          |:material-close:|
| `/nodes/:node/qemu/:vmid/firewall/log`                      |:material-close:|
| `/nodes/:node/qemu/:vmid/firewall/option`                   |:material-close:||:material-close:|
| `/nodes/:node/qemu/:vmid/firewall/refs`                     |:material-close:|
| `/nodes/:node/qemu/:vmid/firewall/aliases`                  |:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid/firewall/aliases/:name`            |:material-close:||:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid/firewall/ipset`                    |:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid/firewall/ipset/:name`              |:material-close:|:material-close:||:material-close:|
| `/nodes/:node/qemu/:vmid/firewall/ipset/:name/:cidr`        |:material-close:||:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid/firewall/rules`                    |:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid/firewall/rules/:pos`               |:material-close:||:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid/snapshot`                          |:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid/snapshot/:snapname`                |:material-close:|||:material-close:|
| `/nodes/:node/qemu/:vmid/snapshot/:snapname/config`         |:material-close:||:material-close:|
| `/nodes/:node/qemu/:vmid/snapshot/:snapname/rollback`       ||:material-close:|
| `/nodes/:node/qemu/:vmid/status`                            |:material-close:|
| `/nodes/:node/qemu/:vmid/status/current`                    |:material-close:|
| `/nodes/:node/qemu/:vmid/status/reboot`                     ||:material-close:|
| `/nodes/:node/qemu/:vmid/status/reset`                      ||:material-close:|
| `/nodes/:node/qemu/:vmid/status/resume`                     ||:material-close:|
| `/nodes/:node/qemu/:vmid/status/shutdown`                   ||:material-close:|
| `/nodes/:node/qemu/:vmid/status/start`                      ||:material-close:|
| `/nodes/:node/qemu/:vmid/status/stop`                       ||:material-close:|
| `/nodes/:node/qemu/:vmid/status/suspend`                    ||:material-close:|
| `/nodes/:node/qemu/:vmid/clone`                             ||:material-close:|
| `/nodes/:node/qemu/:vmid/config`                            |:material-close:|:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid/feature`                           |:material-close:|
| `/nodes/:node/qemu/:vmid/migrate`                           |:material-close:|:material-close:|
| `/nodes/:node/qemu/:vmid/monitor`                           ||:material-close:|
| `/nodes/:node/qemu/:vmid/move_disk`                         ||:material-close:|
| `/nodes/:node/qemu/:vmid/mtunnel`                           ||:material-close:|
| `/nodes/:node/qemu/:vmid/mtunnelwebsocket`                  |:material-close:|
| `/nodes/:node/qemu/:vmid/pending`                           |:material-close:|
| `/nodes/:node/qemu/:vmid/remote_migrate`                    ||:material-close:|
| `/nodes/:node/qemu/:vmid/resize`                            |||:material-close:|
| `/nodes/:node/qemu/:vmid/rrd`                               |:material-close:|
| `/nodes/:node/qemu/:vmid/rrddata`                           |:material-close:|
| `/nodes/:node/qemu/:vmid/sendkey`                           |||:material-close:|
| `/nodes/:node/qemu/:vmid/spicyproxy`                        ||:material-close:|
| `/nodes/:node/qemu/:vmid/template`                          ||:material-close:|
| `/nodes/:node/qemu/:vmid/termproxy`                         ||:material-close:|
| `/nodes/:node/qemu/:vmid/unlink`                            |||:material-close:|
| `/nodes/:node/qemu/:vmid/vncproxy`                          ||:material-close:|
| `/nodes/:node/qemu/:vmid/vncwebsocket`                      |:material-close:|

### Node: Replication
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/replication`                                  |:material-close:|
| `/nodes/:node/replication/:id`                              |:material-close:|
| `/nodes/:node/replication/:id/log`                          |:material-close:|
| `/nodes/:node/replication/:id/schedule_now`                 ||:material-close:|
| `/nodes/:node/replication/:id/status`                       |:material-close:|


### Node: Scan
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/scan`                                         |:material-close:|
| `/nodes/:node/scan/cifs`                                    |:material-close:|
| `/nodes/:node/scan/gluterfs`                                |:material-close:|
| `/nodes/:node/scan/iscsi`                                   |:material-close:|
| `/nodes/:node/scan/lvm`                                     |:material-close:|
| `/nodes/:node/scan/lvmthin`                                 |:material-close:|
| `/nodes/:node/scan/nfs`                                     |:material-close:|
| `/nodes/:node/scan/pbs`                                     |:material-close:|
| `/nodes/:node/scan/zfs`                                     |:material-close:|

### Node: sdn
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/sdn`                                          |:material-close:|
| `/nodes/:node/sdn/zones`                                    |:material-close:|
| `/nodes/:node/sdn/zones/:zone`                              |:material-close:|
| `/nodes/:node/sdn/zones/:zone/content`                      |:material-close:|

### Node: Services
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/services`                                     |:material-close:|
| `/nodes/:node/services/:service`                            |:material-close:|
| `/nodes/:node/services/:service/reload`                     ||:material-close:|
| `/nodes/:node/services/:service/restart`                    ||:material-close:|
| `/nodes/:node/services/:service/start`                      ||:material-close:|
| `/nodes/:node/services/:service/state`                      |:material-close:|
| `/nodes/:node/services/:service/stop`                       ||:material-close:|


### Node: Storage
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/storage`                                      |:material-check-all:|
| `/nodes/:node/storage/:storage`                             |:material-close:|
| `/nodes/:node/storage/:storage/content`                     |:material-close:|:material-check-all:|
| `/nodes/:node/storage/:storage/content/:volume`             |:material-close:|:material-close:|:material-close:|:material-close:|
| `/nodes/:node/storage/:storage/file-restore`                |:material-close:|:material-close:|:material-close:|:material-close:|
| `/nodes/:node/storage/:storage/file-restore/download`       |:material-close:|
| `/nodes/:node/storage/:storage/file-restore/list`           |:material-close:|
| `/nodes/:node/storage/:storage/download-url`                ||:material-check:|
| `/nodes/:node/storage/:storage/import-metadata`             |:material-close:|
| `/nodes/:node/storage/:storage/prunebackups`                |:material-close:|||:material-close:|
| `/nodes/:node/storage/:storage/rdd`                         |:material-close:|
| `/nodes/:node/storage/:storage/rdddata`                     |:material-close:|
| `/nodes/:node/storage/:storage/status`                      |:material-close:|
| `/nodes/:node/storage/:storage/upload`                      ||:material-close:|

### Node: Tasks
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/tasks`                                        |:material-close:|
| `/nodes/:node/tasks/:upid`                                  |:material-close:|||:material-close:|
| `/nodes/:node/tasks/:upid/log`                              |:material-close:|
| `/nodes/:node/tasks/:upid/status`                           |:material-close:|

### Node: VZDump
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/nodes/:node/vzdump`                                       ||:material-close:|
| `/nodes/:node/vzdump/defaults`                              |:material-close:|
| `/nodes/:node/vzdump/extractconfig`                         |:material-close:|

## Pools
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/pools`                                                    |:material-close:|:material-close:|:material-close:|:material-close:|
| `/pools/:poolid`                                            |:material-close:||:material-close:|:material-close:|

## Storage
| Path                                                        | GET                | POST           | PUT | DELETE |
|-------------------------------------------------------------|:------------------:|:--------------:|:-----:|:-:|
| `/storage`                                                  |:material-close:|:material-close:|
| `/storage/:storage`                                         |:material-close:||:material-close:|:material-close:|

