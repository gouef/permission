# `AccessControl`

The main structure managing entities and resources.

```go
ac := permission.NewAccessControl()
```

- `CreateEntity(id string) *Entity` - Creates a new entity.
- `CreateResource(id string) *Resource` - Creates a new resource.
- `Allow(entity, resource, permission)` - Grants permission to an entity for a resource.
- `Deny(entity, resource, permission)` - Revokes permission.
- `Can(entity, resource, permission) bool` - Checks permission.
- `AddEntities(entities ...*Entity)` - Adds multiple entities.
- `AddResources(resources ...*Resource)` - Adds multiple resources.

## Example Usage

```go
ac := permission.NewAccessControl()
user := ac.CreateEntity("user")
doc := ac.CreateResource("document")
ac.Allow(user, doc, permission.Read)
if ac.Can(user, doc, permission.Read) {
    fmt.Println("User can read document.")
}
```

