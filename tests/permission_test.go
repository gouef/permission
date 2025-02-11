package tests

import (
	"github.com/gouef/permission"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermissions(t *testing.T) {

	t.Run("Test 1", func(t *testing.T) {

		// Inicializace AccessControl
		ac := permission.NewAccessControl()

		// Vytvoření resource
		webResource := permission.NewResource("Web").AddSubs(
			permission.NewResource("comments").AddSubs(
				permission.NewResource("comment1"),
				permission.NewResource("comment2"),
			),
		)

		// Vytvoření entity (uživatelé a skupiny)
		user1 := permission.NewEntity("user1")
		user2 := permission.NewEntity("user2")
		group1 := permission.NewEntity("group1")

		// Přiřazení oprávnění pro skupinu
		group1.AddPermRead(webResource, true)
		group1.AddPermCreate(webResource, true)

		// Přiřazení oprávnění pro uživatele
		user1.AddPermRead(webResource, true)
		user1.AddPermCreate(webResource, false)
		user2.AddPermRead(webResource, false)

		// Přiřazení uživatelů do skupiny
		group1.AddChildren(user1, user2)
		user1.AddParents(group1)
		user2.AddParents(group1)

		// Přiřazení resource do AccessControl
		ac.AddResources(webResource)
		ac.AddEntities(group1, user1, user2)

		// Testování oprávnění

		assert.False(t, ac.CanCreate(user1, webResource), "Can user1 create Web?")
		assert.True(t, ac.CanRead(user1, webResource), "Can user1 read Web?")
		assert.False(t, ac.CanRead(user2, webResource), "Can user2 read Web?")
		assert.True(t, ac.CanCreate(group1, webResource), "Can group1 create Web?")
	})

	t.Run("Simple Test", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		user := ac.CreateEntity("user")

		assert.False(t, ac.CanCreate(user, res))
		assert.False(t, ac.CanRead(user, res))
		assert.False(t, ac.CanUpdate(user, res))
		assert.False(t, ac.CanDelete(user, res))
	})

	t.Run("Simple Test All", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		user := ac.CreateEntity("user")
		user.AddPermAll(res, true)

		assert.True(t, ac.CanCreate(user, res))
		assert.True(t, ac.CanRead(user, res))
		assert.True(t, ac.CanUpdate(user, res))
		assert.True(t, ac.CanDelete(user, res))
	})

	t.Run("Simple Test Update", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		user := ac.CreateEntity("user")
		user.AddPermUpdate(res, true)

		assert.False(t, ac.CanCreate(user, res))
		assert.False(t, ac.CanRead(user, res))
		assert.True(t, ac.CanUpdate(user, res))
		assert.False(t, ac.CanDelete(user, res))
	})

	t.Run("Simple Test Create", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		user := ac.CreateEntity("user")
		user.AddPermCreate(res, true)

		assert.True(t, ac.CanCreate(user, res))
		assert.False(t, ac.CanRead(user, res))
		assert.False(t, ac.CanUpdate(user, res))
		assert.False(t, ac.CanDelete(user, res))
	})

	t.Run("Simple Test Read", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		user := ac.CreateEntity("user")
		user.AddPermRead(res, true)

		assert.False(t, ac.CanCreate(user, res))
		assert.True(t, ac.CanRead(user, res))
		assert.False(t, ac.CanUpdate(user, res))
		assert.False(t, ac.CanDelete(user, res))
	})

	t.Run("Simple Test Delete", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		user := ac.CreateEntity("user")
		user.AddPermDelete(res, true)

		assert.False(t, ac.CanCreate(user, res))
		assert.False(t, ac.CanRead(user, res))
		assert.False(t, ac.CanUpdate(user, res))
		assert.True(t, ac.CanDelete(user, res))
	})

	t.Run("Simple Test Owner", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		user := ac.CreateEntity("user")
		res.AddOwners(user)

		assert.True(t, ac.CanCreate(user, res))
		assert.True(t, ac.CanRead(user, res))
		assert.True(t, ac.CanUpdate(user, res))
		assert.True(t, ac.CanDelete(user, res))
	})

	t.Run("Simple Test Allow", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		user := ac.CreateEntity("user")
		ac.Allow(user, res, permission.Read)

		assert.False(t, ac.CanCreate(user, res))
		assert.True(t, ac.CanRead(user, res))
		assert.False(t, ac.CanUpdate(user, res))
		assert.False(t, ac.CanDelete(user, res))
	})

	t.Run("Simple Test Deny", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		user := ac.CreateEntity("user")
		user.AddPermAll(res, true)
		ac.Deny(user, res, permission.Read)

		assert.True(t, ac.CanCreate(user, res))
		assert.False(t, ac.CanRead(user, res))
		assert.True(t, ac.CanUpdate(user, res))
		assert.True(t, ac.CanDelete(user, res))
	})

	t.Run("Entity with children", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		users := ac.CreateEntity("users")
		warfacez := users.CreateChild("warfacez")
		user := ac.CreateEntity("user")
		warfacez.AddPermAll(res, true)
		ac.Deny(warfacez, res, permission.Read)

		assert.False(t, ac.CanCreate(user, res))
		assert.False(t, ac.CanRead(user, res))
		assert.False(t, ac.CanUpdate(user, res))
		assert.False(t, ac.CanDelete(user, res))

		assert.True(t, ac.CanCreate(warfacez, res))
		assert.False(t, ac.CanRead(warfacez, res))
		assert.True(t, ac.CanUpdate(warfacez, res))
		assert.True(t, ac.CanDelete(warfacez, res))
	})

	t.Run("Entity with children 2", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		users := ac.CreateEntity("users")
		warfacez := users.CreateChild("warfacez")
		users.AddChildren(warfacez)
		warfacez.AddParents(users)
		user := ac.CreateEntity("user")
		users.AddPermAll(res, true)
		ac.Deny(warfacez, res, permission.Read)

		assert.False(t, ac.CanCreate(user, res))
		assert.False(t, ac.CanRead(user, res))
		assert.False(t, ac.CanUpdate(user, res))
		assert.False(t, ac.CanDelete(user, res))

		assert.True(t, ac.CanCreate(warfacez, res))
		assert.False(t, ac.CanRead(warfacez, res))
		assert.True(t, ac.CanUpdate(warfacez, res))
		assert.True(t, ac.CanDelete(warfacez, res))
	})

	t.Run("Entity add parents", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")
		users := ac.CreateEntity("users")
		warfacez := ac.CreateEntity("warfacez")
		warfacez.AddParents(users)
		user := ac.CreateEntity("user")
		users.AddPermAll(res, true)
		ac.Deny(warfacez, res, permission.Read)

		warfacez.Deny(res, permission.Create)
		user.Allow(res, permission.Read)

		assert.False(t, ac.CanCreate(user, res))
		assert.True(t, ac.CanRead(user, res))
		assert.False(t, ac.CanUpdate(user, res))
		assert.False(t, ac.CanDelete(user, res))

		assert.False(t, ac.CanCreate(warfacez, res))
		assert.False(t, ac.CanRead(warfacez, res))
		assert.True(t, ac.CanUpdate(warfacez, res))
		assert.True(t, ac.CanDelete(warfacez, res))
	})

	t.Run("Entity as users and groups", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("news")

		users := ac.CreateEntity("users")
		groups := ac.CreateEntity("groups")

		adminGroup := groups.CreateChild("admin")
		userGroup := groups.CreateChild("user")
		moderatorGroup := groups.CreateChild("moderator")

		warfacez := users.CreateChild("warfacez")
		basicUser := users.CreateChild("basicUser")
		modUser := users.CreateChild("modUser")

		adminGroup.AddChildren(warfacez)
		userGroup.AddChildren(basicUser)
		moderatorGroup.AddChildren(modUser)

		userGroup.Allow(res, permission.Read)
		moderatorGroup.Allow(res, permission.Read, permission.Create, permission.Update)
		adminGroup.Allow(res, permission.All)

		assert.False(t, ac.CanCreate(userGroup, res))
		assert.True(t, ac.CanRead(userGroup, res))
		assert.False(t, ac.CanUpdate(userGroup, res))
		assert.False(t, ac.CanDelete(userGroup, res))

		assert.True(t, ac.CanCreate(moderatorGroup, res))
		assert.True(t, ac.CanRead(moderatorGroup, res))
		assert.True(t, ac.CanUpdate(moderatorGroup, res))
		assert.False(t, ac.CanDelete(moderatorGroup, res))

		assert.True(t, ac.CanCreate(adminGroup, res))
		assert.True(t, ac.CanRead(adminGroup, res))
		assert.True(t, ac.CanUpdate(adminGroup, res))
		assert.True(t, ac.CanDelete(adminGroup, res))

		assert.False(t, ac.CanCreate(basicUser, res))
		assert.True(t, ac.CanRead(basicUser, res))
		assert.False(t, ac.CanUpdate(basicUser, res))
		assert.False(t, ac.CanDelete(basicUser, res))

		assert.True(t, ac.CanCreate(modUser, res))
		assert.True(t, ac.CanRead(modUser, res))
		assert.True(t, ac.CanUpdate(modUser, res))
		assert.False(t, ac.CanDelete(modUser, res))

		assert.True(t, ac.CanCreate(warfacez, res))
		assert.True(t, ac.CanRead(warfacez, res))
		assert.True(t, ac.CanUpdate(warfacez, res))
		assert.True(t, ac.CanDelete(warfacez, res))
	})

	t.Run("Entity as users and groups with subEntity", func(t *testing.T) {
		ac := permission.NewAccessControl()

		res := ac.CreateResource("website")
		newsRes := res.CreateSub("news")
		newsRes.CreateSubs("1", "2", "3")

		users := ac.CreateEntity("users")
		groups := ac.CreateEntity("groups")

		adminGroup := groups.CreateChild("adminGroup")
		userGroup := groups.CreateChild("userGroup")
		moderatorGroup := groups.CreateChild("moderatorGroup")

		warfacez := users.CreateChild("warfacez")
		basicUser := users.CreateChild("basicUser")
		modUser := users.CreateChild("modUser")

		adminGroup.AddChildren(warfacez)
		userGroup.AddChildren(basicUser)
		moderatorGroup.AddChildren(modUser)

		userGroup.Allow(res, permission.Read)
		moderatorGroup.Allow(res, permission.Read, permission.Create, permission.Update)
		adminGroup.Allow(res, permission.All)
		moderatorGroup.Deny(newsRes)

		news3 := newsRes.GetSub("3")
		basicUser.Allow(news3, permission.Update)

		news1 := newsRes.GetSub("1").AddOwners(modUser)

		assert.True(t, ac.CanCreate(moderatorGroup, news1))
		assert.True(t, ac.CanRead(moderatorGroup, news1))
		assert.True(t, ac.CanUpdate(moderatorGroup, news1))
		assert.False(t, ac.CanDelete(moderatorGroup, news1))

		assert.True(t, ac.CanRead(basicUser, news3))
		assert.False(t, ac.CanCreate(basicUser, news3))
		assert.True(t, ac.CanUpdate(basicUser, news3))
		assert.False(t, ac.CanDelete(basicUser, news3))

		assert.True(t, ac.CanCreate(moderatorGroup, newsRes))
		assert.True(t, ac.CanRead(moderatorGroup, newsRes))
		assert.True(t, ac.CanUpdate(moderatorGroup, newsRes))
		assert.False(t, ac.CanDelete(moderatorGroup, newsRes))

		assert.False(t, ac.CanCreate(userGroup, res))
		assert.True(t, ac.CanRead(userGroup, res))
		assert.False(t, ac.CanUpdate(userGroup, res))
		assert.False(t, ac.CanDelete(userGroup, res))

		assert.True(t, ac.CanCreate(moderatorGroup, res))
		assert.True(t, ac.CanRead(moderatorGroup, res))
		assert.True(t, ac.CanUpdate(moderatorGroup, res))
		assert.False(t, ac.CanDelete(moderatorGroup, res))

		assert.True(t, ac.CanCreate(adminGroup, res))
		assert.True(t, ac.CanRead(adminGroup, res))
		assert.True(t, ac.CanUpdate(adminGroup, res))
		assert.True(t, ac.CanDelete(adminGroup, res))

		assert.False(t, ac.CanCreate(basicUser, res))
		assert.True(t, ac.CanRead(basicUser, res))
		assert.False(t, ac.CanUpdate(basicUser, res))
		assert.False(t, ac.CanDelete(basicUser, res))

		assert.True(t, ac.CanCreate(modUser, res))
		assert.True(t, ac.CanRead(modUser, res))
		assert.True(t, ac.CanUpdate(modUser, res))
		assert.False(t, ac.CanDelete(modUser, res))

		assert.True(t, ac.CanCreate(warfacez, res))
		assert.True(t, ac.CanRead(warfacez, res))
		assert.True(t, ac.CanUpdate(warfacez, res))
		assert.True(t, ac.CanDelete(warfacez, res))

		assert.True(t, ac.CanCreate(modUser, news1))
		assert.True(t, ac.CanRead(modUser, news1))
		assert.True(t, ac.CanUpdate(modUser, news1))
		assert.True(t, ac.CanDelete(modUser, news1))
	})
}
