package config_test

import (
	"testing"

	"github.com/just-hms/large-scale-multistructure-db/be/config"
	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
}

func TestCondigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (s *ConfigSuite) TestConfig() {
	cfg, err := config.NewConfig()
	s.Require().NoError(err)
	s.Require().NotEqual("", cfg.Mongo.Host)
	s.Require().NotEqual("", cfg.Redis.Host)
}
