package swagger

import (
	"github.com/joho/godotenv"
	"os"
)

type env struct {
	DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT string
}

func (i *env) Init () error {

	if _, err := os.Stat("./.env"); err == nil {

		if (!os.IsNotExist(err)) {
			return  err
		}

		err := godotenv.Load()

		if err != nil {
			return err;
		}
	}

	i.DB_HOST = os.Getenv("DB_HOST");
	i.DB_NAME = os.Getenv("DB_NAME");
	i.DB_USER = os.Getenv("DB_USER");
	i.DB_PASS = os.Getenv("DB_PASS");
	i.DB_PORT = os.Getenv("DB_PORT");

	return nil;
}

var ENV env

