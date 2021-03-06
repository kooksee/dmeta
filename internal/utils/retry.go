package utils

import (
	"github.com/kooksee/dmeta/internal/errs"
	"io"
	"strings"
	"time"
)

func fibonacci() func() int {
	a1, a2 := 0, 1
	return func() int {
		a1, a2 = a2, a1+a2
		return a1
	}
}

func Retry(num int, fn func() error) (err error) {
	t := fibonacci()
	for i := 0; ; i++ {

		if err = fn(); err == nil || err == errs.NotFound || err == io.EOF {
			return err
		}

		sleepTime := t()

		if strings.Contains(err.Error(), "timeout") {
			time.Sleep(time.Second * time.Duration(sleepTime))
			continue
		}

		if i > num {
			return err
		}

		time.Sleep(time.Second * time.Duration(sleepTime))
	}
}
