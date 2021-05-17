package config

import (
	"database/sql"
	"fmt"
	"strconv"

	// _ mysql driver
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type mySQL struct {
	Database       string
	ReaderHost     string
	ReaderPort     string
	ReaderUser     string
	ReaderPassword string
	WriterHost     string
	WriterPort     string
	WriterUser     string
	WriterPassword string
}

func setupMySQL() *mySQL {
	v := &mySQL{}

	v.setupGeneral()
	v.setupReader()
	v.setupWriter()

	return v
}

func (v *mySQL) setupGeneral() {
	if v.Database = os.Getenv("MYSQL_DATABASE_NAME"); v.Database == "" {
		v.Database = "shop"
		fmt.Printf("using default %s as MySQL's database name\n", v.Database)
	}
}

func (v *mySQL) setupReader() {
	if v.ReaderHost = os.Getenv("READER_HOST"); v.ReaderHost == "" {
		v.ReaderHost = "127.0.0.1"
		fmt.Printf("using default %s as MySQL's reader host\n", v.ReaderHost)
	}

	if v.ReaderPort = os.Getenv("READER_PORT"); v.ReaderPort == "" {
		v.ReaderPort = "3306"
		fmt.Printf("using default %s as MySQL's reader port\n", v.ReaderPort)
	}

	if v.ReaderUser = os.Getenv("READER_USER"); v.ReaderUser == "" {
		v.ReaderUser = "sammidev"
		fmt.Printf("using default %s as MySQL's reader user\n", v.ReaderUser)
	}

	if v.ReaderPassword = os.Getenv("READER_PASSWORD"); v.ReaderPassword == "" {
		v.ReaderPassword = "sammidev"
		fmt.Printf("using default MySQL's reader password\n")
	}
}

func (v *mySQL) setupWriter() {
	if v.WriterHost = os.Getenv("WRITER_HOST"); v.WriterHost == "" {
		v.WriterHost = "127.0.0.1"
		fmt.Printf("using default %s as MySQL's writer host\n", v.WriterHost)
	}

	if v.WriterPort = os.Getenv("WRITER_PORT"); v.WriterPort == "" {
		v.WriterPort = "3306"
		fmt.Printf("using default %s as MySQL's writer port\n", v.WriterPort)
	}

	if v.WriterUser = os.Getenv("WRITER_USER"); v.WriterUser == "" {
		v.WriterUser = "sammidev"
		fmt.Printf("using default %s as MySQL's writer user\n", v.WriterUser)
	}

	if v.WriterPassword = os.Getenv("WRITER_PASSWORD"); v.WriterPassword == "" {
		v.WriterPassword = "sammidev"
		fmt.Printf("using default MySQL's writer password\n")
	}
}


// Option can be used to configure mysql connection
type Option struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

// SetupDatabase will prepare mysql connection
func SetupDatabase(readerConfig, writerConfig Option) (reader, writer *sql.DB, err error) {
	reader, err = createConnection(readerConfig)
	if err != nil {
		return nil, nil, err
	}

	writer, err = createConnection(writerConfig)
	if err != nil {
		return nil, nil, err
	}

	return reader, writer, nil
}

func createConnection(config Option) (db *sql.DB, err error) {
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}

	if config.Port == "" {
		config.Port = "3306"
	}

	auth := config.User + ":" + config.Password
	uri := "tcp(" + config.Host + ":" + config.Port + ")"
	dsn := auth + "@" + uri + "/" + config.Database + "?parseTime=true"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	maxConn, err := strconv.Atoi("DB_MAX_CONNECTIONS")
	db.SetMaxOpenConns(maxConn)

	err = db.Ping()

	return
}
