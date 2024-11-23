package seeders

import (
	"eduflow/config"
	"eduflow/pkg/helper"

	"github.com/jmoiron/sqlx"
)

func SeedSuperAdmin(cfg *config.Config, db *sqlx.DB) error {
	password, _ := helper.GenerateHash(cfg.SuperAdminPassword)

	superAdmin := struct {
		Id        string
		RoleId    string
		FirstName string
		Username  string
		Password  string
	}{
		Id:        config.UserSuperAdminId,
		RoleId:    config.RoleSuperAdminId,
		FirstName: "Super Admin",
		Username:  cfg.SuperAdminUsername,
	}

	query := `
		INSERT INTO users (id, role_id, first_name, username, password, status) 
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (username) DO NOTHING;`

	_, err := db.Exec(query, superAdmin.Id, superAdmin.RoleId, superAdmin.FirstName, superAdmin.Username, password, true)
	if err != nil {
		return err
	}

	return nil
}
