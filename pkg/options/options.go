package options

import (
	"errors"
	"flag"
)

// Options is struct for defining a global preset of work settings
type Options struct {
	Port             string
	PathServerConfig string
}

// GetOptions is function for getting startup parameters from console arguments
func GetOptions() (Options, error) {
	var o Options

	flag.Parse()

	o.Port = flag.Arg(0)
	o.PathServerConfig = flag.Arg(1)

	resCheck := o.checkOptions()
	if resCheck != nil {
		return Options{}, resCheck
	}

	return o, nil
}

// checkOptions is method for validation parameters
func (o *Options) checkOptions() error {
	if o.PathServerConfig == "" {
		return errors.New("the path to the work configuration is not specified")
	}

	if o.Port == "" {
		return errors.New("port is not specified")
	}

	return nil
}
