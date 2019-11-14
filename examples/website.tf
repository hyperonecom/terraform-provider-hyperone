resource "hyperone_website" "my_website_1" {
    name = "tf-website-1"
    type = "website"
    image = "h1cr.io/website/nginx-static:latest"
}

resource "hyperone_website_credential" "my_website_1_credential_password" {
    website_id = "${hyperone_website.my_website_1.id}"
    name = "tf-website-1-credential-password"
    type = "sha512"
    value = "${base64encode("mysalt_")} ${base64sha512("mysalt_fred")}"
}

resource "hyperone_website_credential" "my_website_1_credential_ssh" {
    website_id = "${hyperone_website.my_website_1.id}"
    name = "tf-website-1-credential-ssh"
    type = "ssh"
    value = file("~/.ssh/id_rsa.pub")
}
