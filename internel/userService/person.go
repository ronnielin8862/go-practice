package person

type Person struct {
	ID        string   `json:"id,omitemty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
}