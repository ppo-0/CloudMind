package svc

import (
	"CloudMind/app/es/cmd/rpc/internal/config"
	"github.com/elastic/go-elasticsearch/v7"
)

type ServiceContext struct {
	Config config.Config
	Es     *elasticsearch.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	es, _ := elasticsearch.NewDefaultClient()
	return &ServiceContext{
		Config: c,
		Es:     es,
	}
}
