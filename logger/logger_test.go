package logger_test

import (
	"errors"
	"github.com/wilsontamarozzi/panda-api/logger"
	"testing"
)

func TestFatal(t *testing.T) {
	err := errors.New("pq: erro postgresql")

	logger.Fatal(err)
}
