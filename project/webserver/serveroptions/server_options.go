package serveroptions

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/wonderivan/logger"
)

// ServerOptions is struct for defining a server preset of work settings
type ServerOptions struct {
	IP           string
	Port         string
	ReadTimeout  int
	WriteTimeout int
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

	o.IP = keyValue[0][1]

	o.Port = keyValue[1][1]

	o.ReadTimeout, err = strconv.Atoi(keyValue[2][1])
	if err != nil {
		return ServerOptions{}, errors.New("conversion error for the parameter - " + keyValue[1][0])
	}

	o.WriteTimeout, err = strconv.Atoi(keyValue[3][1])
	if err != nil {
		return ServerOptions{}, errors.New("conversion error for the parameter - " + keyValue[2][0])
	}

	err = o.checkServerOptions()
	if err != nil {
		return ServerOptions{}, err
	}

	return o, nil
}

func (o *ServerOptions) checkServerOptions() error {
	if o.IP == "" {
		return errors.New("IP not found")
	}

	if o.Port == "" {
		return errors.New("port not found")
	}

	if o.ReadTimeout <= MinTimeout || o.ReadTimeout > MaxTimeout {
		return errors.New("read timeout cannot be 0")
	}

	if o.WriteTimeout <= MinTimeout || o.ReadTimeout > MaxTimeout {
		return errors.New("write timeout cannot be 0")
	}

	return nil
}