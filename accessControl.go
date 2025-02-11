package permission

type AccessControl struct {
	Entities  []*Entity
	Resources []*Resource
}

func NewAccessControl() *AccessControl {
	return &AccessControl{
		Entities:  []*Entity{},
		Resources: []*Resource{},
	}
}

func (ac *AccessControl) CreateResource(id string) *Resource {
	resource := NewResource(id)
	ac.AddResource(resource)

	return resource
}

func (ac *AccessControl) AddResource(resource *Resource) *AccessControl {
	ac.Resources = append(ac.Resources, resource)
	return ac
}

func (ac *AccessControl) CreateEntity(id string) *Entity {
	entity := NewEntity(id)
	ac.AddEntity(entity)
	return entity
}

func (ac *AccessControl) AddEntity(entity *Entity) *AccessControl {
	ac.Entities = append(ac.Entities, entity)
	return ac
}

func (ac *AccessControl) Allow(entity *Entity, resource *Resource, permission Permission) *AccessControl {
	entity.AddPerm(permission, resource, true)
	return ac
}

func (ac *AccessControl) Deny(entity *Entity, resource *Resource, permission Permission) *AccessControl {
	entity.AddPerm(permission, resource, false)
	return ac
}

func (ac *AccessControl) AddEntities(entities ...*Entity) {
	ac.Entities = append(ac.Entities, entities...)
}

func (ac *AccessControl) AddResources(resources ...*Resource) {
	ac.Resources = append(ac.Resources, resources...)
}

// Ověření oprávnění entity na konkrétní resource
func (ac *AccessControl) HasPermission(entity *Entity, resource *Resource, permission Permission) bool {

	for _, owner := range resource.Owners {
		if owner == entity {
			return true
		}
	}

	// 1. Pokud má entita explicitně dané oprávnění, použijeme ho
	if perms, exists := entity.Permission[permission]; exists {
		if val, ok := perms[resource]; ok {
			return val
		}
	}

	// 2. Pokud má `ALL` oprávnění, vracíme true
	if perms, exists := entity.Permission[All]; exists {
		if val, ok := perms[resource]; ok && val {
			return true
		}
	}

	// 3. Procházíme všechny rodičovské entity (skupiny, nadřazené skupiny)
	for _, parent := range entity.Parents {
		if ac.HasPermission(parent, resource, permission) {
			return true
		}
	}

	// Kontrola parents od resource
	if resource.Parent != nil {
		if ac.HasPermission(entity, resource.Parent, permission) {
			return true
		}
	}

	return false
}

func (ac *AccessControl) CanCreate(entity *Entity, resource *Resource) bool {
	return ac.HasPermission(entity, resource, Create)
}

func (ac *AccessControl) CanRead(entity *Entity, resource *Resource) bool {
	return ac.HasPermission(entity, resource, Read)
}

func (ac *AccessControl) CanUpdate(entity *Entity, resource *Resource) bool {
	return ac.HasPermission(entity, resource, Update)
}

func (ac *AccessControl) CanDelete(entity *Entity, resource *Resource) bool {
	return ac.HasPermission(entity, resource, Delete)
}
