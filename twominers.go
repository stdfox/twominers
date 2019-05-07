package twominers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	BTG        = "https://btg.2miners.com/api"
	SOLO_BTG   = "https://solo-btg.2miners.com/api"
	ZEC        = "https://zec.2miners.com/api"
	SOLO_ZEC   = "https://solo-zec.2miners.com/api"
	ZCL        = "https://zcl.2miners.com/api"
	SOLO_ZCL   = "https://solo-zcl.2miners.com/api"
	ZEN        = "https://zen.2miners.com/api"
	SOLO_ZEN   = "https://solo-zen.2miners.com/api"
	ETH        = "https://eth.2miners.com/api"
	SOLO_ETH   = "https://solo-eth.2miners.com/api"
	ETC        = "https://etc.2miners.com/api"
	SOLO_ETC   = "https://solo-etc.2miners.com/api"
	EXP        = "https://exp.2miners.com/api"
	SOLO_EXP   = "https://solo-exp.2miners.com/api"
	MUSIC      = "https://music.2miners.com/api"
	SOLO_MUSIC = "https://solo-music.2miners.com/api"
	ETP        = "https://etp.2miners.com/api"
	SOLO_ETP   = "https://solo-etp.2miners.com/api"
	PIRL       = "https://pirl.2miners.com/api"
	SOLO_PIRL  = "https://solo-pirl.2miners.com/api"
	SOLO_WHL   = "https://solo-whl.2miners.com/api"
	SOLO_DBIX  = "https://solo-dbix.2miners.com/api"
	BTCZ       = "https://btcz.2miners.com/api"
	SOLO_BTCZ  = "https://solo-btcz.2miners.com/api"
	CLO        = "https://clo.2miners.com/api"
	SOLO_CLO   = "https://solo-clo.2miners.com/api"
	AKA        = "https://aka.2miners.com/api"
	SOLO_AKA   = "https://solo-aka.2miners.com/api"
	MOAC       = "https://moac.2miners.com/api"
	SOLO_MOAC  = "https://solo-moac.2miners.com/api"
	XMR        = "https://xmr.2miners.com/api"
	SOLO_XMR   = "https://solo-xmr.2miners.com/api"
	XZC        = "https://xzc.2miners.com/api"
	SOLO_XZC   = "https://solo-xzc.2miners.com/api"
	ZEL        = "https://zel.2miners.com/api"
	SOLO_ZEL   = "https://solo-zel.2miners.com/api"
	GRIN       = "https://grin.2miners.com/api"
	SOLO_GRIN  = "https://solo-grin.2miners.com/api"
)

type Twominers struct {
	httpClient *http.Client
}

type Payment struct {
	Amount    int64  `json:"amount"`
	Timestamp string `json:"timestamp"`
	TX        string `json:"tx"`
}

type Reward struct {
	Blockheight int64   `json:"blockheight"`
	Timestamp   string  `json:"timestamp"`
	Blockhash   string  `json:"blockhash"`
	Reward      int64   `json:"reward"`
	Percent     float64 `json:"percent"`
	Immature    bool    `json:"immature"`
}

type Stats struct {
	Balance     int32 `json:"balance"`
	BlocksFound int32 `json:"blocksFound"`
	Immature    int32 `json:"immature"`
	LastShare   int32 `json:"lastShare"`
	Paid        int64 `json:"paid"`
	Pending     bool  `json:"pending"`
}

type Sumrewards struct {
	Inverval  int32  `json:"inverval"`
	Reward    int32  `json:"reward"`
	Numreward int32  `json:"numreward"`
	Name      string `json:"name"`
	Offset    int32  `json:"offset"`
}

type WorkerGroup struct {
	LastBeat string  `json:"lastBeat"`
	HR       float64 `json:"hr"`
	Offline  bool    `json:"offline"`
	HR2      float64 `json:"hr2"`
}

type Worker struct {
	WorkerGroup WorkerGroup `json:"workerGroup"`
}

type Account struct {
	CurrentHashrate int32        `json:"currentHashrate"`
	CurrentLuck     string       `json:"currentLuck"`
	Hashrate        int64        `json:"hashrate"`
	PageSize        int32        `json:"pageSize"`
	Payments        []Payment    `json:"payments"`
	PaymentsTotal   int32        `json:"paymentsTotal"`
	Rewards         []Reward     `json:"rewards"`
	RoundShares     int32        `json:"roundShares"`
	Shares          []string     `json:"shares"`
	Stats           Stats        `json:"stats"`
	SumRewards      []Sumrewards `json:"sumRewards"`
	Workers         Worker       `json:"workers"`
	WorkersOffline  int32        `json:"workersOffline"`
	WorkersOnline   int32        `json:"workersOnline"`
	WorkersTotal    int32        `json:"workersTotal"`
	Reward24h       int64        `json:"24hreward"`
	NumReward24h    int64        `json:"24hnumReward"`
}

func New() *Twominers {
	client := &Twominers{
		httpClient: &http.Client{},
	}

	return client
}

func (t *Twominers) GetAccount(server string, walletId string) (*Account, error) {
	req, err := http.NewRequest(http.MethodGet, server+"/accounts/"+walletId, nil)
	if err != nil {
		return nil, err
	}

	res, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	account := &Account{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}

	return account, err
}
