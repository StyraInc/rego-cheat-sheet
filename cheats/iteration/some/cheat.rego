package cheat

import future.keywords

all_regions := {
	"emea": {"west", "east"},
	"na": {"west", "east", "central"},
	"latam": {"west", "east"},
	"apac": {"north", "south"},
}

allowed_regions contains region_id if {
	some area, regions in all_regions

	some region in regions
	region_id := sprintf("%s_%s", [area, region])
}
