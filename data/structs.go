package data

type Recipe struct { // template struct to fill in later in the same layout / field name and type
	Id           string            `json:"Id"`				// to call the json files, take the information from this name into this field
	RecipeName   string            `json:"Recipe Name"`
	Source       string            `json:"Source"`
	PrepTime     string            `json:"Preperation Time"`
	CookTime     string            `json:"Cooking Time"`
	ServingSize  int               `json:"Serving Size"`
	Ingredients  map[string]string `json:"Ingredients"`
	Instructions []string          `json:"Directions"`
	Tag          []string          `json:"tags"`
}

 var AllRecipes []Recipe