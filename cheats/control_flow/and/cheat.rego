package cheat

import future.keywords

valid_staff_email if {
	regex.match(`^\S+@\S+\.\S+$`, input.email)

	endswith(input.email, "example.com")
}
