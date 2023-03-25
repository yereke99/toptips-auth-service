package dto

type ResponseDTO struct {
	Token   string  `json:"token"`
	Profile Profile `json:"profile"`
}

// Must to change
type Profile struct {
	Id        int64  `json:"id"`
	FirtsName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Status    string `json:"status"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	UserId    int64  `json:"userId"`
	Token     string `json:"token"`
}

type Bimetrics struct {
	Id              int64  `json:"id"`
	Inn             string `json:"inn"`
	Phone           string `json:"phone"`
	FirtsName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	MiddleName      string `json:"middleName"`
	PassPortNumber  string `json:"passportNumber"`
	PassPortIssue   string `json:"passportIssue"`
	PassPortIssueBy string `json:"passportIssueBy"`
	BirthDay        string `json:"birthday"`
}

type Waiter struct {
	Id      int64     `json:"id"`
	Profile Profile   `json:"profile"`
	Bio     Bimetrics `json:"biometrics"`
	Token   string    `json:"token"`
}
