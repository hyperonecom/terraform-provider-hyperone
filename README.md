HyperOne Terraform Provider
==================


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
```

Using H1 cli to set the environment variables
----------------------

```
$ eval $(h1 env)
```

- [Reference documentation for `h1 env`](https://www.hyperone.com/tools/cli/resources/reference/env.html#syntax)
- [Documentation for h1-cli](https://www.hyperone.com/tools/cli/)