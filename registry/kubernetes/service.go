package kubernetes

import (
  "log"
  "net/http"
  "io/ioutil"
  "sort"
  "strings"
  "time"
  "github.com/tidwall/gjson"
)

func getServices(url string, authToken string, extDomain string, intDomain string) (string, error) {
  var config []string
  client := &http.Client{ Timeout: 5 * time.Second }
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    return "", nil
  }

  req.Header.Add("Authorization", authToken)

  resp, err := client.Do(req)
  if err != nil {
    return "", nil
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", nil
  }

  s := gjson.Get(string(body), "items")

  for _, svc := range s.Array() {
    name := gjson.Get(svc.String(), "metadata.name").String()
    ns := gjson.Get(svc.String(), "metadata.namespace").String()
    port := gjson.Get(svc.String(), "spec.ports.0.port").String()

    cfg := buildConfig(name, ns, port, extDomain, intDomain)
    config = append(config, cfg...)
  }

  sort.Sort(sort.Reverse(sort.StringSlice(config)))
  return strings.Join(config, "\n"), nil
}

func watchServices(url string, authToken string, extDomain string, intDomain string, checkInt time.Duration, config chan string) {
  var lastIndex uint64
  lastIndex = 1

  for {
    routes, err := getServices(url, authToken, extDomain, intDomain)
    if err != nil {
      log.Printf("[WARN] kubernetes: Error loading services: %s", err)
      continue
    }

    lastIndex = lastIndex + 1
    config <- routes
    time.Sleep(checkInt)
  }
}

func buildConfig(name string, ns string, port string, extDomain string, intDomain string) (config []string) {
  if name == "" || ns == "" {
   return nil
  }

  cfg := "route add"
  cfg += " " + name + "-" + ns
  cfg += " " + name + "." + ns + "." + extDomain
  cfg += " " + name + "." + ns + "." + intDomain + ":" + port
  cfg += " weight 1"
  cfg += ` opts "proto=tcp+sni"`

  log.Printf("[INFO] %s", cfg)
  config = append(config, cfg)
  return config
}
