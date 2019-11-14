
resource "hyperone_ip" "my_netgw_1_ip" {
}

resource "hyperone_netgw" "my_netgw_1" {
    name = "tf-netgw-1"
    ip = "${hyperone_ip.my_netgw_1_ip.id}"
}