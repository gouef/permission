package permission

import "fmt"

// AccessControl manage permissions, users, groups and section
type AccessControl struct {
	users    map[string]*User
	groups   map[string]*Group
	sections map[string]*Section // Sekce v systému
}

func (ac *AccessControl) CreateSection(name string, perms map[Permission]bool) *Section {
	section := &Section{
		Name: name,
		Permissions: &SectionPermissions{
			Permissions: perms,
		},
	}
	ac.AddSection(section)
	return section
}

func (ac *AccessControl) CreateGroup(name string, sections map[string]map[Permission]bool) *Group {
	group := &Group{
		Name:     name,
		Sections: make(map[string]*SectionPermissions),
	}

	for secName, perms := range sections {
		group.Sections[secName] = &SectionPermissions{Permissions: perms}
	}

	ac.AddGroup(group)
	return group
}

func (ac *AccessControl) CreateUser(name string, groups []*Group, personalPerm map[string]map[Permission]bool, ownerOf map[string]bool) *User {
	user := &User{
		Name:         name,
		Groups:       groups,
		PersonalPerm: make(map[string]*SectionPermissions),
		OwnerOf:      ownerOf,
	}

	for secName, perms := range personalPerm {
		user.PersonalPerm[secName] = &SectionPermissions{Permissions: perms}
	}

	ac.AddUser(user)
	return user
}

func (ac *AccessControl) AddUser(user *User) {
	if ac.users == nil {
		ac.users = make(map[string]*User)
	}
	ac.users[user.Name] = user
}

func (ac *AccessControl) AddGroup(group *Group) {
	if ac.groups == nil {
		ac.groups = make(map[string]*Group)
	}
	ac.groups[group.Name] = group
}

func (ac *AccessControl) AddSection(section *Section) {
	if ac.sections == nil {
		ac.sections = make(map[string]*Section)
	}
	ac.sections[section.Name] = section
}

func (ac *AccessControl) AddUserToGroup(userName, groupName string) {
	user, userExists := ac.users[userName]
	group, groupExists := ac.groups[groupName]

	if userExists && groupExists {
		user.Groups = append(user.Groups, group)
	}
}

func (ac *AccessControl) SetUserPermissions(userName, sectionName string, perms map[Permission]bool) {
	user, exists := ac.users[userName]
	if !exists {
		fmt.Println("User not found")
		return
	}

	if user.PersonalPerm == nil {
		user.PersonalPerm = make(map[string]*SectionPermissions)
	}

	user.PersonalPerm[sectionName] = &SectionPermissions{
		Permissions: perms,
	}

	fmt.Printf("Permissions for user %s in section %s updated.\n", userName, sectionName)
}

func (ac *AccessControl) Can(userName, section string, permission Permission) bool {
	user, exists := ac.users[userName]
	if !exists {
		fmt.Println("User not found")
		return false
	}

	if user.HasPermission(section, permission, ac) {
		return true
	}

	for _, group := range user.Groups {
		if group.HasPermission(section, permission, ac) {
			return true
		}
	}

	if sectionPerms, exists := ac.sections[section]; exists {
		if sectionPerms.Permissions.Permissions[permission] {
			return true
		}
	}

	return false
}

func (ac *AccessControl) CanCreate(userName, section string) bool {
	return ac.Can(userName, section, Create)
}

func (ac *AccessControl) CanRead(userName, section string) bool {
	return ac.Can(userName, section, Read)
}

func (ac *AccessControl) CanUpdate(userName, section string) bool {
	return ac.Can(userName, section, Update)
}

func (ac *AccessControl) CanDelete(userName, section string) bool {
	return ac.Can(userName, section, Delete)
}
