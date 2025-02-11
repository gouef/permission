package permission

type Resource struct {
	ID           string
	Parent       *Resource
	SubResources map[string]*Resource // Podresource podle názvu
	Owners       []*Entity            // Vlastníci resource
}

func NewResource(id string) *Resource {
	return &Resource{
		ID:           id,
		SubResources: make(map[string]*Resource),
		Owners:       make([]*Entity, 0),
	}
}

func (r *Resource) CreateSub(id string) *Resource {
	sub := NewResource(id)

	r.AddSubs(sub)

	return sub
}

func (r *Resource) GetSub(id string) *Resource {
	return r.SubResources[id]
}

func (r *Resource) CreateSubs(ids ...string) *Resource {
	for _, id := range ids {
		r.CreateSub(id)
	}
	return r
}

func (r *Resource) AddSubs(resources ...*Resource) *Resource {
	for _, resource := range resources {
		resource.Parent = r
		r.SubResources[resource.ID] = resource
	}

	return r
}

func (r *Resource) AddOwners(owners ...*Entity) *Resource {
	r.Owners = append(r.Owners, owners...)
	return r
}
