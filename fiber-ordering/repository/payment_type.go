package repository

import (
	"context"
	"database/sql"
	"fiber-ordering/entity"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

type paymentTypeRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

// NewPaymentTypeRepo will initiate repo
func NewPaymentTypeRepo(reader, writer *sql.DB) entity.PaymentTypeRepository {
	return &paymentTypeRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (t *paymentTypeRepo) GetAll() (res []entity.PaymentType, err error) {
	table := "payment_type"

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
		var r entity.PaymentType
		err = rows.Scan(
			&r.PaymentTypeID,
			&r.Method,
			&r.Company,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (t *paymentTypeRepo) GetByID(paymentTypeID uint32) (res *entity.PaymentType, err error) {
	table := "payment_type"

	query := sq.Select("*").
		From(table).
		RunWith(t.Reader).
		PlaceholderFormat(sq.Question).
		Where(sq.Eq{
			"payment_type_id": paymentTypeID,
		})

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r entity.PaymentType
		err = rows.Scan(
			&r.PaymentTypeID,
			&r.Method,
			&r.Company,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (t *paymentTypeRepo) DeleteByID(ctx context.Context, paymentTypeID uint32) error {
	table := "payment_type"

	query := sq.Delete("").
		From(table).
		RunWith(t.Reader).
		PlaceholderFormat(sq.Question).
		Where(sq.Eq{
			"payment_type_id": paymentTypeID,
		})
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (t *paymentTypeRepo) Store(ctx context.Context, paymentType *entity.PaymentType) (ptID uint32, err error) {
	table := "payment_type"

	query := sq.Insert(table).
		Columns(
			"payment_method",
			"payment_company",
		).
		Values(
			paymentType.Method,
			paymentType.Company,
		).
		RunWith(t.Writer).
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
	ptID = uint32(id)

	return
}