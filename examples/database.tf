resource "hyperone_database" "my_database_1" {
    name = "tf-database-1"
    type = "mysql:5.7"
}

resource "hyperone_database_credential" "my_database_1_credential_password" {
    database_id = "${hyperone_database.my_database_1.id}"
    name = "tf-database-1-credential-password"
    type = "mysql"
    value = "d85f49bf0c69f166401e62269534f54e325937f3" # sha1(sha1("mySecretPassword"))
}