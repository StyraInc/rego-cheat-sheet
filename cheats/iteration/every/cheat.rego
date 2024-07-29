package cheat

import rego.v1

allow if {
	prefix := sprintf("/docs/%s/", [input.userID])
	every path in input.paths {
		startswith(path, prefix)
	}
}
