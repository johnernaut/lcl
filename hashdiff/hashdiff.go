package hashdiff

// DiffMapKeys returns the original diff'd keys/values that differ from the first maps keys
func DiffMapKeys(firstMappings map[string]string,
	secondMappings map[string]string) map[string]string {
	diffMap := map[string]string{}

	for k, v := range firstMappings {
		if _, ok := secondMappings[k]; !ok {
			diffMap[k] = v
		}
	}

	return diffMap
}

// DiffMapValues returns the original diff'd keys/values that differ from the first maps values
func DiffMapValues(firstMappings map[string]string,
	secondMappings map[string]string) map[string]string {
	diffMap := map[string]string{}

	for k, v := range firstMappings {
		if val, ok := secondMappings[k]; ok {
			if val != v {
				diffMap[k] = v
			}
		}
	}

	return diffMap
}
