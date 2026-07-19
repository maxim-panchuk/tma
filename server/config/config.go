package config

import (
	"log"

	"github.com/caarlos0/env/v6"
)

var Config = struct {
	Port  int `env:"PORT" envDefault:"443"`
	Proof struct {
		PayloadSignatureKey string `env:"TONPROOF_PAYLOAD_SIGNATURE_KEY,required"`
		PayloadLifeTimeSec  int64  `env:"TONPROOF_PAYLOAD_LIFETIME_SEC" envDefault:"300"`
		ProofLifeTimeSec    int64  `env:"TONPROOF_PROOF_LIFETIME_SEC" envDefault:"300"`
	}
	Market struct {
		BankAddr   string `env:"BANK_ADDR,required"`
		WalletSeed string `env:"WALLET_SEED,required"`
	}
	AdminSecretKey string `env:"ADMIN_SECRET_KEY,required"`
	DatabaseURL    string `env:"DATABASE_URL" envDefault:"postgresql://postgres:password@localhost:5432/tma"`
}{}

func LoadConfig() {
	if err := env.Parse(&Config); err != nil {
		log.Fatalf("config parsing failed: %v\n", err)
	}
}
