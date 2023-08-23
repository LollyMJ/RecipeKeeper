package data

type Recipe struct {
	Id           string
	RecipeName   string
	Source       string
	PrepTime     string
	CooTime      string
	ServingSize  int
	Ingredients  map[string]string
	Instructions []string
	Tags         []string
}
