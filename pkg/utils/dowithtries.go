package utils

import (
	"fmt"
	"time"
)

func DoWithTries(fn func() error, attemps int64, delay time.Duration) (err error) {
	for i := attemps; i > 0; i-- {
		if err = fn(); err != nil {
			time.Sleep(delay)
			continue
		}

		return nil
	}

	return fmt.Errorf("utils: doWithTries: error after %d attemps", attemps)
}
