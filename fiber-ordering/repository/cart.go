package repository

import (
	"context"
	"database/sql"
	"fiber-ordering/entity"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

type cartRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

func (r cartRepo) GetAll() (res []entity.Cart, err error) {
	table := "cart"
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
		var r entity.Cart
		err = rows.Scan(
			&r.CartID,
			&r.Quantity,
			&r.TotalPrice,
			&r.ProductID,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}

func (r cartRepo) GetByID(cartID uint32) (res *entity.Cart, err error) {
	table := "cart"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"cart_id": cartID,
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
		var r entity.Cart
		err = rows.Scan(
			&r.CartID,
			&r.Quantity,
			&r.TotalPrice,
			&r.ProductID,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}
	return
}

func (r cartRepo) UpdateByID(c context.Context, cartID uint32, cart *entity.Cart) (err error) {
	table := "cart"
	query := sq.Update(table).
		Where(sq.Eq{
			"cart_id": cartID,
		}).
		SetMap(map[string]interface{}{
			"quantity":		cart.Quantity,
			"total_price":	cart.TotalPrice,
			"product_id":	cart.ProductID,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err = query.ExecContext(c)

	if err != nil {
		return err
	}

	return nil
}

func (r cartRepo) DeleteByID(c context.Context, cartID uint32) error {
	table := "cart"
	query := sq.Delete("").
		From(table).
		Where(
			sq.Eq{
				"cart_id": cartID,
			},
		).
		RunWith(r.Writer)
	_, err := query.ExecContext(c)
	if err != nil {
		return err
	}
	return nil
}

func (r cartRepo) Store(c context.Context, ord *entity.Cart) (cartID uint32, err error) {
	table := "cart"
	query := sq.Insert(table).
		Columns(
			"cart_id",
			"quantity",
			"total_price",
			"product_id",
		).
		Values(
			ord.CartID,
			ord.Quantity,
			ord.TotalPrice,
			ord.ProductID,
		).
		PlaceholderFormat(sq.Question)

	sqlInsert, argsInsert, err := query.ToSql()
	res, err := r.Writer.ExecContext(
		c,
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
	cartID = uint32(id)

	return
}

func (r cartRepo) UpdateTotalPrice(c context.Context, cartID uint32, updatePrice float32) (err error) {
	table := "cart"
	query := sq.Update(table).
		Where(sq.Eq{
			"cart_id": cartID,
		}).
		SetMap(map[string]interface{}{
			"total_price":	updatePrice,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err = query.ExecContext(c)

	if err != nil {
		return err
	}

	return nil
}

// NewCartRepo create new cart repo
func NewCartRepo(reader, writer *sql.DB) entity.CartRepository {
	return &cartRepo{
		Reader: reader,
		Writer: writer,
	}
}