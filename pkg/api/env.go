package api

import (
	"slices"
)

var baseUrl = map[string]string{
	"ito":  "http://10.1.65.31",
	"eto":  "http://test.purs.gov.rs",
	"test": "http://test.purs.gov.rs",
	"prod": "http://eporezi.purs.gov.rs",
}

var validEnvs = []string{"ito", "eto", "test", "prod"}

func isValidEnv(env string) bool {
	return slices.Contains(validEnvs, env)
}
