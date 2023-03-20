package config

import "productAccounting-v1/internal/common"

type Config struct {
	Server common.ServerConfig
	Neo4j  common.Neo4jConfig
}
