package entity

type User struct {
	Username    string       `gorm:"primaryKey;column:username;type:varchar(100)"`
	Password    string       `gorm:"column:password;type:varchar(200)"`
	IsActive    bool         `gorm:"column:is_active;type:boolean"`
	Roles       []Role       `gorm:"many2many:role_user;"`
	Permissions []Permission `gorm:"many2many:permission_user;"`
	Todos       []Todo       `gorm:"foreignKey:Username"`
}

func (user *User) AttachRole(role Role) User {
	user.Roles = append(user.Roles, role)
	return *user
}

func (user User) HasRole(name string) bool {
	for _, v := range user.Roles {
		if v.Name == name {
			return true
		}
	}
	return false
}

func (user User) GetRoles() []string {
	var roles []string
	for _, v := range user.Roles {
		roles = append(roles, v.Name)
	}
	return roles
}

func (user *User) DetachRole(role Role) User {

	for index, v := range user.Roles {
		if v.Name == role.Name {
			// user.Roles = append(user.Roles[:index], user.Roles[index+1]...)
			user.Roles = user.Roles[:index+copy(user.Roles[index:], user.Roles[index+1:])]
			return *user
		}
	}
	return *user
}

func (user *User) AttachPemission(permission Permission) User {
	user.Permissions = append(user.Permissions, permission)
	return *user
}

func (user User) HasPermission(name string) bool {
	for _, v := range user.Permissions {
		if v.Name == name {
			return true
		}
	}
	return false
}

func (user User) GetPermissions() []string {
	var permissions []string
	for _, v := range user.Permissions {
		permissions = append(permissions, v.Name)
	}
	return permissions
}

func (user *User) DetachPermission(permission Permission) User {

	for index, v := range user.Permissions {
		if v.Name == permission.Name {
			user.Permissions = user.Permissions[:index+copy(user.Permissions[index:], user.Permissions[index+1:])]
			return *user
		}
	}
	return *user
}
