package config

import (
	"os"
	"runtime"
	"time"
)

var defaultValues = struct {
	ListenerValue         string
	CertSourcesValue      string
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	UIListenerValue       string
	GZIPContentTypesValue string
}{
	ListenerValue:   ":9999",
	UIListenerValue: ":9998",
}

var defaultConfig = &Config{
	ProfilePath: os.TempDir(),
	Log: Log{
		AccessFormat: "common",
		RoutesFormat: "delta",
		Level:        "INFO",
	},
	Metrics: Metrics{
		Prefix:   "{{clean .Hostname}}.{{clean .Exec}}",
		Names:    "{{clean .Service}}.{{clean .Host}}.{{clean .Path}}.{{clean .TargetURL.Host}}",
		Interval: 30 * time.Second,
		Timeout:  10 * time.Second,
		Retry:    500 * time.Millisecond,
		Circonus: Circonus{
			APIApp: "fabio",
		},
	},
	Proxy: Proxy{
		MaxConn:       10000,
		Strategy:      "rnd",
		Matcher:       "prefix",
		NoRouteStatus: 404,
		DialTimeout:   30 * time.Second,
		FlushInterval: time.Second,
		LocalIP:       LocalIPString(),
	},
	Registry: Registry{
		Backend: "consul",
		Consul: Consul{
			Addr:                                "localhost:8500",
			Scheme:                              "http",
			KVPath:                              "/fabio/config",
			NoRouteHTMLPath:                     "/fabio/noroute.html",
			TagPrefix:                           "urlprefix-",
			Register:                            true,
			ServiceAddr:                         ":9998",
			ServiceName:                         "fabio",
			ServiceStatus:                       []string{"passing"},
			CheckInterval:                       time.Second,
			CheckTimeout:                        3 * time.Second,
			CheckScheme:                         "http",
			CheckDeregisterCriticalServiceAfter: "90m",
		},
		Kubernetes: Kubernetes{
			ApiUrl:                              "http://localhost:8080",
      Token:                               "",
			ServicePath:                         "/api/v1/services",
      ExtDomain:                           "com",
      IntDomain:                           "svc.cluster.local",
			NoRouteHTML:                         "No Route",
			CheckInterval:                       10 * time.Second,
			CheckTimeout:                        5 * time.Second,
			CheckScheme:                         "http",
		},
		Timeout: 10 * time.Second,
		Retry:   500 * time.Millisecond,
	},
	Runtime: Runtime{
		GOGC:       800,
		GOMAXPROCS: runtime.NumCPU(),
	},
	UI: UI{
		Listen: Listen{
			Addr:  ":9998",
			Proto: "http",
		},
		Color:  "light-green",
		Access: "rw",
	},
}
