# `Entity`

Represents a user, group, or role with permissions.

```go
user := permission.NewEntity("user1")
```

- `AddParents(parents ...*Entity)` - Adds parent entities.
- `AddChildren(children ...*Entity)` - Adds child entities.
- `Allow(resource, permissions...)` - Grants multiple permissions.
- `Deny(resource, permissions...)` - Denies permissions.
- `CreateChild(id string) *Entity` - Creates a child entity.
- `AddPerm(permission Permission, resource *Resource, enabled bool)` - Grants or revokes specific permissions.
- `AddPermAll(resource *Resource, enabled bool)` - Grants or revokes all permissions.
- `AddPermCreate(resource *Resource, enabled bool)` - Grants or revokes create permissions.
- `AddPermRead(resource *Resource, enabled bool)` - Grants or revokes read permissions.
- `AddPermUpdate(resource *Resource, enabled bool)` - Grants or revokes update permissions.
- `AddPermDelete(resource *Resource, enabled bool)` - Grants or revokes delete permissions.
