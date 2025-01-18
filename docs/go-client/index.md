# Go client
## Install
```bash 
go get github.com/iolave/go-proxmox@v0.5.0
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
First, import the [pve package]:
```go
import "github.com/iolave/go-proxmox/pkg/pve"
```

In order to create a new pve api you can either use environment variables or the built-in [credentials] constructors.

=== "Using environment variables"
    ```go
    config := pve.Config{
	    Host:               "pve.example.com",
	    Port:               8006,
	    InsecureSkipVerify: true,
    }

    api, err := pve.New()
    ```

=== "Using the built-in credentials constructors"
    ```go
    creds, err := pve.NewEnvCreds()
    // Or:
    // creds := pve.NewTokenCreds("root@pam", "TOKEN_NAME", "UUID_TOKEN")

    if err != nil {
	    // handle the error as you please
    }
	
    config := pve.Config{
	    Host:               "pve.example.com",
	    Port:               8006,
	    InsecureSkipVerify: true,
    }

    api, err := pve.NewWithCredentials(config, creds)
    ```

### Connecting to a proxmox instance secured by Cloudflare 
If you are exposing your pve instance through proxmox zero trust, you need to setup an [application] within cloudflare and generate a [service token] for that application.

Once you have your service token id and secret, add them your `pve.Config`:
```go
import "github.com/iolave/go-proxmox/pkg/cloudflare"

//...

config := pve.Config{
    Host:               "pve.example.com",
	Port:               8006,
	InsecureSkipVerify: true,
	CfServiceToken:     cloudflare.NewServiceToken("token-id.access", "token-secret"),
}
```


[pve package]: https://go-proxmox.iolave.com/go-client/pkg/pve/
[credentials]: https://go-proxmox.iolave.com/go-client/pkg/pve/#type-credentials
[service token]: https://developers.cloudflare.com/cloudflare-one/identity/service-tokens
[application]: https://developers.cloudflare.com/cloudflare-one/applications/
