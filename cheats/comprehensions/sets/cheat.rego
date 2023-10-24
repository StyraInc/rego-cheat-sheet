package cheat

import future.keywords

default unique := false

unique if {
	numbers := [1, 2, 3, 3, 4, 5]
	unique_numbers := {n |
		some n in numbers
	}
	count(unique_numbers) == count(numbers)
}
