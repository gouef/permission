package permission

// Permission represents different levels of access control.
type Permission string

const (
	// Create allows an entity to create a resource.
	Create Permission = "CREATE"
	// Read allows an entity to read a resource.
	Read Permission = "READ"
	// Update allows an entity to modify a resource.
	Update Permission = "UPDATE"
	// Delete allows an entity to remove a resource.
	Delete Permission = "DELETE"
	// All grants full access to a resource.
	All Permission = "ALL"
)
