package cheat

import rego.v1

unique_doubled contains m if {
	some n in [10, 20, 30, 20, 10]
	m := n * 2
}
