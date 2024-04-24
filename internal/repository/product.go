package repository

import (
	"context"
	"database/sql"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/irvankadhafi/go-boilerplate/database"
	"github.com/irvankadhafi/go-boilerplate/internal/model"
	"github.com/irvankadhafi/go-boilerplate/pkg/pagination"
)

//go:generate mockgen -source=product.go -destination=mock/product.go -package=repository
type Product interface {
	// Create inserts product to db, return productID and error
	Create(ctx context.Context, product *model.Product) (productID int64, err error)
	// GetList return products filtered using pagination
	GetList(ctx context.Context, pagination *pagination.Pagination) (products []*model.Product, err error)
	GetDetail(ctx context.Context, id int) (product *model.Product, err error)
	// Update will update product by id for every field that is not default value
	Update(ctx context.Context, product *model.Product) (productID int64, err error)
}

type productRepository struct {
	dbSql *database.DbSql
}

func NewProductRepository(dbSql *database.DbSql) Product {
	return &productRepository{
		dbSql: dbSql,
	}
}

func (p *productRepository) GetDbSql() *database.DbSql {
	return p.dbSql
}

const (
	insertProductQuery = `INSERT INTO products 
	("name", price, stock, description, image_url, created_at, updated_at, deleted_at) 
	VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id;`
	getProductQuery = `WITH FinalTable AS (
		SELECT
			p.id,
			p.name,
			p.price,
			p.stock,
			p.description,
			p.image_url,
			p.created_at,
			p.updated_at,
			p.deleted_at
		FROM
			products p
		WHERE
			p.deleted_at IS NULL
	)`
	updateProductQuery = `UPDATE products SET 
	name=CASE WHEN $1 <> '' THEN $1 ELSE name END, 
	price=CASE WHEN $2 <> 0 THEN $2 ELSE price END, 
	stock=CASE WHEN $3 <> 0 THEN $3 ELSE stock END, 
	description=CASE WHEN $4 <> '' THEN $4 ELSE description END, 
	image_url=CASE WHEN $5 <> '' THEN $5 ELSE image_url END, 
	updated_at=$6,
	deleted_at=CASE WHEN $7::TEXT <> '' THEN TO_TIMESTAMP($7, 'YYYY-MM-DD HH24:MI:SS') ELSE deleted_at END 
	WHERE id=$8`
)

// Create inserts product to db, return productID and error
func (pr *productRepository) Create(ctx context.Context, product *model.Product) (int64, error) {

	sqlStmt, stmtErr := pr.GetDbSql().SqlDb.Prepare(insertProductQuery)
	if stmtErr != nil {
		return 0, stmtErr
	}
	defer sqlStmt.Close()

	ctxTimeout, ctxCancel := context.WithTimeout(ctx, pr.GetDbSql().GetTimeout())
	defer ctxCancel()

	var lastInsertedId int64

	now := time.Now().UTC()

	execErr := sqlStmt.QueryRowContext(ctxTimeout,
		product.Name,
		product.Price,
		product.Stock,
		product.Description,
		product.ImageUrl,
		now,
		now,
		nil,
	).Scan(&lastInsertedId)
	if execErr != nil {
		return 0, execErr
	}

	return lastInsertedId, nil
}

func (pr *productRepository) GetDetail(ctx context.Context, id int) (product *model.Product, err error) {
	query := fmt.Sprintf(`%s SELECT
							ft.id,
							ft.name,
							ft.price,
							ft.stock,
							ft.description,
							ft.image_url,
							ft.created_at,
							ft.updated_at,
							ft.deleted_at
						FROM
							FinalTable ft
						WHERE ft.id = $1
						LIMIT 1`, getProductQuery)

	sqlStmt, stmtErr := pr.GetDbSql().SqlDb.Prepare(query)
	if stmtErr != nil {
		return nil, stmtErr
	}
	defer sqlStmt.Close()

	ctxTimeout, ctxCancel := context.WithTimeout(ctx, pr.GetDbSql().GetTimeout())
	defer ctxCancel()

	result := sqlStmt.QueryRowContext(ctxTimeout, id)

	var createdAt, updatedAt, deletedAt sql.NullTime

	if scanErr := result.Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.Description,
		&product.ImageUrl,
		&createdAt,
		&updatedAt,
		&deletedAt,
	); scanErr != nil {
		return nil, scanErr
	}

	product.CreatedAt = createdAt.Time.In(time.Local).Format(time.DateTime)
	product.UpdatedAt = updatedAt.Time.In(time.Local).Format(time.DateTime)
	if deletedAt.Valid{
		product.DeletedAt = deletedAt.Time.In(time.Local).Format(time.DateTime)
	}

	return product, nil
}

// GetList return products filtered using pagination
func (pr *productRepository) GetList(ctx context.Context, pagination *pagination.Pagination) ([]*model.Product, error) {
	sort := `id`

	allowedOrderBy := []string{
		"id",
		"name",
		"price",
		"stock",
		"description",
	}

	if len(pagination.Sort) > 0 {
		if slices.Contains(allowedOrderBy, pagination.Sort) {
			sort = pagination.Sort
		}
	}

	query := fmt.Sprintf(`%s SELECT ft.*, COUNT(1) OVER() AS total_rows FROM FinalTable ft 
	WHERE 
		CASE 
			WHEN $3 = '' THEN TRUE 
			ELSE (
				ft.name ILIKE $3 
				OR ft.price::TEXT ILIKE $3 
				OR ft.stock::TEXT ILIKE $3 
				OR ft.description ILIKE $3
			)
		END`, getProductQuery)

	if strings.EqualFold(pagination.Dir, "DESC") {
		query = fmt.Sprintf("%s ORDER BY ft.%s %s", query, sort, "DESC")
	} else {
		query = fmt.Sprintf("%s ORDER BY ft.%s %s", query, sort, "ASC")
	}

	query = fmt.Sprintf("%s LIMIT $1", query)
	query = fmt.Sprintf("%s OFFSET $2", query)

	var filter string
	if pagination.Sort != "" {
		filter = fmt.Sprint("%", pagination.Sort, "%")
	}

	var totalRows, totalPages int64
	sqlStmt, stmtErr := pr.GetDbSql().SqlDb.Prepare(query)
	if stmtErr != nil {
		return nil, stmtErr
	}
	defer sqlStmt.Close()

	ctxTimeout, ctxCancel := context.WithTimeout(ctx, pr.GetDbSql().GetTimeout())
	defer ctxCancel()

	result := make([]*model.Product, 0)
	dataRow, execErr := sqlStmt.QueryContext(ctxTimeout,
		pagination.Limit,
		pagination.Offset,
		filter,
	)
	if execErr != nil {
		return nil, execErr
	}
	for dataRow.Next() {
		product := &model.Product{}
		var createdAt, updatedAt, deletedAt sql.NullTime
		errScan := dataRow.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.Description,
			&product.ImageUrl,
			&createdAt,
			&updatedAt,
			&deletedAt,
			&totalRows,
		)
		if errScan != nil {
			return nil, errScan
		}

		product.CreatedAt = createdAt.Time.In(time.Local).Format(time.DateTime)
		product.UpdatedAt = updatedAt.Time.In(time.Local).Format(time.DateTime)
		if deletedAt.Valid{
			product.DeletedAt = deletedAt.Time.In(time.Local).Format(time.DateTime)
		}

		result = append(result, product)
	}

	totalPages = pagination.CalculateTotalPage(int64(pagination.Limit), totalRows)

	pagination.TotalPages = totalPages
	pagination.TotalRows = totalRows

	return result, nil
}

// Update will update product by id for every field that is not default value
func (pr *productRepository) Update(ctx context.Context, product *model.Product) (productID int64, err error) {
	sqlStmt, stmtErr := pr.GetDbSql().SqlDb.Prepare(updateProductQuery)
	if stmtErr != nil {
		return 0, stmtErr
	}
	defer sqlStmt.Close()

	ctxTimeout, ctxCancel := context.WithTimeout(ctx, pr.GetDbSql().GetTimeout())
	defer ctxCancel()

	now := time.Now().UTC()
	_, execErr := sqlStmt.ExecContext(ctxTimeout,
		product.Name,
		product.Price,
		product.Stock,
		product.Description,
		product.ImageUrl,
		now,
		product.DeletedAt,
		product.Id,
	)
	if execErr != nil {
		return 0, execErr
	}

	return product.Id, nil
}
