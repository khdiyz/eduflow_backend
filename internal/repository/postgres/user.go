package postgres

import (
	"database/sql"
	"eduflow/internal/models"
	"eduflow/pkg/logger"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserRepository struct {
	db     *sqlx.DB
	logger *logger.Logger
}

func NewUserRepository(db *sqlx.DB, logger *logger.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) Create(input models.CreateUser) (uuid.UUID, error) {
	id := uuid.New()

	query := `
	INSERT INTO users (
		id,
		role_id,
		first_name,
		last_name,
		phone_numbers,
		username,
		password,
		branch_id
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`

	if _, err := r.db.Exec(query, id,
		input.RoleId,
		input.FirstName,
		input.LastName,
		pq.Array(input.PhoneNumbers),
		input.Username,
		input.Password,
		input.BranchId,
	); err != nil {
		r.logger.Error(err)
		return uuid.Nil, err
	}

	return id, nil
}

func (r *UserRepository) GetList(filter models.UserFilter) ([]models.User, int, error) {
	// Squirrel query builder
	query := squirrel.Select(
		"id",
		"role_id",
		"first_name",
		"last_name",
		"phone_numbers",
		"username",
		"password",
		"status",
		"branch_id",
		"created_at",
		"updated_at",
	).
		From("users").
		Where("TRUE")

	// Filter parameters (conditions)
	var conditions []string
	params := make(map[string]interface{})

	// Search filter
	if filter.Search != "" {
		conditions = append(conditions, "(first_name || last_name || username) ILIKE :search")
		params["search"] = "%" + filter.Search + "%"
	}

	// Role filter
	if filter.RoleId != uuid.Nil {
		conditions = append(conditions, "role_id = :role_id")
		params["role_id"] = filter.RoleId
	}

	// Status filter
	if filter.Status != nil {
		conditions = append(conditions, "status = :status")
		params["status"] = *filter.Status
	}

	// Branch filter
	if filter.BranchId != uuid.Nil {
		conditions = append(conditions, "branch_id = :branch_id")
		params["branch_id"] = filter.BranchId
	}

	// Add the conditions to the query if they exist
	if len(conditions) > 0 {
		query = query.Where(squirrel.Expr(strings.Join(conditions, " AND ")))
	}

	// Add pagination (limit and offset)
	query = query.Limit(uint64(filter.Limit)).Offset(uint64(filter.Offset))

	// Get the query string and parameters
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}

	// Execute the main query
	users := []models.User{}
	rows, err := r.db.Query(sqlQuery, args...)
	if err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}
	defer rows.Close()

	// Scan rows into users
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.Id,
			&user.RoleId,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumbers,
			&user.Username,
			&user.Password,
			&user.Status,
			&user.BranchId,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			r.logger.Error(err)
			return nil, 0, err
		}
		users = append(users, user)
	}

	// Get the total count
	countQuery := squirrel.Select("COUNT(id)").
		From("users").
		Where("TRUE")

	// Add the same conditions to the count query
	if len(conditions) > 0 {
		countQuery = countQuery.Where(squirrel.Expr(strings.Join(conditions, " AND ")))
	}

	countSqlQuery, countArgs, err := countQuery.ToSql()
	if err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}

	var total int
	if err := r.db.Get(&total, countSqlQuery, countArgs...); err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}

	return users, total, nil
}

func (r *UserRepository) GetById(id uuid.UUID) (models.User, error) {
	var user models.User

	// Squirrel query builder
	query, args, err := squirrel.
		Select(
			"id",
			"role_id",
			"first_name",
			"last_name",
			"phone_numbers",
			"username",
			"password",
			"status",
			"branch_id",
			"created_at",
			"updated_at",
		).
		From("users").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		r.logger.Error(err)
		return models.User{}, err
	}

	// Execute the query
	if err := r.db.QueryRow(query, append(args, id)...).Scan(
		&user.Id,
		&user.RoleId,
		&user.FirstName,
		&user.LastName,
		pq.Array(&user.PhoneNumbers),
		&user.Username,
		&user.Password,
		&user.Status,
		&user.BranchId,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		r.logger.Error(err)
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetByUsername(username string) (models.User, error) {
	var user models.User

	var (
		lastName sql.NullString
		branchId sql.NullString
	)

	// Squirrel query builder
	query, args, err := squirrel.
		Select(
			"id",
			"role_id",
			"first_name",
			"last_name",
			"phone_numbers",
			"username",
			"password",
			"status",
			"branch_id",
			"created_at",
			"updated_at",
		).
		From("users").
		Where(squirrel.Eq{"username": username}). // Use squirrel.Eq for equality comparison
		PlaceholderFormat(squirrel.Dollar).       // Use $ placeholders for PostgreSQL
		ToSql()

	if err != nil {
		r.logger.Error("Failed to build query:", err)
		return models.User{}, err
	}

	// Execute the query
	if err := r.db.QueryRow(query, args...).Scan(
		&user.Id,
		&user.RoleId,
		&user.FirstName,
		&lastName,
		pq.Array(&user.PhoneNumbers),
		&user.Username,
		&user.Password,
		&user.Status,
		&branchId,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		r.logger.Error("Failed to execute query:", err)
		return models.User{}, err
	}

	user.LastName = lastName.String
	user.BranchId = branchId.String

	return user, nil
}
