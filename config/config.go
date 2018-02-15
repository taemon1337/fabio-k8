package config

import (
	"net/http"
	"regexp"
	"time"
)

type Config struct {
	Proxy       Proxy
	Registry    Registry
	Listen      []Listen
	Log         Log
	Metrics     Metrics
	UI          UI
	Runtime     Runtime
	ProfileMode string
	ProfilePath string
	Insecure    bool
}

type CertSource struct {
	Name         string
	Type         string
	CertPath     string
	KeyPath      string
	ClientCAPath string
	CAUpgradeCN  string
	Refresh      time.Duration
	Header       http.Header
}

type Listen struct {
	Addr          string
	Proto         string
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	CertSource    CertSource
	StrictMatch   bool
	TLSMinVersion uint16
	TLSMaxVersion uint16
	TLSCiphers    []uint16
}

type UI struct {
	Listen Listen
	Color  string
	Title  string
	Access string
}

type Proxy struct {
	Strategy              string
	Matcher               string
	NoRouteStatus         int
	MaxConn               int
	ShutdownWait          time.Duration
	DialTimeout           time.Duration
	ResponseHeaderTimeout time.Duration
	KeepAliveTimeout      time.Duration
	FlushInterval         time.Duration
	LocalIP               string
	ClientIPHeader        string
	TLSHeader             string
	TLSHeaderValue        string
	GZIPContentTypes      *regexp.Regexp
	RequestID             string
	STSHeader             STSHeader
}

type STSHeader struct {
	MaxAge     int
	Subdomains bool
	Preload    bool
}

type Runtime struct {
	GOGC       int
	GOMAXPROCS int
}

type Circonus struct {
	APIKey   string
	APIApp   string
	APIURL   string
	CheckID  string
	BrokerID string
}

type Log struct {
	AccessFormat string
	AccessTarget string
	RoutesFormat string
	Level        string
}

type Metrics struct {
	Target       string
	Prefix       string
	Names        string
	Interval     time.Duration
	Timeout      time.Duration
	Retry        time.Duration
	GraphiteAddr string
	StatsDAddr   string
	Circonus     Circonus
}

type Registry struct {
	Backend string
	Static  Static
	File    File
	Consul  Consul
  Kubernetes Kubernetes
	Timeout time.Duration
	Retry   time.Duration
}

type Static struct {
	NoRouteHTML string
	Routes      string
}

type File struct {
	NoRouteHTMLPath string
	RoutesPath      string
}

type Consul struct {
	Addr                                string
	Scheme                              string
	Token                               string
	KVPath                              string
	NoRouteHTMLPath                     string
	TagPrefix                           string
	Register                            bool
	ServiceAddr                         string
	ServiceName                         string
	ServiceTags                         []string
	ServiceStatus                       []string
	CheckInterval                       time.Duration
	CheckTimeout                        time.Duration
	CheckScheme                         string
	CheckTLSSkipVerify                  bool
	CheckDeregisterCriticalServiceAfter string
}

type Kubernetes struct {
	ApiUrl                              string
	Token                               string
	ServicePath                         string
  ExtDomain                           string
  IntDomain                           string
	NoRouteHTML                         string
	CheckInterval                       time.Duration
	CheckTimeout                        time.Duration
	CheckScheme                         string
	CheckTLSSkipVerify                  bool
}

