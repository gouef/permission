# `Permission`

Defines types of permissions:

```go
type Permission string
const (
    Create Permission = "CREATE"
    Read   Permission = "READ"
    Update Permission = "UPDATE"
    Delete Permission = "DELETE"
    All    Permission = "ALL"
)
```

You can add custom Permission types

## Example Usage

```go
var VotePerm permission.Permission = "vote"

ac := permission.NewAccessControl()
res := ac.CreateResource("news")
user := ac.CreateEntity("user")

user.Allow(res, VotePerm)

if ac.Can(user, doc, VotePerm) {
    fmt.Println("User can vote .")
}
```
