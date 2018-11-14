package logger

import (
	"errors"
	"fmt"
	"testing"
)

func TestFLogger(t *testing.T) {
	FLogger()
	SetLogLevel(DEBUG_LEVEL)
	Logger.Debug("debug")
	Logger.Info("info")
	Logger.Warn("warn")
	Logger.Error(errors.New("error"))
	defer func() {
		if err := recover(); err != nil {
			Logger.Fatal(err)
		}

	}()
	a := make([]int, 0)
	fmt.Println(a[1])
}
func TestConsoleLogger(t *testing.T) {
	CLogger()
	SetLogLevel(INFO_LEVEL)
	Logger.Debug("debug")
	Logger.Info("info")
	Logger.Warn("warn")
	Logger.Error(errors.New("error"))
	defer func() {
		if err := recover(); err != nil {
			Logger.Fatal(err)
		}

	}()
	a := make([]int, 0)
	fmt.Println(a[1])
}

func TestFCLogger(t *testing.T) {
	FCLogger()
	SetLogLevel(ERR_LEVEL)
	Logger.Debug("debug")
	Logger.Info("info")
	Logger.Warn("warn")
	Logger.Error(errors.New("error"))
	defer func() {
		if err := recover(); err != nil {
			Logger.Fatal(err)
		}

	}()
	a := make([]int, 0)
	fmt.Println(a[1])
}