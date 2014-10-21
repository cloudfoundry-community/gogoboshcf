# Cloud Foundry enhancements for gogobosh

The `gogobosh` golang package is a client library for a BOSH director. This `gogoboshcf` package enhances `gogobosh` for Cloud Foundry specific deployments.

* Godocs https://godoc.org/github.com/cloudfoundry-community/gogoboshcf
* Test status ![ci](https://travis-ci.org/cloudfoundry-community/gogoboshcf.svg?branch=master)

## Examples

```
$ go run examples/extractmanifest.go samples/bosh-lite-cf.yml
gogoboshcf.PropertiesManifest{
    NATS: gogoboshcf.NATS{
        MachinesHostnames: {"10.244.0.6"},
        Port:              4222,
        Username:          "nats",
        Password:          "nats",
    },
    UAA: gogoboshcf.UAA{
        URI:   "https://uaa.10.244.0.34.xip.io",
        Admin: gogoboshcf.ClientIDSecret{ClientID:"", ClientSecret:"admin-secret"},
    },
    RootDomain:       "10.244.0.34.xip.io",
    SystemDomain:     "10.244.0.34.xip.io",
    AppDomains:       {"10.244.0.34.xip.io"},
    SSL:              gogoboshcf.SSL{SkipCertificateVerify:true},
    SyslogAggregator: {},
}
```

```
$ go run examples/extractmanifest.go samples/cf-ec2-tiny.yml
gogoboshcf.PropertiesManifest{
    NATS: gogoboshcf.NATS{
        MachinesHostnames: {"0.data.cf1.my-aws-ec2.microbosh"},
        Port:              4222,
        Username:          "nats",
        Password:          "c1oudc0w",
    },
    UAA: gogoboshcf.UAA{
        URI:   "https://uaa.54.163.246.230.xip.io",
        Admin: gogoboshcf.ClientIDSecret{ClientID:"", ClientSecret:"c1oudc0w"},
    },
    RootDomain:       "54.163.246.230.xip.io",
    SystemDomain:     "54.163.246.230.xip.io",
    AppDomains:       {"54.163.246.230.xip.io"},
    SSL:              gogoboshcf.SSL{SkipCertificateVerify:true},
    SyslogAggregator: {},
}
```
