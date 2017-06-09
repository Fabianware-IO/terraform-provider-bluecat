provider "blah" {
	login = "foo"
	password = "bar"
}

resource "bluecat_dns" "client" {
	name = "foo"
	content = "bar"
}
