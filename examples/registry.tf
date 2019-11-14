resource "hyperone_registry" "my_registry_1" {
    name = "tf-registry-1"
    type = "container"
}

resource "hyperone_registry_credential" "my_registry_1_credential_password" {
    registry_id = "${hyperone_registry.my_registry_1.id}"
    name = "tf-registry-1-credential-password"
    type = "sha512"
    value = "${base64encode("mysalt_")} ${base64sha512("mysalt_mySecretPassword")}"
}