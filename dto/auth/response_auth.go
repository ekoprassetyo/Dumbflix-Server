package authdto

type RegisterResponse struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Gender   string `json:"gender" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" gorm:"type: varchar(255)"`
	Address  string `json:"address" gorm:"type: varchar(255)"`
	Status   bool   `json:"status" gorm:"type: boolean"`
}

type LoginResponse struct {
	Name      string `json:"name" form:"name"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Token     string `gorm:"type: varchar(255)" json:"token"`
	Gender    string `json:"gender" gorm:"type: varchar(255)"`
	Phone     string `json:"phone" gorm:"type: varchar(255)"`
	Address   string `json:"address" gorm:"type: varchar(255)"`
	Subscribe bool   `json:"subscribe" form:"subscribe"`
	Role      string `gorm:"type: varchar(255)" json:"role"`
}

type CheckAuthResponse struct {
	ID        int    `json:"id"`
	Name      string `gorm:"type: varchar(255)" json:"name"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Gender    string `gorm:"type: varchar(255)" json:"gender"`
	Phone     string `gorm:"type: varchar(255)" json:"phone"`
	Address   string `gorm:"type: varchar(255)" json:"address"`
	Token     string `gorm:"type: varchar(255)" json:"token"`
	Subscribe bool   `json:"subscribe" form:"subscribe"`
	Role      string `gorm:"type: varchar(100)" json:"role"`
}
