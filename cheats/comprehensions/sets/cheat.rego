package cheat

import future.keywords

unique_doubled := {m |
	some n in [10, 20, 30, 20, 10]
	m := n * 2
}
