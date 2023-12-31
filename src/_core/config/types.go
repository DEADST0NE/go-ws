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

type Redis struct {
	Host string `json:"host"`
}

type WsDefault struct {
	Rsi string `json:"rsi"`
}

type Ws struct {
	Port    int       `json:"port"`
	Default WsDefault `json:"default"`
}

type Pg struct {
	Host string `json:"host"`
}

type Api struct {
	Port string `json:"port"`
}

type Config struct {
	Ws      Ws       `json:"ws"`
	Redis   Redis    `json:"redis"`
	Broker  Broker   `json:"broker"`
	Symbols []string `json:"symbols"`
	Candle  Candle   `json:"candle"`
	Rsi     []Rsi    `json:"rsi"`
	Storage Storage  `json:"storage"`
	Pg      Pg       `json:"pg"`
	Api     Api      `json:api`
}
