package cheat

import rego.v1

# Creates an object with sets as the values.
paths_by_prefix[prefix] contains path if {
	some path in input.paths
	parts := split(path, "/")
	prefix := parts[0]
}
