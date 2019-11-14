resource "hyperone_zone" "my_zone_1" {
    name = "my-zone-1"
    dns_name = "myzone1.com."
    type = "public"
}

resource "hyperone_zone_recordset" "my_zone_1_recordset" {
    zone_id = "${hyperone_zone.my_zone_1.id}"
    type = "A"
    name = "test.myzone1.com."
    ttl = 60

    record {
        content = "1.2.3.4"
    }
}