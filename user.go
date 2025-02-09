package permission

type User struct {
	Name         string
	Groups       []*Group
	PersonalPerm map[string]*SectionPermissions
	OwnerOf      map[string]bool
}

func (u *User) CanCreate(section string, ac *AccessControl) bool {
	return u.HasPermission(section, Create, ac)
}

func (u *User) CanRead(section string, ac *AccessControl) bool {
	return u.HasPermission(section, Read, ac)
}

func (u *User) CanUpdate(section string, ac *AccessControl) bool {
	return u.HasPermission(section, Update, ac)
}

func (u *User) CanDelete(section string, ac *AccessControl) bool {
	return u.HasPermission(section, Delete, ac)
}

func (u *User) HasPermission(section string, permission Permission, ac *AccessControl) bool {
	if u.OwnerOf[section] {
		return true
	}

	if secPerms, exists := u.PersonalPerm[section]; exists {
		return secPerms.Permissions[permission]
	}

	for _, group := range u.Groups {
		if group.HasPermission(section, permission, ac) {
			return true
		}
	}

	if sectionPerms, exists := ac.sections[section]; exists {
		return sectionPerms.Permissions.Permissions[permission]
	}

	return false
}
