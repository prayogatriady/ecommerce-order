package model

import (
	"time"

	"gorm.io/gorm"
)

// t_orders
type Order struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	OrderedAt   time.Time     `gorm:"column:ordered_at;autoCreateTime"`
	TotalPrice  int64         `gorm:"column:total_price"`
	UserId      int64         `gorm:"column:user_id"`
	VoucherId   int64         `gorm:"column:voucher_id"`
	OrderDetail []OrderDetail `gorm:"foreignKey:order_id;references:id"` // one to many relationship
	User        *User         `gorm:"foreignKey:user_id;references:id"`  // many to one relationship (belongs to)
}

func (o *Order) TableName() string {
	return "t_orders"
}

// t_order_details
type OrderDetail struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	OrderId   int64    `gorm:"column:order_id"`
	ProductId int64    `gorm:"column:product_id"`
	Price     int64    `gorm:"column:price"`
	Quantity  int64    `gorm:"column:quantity"`
	Order     *Order   `gorm:"foreignKey:order_id;references:id"`   // many to one relationship (belongs to)
	Product   *Product `gorm:"foreignKey:product_id;references:id"` // many to one relationship (belongs to)
}

func (o *OrderDetail) TableName() string {
	return "t_order_details"
}

// m_products
type Product struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	ProductName string        `gorm:"column:product_name"`
	Description string        `gorm:"column:description"`
	Quantity    int64         `gorm:"column:quantity"`
	Price       int64         `gorm:"column:price"`
	UserId      int64         `gorm:"column:user_id"`
	IsActive    bool          `gorm:"column:is_active"`
	OrderDetail []OrderDetail `gorm:"foreignKey:product_id;references:id"` // one to many relationship
	User        *User         `gorm:"foreignKey:user_id;references:id"`    // many to one relationship (belongs to)
}

func (p *Product) TableName() string {
	return "m_products"
}

// m_vouchers
type Voucher struct {
	ID           int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Description  string    `gorm:"column:description"`
	MinPrice     int64     `gorm:"column:min_price"`
	QuantityAll  int64     `gorm:"column:quantity_all"`
	QuantityUser int64     `gorm:"column:quantity_user"`
	ExpiredAt    time.Time `gorm:"column:expired_at"`
	IsActive     bool      `gorm:"column:is_active"`
	Order        []Order   `gorm:"foreignKey:voucher_id;references:id"` // one to many relationship
}

func (v *Voucher) TableName() string {
	return "m_vouchers"
}

// m_users
type User struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Username  string    `gorm:"column:username"`
	FullName  string    `gorm:"column:full_name"`
	Password  string    `gorm:"column:password"`
	RoleId    int64     `gorm:"column:role_id"`
	Balance   int64     `gorm:"column:balance"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	IsActive  bool      `gorm:"column:is_active"`
	Order     []Order   `gorm:"foreignKey:user_id;references:id"` // one to many relationship
	Product   []Product `gorm:"foreignKey:user_id;references:id"` // one to many relationship
	Role      *Role     `gorm:"foreignKey:role_id;references:id"` // many to one relationship (belongs to)
}

func (u *User) TableName() string {
	return "m_users"
}

// m_roles
type Role struct {
	ID          int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	RoleName    string `gorm:"column:role_name"`
	Description string `gorm:"column:description"`
	IsActive    bool   `gorm:"column:is_active"`
	User        []User `gorm:"foreignKey:role_id;references:id"` // one to many relationship
}

func (r *Role) TableName() string {
	return "m_roles"
}
