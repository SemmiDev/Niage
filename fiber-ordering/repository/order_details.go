package repository

import (
	"context"
	"database/sql"
	"fiber-ordering/entity"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

type orderDetailsRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewOrderDetailsRepo create new order repo
func NewOrderDetailsRepo(reader, writer *sql.DB) entity.OrderDetailsRepository {
	return &orderDetailsRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *orderDetailsRepo) UpdateTotalPrice(ctx context.Context, orderDetailsID uint32) error {
	orderDetail, err := r.GetByID(orderDetailsID)
	if err != nil {
		return err
	}

	var menu entity.Product
	menuQuery := sq.Select("*").
		From("menu").
		Where(
			sq.Eq{
				"menu_id": orderDetail.ProductID,
			},
		).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)
	rows, err := menuQuery.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&menu.ProductID,
			&menu.Code,
			&menu.Name,
			&menu.ProductImage,
			&menu.Price,
			&menu.CategoryID,
			&menu.ProductStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
	}

	newPrice := float32(orderDetail.Quantity) * menu.Price
	err = r.UpdateArbitrary(
		ctx,
		orderDetailsID,
		"total_price",
		newPrice,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderDetailsRepo) GetOrderDetailsByOrderID(
	orderID uint32,
) (res []entity.OrderDetails, err error) {
	table := "order_details"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"order_id": orderID,
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
		var r entity.OrderDetails
		err = rows.Scan(
			&r.OrderDetailsID,
			&r.OrderID,
			&r.ProductID,
			&r.Quantity,
			&r.TotalPrice,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *orderDetailsRepo) GetAll() (res []entity.OrderDetails, err error) {
	table := "order_details"

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
		var r entity.OrderDetails
		err = rows.Scan(
			&r.OrderDetailsID,
			&r.OrderID,
			&r.ProductID,
			&r.Quantity,
			&r.TotalPrice,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *orderDetailsRepo) GetByID(orderDetailsID uint32) (res *entity.OrderDetails, err error) {
	table := "order_details"

	query := sq.Select("*").
		From(table).
		Where(
			sq.Eq{
				"order_details_id": orderDetailsID,
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
		var r entity.OrderDetails
		err = rows.Scan(
			&r.OrderDetailsID,
			&r.OrderID,
			&r.ProductID,
			&r.Quantity,
			&r.TotalPrice,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *orderDetailsRepo) UpdateArbitrary(
	ctx context.Context,
	orderDetailsID uint32,
	columnName string,
	value interface{},
) error {
	table := "order_details"
	query := sq.Update(table).
		Where(sq.Eq{
			"order_details_id": orderDetailsID,
		}).
		Set(columnName, value).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderDetailsRepo) UpdateByID(
	ctx context.Context,
	orderDetailsID uint32,
	order *entity.OrderDetails,
) error {
	table := "order_details"
	query := sq.Update(table).
		Where(sq.Eq{
			"order_details_id": orderDetailsID,
		}).
		SetMap(map[string]interface{}{
			"order_id":    order.OrderID,
			"menu_id":     order.ProductID,
			"quantity":    order.Quantity,
			"total_price": order.TotalPrice,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *orderDetailsRepo) DeleteByID(ctx context.Context, orderDetailsID uint32) error {
	table := "order_details"
	query := sq.Delete("").
		From(table).
		Where(
			sq.Eq{
				"order_details_id": orderDetailsID,
			},
		).
		RunWith(r.Writer)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *orderDetailsRepo) Store(
	ctx context.Context,
	ord *entity.OrderDetails,
) (orddID uint32, err error) {
	table := "order_details"
	query := sq.Insert(table).
		Columns(
			"order_id",
			"product_id",
			"quantity",
			"total_price",
		).
		Values(
			ord.OrderID,
			ord.ProductID,
			ord.Quantity,
			ord.TotalPrice,
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
	orddID = uint32(id)

	return
}