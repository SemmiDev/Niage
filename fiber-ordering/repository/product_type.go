package repository

import (
	"context"
	"database/sql"
	"fiber-ordering/entity"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

type productTypeRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewProductCategoryRepo create new order repo
func NewProductCategoryRepo(reader, writer *sql.DB) entity.ProductCategoryRepository {
	return &productTypeRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *productTypeRepo) GetAll() (res []entity.ProductCategory, err error) {
	table := "product_category"

	query := sq.Select("*").
		From(table).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r entity.ProductCategory
		err = rows.Scan(
			&r.ProductCategoryID,
			&r.CategoryName,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}

func (r *productTypeRepo) GetByID(productCategoryID uint32) (res *entity.ProductCategory, err error) {
	table := "product_category"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"product_category_id": productCategoryID,
			},
		).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r entity.ProductCategory
		err = rows.Scan(
			&r.ProductCategoryID,
			&r.CategoryName,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *productTypeRepo) UpdateByID(ctx context.Context, productCategoryID uint32, order *entity.ProductCategory) error {
	table := "product_category"
	query := sq.Update(table).
		Where(sq.Eq{
			"product_category_id": productCategoryID,
		}).
		SetMap(map[string]interface{}{
			"product_category_name":   order.CategoryName,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *productTypeRepo) DeleteByID(ctx context.Context, productCategoryID uint32) error {
	table := "product_category"
	query := sq.Delete("").
		From(table).
		Where(
			sq.Eq{
				"product_category_id": productCategoryID,
			},
		).
		RunWith(r.Writer)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *productTypeRepo) Store(ctx context.Context, ord *entity.ProductCategory) (productCategoryID uint32, err error) {
	table := "product_category"
	query := sq.Insert(table).
		Columns("product_category_name").
		Values(ord.CategoryName).
		PlaceholderFormat(sq.Question)

	sqlInsert, argsInsert, err := query.ToSql()
	res, err := r.Writer.ExecContext(
		ctx,
		sqlInsert,
		argsInsert...,
	)
	if err != nil {
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		return
	}
	productCategoryID = uint32(id)

	return
}