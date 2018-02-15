package kubernetes

import (
  "log"
  "time"
  "net/http"
  "os"
  "io/ioutil"

	"github.com/fabiolb/fabio/config"
	"github.com/fabiolb/fabio/registry"
)

type be struct {
	cfg *config.Kubernetes
}

func NewBackend(cfg *config.Kubernetes) (registry.Backend, error) {
  client := &http.Client{ Timeout: 5 * time.Second }

  _, err := os.Stat(cfg.Token)
  if err == nil {
    tok, err := ioutil.ReadFile(cfg.Token)
    if err == nil {
      log.Printf("[INFO] Token read from '%s'.", cfg.Token)
      cfg.Token = string(tok)
    }
  }

  req, err := http.NewRequest("GET", cfg.ApiUrl, nil)
  if err != nil {
    log.Fatal(err)
  }

  req.Header.Add("Authorization", "Bearer " + cfg.Token)

  resp, err := client.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()

  _, err = ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("[INFO] Connected to Kubernetes API at %s.", cfg.ApiUrl)
  return &be{cfg: cfg}, nil
}

func (b *be) Register(services []string) error {
	return nil
}

func (b *be) Deregister(serviceName string) error {
	return nil
}

func (b *be) DeregisterAll() error {
	return nil
}

func (b *be) ManualPaths() ([]string, error) {
	return nil, nil
}

func (b *be) ReadManual(string) (value string, version uint64, err error) {
	return "", 0, nil
}

func (b *be) WriteManual(path string, value string, version uint64) (ok bool, err error) {
	return false, nil
}

func (b *be) WatchServices() chan string {
  ch := make(chan string, 1)
  go watchServices(b.cfg.ApiUrl + b.cfg.ServicePath, "Bearer " + b.cfg.Token, b.cfg.ExtDomain, b.cfg.IntDomain, b.cfg.CheckInterval, ch)
  return ch
}

func (b *be) WatchManual() chan string {
	return make(chan string)
}

func (b *be) WatchNoRouteHTML() chan string {
	ch := make(chan string, 1)
	ch <- b.cfg.NoRouteHTML
	return ch
}
