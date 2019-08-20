HyperOne Terraform Provider
==================

[![Build Status](https://travis-ci.com/hyperonecom/terraform-provider-hyperone.svg?branch=master)](https://travis-ci.com/hyperonecom/terraform-provider-hyperone)
[![Go Report Card](https://goreportcard.com/badge/github.com/hyperonecom/terraform-provider-hyperone)](https://goreportcard.com/report/github.com/hyperonecom/terraform-provider-hyperone)

Getting The Provider
---------------------

```sh
# install latest (git) version of terraform-provider-hyperone in your $GOPATH/bin
$ go get -u github.com/hyperonecom/terraform-provider-hyperone
```

Using the provider
----------------------

```
provider "hyperone" {
    token   = "..."      # you can also use HYPERONE_ACCESS_TOKEN_SECRET enviroment variable
    project = "..."      # you can also use HYPERONE_PROJECT enviroment variable
}

resource "hyperone_disk" "disk" {
    size = 200
    type = "volume"
    name = "nginx-logs"
}

resource "hyperone_firewall" "firewall" {
    name = "nginx-firewall"

    ingress {
        name = "http"
        action = "allow"
        filter = [ "tcp:80" ]
        external = [ "0.0.0.0/0" ]
        internal = [ "*" ]
        priority = 300
    }

    egress {
        name = "all"
        action = "allow"
        filter = [ "tcp", "udp" ]
        external = [ "0.0.0.0/0" ]
        internal = [ "*" ]
        priority = 100
    }
}
```

Using H1 cli to set the environment variables
----------------------

```
$ eval $(h1 env)
```

- [Reference documentation for `h1 env`](https://www.hyperone.com/tools/cli/resources/reference/env.html#syntax)
- [Documentation for h1-cli](https://www.hyperone.com/tools/cli/)
