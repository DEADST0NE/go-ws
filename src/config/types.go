package config

type BrokerTransportSS struct {
	Scheme string `json:"scheme"`
	Path   string `json:"path"`
	Host   string `json:"host"`
}

type BrokerTransportCore struct {
	Scheme string `json:"scheme"`
	Path   string `json:"path"`
	Host   string `json:"host"`
	Env    string `json:"env"`
}

type Broker struct {
	Ss   BrokerTransportSS   `json:"ss"`
	Core BrokerTransportCore `json:"core"`
}

type Candle struct {
	Periods    []string        `json:"periods"`
	Rsi_events map[string]bool `json:"rsi_events"`
}

type Rsi struct {
	Period     string `json:"period"`
	Interval   int    `json:"interval"`
	FillBlanks bool   `json:"fillBlanks"`
}

type Storage struct {
	Trade_history_limit int `json:"trade_history_limit"`
}

type Config struct {
	Broker  Broker   `json:"broker"`
	Symbols []string `json:"symbols"`
	Candle  Candle   `json:"candle"`
	Rsi     []Rsi    `json:"rsi"`
	Storage Storage  `json:"storage"`
}
