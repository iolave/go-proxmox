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

