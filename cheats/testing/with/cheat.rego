package cheat

import future.keywords

allow if {
	input.admin == true
}

test_allow_when_admin if {
	allow with input as {"admin": true}
}
