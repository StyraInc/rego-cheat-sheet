package cheat

import future.keywords

commit_list := [
	{"sha": "701b7f", "message": "WIP"},
	{"sha": "badf2b", "message": "Build Feature A"},
]

commit_invalid contains "WIP commits are not allowed" if {
	commits := {commit.sha: commit.message |
		some commit in commit_list
	}
	commits[input.sha] == "WIP"
}
