package dto

type CreateUserDTO struct {
	TotalPrice  int64 `json:"total_price"`
	UserId      int64 `json:"user_id"`
	VoucherId   int64 `json:"voucher_id"`
	OrderDetail []struct {
		ProductId int64 `json:"product_id"`
		Price     int64 `json:"price"`
		Quantity  int64 `json:"quantity"`
	} `json:"order_detail"`
}
