package permission

type Group struct {
	Name     string
	Sections map[string]*SectionPermissions
}

func (g *Group) HasPermission(section string, permission Permission, ac *AccessControl) bool {
	if secPerms, exists := g.Sections[section]; exists {
		return secPerms.Permissions[permission]
	}
	return false
}
