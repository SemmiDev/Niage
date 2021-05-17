package repository

import (
	"context"
	"database/sql"
	"fiber-ordering/entity"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
	"log"
)

type paymentRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewPaymentRepo will initiate repo
func NewPaymentRepo(reader, writer *sql.DB) entity.PaymentRepository {
	return &paymentRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (t *paymentRepo) UpdateStatus(ctx context.Context, paymentID uint32, status int32) error {
	table := "payment"

	query := sq.Update(table).
		Set("payment_status", status).
		RunWith(t.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *paymentRepo) GetAll() (res []entity.Payment, err error) {
	table := "payment"

	query := sq.Select("*").
		From(table).
		RunWith(t.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r entity.Payment
		err = rows.Scan(
			&r.PaymentID,
			&r.OrderID,
			&r.Amount,
			&r.PaymentTypeID,
			&r.PaymentDate,
			&r.PaymentStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (t *paymentRepo) GetByID(paymentID uint32) (res *entity.Payment, err error) {
	table := "payment"

	query := sq.Select("*").
		From(table).
		Where(sq.Eq{
			"payment_id": paymentID,
		}).
		RunWith(t.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r entity.Payment
		err = rows.Scan(
			&r.PaymentID,
			&r.OrderID,
			&r.Amount,
			&r.PaymentTypeID,
			&r.PaymentDate,
			&r.PaymentStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (t *paymentRepo) DeleteByID(ctx context.Context, paymentID uint32) error {
	table := "payment"

	query := sq.Delete("").
		From(table).
		Where(sq.Eq{
			"payment_id": paymentID,
		}).
		RunWith(t.Reader)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *paymentRepo) Store(
	ctx context.Context,
	payment *entity.Payment,
) (paymentID uint32, err error) {
	table := "payment"

	query := sq.Insert(table).
		Columns(
			"order_id",
			"amount",
			"payment_type_id",
			"payment_date",
		).
		Values(
			payment.OrderID,
			payment.Amount,
			payment.PaymentTypeID,
			sq.Expr("NOW()"),
		).RunWith(t.Writer).
		PlaceholderFormat(sq.Question)
	sqlInsert, argsInsert, err := query.ToSql()
	res, err := t.Writer.ExecContext(
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
	paymentID = uint32(id)

	return
}

func (t *paymentRepo) GetListOfPaymentsByCustomerID(ctx context.Context, customerID uint32) (res []entity.Payment, err error) {

	rows, err := t.Reader.QueryContext(ctx, "SELECT p.payment_id, p.order_id, p.amount, p.payment_type_id, p.payment_date, p.payment_status FROM payment p JOIN orders on p.order_id = orders.order_id JOIN customer on orders.customer_id = customer.customer_id WHERE customer.customer_id = ? and p.payment_status = 1", customerID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var r entity.Payment
		err = rows.Scan(
			&r.PaymentID,
			&r.OrderID,
			&r.Amount,
			&r.PaymentTypeID,
			&r.PaymentDate,
			&r.PaymentStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}
	return
}
