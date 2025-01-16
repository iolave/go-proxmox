# helpers

```go
import "github.com/iolave/go-proxmox/pkg/helpers"
```





<a name="BoolToInt"></a>
## func [BoolToInt](<https://github.com/iolave/go-proxmox/blob/master/pkg/helpers/bool.go#L6>)

```go
func BoolToInt(b bool) int
```

BoolToInt converts a bool to an int.

It returns "1" if b == true and "0" otherwise.

<a name="NewBool"></a>
## func [NewBool](<https://github.com/iolave/go-proxmox/blob/master/pkg/helpers/pointers.go#L17>)

```go
func NewBool(v bool) *bool
```

NewBool returns bool pointer of the passed value.

<a name="NewInt"></a>
## func [NewInt](<https://github.com/iolave/go-proxmox/blob/master/pkg/helpers/pointers.go#L5>)

```go
func NewInt(v int) *int
```

NewInt returns an int pointer of the passed value.

<a name="NewStr"></a>
## func [NewStr](<https://github.com/iolave/go-proxmox/blob/master/pkg/helpers/pointers.go#L11>)

```go
func NewStr(v string) *string
```

NewStr returns a string pointer of the passed value.