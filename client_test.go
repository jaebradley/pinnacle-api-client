package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchingApplicationDetails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != "https://pinnacle.com/config.app.json" {
			t.Fail()
		}

		rw.Write([]byte(`{
			"version": "6.6.65-rel",
			"commitId": "b8a4b4edd54199f23f7b89246f09bf0f18d19f98",
			"environment": "PRODUCTION",
			"originalDomain": "www.pinnacle.com",
			"cookieDomainSetExplicitly": true,
			"integration": {
				"logoutRedirect": false,
				"domains": {
					"forceSsl": true,
					"replaceLocal": true,
					"guest": "www.pinnacle.com",
					"classic": "www1.pinnacle.com",
					"beta": "beta.pinnacle.com",
					"members": "www1.pinnacle.com",
					"asian": "www1.pinnacle.com",
					"compact": "www1.pinnacle.com",
					"cashier": "cashier.pinnacle.com",
					"status": "status.pinnacle.com",
					"help": "help.future.pinnacle.com",
					"stats": "www.pinnacle-stats.com",
					"replaceKey": ".pinnacle.com"
				},
				"routes": {
					"asianView": "/Sportsbook/Asia",
					"betaView": "/",
					"classicView": "/members/canvas.asp",
					"compactView": "/Sportsbook/Asia",
					"guestSite": "/"
				},
				"redirectToPreferredView": true
			},
			"domainsByLicenses": [
				{
					"type": "malta",
					"domains": [
						".pinnacle.bet"
					]
				},
				{
					"type": "curacao",
					"domains": [
						".pinnacle.com",
						".pinnaclesports.com"
					]
				},
				{
					"type": "australia",
					"domains": [
						".pinnacleoz.com.au"
					]
				},
				{
					"type": "sweden",
					"domains": [
						".pinnacle.se"
					]
				}
			],
			"payments": {
				"replaceLocal": true,
				"baseUrl": "future",
				"netellerSuccess": "/account/deposit/Neteller/completed/",
				"moneyBookersSuccess": "/account/deposit/Skrill/success/",
				"moneyBookersFail": "/account/deposit/Skrill/fail/",
				"classicCashier": "/members/cashier.asp",
				"creditCard3DSecureSuccess": "/account/deposit/CreditCard/success/"
			},
			"api": {
				"haywire": {
					"root": "https://api.arcadia.pinnacle.com",
					"guestRoot": "https://guest.api.arcadia.pinnacle.com",
					"apiVersion": "0.1",
					"websockets": "wss://api.arcadia.pinnacle.com/ws",
					"apiKey": "CmX2KcMrXuFmNg6YFbmTxE0y9CIrOi0R"
				},
				"betPollInterval": 1000,
				"checkBalanceForMembers": true,
				"enableWebsockets": true,
				"ignoreInvalidMarketCounts": false,
				"websocketOptions": {
					"refreshTime": 2000,
					"refreshOnUpdate": false,
					"sportBlacklist": []
				},
				"features": {
					"sessionTimeLimits": true,
					"matchupLevelLeagueProps": true
				}
			},
			"features": {
				"newCasino": {
					"enabled": false,
					"apiKey": "TBD"
				},
				"themes": {
					"enabled": true,
					"guestMode": false
				},
				"liveCentre": {
					"guestMode": false,
					"showPeriodMatchups": false
				},
				"openGraph": {
					"useImage": true
				},
				"twitter": {
					"useCard": true,
					"useTitle": true
				},
				"accountRecoveryWithCaptcha": true,
				"loginWithCaptcha": true,
				"legacyCashier": {
					"useIframe": false
				},
				"betshare": {
					"enabled": true
				},
				"fixtureTranslations": {
					"enabled": true
				},
				"periodMarketFlags": {
					"enabled": false
				},
				"betslip": {
					"tabs": true,
					"unified": false
				},
				"anonGeoLocationRedirect": {
					"enabled": true
				},
				"nhlOverride": {
					"enabled": false
				},
				"showNewMarketFilter": {
					"enabled": true
				},
				"gamingPromptCumulativeProfit": {
					"enabled": false
				},
				"esportsGamesPage": {
					"enabled": true
				},
				"multiView": {
					"primaryOnly": false
				},
				"productAccess": {
					"enabled": false
				}
			},
			"polling": {
				"guest_multiplier": 3,
				"events": {
					"matchups": 5000,
					"markets": 3000
				},
				"matchup": 3000,
				"teasers": 7500,
				"betslip": {
					"quotes": {
						"singles": 1000,
						"multiples": 1000,
						"teasers": 10000
					},
					"event": 1500
				}
			},
			"registration": {
				"registeredBy": {
					"desktop": "ArcadiaDesktop",
					"mobile": "ArcadiaMobile"
				}
			},
			"vendors": {
				"navEgg": {
					"enabled": true
				},
				"hotjar": {
					"enabled": true,
					"key": 1083293
				},
				"gtm": {
					"enabled": true,
					"auth": "RSYvLtUxCh8o0eYkz0CGig",
					"env": "env-2"
				},
				"otherlevels": {
					"enabled": true,
					"key": "4d21c8ceef2c0fa35ebee38f497184cf"
				},
				"touchline": {
					"enabled": true,
					"features": {
						"statsBanner": false,
						"statsLink": true,
						"soccerTips": false,
						"predictions": true
					}
				},
				"tomorrowTTH": {
					"enabled": true
				},
				"sumsub": {
					"domain": "https://api.sumsub.com",
					"clientid": "pinnacle"
				},
				"sentry": {
					"enabled": true,
					"dsn": "https://80eedce15edb4644a5bac761265e091c@o417691.ingest.sentry.io/5339569",
					"debug": false,
					"environment": "production",
					"sampleRate": 0.25
				}
			}
		}
		`))
	}))

	defer server.Close()

	client := Client{Client: server.Client()}
	body, err := fetchApplicationDetails(&client)

	if nil != err {
		t.Fail()
	}

	expected := ApplicationDetails{
		ApiDetails: ApiDetails{
			HaywireDetails: HaywireDetails{
				ApiKey: "CmX2KcMrXuFmNg6YFbmTxE0y9CIrOi0R",
			},
		},
	}

	if expected != body {
		t.Fail()
	}
}

func TestFetchingStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != "https://guest.api.arcadia.pinnacle.com/0.1/status" {
			t.Fail()
		}

		rw.Write([]byte(`{
			"code": "ONLINE",
			"description": "System is operating correctly, no known issues.",
			"services": [
				{
					"name": "api",
					"status": "ONLINE"
				},
				{
					"name": "etl",
					"status": "ONLINE"
				},
				{
					"name": "search",
					"status": "ONLINE"
				},
				{
					"name": "websocket",
					"status": "OFFLINE"
				}
			],
			"upstream": [
				{
					"health": "ONLINE",
					"name": "account",
					"status": ""
				},
				{
					"health": "ONLINE",
					"name": "betting",
					"status": "ENABLED"
				},
				{
					"health": "ONLINE",
					"name": "cashier",
					"status": ""
				},
				{
					"health": "ONLINE",
					"name": "casino",
					"status": ""
				},
				{
					"health": "ONLINE",
					"name": "login",
					"status": ""
				},
				{
					"health": "ONLINE",
					"name": "lookup",
					"status": ""
				},
				{
					"health": "ONLINE",
					"name": "notification",
					"status": ""
				},
				{
					"health": "ONLINE",
					"name": "profile",
					"status": ""
				},
				{
					"health": "ONLINE",
					"name": "registration",
					"status": ""
				},
				{
					"health": "ONLINE",
					"name": "responsible_gaming",
					"status": ""
				},
				{
					"health": "ONLINE",
					"name": "transaction",
					"status": ""
				}
			]
		}
		`))
	}))

	defer server.Close()

	client := Client{Client: server.Client()}
	body, err := FetchStatus(&client)

	if nil != err {
		t.Fail()
	}

	if body.Code != "ONLINE" {
		t.Log("fucking code:", body.Code)
		t.Fail()
	}
}

func TestFetchingLeagueStraightOdds(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.String() != "https://guest.api.arcadia.pinnacle.com/0.1/leagues/487/markets/straight" {
			t.Fail()
		}

		rw.Write([]byte(`[
			{
				"cutoffAt": "2021-04-04T01:10:00+00:00",
				"key": "s;0;ou",
				"limits": [
					{
						"amount": 600,
						"type": "maxRiskStake"
					}
				],
				"matchupId": 1297204758,
				"period": 0,
				"prices": [
					{
						"participantId": 1297204759,
						"points": 2.5,
						"price": 144
					},
					{
						"participantId": 1297204760,
						"points": 2.5,
						"price": -199
					}
				],
				"type": "total",
				"version": 1077148336
			},
			{
				"cutoffAt": "2021-04-04T00:10:00+00:00",
				"key": "s;0;ou",
				"limits": [
					{
						"amount": 600,
						"type": "maxRiskStake"
					}
				],
				"matchupId": 1297800427,
				"period": 0,
				"prices": [
					{
						"participantId": 1297800428,
						"points": 10.5,
						"price": 112
					},
					{
						"participantId": 1297800429,
						"points": 10.5,
						"price": -150
					}
				],
				"type": "total",
				"version": 1077057688
			}
		]`))
	}))

	defer server.Close()

	client := Client{Client: server.Client()}
	body, err := FetchLeagueStraightOdds(&client, 487)

	if nil != err {
		t.Fail()
	}

	if len(body) < 1 {
		t.Fail()
	}
}
