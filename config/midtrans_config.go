package config

import (
	"github.com/veritrans/go-midtrans"
)

var (
	ServerKey = "SB-Mid-server-Rak2zKFmOgBuxALhKyDLFMiD"
	ClientKey = "SB-Mid-client-4EmZKNF-SvNAWczT"
)

func InitMidtrans() midtrans.Client {
	m := midtrans.NewClient()
	m.ServerKey = ServerKey
	m.ClientKey = ClientKey
	m.APIEnvType = midtrans.Sandbox
	return m
}
