package repository

import (
	"context"
	"database/sql"
	"fiber-ordering/entity"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

type productRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewProductRepo create new product repo
func NewProductRepo(reader, writer *sql.DB) entity.ProductRepository {
	return &productRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *productRepo) GetAll() (res []entity.Product, err error) {
	table := "product"

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
		var r entity.Product
		err = rows.Scan(
			&r.ProductID,
			&r.CategoryID,
			&r.Code,
			&r.Name,
			&r.ProductImage,
			&r.Price,
			&r.ProductStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *productRepo) GetByID(productID uint32) (res *entity.Product, err error) {
	table := "product"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"product_id": productID,
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
		var r entity.Product
		err = rows.Scan(
			&r.ProductID,
			&r.CategoryID,
			&r.Code,
			&r.Name,
			&r.ProductImage,
			&r.Price,
			&r.ProductStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *productRepo) UpdateByID(ctx context.Context, productID uint32, order *entity.Product) error {
	table := "product"
	query := sq.Update(table).
		Where(sq.Eq{
			"product_id": productID,
		}).
		SetMap(map[string]interface{}{
			"product_name":    		order.Name,
			"product_code":    		order.Code,
			"product_price":    	order.Price,
			"product_image": 		order.ProductImage,
			"product_category_id":  order.CategoryID,
			"product_status":   	order.ProductStatus,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *productRepo) DeleteByID(ctx context.Context, productID uint32) error {
	table := "product"
	query := sq.Delete("").
		From(table).
		Where(
			sq.Eq{
				"product_id": productID,
			},
		).
		RunWith(r.Writer)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *productRepo) Store(ctx context.Context, ord *entity.Product) (productID uint32, err error) {
	table := "product"
	query := sq.Insert(table).
		Columns(
			"product_name",
			"product_code",
			"product_image",
			"product_price",
			"product_category_id",
			"product_status",
		).
		Values(
			ord.Name,
			ord.Code,
			ord.ProductImage,
			ord.Price,
			ord.CategoryID,
			ord.ProductStatus,
		).
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
	productID = uint32(id)

	return
}