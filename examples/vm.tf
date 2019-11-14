
resource "hyperone_network" "my_network_vm_1" {
    name = "tf-network-vm-1"
    address = "10.1.0.0/24"
}
resource "hyperone_vm" "my_vm_1" {
    name = "tf-vm-1"
    image = "alpine"
    type = "m2.tiny"
    ssh_keys = [ "mac" ]
    password = "mySecretPassword"

    disk {
        name = "os-disk"
        type = "ssd"
        size = 10
    }

    netadp {
        network = "${hyperone_network.my_network_vm_1.id}"
        type = "private"
    }

    user_metadata = "${base64encode("my-config")}"
}