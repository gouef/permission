# `Resource`

Represents an object with access control.

```go
res := permission.NewResource("document")
```

- `CreateSub(id string) *Resource` - Creates a sub-resource.
- `GetSub(id string) *Resource` - Retrieves a sub-resource.
- `CreateSubs(ids ...string) *Resource` - Creates multiple sub-resources.
- `AddSubs(resources ...*Resource) *Resource` - Adds multiple sub-resources.
- `AddOwners(owners ...*Entity)` - Sets owners of the resource.
