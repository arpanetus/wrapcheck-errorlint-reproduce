package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrSomeStuff = errors.New("i `has` occurred")

func hasToReturnErr() error {
	return ErrSomeStuff //nolint:wrapcheck // i have no idea why wrapchek whines about it
}

func desperateWrapWithWrapcheck() error {
	if err := hasToReturnErr(); err != nil {
		return fmt.Errorf("wrapcheck, do it again: %w", err)
	}

	return nil
}

func desperateWrapWithErrorlint() error {
	if err := hasToReturnErr(); err != nil {
		return fmt.Errorf("errorlint, do it again: %v", err)
	}

	return nil
}

func main() {
	log.Println(desperateWrapWithErrorlint())
	log.Println(desperateWrapWithWrapcheck())
}
