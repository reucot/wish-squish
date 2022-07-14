package config

import (
	"fmt"
	"path/filepath"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
	"github.com/gookit/ini/v2/dotenv"
)

var conf Config

func Load() error {
	var filename string
	var err error

	config.WithOptions(config.ParseEnv, func(o *config.Options) { o.DecoderConfig.TagName = "json" })

	config.AddDriver(json.Driver)

	if filename, err = filepath.Abs(filepath.Join("./", ".env")); err != nil {
		fmt.Println(filename)
		return err
	}

	if err = dotenv.LoadFiles(filename); err != nil {
		return err
	}

	if filename, err = filepath.Abs(filepath.Join("./", "config.json")); err != nil {
		return err
	}

	if err = config.LoadFiles(filename); err != nil {
		return err
	}

	if err = config.BindStruct("", &conf); err != nil {
		return err
	}

	return nil
}

func Get() Config {
	return conf
}
