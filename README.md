HyperOne Terraform Provider
==================


Getting The Provider
---------------------

```sh
# install latest (git) version of docker-machine-driver-hyperone in your $GOPATH/bin
$ go get -u github.com/hyperonecom/terraform-provider-hyperone
```

Using the provider
----------------------

```
provider "hyperone" {
    token   = "..."      # you can also use HYPERONE_TOKEN enviroment variable
    project = "..."      # you can also use HYPERONE_PROJECT enviroment variable
}

resource "hyperone_disk" "disk" {
    size = 200
    type = "volume"
    name = "nginx-logs"
}
```
