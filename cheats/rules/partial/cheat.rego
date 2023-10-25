package cheat

import future.keywords

paths contains path if {
	path := "/handbook/*"
}

paths contains path if {
	some team in input.user.teams
	path := sprintf("/teams/%v/*", [team])
}
