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

		variables := getPathVariables(endpoint.Path)

		for _, variable := range variables {
			endpoint.Path = strings.ReplaceAll(endpoint.Path, "{"+variable+"}", ":"+variable)
		}

		ep := endpoint
		r.Handle(method, endpoint.Path, func(c *gin.Context) {
			body := replaceVariables(ep.Response.Body, c)
			c.JSON(ep.Response.Status, body)
		})
	}
}

func getPathVariables(path string) []string {
	var variables []string

	segments := strings.Split(path, "/")
	for _, segment := range segments {
		if strings.HasPrefix(segment, "{") && strings.HasSuffix(segment, "}") {
			varName := strings.TrimSuffix(strings.TrimPrefix(segment, "{"), "}")
			variables = append(variables, varName)
		}
	}

	return variables
}

func replaceVariables(body map[string]interface{}, c *gin.Context) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range body {
		if str, ok := value.(string); ok {
			for _, param := range c.Params {
				str = strings.ReplaceAll(str, "{"+param.Key+"}", param.Value)
			}
			result[key] = str
		} else {
			result[key] = value
		}
	}
	return result
}
