package swagger

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB // global variable to share it between main and the HTTP handler
)


func DatabaseInit () error {
	var (
		err error
		connectionString string = ENV.DB_USER +":"+ ENV.DB_PASS +"@tcp("+ ENV.DB_HOST +":"+ ENV.DB_PORT +")/"+ ENV.DB_NAME
	)

	// More info https://github.com/go-sql-driver/mysql
	DB, err = sql.Open("mysql", connectionString)

	if err != nil {
	return err;
	}

	DB.SetMaxIdleConns(100)

	return  DB.Ping();
}
