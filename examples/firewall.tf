resource "hyperone_firewall" "my_firewall_1" {
    name = "tf-firewall-1"
    ingress {
        name = "ssh"
        action = "allow"
        filter = [ "tcp:22" ]
        external = [ "62.181.3.200/32" ]
        internal = [ "*" ]
        priority = 200
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