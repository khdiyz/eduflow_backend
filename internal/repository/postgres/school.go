package postgres

import (
	"eduflow/internal/models"
	"eduflow/pkg/logger"
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	errNoRowsAffected = errors.New("no rows affected")
)

const (
	schoolsTable = "schools"
)

type SchoolRepository struct {
	db     *sqlx.DB
	logger *logger.Logger
}

func NewSchoolRepository(db *sqlx.DB, logger *logger.Logger) *SchoolRepository {
	return &SchoolRepository{
		db:     db,
		logger: logger,
	}
}

func (r *SchoolRepository) Create(input models.CreateSchool) (uuid.UUID, error) {
	id := uuid.New()

	query, args, err := sq.Insert(schoolsTable).Columns(
		"id",
		"name",
		"address",
		"email",
		"phone_number",
		"currency",
		"timezone",
		"status",
	).Values(
		id,
		input.Name,
		input.Address,
		input.Email,
		input.PhoneNumber,
		input.Currency,
		input.Timezone,
		input.Status,
	).PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		r.logger.Error(err)
		return uuid.Nil, err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		r.logger.Error(err)
		return uuid.Nil, err
	}

	return id, nil
}

func (r *SchoolRepository) GetList(filter models.SchoolFilter) ([]models.School, int, error) {
	// Squirrel query builder
	query := sq.Select(
		"id",
		"name",
		"address",
		"email",
		"phone_number",
		"currency",
		"timezone",
		"status",
		"created_at",
		"updated_at",
	).From(schoolsTable).Where("TRUE")

	// Get the total count
	countQuery := sq.Select("COUNT(id)").
		From(schoolsTable).
		Where("TRUE")

	// Filter parameters (conditions)
	if filter.Search != "" {
		query = query.Where(
			"(name || address || email || phone_number) ILIKE ?",
			"%"+filter.Search+"%",
		)
		countQuery = countQuery.Where(
			"(name || address || email || phone_number) ILIKE ?",
			"%"+filter.Search+"%",
		)
	}

	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
		countQuery = countQuery.Where("status = ?", *filter.Status)
	}

	// Add pagination (limit and offset)
	if filter.Limit > 0 {
		query = query.Limit(uint64(filter.Limit))
	}
	query = query.Offset(uint64(filter.Offset))

	// Get the query string and parameters
	sqlQuery, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}

	// Execute the main query
	schools := []models.School{}
	rows, err := r.db.Query(sqlQuery, args...)
	if err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}
	defer rows.Close()

	// Scan rows into schools
	for rows.Next() {
		var school models.School
		if err := rows.Scan(
			&school.Id,
			&school.Name,
			&school.Address,
			&school.Email,
			&school.PhoneNumber,
			&school.Currency,
			&school.Timezone,
			&school.Status,
			&school.CreatedAt,
			&school.UpdatedAt,
		); err != nil {
			r.logger.Error(err)
			return nil, 0, err
		}
		schools = append(schools, school)
	}

	countSqlQuery, countArgs, err := countQuery.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}

	var total int
	if err := r.db.QueryRow(countSqlQuery, countArgs...).Scan(&total); err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}

	return schools, total, nil
}

func (r *SchoolRepository) GetById(id uuid.UUID) (models.School, error) {
	var school models.School

	// Squirrel query builder
	query, args, err := sq.Select(
		"id",
		"name",
		"address",
		"email",
		"phone_number",
		"currency",
		"timezone",
		"status",
		"created_at",
		"updated_at",
	).From(schoolsTable).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		r.logger.Error(err)
		return models.School{}, err
	}

	// Execute the query
	if err := r.db.QueryRow(query, args...).Scan(
		&school.Id,
		&school.Name,
		&school.Address,
		&school.Email,
		&school.PhoneNumber,
		&school.Currency,
		&school.Timezone,
		&school.Status,
		&school.CreatedAt,
		&school.UpdatedAt,
	); err != nil {
		r.logger.Error(err)
		return models.School{}, err
	}

	return school, nil
}

func (r *SchoolRepository) Update(input models.UpdateSchool) error {
	query := sq.Update(schoolsTable).PlaceholderFormat(sq.Dollar)

	query = query.Set("name", input.Name)
	query = query.Set("address", input.Addess)
	query = query.Set("email", input.Email)
	query = query.Set("phone_number", input.PhoneNumber)
	query = query.Set("currency", input.Currency)
	query = query.Set("timezone", input.Timezone)
	query = query.Set("status", input.Status)
	query = query.Set("updated_at", time.Now())

	query = query.Where(sq.Eq{"id": input.Id})

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		r.logger.Error(err)
		return err
	}

	row, err := r.db.Exec(sqlQuery, args...)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		r.logger.Error(err)
		return err
	}

	if rowAffected == 0 {
		return errNoRowsAffected
	}

	return nil
}

func (r *SchoolRepository) Delete(id uuid.UUID) error {
	query := sq.Delete(schoolsTable).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		r.logger.Error(err)
		return err
	}

	result, err := r.db.Exec(sqlQuery, args...)
	if err != nil {
		r.logger.Error(err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		r.logger.Error(err)
		return err
	}

	if rowsAffected == 0 {
		return errNoRowsAffected
	}

	return nil
}
