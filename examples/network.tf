resource "hyperone_network" "my_network_1" {
    name = "tf-network-1"
    address = "10.0.0.0/24"
}

# resource "hyperone_firewall" "my_network_1_firewall" {
#     name = "tf-firewall-2"
# }

# resource "hyperone_network" "my_network_2" {
#     name = "tf-network-2"
#     firewall = "${hyperone_firewall.my_network_1_firewall.id}"
# }
