package main

import (
	"context"
	"fmt"
	"github.com/TON-Market/tma/server/config"
	"github.com/TON-Market/tma/server/datatype/market"
	"github.com/TON-Market/tma/server/datatype/token"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/tonkeeper/tongo/liteapi"
	"github.com/tonkeeper/tongo/tonconnect"
	_ "net/http/pprof"
)

func main() {
	log.Info("Tonproof is running")
	config.LoadConfig()

	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		Skipper:           nil,
		DisableStackAll:   true,
		DisablePrintStack: false,
	}))
	e.Use(middleware.Logger())
	e.Static("/", "./static")

	mainNetClient, err := liteapi.NewClientWithDefaultMainnet()
	if err != nil {
		log.Fatalf("failed init mainnet liteapi client")
	}
	networks["-239"] = mainNetClient

	payloadLifeTime := config.Config.Proof.PayloadLifeTimeSec
	proofLifeTime := config.Config.Proof.ProofLifeTimeSec
	tonConnectMainNet, err := tonconnect.NewTonConnect(mainNetClient, config.Config.Proof.PayloadSignatureKey,
		tonconnect.WithLifeTimePayload(payloadLifeTime), tonconnect.WithLifeTimeProof(proofLifeTime))

	h := newHandler(tonConnectMainNet)
	w := newSocket()

	market.GetMarket().Start(context.TODO())
	registerHandlers(e, h, w)

	testData()

	log.Fatal(e.StartTLS(fmt.Sprintf(":%v", config.Config.Port), "fullchain.pem", "privkey.pem"))
}

func testData() {
	e := &market.Event{
		Tag:      market.Crypto,
		LogoLink: "/img/bootcamp.jpg",
		Title:    "Will TonMarket be in top-3 of Moscow Bootcamp?",
		BetMap: map[token.Token]*market.Bet{
			token.A: {
				Token:    token.A,
				Title:    "Yes",
				LogoLink: "/img/Ton_Market_YES.png",
			},
			token.B: {
				Token:    token.B,
				Title:    "No",
				LogoLink: "/img/Ton_Market_No.png",
			},
		},
	}
	elections := &market.Event{
		Tag:      market.Politic,
		LogoLink: "/img/Elections.png",
		Title:    "USA Elections 2024",
		BetMap: map[token.Token]*market.Bet{
			token.A: {
				Token:    token.A,
				Title:    "Trump",
				LogoLink: "/img/trump.jpg",
			},
			token.B: {
				Token:    token.B,
				Title:    "Harris",
				LogoLink: "/img/harris.jpg",
			},
		},
	}
	err := market.GetMarket().AddEvent(context.Background(), e)
	if err != nil {
		log.Fatalln(err)
	}
	err = market.GetMarket().AddEvent(context.Background(), elections)
	if err != nil {
		log.Fatalln(err)
	}
}
