package configuration

import (
	"golang-todo-app/exception"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(ket string) string
}

type confogImpl struct {
}

func (config confogImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicLogging(err)
	return &confogImpl{}
}
