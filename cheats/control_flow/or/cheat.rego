package cheat

import future.keywords

# using multiple rules
valid_email if endswith(input.email, "@example.com")
valid_email if endswith(input.email, "@example.org")
valid_email if endswith(input.email, "@example.net")

# using functions
allowed_firstname(name) if name == "joe"
allowed_firstname(name) if name == "jane"

valid_name if {
	allowed_firstname(input.name)
}

# using `in`
valid_request if {
	input.method in {"GET", "POST"}
}
