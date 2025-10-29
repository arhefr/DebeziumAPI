package utils

import (
	"fmt"
	"time"
)

func DoWithTries(fn func() error, attemps int, delay time.Duration) error {
	for i := 0; i < attemps; i++ {
		if err := fn(); err != nil {
			time.Sleep(delay)
			continue
		}

		return nil
	}

	return fmt.Errorf("utils: doWithTries: error after %d attemps", attemps)
}
