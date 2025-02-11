package permission

import "github.com/gouef/utils"

// Entity represents a user, group, role (or what you want) with specific permissions.
type Entity struct {
	ID         string
	Parents    []*Entity
	Children   []*Entity
	Permission map[Permission]map[*Resource]bool
}

// NewEntity creates a new entity with default permission sets.
func NewEntity(id string) *Entity {
	perms := make(map[Permission]map[*Resource]bool)
	perms[All] = make(map[*Resource]bool)
	perms[Create] = make(map[*Resource]bool)
	perms[Read] = make(map[*Resource]bool)
	perms[Update] = make(map[*Resource]bool)
	perms[Delete] = make(map[*Resource]bool)

	return &Entity{
		ID:         id,
		Parents:    make([]*Entity, 0),
		Children:   make([]*Entity, 0),
		Permission: perms,
	}
}

// CreateChild creates a child entity and assigns it as a descendant.
func (e *Entity) CreateChild(id string) *Entity {
	child := NewEntity(id)
	e.AddChildren(child)

	return child
}

// AddChildren associates child entities with the current entity.
func (e *Entity) AddChildren(children ...*Entity) {
	e.Children = append(e.Children, children...)
	for _, child := range children {
		if !child.parentExists(e) {
			child.AddParents(e)
		}
	}
}

// AddParents associates parent entities with the current entity.
func (e *Entity) AddParents(parents ...*Entity) {
	e.Parents = append(e.Parents, parents...)
	for _, parent := range parents {
		if !parent.childExists(e) {
			parent.AddChildren(e)
		}
	}
}

// Allow grants specified permissions for a resource to the entity.
func (e *Entity) Allow(resource *Resource, permissions ...Permission) {
	for _, permission := range permissions {
		e.AddPerm(permission, resource, true)
	}
}

// Deny revokes specified permissions for a resource from the entity.
func (e *Entity) Deny(resource *Resource, permissions ...Permission) {
	for _, permission := range permissions {
		e.AddPerm(permission, resource, false)
	}
}

// AddPerm sets or removes a specific permission for a resource.
func (e *Entity) AddPerm(permission Permission, resource *Resource, enabled bool) {
	if _, ok := e.Permission[permission]; !ok {
		e.Permission[permission] = make(map[*Resource]bool)
	}
	e.Permission[permission][resource] = enabled
}

func (e *Entity) AddPermAll(resource *Resource, enabled bool) {
	e.AddPerm(All, resource, enabled)
}

func (e *Entity) AddPermCreate(resource *Resource, enabled bool) {
	e.AddPerm(Create, resource, enabled)
}

func (e *Entity) AddPermRead(resource *Resource, enabled bool) {
	e.AddPerm(Read, resource, enabled)
}

func (e *Entity) AddPermUpdate(resource *Resource, enabled bool) {
	e.AddPerm(Update, resource, enabled)
}

func (e *Entity) AddPermDelete(resource *Resource, enabled bool) {
	e.AddPerm(Delete, resource, enabled)
}

func (e *Entity) parentExists(parent *Entity) bool {
	return utils.InArray(parent, e.Parents)
}

func (e *Entity) childExists(child *Entity) bool {
	return utils.InArray(child, e.Children)
}
