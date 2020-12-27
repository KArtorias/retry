package retry

import (
	"errors"
	"log"
	"time"
)

var (
	retryTimesError   = errors.New("retryTimes must be greater than zero")
	retryTimeoutError = errors.New("retry timeout")
)

func DoFunc(retryTimes uint, retryTimeout, waitTime time.Duration, funcName string, f func() error) error {
	if retryTimes == 0 {
		log.Printf("ERROE: [%s] %v\n", funcName, retryTimesError)
		return retryTimesError
	}
	for i := uint(1); i <= retryTimes; i++ {
		if i > 1 && waitTime > 0 {
			time.Sleep(waitTime)
		}
		var err error
		ch1 := make(chan int)
		ch2 := make(chan int)

		if retryTimeout > 0 {
			go func() {
				time.Sleep(retryTimeout)
				ch1 <- 1
			}()
		}

		go func() {
			err = f()
			ch2 <- 1
		}()
		select {
		case <-ch1:
			log.Printf("ERROE: [%s] retry %d times error: %v\n", funcName, i, retryTimeoutError)
			if i == retryTimes {
				return retryTimeoutError
			}
			continue
		case <-ch2:
			if err != nil {
				log.Printf("ERROE: [%s] retry %d times error: %v\n", funcName, i, err)
				if i == retryTimes {
					return err
				}
				continue
			}
			return nil
		}
	}
	return nil
}
