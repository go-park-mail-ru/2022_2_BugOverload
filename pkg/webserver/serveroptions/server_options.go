package serveroptions

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/wonderivan/logger"
)

// ServerOptions is struct for defining a server preset of work settings
type ServerOptions struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// GetServerOptions is function for getting startup parameters from global options
func GetServerOptions(pathConfig string) (ServerOptions, error) {
	var o ServerOptions

	stream, err := os.Open(pathConfig)
	if err != nil {
		return ServerOptions{}, err
	}
	defer func() {
		err = stream.Close()
		if err != nil {
			logger.Error(err)
		}
	}()

	bytes, err := io.ReadAll(stream)
	if err != nil {
		return ServerOptions{}, err
	}

	textConfig := string(bytes)

	settings := strings.Split(textConfig, "\n")

	var keyValue [][]string

	for _, val := range settings {
		split := strings.Split(val, ": ")
		keyValue = append(keyValue, split)
	}

	number, err := strconv.Atoi(keyValue[0][1])
	if err != nil {
		return ServerOptions{}, errors.New("conversion error for the parameter - " + keyValue[0][0])
	}
	o.ReadTimeout = time.Duration(number) * time.Second

	number, err = strconv.Atoi(keyValue[1][1])
	if err != nil {
		return ServerOptions{}, errors.New("conversion error for the parameter - " + keyValue[1][0])
	}
	o.WriteTimeout = time.Duration(number) * time.Second

	err = o.checkServerOptions()
	if err != nil {
		return ServerOptions{}, err
	}

	return o, nil
}

func (o *ServerOptions) checkServerOptions() error {
	if o.ReadTimeout <= MinTimeout || o.ReadTimeout > MaxTimeout {
		return errors.New("read timeout cannot be 0")
	}

	if o.WriteTimeout <= MinTimeout || o.ReadTimeout > MaxTimeout {
		return errors.New("write timeout cannot be 0")
	}

	return nil
}
