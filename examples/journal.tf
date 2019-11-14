resource "hyperone_journal" "my_journal_1" {
    name = "tf-journal-1"
    retention = 30
}

resource "hyperone_journal_credential" "my_journal_1_credential_password" {
    journal_id = "${hyperone_journal.my_journal_1.id}"
    name = "tf-journal-1-credential-password"
    type = "sha512"
    value = "${base64encode("mysalt_")} ${base64sha512("mysalt_mySecretPassword")}"
}