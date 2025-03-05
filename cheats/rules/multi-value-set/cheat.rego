package cheat

paths contains "/handbook/*"

paths contains path if {
	some team in input.user.teams
	path := sprintf("/teams/%v/*", [team])
}
