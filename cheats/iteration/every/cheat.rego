package cheat

import future.keywords

allow if {
	required_prefix := sprintf("/docs/%s/", [input.userID])
	every path in input.paths {
		startswith(path, required_prefix)
	}
}
