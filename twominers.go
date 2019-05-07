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

type PaymentReturn struct {
	Payments      []Payments `json:"payments"`
	PaymentsTotal int32      `json:"paymentsTotal"`
}

type StatsReturn struct {
	CandidatesTotal int32        `json:"candidatesTotal"`
	Hashrate        float32      `json:"hashrate"`
	ImmatureTotal   int32        `json:"immatureTotal"`
	Luck            float32      `json:"luck"`
	MaturedTotal    int32        `json:"maturedTotal"`
	MinersTotal     int32        `json:"minersTotal"`
	Nodes           []Node       `json:"nodes"`
	Now             int64        `json:"now"`
	PaymentsTotal   int32        `json:"paymentsTotal"`
	PoolCharts      []PoolCharts `json:"poolCharts"`
	Stats           Starts       `json:"stats"`
}

type Sumrewards struct {
	Inverval  int32  `json:"inverval"`
	Reward    int64  `json:"reward"`
	Numreward int32  `json:"numreward"`
	Name      string `json:"name"`
	Offset    int32  `json:"offset"`
}

type WorkerGroup struct {
	LastBeat string  `json:"lastBeat"`
	Hr       float32 `json:"hr"`
	Offline  bool    `json:"offline"`
	Hr2      float32 `json:"hr2"`
}

type AccountReturn struct {
	CurrentHashrate int32        `json:"currentHashrate"`
	CurrentLuck     string       `json:"currentLuck"`
	Hashrate        int64        `json:"hashrate"`
	PageSize        int32        `json:"pageSize"`
	Payments        []Payment    `json:"payments"`
	PaymentsTotal   int32        `json:"paymentsTotal"`
	Rewards         []Rewards    `json:"rewards"`
	RoundShares     int32        `json:"roundShares"`
	Shares          []string     `json:"shares"`
	Stats           Stats        `json:"stats"`
	Sumrewards      []Sumrewards `json:"sumrewards"`
	Workers         Worker       `json:"workers"`
	WorkersOffline  int32        `json:"workersOffline"`
	WorkersOnline   int32        `json:"workersOnline"`
	WorkersTotal    int32        `json:"workersTotal"`
	Reward24h       int64        `json:"24hreward"`
	Numreward24h    int64        `json:"24hnumreward"`
}

type Payment struct {
	Amount    int64  `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	TX        string `json:"tx"`
}

type Stats struct {
	Balance     int32 `json:"balance"`
	BlocksFound int32 `json:"blocksFound"`
	Immature    int32 `json:"immature"`
	LastShare   int32 `json:"lastShare"`
	Paid        int64 `json:"paid"`
	Pending     byte  `json:"pending"`
}

type MinerReturn struct {
	Hashrate    float32 `json:"hashrate"`
	Miners      Miner   `json:"miners"`
	MinersTotal int32   `json:"minersTotal"`
	Now         int64   `json:"now"`
}

type Miner struct {
	MinerUid MinerUID `json:"minerUid"`
}

type Node struct {
	Difficulty    string  `json:"difficulty"`
	Height        string  `json:"height"`
	LastBeat      string  `json:"lastBeat"`
	Name          string  `json:"name"`
	Networkhashps float32 `json:"networkhashps"`
}

type Luck struct {
	LuckNumber LuckNumber `json:"luckNumber"`
}

type Rewards struct {
	Blockheight int64   `json:"blockheight"`
	Timestamp   int64   `json:"timestamp"`
	Blockhash   string  `json:"blockhash"`
	Reward      int64   `json:"reward"`
	Percent     float32 `json:"percent"`
	Immature    bool    `json:"immature"`
}

type Candidates struct {
	Height      int64   `json:"height"`
	Timestamp   int64   `json:"timestamp"`
	Difficulty  float32 `json:"difficulty"`
	Shares      float32 `json:"shares"`
	Finder      string  `json:"finder"`
	Uncle       bool    `json:"uncle"`
	UncleHeight int64   `json:"uncleHeight"`
	Orphan      bool    `json:"orphan"`
	Hash        string  `json:"hash"`
	Reward      int64   `json:"reward"`
}

type ImMatured struct {
	Height      int64   `json:"height"`
	Timestamp   int64   `json:"timestamp"`
	Difficulty  float32 `json:"difficulty"`
	Shares      float32 `json:"shares"`
	Finder      string  `json:"finder"`
	Uncle       bool    `json:"uncle"`
	UncleHeight int64   `json:"uncleHeight"`
	Orphan      bool    `json:"orphan"`
	Hash        string  `json:"hash"`
	Reward      int64   `json:"reward"`
}

type Matured struct {
	Height      int64   `json:"height"`
	Timestamp   int64   `json:"timestamp"`
	Difficulty  float32 `json:"difficulty"`
	Shares      float32 `json:"shares"`
	Finder      string  `json:"finder"`
	Uncle       bool    `json:"uncle"`
	UncleHeight int64   `json:"uncleHeight"`
	Orphan      bool    `json:"orphan"`
	Hash        string  `json:"hash"`
	Reward      int64   `json:"reward"`
}

type BlockReturn struct {
	Candidates      []Candidates `json:"candidates"`
	CandidatesTotal int32        `json:"candidatesTotal"`
	Immature        []ImMatured  `json:"immature"`
	ImmatureTotal   int32        `json:"immatureTotal"`
	Luck            Luck         `json:"luck"`
	Matured         []Matured    `json:"matured"`
	MaturedTotal    int32        `json:"maturedTotal"`
}

type Starts struct {
	LastBlockFound int64   `json:"lastBlockFound"`
	RoundShares    float32 `json:"roundShares"`
	Nshares        int64   `json:"nshares"`
}

type PoolCharts struct {
	X          int64   `json:"x"`
	TimeFormat string  `json:"timeFormat"`
	Y          float32 `json:"y"`
	Netdiff    float32 `json:"netdiff"`
	Nethr      float32 `json:"nethr"`
}

type MinerUID struct {
	LastBeat int64 `json:"lastBeat"`
	Height   int64 `json:"height"`
	Offline  bool  `json:"offline"`
}

type Payments struct {
	Amount      int64  `json:"amount"`
	Timestamp   int64  `json:"timestamp"`
	TotalPayees int32  `json:"totalPayees"`
	TX          string `json:"tx"`
}

type Worker struct {
	WorkerGroup WorkerGroup `json:"workerGroup"`
}

type LuckNumber struct {
	Luck       float32 `json:"luck"`
	OrphanRate float64 `json:"orphanRate"`
}

func New() *Twominers {
	client := &Twominers{
		httpClient: &http.Client{},
	}

	return client
}

func (t *Twominers) do(method string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
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

	return body, err
}

func (t *Twominers) GetAccounts(server string, walletId string) (*AccountReturn, error) {
	res, err := t.do(http.MethodGet, server+"/accounts/"+walletId)
	if err != nil {
		return nil, err
	}

	account := &AccountReturn{}
	err = json.Unmarshal(res, &account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (t *Twominers) GetBlocks(server string) (*BlockReturn, error) {
	res, err := t.do(http.MethodGet, server+"/blocks")
	if err != nil {
		return nil, err
	}

	block := &BlockReturn{}
	err = json.Unmarshal(res, &block)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func (t *Twominers) GetMiners(server string) (*MinerReturn, error) {
	res, err := t.do(http.MethodGet, server+"/miners")
	if err != nil {
		return nil, err
	}

	miner := &MinerReturn{}
	err = json.Unmarshal(res, &miner)
	if err != nil {
		return nil, err
	}

	return miner, nil
}

func (t *Twominers) GetPayments(server string) (*PaymentReturn, error) {
	res, err := t.do(http.MethodGet, server+"/payments")
	if err != nil {
		return nil, err
	}

	payment := &PaymentReturn{}
	err = json.Unmarshal(res, &payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (t *Twominers) GetStats(server string) (*StatsReturn, error) {
	res, err := t.do(http.MethodGet, server+"/stats")
	if err != nil {
		return nil, err
	}

	stats := &StatsReturn{}
	err = json.Unmarshal(res, &stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
