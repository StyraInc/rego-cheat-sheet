package cheat

import rego.v1

allow if input.admin == true

test_allow_when_admin if {
	allow with input as {"admin": true}
}
