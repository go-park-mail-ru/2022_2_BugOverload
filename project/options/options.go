package options

import (
	"errors"
	"flag"
)

type Options struct {
	PathServerConfig string
}

func GetOptions() (Options, error) {
	var o Options

	flag.Parse()

	o.PathServerConfig = flag.Arg(0)

	resCheck := o.checkOptions()
	if resCheck != nil {
		return Options{}, resCheck
	}

	return o, nil
}

func (o *Options) checkOptions() error {
	if o.PathServerConfig == "" {
		return errors.New("the path to the work configuration is not specified")
	}

	return nil
}
