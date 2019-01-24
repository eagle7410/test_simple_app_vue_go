package swagger

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path"
)

type env struct {
	DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT string
}

func (i *env) Init () error {

	pwd, err := os.Getwd()

	fmt.Println(pwd)

	if err != nil {
		return err
	}

	envPath := path.Join(pwd, ".env")

	if _, err := os.Stat(envPath); err == nil {

		err := godotenv.Load(envPath)

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

