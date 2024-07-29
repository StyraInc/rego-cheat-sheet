package cheat

import rego.v1

doubled := [m |
	some n in [1, 2, 3, 3]
	m := n * 2
]
