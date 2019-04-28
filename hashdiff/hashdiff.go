package hashdiff

// DiffMap returns the original diff'd keys/values that differ from the second maps keys/values
func DiffMap(firstMappings map[string]string,
	secondMappings map[string]string) map[string]string {
	diffMap := map[string]string{}

	for k, originalVal := range firstMappings {
		secondVal, ok := secondMappings[k]

		// key not found in second map
		if !ok {
			diffMap[k] = originalVal
		} else {
			// key found in second map but the values differ
			if secondVal != originalVal {
				diffMap[k] = originalVal
			}
		}
	}

	return diffMap
}
