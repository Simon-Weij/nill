package router

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

type Config struct {
	Endpoints []Endpoint `mapstructure:"endpoints"`
}
type Endpoint struct {
	Path     string   `mapstructure:"path"`
	Method   string   `mapstructure:"method"`
	Response Response `mapstructure:"response"`
}

type Response struct {
	Status int                    `mapstructure:"status"`
	Body   map[string]interface{} `mapstructure:"body"`
}

func DefineRoutes(cfg *Config) {
	r := gin.Default()
	registerEndpoints(r, cfg)

	port := os.Getenv("PORT")
	if port == "" {
		r.Run()
	} else {
		r.Run(":" + port)
	}
}

func ParseRoutes(Configpath string) (*Config, error) {
	viper.SetConfigFile(Configpath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error reading config file: %w", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("unable to decode config into struct: %w", err)
	}

	return &config, nil
}

func registerEndpoints(r *gin.Engine, cfg *Config) {
	if cfg == nil {
		return
	}

	for _, endpoint := range cfg.Endpoints {
		method := strings.ToUpper(endpoint.Method)
		if method == "" {
			method = "GET"
		}

		r.Handle(method, endpoint.Path, func(c *gin.Context) {
			c.JSON(endpoint.Response.Status, endpoint.Response.Body)
		})
	}
}
