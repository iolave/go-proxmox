# cloudflare

```go
import "github.com/iolave/go-proxmox/pkg/cloudflare"
```

The cloudflare package contains utilities that can be used to connect to services secured by cloudflare through [pkg/http](<https://pkg.go.dev/pkg/http/>) package.





<a name="ServiceToken"></a>
## type [ServiceToken](<https://github.com/iolave/go-proxmox/blob/master/pkg/cloudflare/service_token.go#L8-L11>)



```go
type ServiceToken struct {
    ClientId     string // Client id created when creating a new cloudflare service.
    ClientSecret string // Client secret created when creating a new cloudflare service.
}
```

<a name="NewServiceToken"></a>
### func [NewServiceToken](<https://github.com/iolave/go-proxmox/blob/master/pkg/cloudflare/service_token.go#L15>)

```go
func NewServiceToken(clientId, secret string) *ServiceToken
```

NewServiceToken generates a pointer to a ServiceToken struct with it's properties initialized.

<a name="ServiceToken.Set"></a>
### func \(\*ServiceToken\) [Set](<https://github.com/iolave/go-proxmox/blob/master/pkg/cloudflare/service_token.go#L26>)

```go
func (t *ServiceToken) Set(req *http.Request) error
```

Set adds the corresponding Cloudflare access client headers to the given request.

It returns an error only when nil is passed to the request parameter.