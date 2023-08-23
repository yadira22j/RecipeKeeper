package data

type Recipe struct { // template struct to fill in later in the same layout / field name and type
	Id           string
	RecipeName   string
	Source       string
	PrepTime     string
	CookTime     string
	ServingSize  int
	Ingredients  map[string]string
	Instructions []string
	Tag          []string
}
