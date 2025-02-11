package permission

// AccessControl manages entities and resources, allowing permission assignment.
type AccessControl struct {
	Entities  []*Entity
	Resources []*Resource
}

// NewAccessControl initializes a new AccessControl instance.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	fmt.Println(len(ac.Entities)) // Output: 0
func NewAccessControl() *AccessControl {
	return &AccessControl{
		Entities:  []*Entity{},
		Resources: []*Resource{},
	}
}

// CreateResource creates a new resource and adds it to the system.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	res := ac.CreateResource("document")
//	fmt.Println(res.ID) // Output: document
func (ac *AccessControl) CreateResource(id string) *Resource {
	resource := NewResource(id)
	ac.AddResource(resource)

	return resource
}

// AddResource manually adds a resource to the access control system.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	doc := permission.NewResource("document")
//	ac.AddResource(doc)
func (ac *AccessControl) AddResource(resource *Resource) *AccessControl {
	ac.Resources = append(ac.Resources, resource)
	return ac
}

// CreateEntity creates a new entity and adds it to the system.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user := ac.CreateEntity("user1")
//	fmt.Println(user.ID) // Output: user1
func (ac *AccessControl) CreateEntity(id string) *Entity {
	entity := NewEntity(id)
	ac.AddEntity(entity)
	return entity
}

// AddEntity manually adds an entity to the access control system.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user := permission.NewEntity("user1")
//	ac.AddEntity(user)
func (ac *AccessControl) AddEntity(entity *Entity) *AccessControl {
	ac.Entities = append(ac.Entities, entity)
	return ac
}

// Allow grants a specific permission to an entity for a given resource.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user := ac.CreateEntity("user1")
//	doc := ac.CreateResource("document")
//	ac.Allow(user, doc, pe
func (ac *AccessControl) Allow(entity *Entity, resource *Resource, permission Permission) *AccessControl {
	entity.AddPerm(permission, resource, true)
	return ac
}

// Deny revokes a specific permission from an entity for a given resource.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user := ac.CreateEntity("user1")
//	doc := ac.CreateResource("document")
//	ac.Deny(user, doc, permission.Read)
func (ac *AccessControl) Deny(entity *Entity, resource *Resource, permission Permission) *AccessControl {
	entity.AddPerm(permission, resource, false)
	return ac
}

// AddEntities adds multiple entities at once
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user1 := permission.NewEntity("user1")
//	user2 := permission.NewEntity("user1")
//	ac.AddEntities(user1, user2)
func (ac *AccessControl) AddEntities(entities ...*Entity) {
	ac.Entities = append(ac.Entities, entities...)
}

// AddResources adds multiple resources at once.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	res1 := permission.NewResource("res1")
//	res2 := permission.NewResource("res2")
//	ac.AddResources(res1, res2)
func (ac *AccessControl) AddResources(resources ...*Resource) {
	ac.Resources = append(ac.Resources, resources...)
}

// HasPermission verifies if an entity has permission for a resource.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user := ac.CreateEntity("user1")
//	doc := ac.CreateResource("document")
//	ac.Allow(user, doc, permission.Read)
//	fmt.Println(ac.HasPermission(user, doc, permission.Read)) // Output: true
func (ac *AccessControl) HasPermission(entity *Entity, resource *Resource, permission Permission) bool {
	for _, owner := range resource.Owners {
		if owner == entity {
			return true
		}
	}

	if perms, exists := entity.Permission[permission]; exists {
		if val, ok := perms[resource]; ok {
			return val
		}
	}

	if perms, exists := entity.Permission[All]; exists {
		if val, ok := perms[resource]; ok && val {
			return true
		}
	}

	for _, parent := range entity.Parents {
		if ac.HasPermission(parent, resource, permission) {
			return true
		}
	}

	if resource.Parent != nil {
		if ac.HasPermission(entity, resource.Parent, permission) {
			return true
		}
	}

	return false
}

// Can checks if an entity has a specific permission for a resource.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user := ac.CreateEntity("user1")
//	doc := ac.CreateResource("document")
//	ac.Allow(user, doc, permission.Read)
//	fmt.Println(ac.Can(user, doc, permission.Read)) // Output: true
func (ac *AccessControl) Can(entity *Entity, resource *Resource, permission Permission) bool {
	return ac.HasPermission(entity, resource, permission)
}

// CanCreate checks if an entity has permission to create a resource.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user := ac.CreateEntity("user1")
//	doc := ac.CreateResource("document")
//	ac.Allow(user, doc, permission.Create)
//	fmt.Println(ac.CanCreate(user, doc)) // Output: true
func (ac *AccessControl) CanCreate(entity *Entity, resource *Resource) bool {
	return ac.Can(entity, resource, Create)
}

// CanRead checks if an entity has permission to read a resource.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user := ac.CreateEntity("user1")
//	doc := ac.CreateResource("document")
//	ac.Allow(user, doc, permission.READ)
//	fmt.Println(ac.CanRead(user, doc)) // Output: true
func (ac *AccessControl) CanRead(entity *Entity, resource *Resource) bool {
	return ac.Can(entity, resource, Read)
}

// CanUpdate checks if an entity has permission to update a resource.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user := ac.CreateEntity("user1")
//	doc := ac.CreateResource("document")
//	ac.Allow(user, doc, permission.Update)
//	fmt.Println(ac.CanUpdate(user, doc)) // Output: true
func (ac *AccessControl) CanUpdate(entity *Entity, resource *Resource) bool {
	return ac.Can(entity, resource, Update)
}

// CanDelete checks if an entity has permission to delete a resource.
//
// Example:
//
//	ac := permission.NewAccessControl()
//	user := ac.CreateEntity("user1")
//	doc := ac.CreateResource("document")
//	ac.Allow(user, doc, permission.DELETE)
//	fmt.Println(ac.CanDelete(user, doc)) // Output: true
func (ac *AccessControl) CanDelete(entity *Entity, resource *Resource) bool {
	return ac.Can(entity, resource, Delete)
}
