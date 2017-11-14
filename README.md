# kong-node-exporter
Simple prometheus-compatible metrics API for Kong node status

Kong tracks nginx stats via the status API, but does not ship these metrics to statsd or prometheus via standard plugins. This provides a simple prometheus compatible metrics api which proxies to a specified kong status API.

## Usage

You need to specifiy a Kong IP and Port to his for the Admin API via env variables. You can also specify what port the exporter should listen on.

```
$ export EXPORTER_KONG_IP=1.2.3.4
$ export EXPORTER_KONG_PORT=8001
// optional
$ export EXPORTER_PORT=9100
```


To use the exporter simply hit the `/metrics` endpoint and it will hit the kong admin api `/status` endpoint, and translate the response data so prometheus can parse it.

This app also provides a `/health` endpoint which can be used to determine liveness of this application regardless of how the kong endpoint is doing. 

## Using in Kubernetes

The simplest way to use this in Kubernetes is to set this container up inside the Kong proxy pod. You can then configure it to hit `127.0.0.1:8001`, or whatever port the admin API is listening on. Since generally services round robin between instances of a replica or daemonset that would be running Kong, using a standard service to hit theis endpoint might not work. What I suggest instead is using the kubernetes proxy API to hit the pod directly, e.g. `api/v1/proxy/namespaces/kong-system/pods/kong-rc-1024079112-jwlr9/metrics`. This will only work if you configure the exporter to be the FIRST container within your pod.