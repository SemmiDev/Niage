package repository

import (
	"context"
	"database/sql"
	"errors"
	"fiber-ordering/entity"
	sq "github.com/Masterminds/squirrel"
	logger "github.com/sirupsen/logrus"
)

type customerRepo struct {
	Reader *sql.DB
	Writer *sql.DB
}

func (r *customerRepo) GetByEmail(email string) (res *entity.Customer, err error) {
	table := "customer"

	query := sq.Select("*").
		From(table).
		Where(sq.Eq{
			"customer_email": email,
		}).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return nil, errors.New("error")
	}
	defer rows.Close()

	for rows.Next() {
		var r entity.Customer
		err = rows.Scan(
			&r.CustomerID,
			&r.FullName,
			&r.Email,
			&r.PhoneNumber,
			&r.Username,
			&r.Password,
			&r.AccountStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}
	return
}

func (r *customerRepo) Login(email string, password string) (string, error) {
	panic("implement me")
}

func (r *customerRepo) UpdateByID(c context.Context, custID uint32, cust *entity.Customer) error {
	table := "customer"
	query := sq.Update(table).
		Where(sq.Eq{
			"customer_id": custID,
		}).
		SetMap(map[string]interface{}{
			"customer_full_name": cust.FullName,
			"customer_email": cust.Email,
			"customer_phone_number": cust.PhoneNumber,
		}).
		RunWith(r.Writer).
		PlaceholderFormat(sq.Question)
	_, err := query.ExecContext(c)
	if err != nil {
		return err
	}

	return nil
}

// NewCustomerRepo will create customer repo
func NewCustomerRepo(reader, writer *sql.DB) entity.CustomerRepository {
	return &customerRepo{
		Reader: reader,
		Writer: writer,
	}
}

func (r *customerRepo) GetAll() (res []entity.Customer, err error) {
	table := "customer"

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
		var r entity.Customer
		err = rows.Scan(
			&r.CustomerID,
			&r.FullName,
			&r.Email,
			&r.PhoneNumber,
			&r.Username,
			&r.Password,
			&r.AccountStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = append(res, r)
	}

	return
}

func (r *customerRepo) GetByID(CustomerID uint32) (res *entity.Customer, err error) {
	table := "customer"

	query := sq.Select("*").
		From(table).
		Where(sq.Eq{
			"customer_id": CustomerID,
		}).
		RunWith(r.Reader).
		PlaceholderFormat(sq.Question)

	rows, err := query.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var r entity.Customer
		err = rows.Scan(
			&r.CustomerID,
			&r.FullName,
			&r.Email,
			&r.PhoneNumber,
			&r.Username,
			&r.Password,
			&r.AccountStatus,
		)
		if err != nil {
			logger.Error("Selection Failed: " + err.Error())
		}
		res = &r
	}

	return
}

func (r *customerRepo) DeleteByID(ctx context.Context, CustomerID uint32) error {
	table := "customer"

	query := sq.Delete("").
		From(table).
		Where(sq.Eq{
			"customer_id": CustomerID,
		}).
		RunWith(r.Reader)
	_, err := query.ExecContext(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepo) Store(ctx context.Context, cust *entity.Customer) (custID uint32, err error) {
	table := "customer"

	query := sq.Insert(table).
		Columns(
			"customer_full_name",
			"customer_email",
			"customer_phone_number",
			"customer_username",
			"customer_password",
		).
		Values(
			cust.FullName,
			cust.Email,
			cust.PhoneNumber,
			cust.Username,
			cust.Password,
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
	custID = uint32(id)

	return
}