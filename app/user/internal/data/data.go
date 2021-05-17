package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go-micro-demo-project/app/user/internal/conf"
	"go-micro-demo-project/app/user/internal/data/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	cleanup := func() {
		logger.Log("msg", "closing the data resources")
	}
	return &Data{}, cleanup, nil
}
