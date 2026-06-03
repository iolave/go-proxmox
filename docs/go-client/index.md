# Go client
## Install
```bash 
go get github.com/iolave/go-proxmox@v0.8.0
```

## Environment variables
The following environment variables can be used to interact both with the cli and the golang pkg.

|VARIABLE           |DESCRIPTION                            |TOKEN|USR & PWD|
|-------------------|---------------------------------------|:---:|:--------:|
|PROXMOX_USERNAME   |Proxmox VE linux user (i.e. "root@pam")|X    |X         |
|PROXMOX_PASSWORD   |Proxmox VE linux password              |-    |X         |
|PROXMOX_TOKEN_NAME |Proxmox VE generated token name        |X    |-         |
|PROXMOX_TOKEN      |Proxmox VE generated token             |X    |-         |

## Getting started 
First, import the [api package]:
```go
import "github.com/iolave/go-proxmox/pkg/api"
```

In order to create a new pve api you can either use environment variables or the built-in [credentials] constructors.

=== "Using environment variables"
    ```go
    creds, err := api.NewTokenCreds("root@pam", "TOKEN_NAME", "UUID_TOKEN")
    if err != nil {
        // handle error
        panic(err)
    }

    config := api.Config{
	    Host:               "pve.example.com",
	    Port:               8006,
	    InsecureSkipVerify: true,
        Credentials:        creds,
    }

    c, err := api.New(config)
    ```

=== "Using token credentials"
    ```go
    creds := api.NewTokenCreds("root@pam", "TOKEN_NAME", "UUID_TOKEN")

    config := api.Config{
	    Host:               "pve.example.com",
	    Port:               8006,
	    InsecureSkipVerify: true,
        Credentials:        creds,
    }

    c, err := api.New(config)
    ```

### Connecting to a proxmox instance using custom headers 
If for some reason you need to add custom headers to your requests in order to reach the proxmox instance, you can do so by passing a map of custom headers to the [api.Config] struct.

```go
creds := api.NewTokenCreds("root@pam", "TOKEN_NAME", "UUID_TOKEN")

config := api.Config{
    Host:               "pve.example.com",
    Port:               8006,
    InsecureSkipVerify: true,
    Credentials:        creds,
    CustomHeaders: http.Header{
        "x-api-key": []string{"my-api-key"},
    },
}

c, err := api.New(config)
```

## Using the go-proxmox services
The go-proxmox services are available under the `pkg/pve` package. Next, there's an example of how to use the core service.

```go
package main

import (
    "fmt"
    "github.com/iolave/go-proxmox/pkg/api"
    "github.com/iolave/go-proxmox/pkg/pve/core"
)

func main() {
    creds := api.NewTokenCreds("root@pam", "TOKEN_NAME", "UUID_TOKEN")
    
    config := api.Config{
        Host:               "pve.example.com",
        Port:               8006,
        InsecureSkipVerify: true,
        Credentials:        creds,
        CustomHeaders: http.Header{
            "x-api-key": []string{"my-api-key"},
        },
    }
    
    c, err := api.New(config)

    cs := core.New(c)
    
    // Get the version of the proxmox instance
    res, err := cs.GetVersion()
    if err != nil {
        // handle error
        panic(err)
    }
    
    fmt.Println(res)
}
```

[api package]: https://go-proxmox.iolave.com/go-client/pkg/api/
[credentials]: https://go-proxmox.iolave.com/go-client/pkg/api/#type-credentials
[application]: https://developers.cloudflare.com/cloudflare-one/applications/
