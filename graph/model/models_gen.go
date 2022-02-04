// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}

type Comment struct {
	ID        int     `json:"id"`
	IDEvent   int     `json:"id_event"`
	IDUser    int     `json:"id_user"`
	Comment   *string `json:"comment"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

type Event struct {
	ID         int     `json:"id"`
	IDUser     int     `json:"id_user"`
	IDCategory int     `json:"id_category"`
	Title      string  `json:"title"`
	StartDate  string  `json:"start_date"`
	EndDate    string  `json:"end_date"`
	Location   string  `json:"location"`
	Details    string  `json:"details"`
	Photo      *string `json:"photo"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
	DeletedAt  *string `json:"deleted_at"`
}

type EventDetail struct {
	ID          int          `json:"id"`
	IDUser      int          `json:"id_user"`
	IDCategory  int          `json:"id_category"`
	Title       string       `json:"title"`
	StartDate   string       `json:"start_date"`
	EndDate     string       `json:"end_date"`
	Location    string       `json:"location"`
	Details     string       `json:"details"`
	Photo       *string      `json:"photo"`
	CreatedAt   *string      `json:"created_at"`
	UpdatedAt   *string      `json:"updated_at"`
	DeletedAt   *string      `json:"deleted_at"`
	Comments    *Comment     `json:"comments"`
	Participant *Participant `json:"participant"`
}

type NewComment struct {
	IDEvent   int     `json:"id_event"`
	IDUser    int     `json:"id_user"`
	Comment   *string `json:"comment"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

type NewEvent struct {
	IDUser     int     `json:"id_user"`
	IDCategory int     `json:"id_category"`
	Title      string  `json:"title"`
	StartDate  string  `json:"start_date"`
	EndDate    string  `json:"end_date"`
	Location   string  `json:"location"`
	Details    string  `json:"details"`
	Photo      *string `json:"photo"`
	CreatedAt  *string `json:"created_at"`
	UpdatedAt  *string `json:"updated_at"`
	DeletedAt  *string `json:"deleted_at"`
}

type NewParticipant struct {
	IDEvent   int     `json:"id_event"`
	IDUser    int     `json:"id_user"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

type NewUser struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	BirthDate   *string `json:"birth_date"`
	PhoneNumber *string `json:"phone_number"`
	Photo       *string `json:"photo"`
	Gender      *string `json:"gender"`
	Address     *string `json:"address"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}

type Participant struct {
	ID        int     `json:"id"`
	IDEvent   int     `json:"id_event"`
	IDUser    int     `json:"id_user"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}

type ResponseLogin struct {
	Code  int    `json:"code"`
	Token string `json:"token"`
}

type ResponseMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type User struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	BirthDate   *string `json:"birth_date"`
	PhoneNumber *string `json:"phone_number"`
	Photo       *string `json:"photo"`
	Gender      *string `json:"gender"`
	Address     *string `json:"address"`
	CreatedAt   *string `json:"created_at"`
	UpdatedAt   *string `json:"updated_at"`
	DeletedAt   *string `json:"deleted_at"`
}
