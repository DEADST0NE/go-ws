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

type Config struct {
	Broker  Broker   `json:"broker"`
	Symbols []string `json:"symbols"`
}
