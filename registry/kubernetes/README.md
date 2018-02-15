## Kubernetes Backend (fabio-k8)
This Kubernetes backend is a very simplistic implementation of a Fabio Backend. It simply pulls the provide ApiUrl + ServicePath from the given configuration (-registry.kubernetes.apiurl and -registry.kubernetes.servicepath) and parses the Kubernetes Api response to generate Fabio routes.  It makes some VERY important assumptions (TCP+SNI and {service}.{namespace}) which you will need to understand in order to use this backend.

### Kubernetes L2 TCP+SNI Global Load Balancing
This backend is designed to route external traffic from outside your cluster to the internal services running inside namespaces.  Fabio-k8 will currently only route TCP+SNI like so:
```
  <service>.<namespace>.<external-domain-name>   -----(tcp+sni)---->   <service>.<namespace>.<internal-domain-name>
```

### Testing
The following steps can be taken to test that fabio will add routes from a test url.

1. change into this directory (so that services.json is in current directory)
```
  cd ~/fabio/registry/kubernetes/
```

2. start python webserver
```
  python -m SimpleHTTPServer
```

3. Compile fabio
```
  ./$HOME/fabio/build/docker.sh
```

4. Start fabio with k8 backend pointed to localhost url
```
  ./fabio -registry.backend kubernetes -registry.kubernetes.apiurl http://localhost:8000 -registry.kubernetes.servicepath /services.json
```

5. Watch log messages to verify routes are created
```
+ route add kubernetes-default kubernetes.default.com kubernetes.default.svc.cluster.local:443 weight 1 opts "proto=tcp+sni"
+ route add kubernetes-dashboard-kube-system kubernetes-dashboard.kube-system.com kubernetes-dashboard.kube-system.svc.cluster.local:80 weight 1 opts "proto=tcp+sni"
+ route add kube-dns-kube-system kube-dns.kube-system.com kube-dns.kube-system.svc.cluster.local:53 weight 1 opts "proto=tcp+sni"
2018/02/14 23:36:51 [INFO] HTTP proxy listening on :9999
```
