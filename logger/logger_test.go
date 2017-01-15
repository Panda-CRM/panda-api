package logger_test

import(
	"testing"
	"errors"
	"github.com/wilsontamarozzi/panda-api/logger"
)

func TestFatal(t *testing.T) {
	err := errors.New("pq: erro postgresql")

	logger.Fatal(err)
}