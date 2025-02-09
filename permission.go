package permission

type Permission string

const (
	Create Permission = "CREATE"
	Read   Permission = "READ"
	Update Permission = "UPDATE"
	Delete Permission = "DELETE"
	All    Permission = "ALL"
)

type SectionPermissions struct {
	Permissions map[Permission]bool
}

type Section struct {
	Name        string
	Permissions *SectionPermissions
}
