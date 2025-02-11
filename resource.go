package permission

// Resource represents an entity that can be assigned permissions.
type Resource struct {
	ID           string
	Parent       *Resource
	SubResources map[string]*Resource // Podresource podle názvu
	Owners       []*Entity            // Vlastníci resource
}

// NewResource initializes a new resource with the given ID.
//
// Example:
//
//	doc := permission.NewResource("document")
//	fmt.Println(doc.ID) // Output: document
func NewResource(id string) *Resource {
	return &Resource{
		ID:           id,
		SubResources: make(map[string]*Resource),
		Owners:       make([]*Entity, 0),
	}
}

// CreateSub generates a new sub-resource under the current resource.
//
// Example:
//
//	website := permission.NewResource("website")
//	news := website.CreateSub("news")
//	fmt.Println(news.ID) // Output: news
func (r *Resource) CreateSub(id string) *Resource {
	sub := NewResource(id)

	r.AddSubs(sub)

	return sub
}

// GetSub retrieves a sub-resource by its ID.
//
// Example:
//
//	website := permission.NewResource("website")
//	news := website.CreateSub("news")
//	fmt.Println(website.GetSub("news").ID) // Output: news
func (r *Resource) GetSub(id string) *Resource {
	return r.SubResources[id]
}

// CreateSubs generates multiple sub-resources.
//
// Example:
//
//	website := permission.NewResource("website")
//	website.CreateSub("news", "comments", "reviews")
func (r *Resource) CreateSubs(ids ...string) *Resource {
	for _, id := range ids {
		r.CreateSub(id)
	}
	return r
}

// AddSubs links additional sub-resources to the current resource.
func (r *Resource) AddSubs(resources ...*Resource) *Resource {
	for _, resource := range resources {
		resource.Parent = r
		r.SubResources[resource.ID] = resource
	}

	return r
}

// AddOwners assigns ownership of the resource to specific entities.
//
// Example:
//
//	user := permission.NewEntity("user")
//	doc := permission.NewResource("document")
//	doc.AddOwners(user)
func (r *Resource) AddOwners(owners ...*Entity) *Resource {
	r.Owners = append(r.Owners, owners...)
	return r
}
