package cheat

import future.keywords

allow if {
	prefix := sprintf("/docs/%s/", [input.userID])
	every path in input.paths {
		startswith(path, prefix)
	}
}
