resource "hyperone_vault" "my_vault_1" {
    name = "tf-vault-1"
    size = 10
}

resource "hyperone_vault_credential" "my_vault_1_credential_password" {
    vault_id = "${hyperone_vault.my_vault_1.id}"
    name = "tf-vault-1-credential-password"
    type = "sha512"
    value = "${base64encode("mysalt_")} ${base64sha512("mysalt_fred")}"
}

resource "hyperone_vault_credential" "my_vault_1_credential_ssh" {
    vault_id = "${hyperone_vault.my_vault_1.id}"
    name = "tf-vault-1-credential-ssh"
    type = "ssh"
    value = file("~/.ssh/id_rsa.pub")
}
