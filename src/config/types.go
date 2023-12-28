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
	Periods []string `json:"periods"`
}

type Rsi struct {
	Period        int8   `json:"period"`
	Interval      string `json:"interval"`
	MovingAverage string `json:"movingAverage"`
	FillBlanks    bool   `json:"fillBlanks"`
	Storage       bool   `json:"storage"`
}

type Config struct {
	Broker  Broker   `json:"broker"`
	Symbols []string `json:"symbols"`
	Candle  Candle   `json:"candle"`
	Rsi     []Rsi    `json:"rsi"`
}
