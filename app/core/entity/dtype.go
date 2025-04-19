package entity

type UserRole string

const (
	RoleClient   UserRole = "client"
	RoleOwner    UserRole = "owner"
	RoleDelivery UserRole = "delivery"
	RoleAny      UserRole = "any"
	RoleAdmin    UserRole = "admin"
)
