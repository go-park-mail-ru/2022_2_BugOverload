package server_options

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ServerOptions struct {
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

func GetServerOptions(pathConfig string) (ServerOptions, error) {
	var o ServerOptions

	bytes, err := ioutil.ReadFile(pathConfig)
	if err != nil {
		return ServerOptions{}, err
	}

	textConfig := string(bytes[:])

	settings := strings.Split(textConfig, "\n")

	var keyValue [][]string

	for _, val := range settings {
		fmt.Println("val:", val)
		split := strings.Split(val, ": ")
		keyValue = append(keyValue, split)
	}

	o.Addr = keyValue[0][1]

	o.ReadTimeout, err = strconv.Atoi(keyValue[1][1])
	if err != nil {
		return ServerOptions{}, errors.New("conversion error for the parameter - " + keyValue[1][0])
	}

	o.WriteTimeout, err = strconv.Atoi(keyValue[2][1])
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
	if o.Addr == "" {
		return errors.New("port not found")
	}

	if o.ReadTimeout <= 0 || o.ReadTimeout > 15 {
		return errors.New("read timeout cannot be 0")
	}

	if o.WriteTimeout <= 0 || o.ReadTimeout > 15 {
		return errors.New("write timeout cannot be 0")
	}

	return nil
}
