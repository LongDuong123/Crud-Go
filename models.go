package main

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Age      int    `json:"Age"`
	Gender   string `json:"Gender"`
	Email    string `json:"Email"`
	Cart     int    `json:"Cart"`
	Role     string `json:"Role"`
	Password string `json:"Password"`
}

type Product struct {
	ID        int    `json:"ID"`
	Name      string `json:"Name"`
	Image_url string `json:"Image"`
	Price     string `json:"Price"`
	Create_By int    `json:"Create_UserID"`
}
