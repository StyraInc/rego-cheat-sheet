package cheat

import rego.v1

allowed_users := {"alice", "bob", "charlie"}

allow if {
	some user in allowed_users
	print(user)
	input.user == user
}
