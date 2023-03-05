package objects

var ProxyConfigObj ProxyConfig

type ProxyConfig struct {
	Port           string `json:"port"`
	BinanceBaseURL string `json:"binanceBaseURL"`
	KuCoinBaseURL  string `json:"kuCoinBaseURL"`
	GateIOBaseURL  string `json:"gateIOBaseURL"`
	ClientServer   string `json:"clientServer"`
}
