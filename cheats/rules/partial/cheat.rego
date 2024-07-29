package cheat

import rego.v1

paths contains "/handbook/*"

paths contains path if {
	some team in input.user.teams
	path := sprintf("/teams/%v/*", [team])
}
