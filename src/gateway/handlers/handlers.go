package handlers

import (
)

type ServicesStruct struct {
}

type GatewayService struct {
	Config     ServicesStruct
	// authTokens map[string]time.Time
}

func NewGatewayService(config *ServicesStruct) *GatewayService {
  return new(GatewayService)
}

