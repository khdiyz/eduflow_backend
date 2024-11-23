package postgres

import (
	"eduflow/internal/models"
	"eduflow/pkg/logger"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	branchesTable = "branches"
)

type BranchRepository struct {
	db     *sqlx.DB
	logger *logger.Logger
}

func NewBranchRepository(db *sqlx.DB, logger *logger.Logger) *BranchRepository {
	return &BranchRepository{
		db:     db,
		logger: logger,
	}
}

func (r *BranchRepository) Create(input models.CreateBranch) (uuid.UUID, error) {
	id := uuid.New()

	query, args, err := sq.Insert(branchesTable).Columns(
		"id",
		"school_id",
		"name",
		"address",
		"email",
		"phone_number",
		"opening_hours",
		"status",
	).Values(
		id,
		input.SchoolId,
		input.Name,
		input.Address,
		input.Email,
		input.PhoneNumber,
		input.OpeningHours,
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

func (r *BranchRepository) GetList(filter models.BranchFilter) ([]models.Branch, int, error) {
	// Squirrel query builder
	query := sq.Select(
		"id",
		"school_id",
		"name",
		"address",
		"email",
		"phone_number",
		"opening_hours",
		"status",
		"created_at",
		"updated_at",
	).From(branchesTable).Where("TRUE")

	// Get the total count
	countQuery := sq.Select("COUNT(id)").From(branchesTable).Where("TRUE")

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

	if filter.SchoolId != uuid.Nil {
		query = query.Where("school_id = ?", filter.SchoolId)
		countQuery = countQuery.Where("school_id = ?", filter.SchoolId)
	}

	// Add pagination (limit and offset)
	if filter.Limit > 0 {
		query = query.Limit(uint64(filter.Limit))
		query = query.Offset(uint64(filter.Offset))
	}

	// Get the query string and parameters
	sqlQuery, args, err := query.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}

	// Execute the main query
	branches := []models.Branch{}
	rows, err := r.db.Query(sqlQuery, args...)
	if err != nil {
		r.logger.Error(err)
		return nil, 0, err
	}
	defer rows.Close()

	// Scan rows into branches
	for rows.Next() {
		var branch models.Branch
		if err := rows.Scan(
			&branch.Id,
			&branch.SchoolId,
			&branch.Name,
			&branch.Address,
			&branch.Email,
			&branch.PhoneNumber,
			&branch.OpeningHours,
			&branch.Status,
			&branch.CreatedAt,
			&branch.UpdatedAt,
		); err != nil {
			r.logger.Error(err)
			return nil, 0, err
		}
		branches = append(branches, branch)
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

	return branches, total, nil
}

func (r *BranchRepository) GetById(id uuid.UUID) (models.Branch, error) {
	var branch models.Branch

	// Squirrel query builder
	query, args, err := sq.Select(
		"id",
		"school_id",
		"name",
		"address",
		"email",
		"phone_number",
		"opening_hours",
		"status",
		"created_at",
		"updated_at",
	).From(branchesTable).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		r.logger.Error(err)
		return models.Branch{}, err
	}

	// Execute the query
	if err := r.db.QueryRow(query, args...).Scan(
		&branch.Id,
		&branch.SchoolId,
		&branch.Name,
		&branch.Address,
		&branch.Email,
		&branch.PhoneNumber,
		&branch.OpeningHours,
		&branch.Status,
		&branch.CreatedAt,
		&branch.UpdatedAt,
	); err != nil {
		r.logger.Error(err)
		return models.Branch{}, err
	}

	return branch, nil
}

func (r *BranchRepository) Update(input models.UpdateBranch) error {
	query := sq.Update(branchesTable).PlaceholderFormat(sq.Dollar)

	query = query.Set("school_id", input.SchoolId)
	query = query.Set("name", input.Name)
	query = query.Set("address", input.Address)
	query = query.Set("email", input.Email)
	query = query.Set("phone_number", input.PhoneNumber)
	query = query.Set("opening_hours", input.OpeningHours)
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

func (r *BranchRepository) Delete(id uuid.UUID) error {
	query := sq.Delete(branchesTable).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar)

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
