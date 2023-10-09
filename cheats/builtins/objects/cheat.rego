package play

import future.keywords.if

obj := {"userid": "18472", "roles": [{"name": "admin"}]}

val := object.get(obj, ["roles", 0, "name"], "missing")
defaulted_val := object.get(obj, ["roles", 0, "permissions"], "unknown")

keys := object.keys(obj)