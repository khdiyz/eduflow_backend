package constants

const (
	EnvironmentProduction  = "production"
	EnvironmentDevelopment = "development"

	// token types
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"

	// Pagination
	DefaultPage     = "1"
	DefaultPageSize = "10"

	// Roles
	RoleSuperAdmin = "SUPER ADMIN"
	RoleAdmin      = "ADMIN"
	RoleMentor     = "MENTOR"

	// Token data
	UserCtx = "user_id"
	RoleCtx = "role_name"
)
