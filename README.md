## Alertmanager Woa Webhook

Webhook service support send Prometheus 2.0 alert message to Woa.

## How To Use

```
cd cmd/webhook
go build
webhook -defaultRobot=https://xz.wps.cn/api/v1/webhook/send?key=xxx -cluster xxx
```

```
go run webhook.go -defaultRobot=https://xz.wps.cn/api/v1/webhook/send?key=xxx -cluster xxx
```

* -defaultRobot: default dingtalk webhook url, all notifaction from alertmanager will direct to this webhook address.

Or you can overwrite by add annotations to Prometheus alertrule to special the woa webhook for each alert rule.

```
groups:
- name: hostStatsAlert
  rules:
  - alert: hostCpuUsageAlert
    expr: sum(avg without (cpu)(irate(node_cpu{mode!='idle'}[5m]))) by (instance) > 0.85
    for: 1m
    labels:
      severity: page
    annotations:
      summary: "Instance {{ $labels.instance }} CPU usgae high"
      description: "{{ $labels.instance }} CPU usage above 85% (current value: {{ $value }})"
      woaRobot: "https://xz.wps.cn/api/v1/webhook/send?key=xxx"
```
