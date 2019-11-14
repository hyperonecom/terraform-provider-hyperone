resource "hyperone_vm" "my_vm_image_1" {
    name = "tf-vm-image-1"
    type = "m2.tiny"
    ssh_keys = [ "mac" ]

    disk {
        name = "os-disk"
        type = "ssd"
        size = 10
    }
}

resource "hyperone_image" "my_image_1" {
    name = "tf-image-1"
    vm = "${hyperone_vm.my_vm_image_1.id}"
}
