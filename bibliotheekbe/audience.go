package bibliotheekbe

var audienceMapping = map[string]string{
	"age0-2":   "ages 0-2",
	"age3-5":   "ages 3-5",
	"age6-8":   "ages 6-8",
	"age9-11":  "ages 9-11",
	"age12-14": "ages 12-14",
	"ageYouth": "youth",
}

func getAudience(in string) string {
	new, ok := audienceMapping[in]
	if !ok {
		return in
	}
	return new
}
