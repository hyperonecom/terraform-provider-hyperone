resource "hyperone_ip" "my_ip_1" {
}

resource "hyperone_ip" "my_ip_2" {
    ptr_record = "test.hyperone.com"
}

resource "hyperone_network" "my_network_ip_3" {
    name = "tf-network-ip-3"
    address = "10.1.0.0/24"
}

resource "hyperone_ip" "my_ip_3" {
    network = "${hyperone_network.my_network_ip_3.id}"
    address = "10.1.0.20"
}
