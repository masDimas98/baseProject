package baseModel

type Session struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	LoginTime int64  `json:"login_time"`
	ExpiresIn int64  `json:"expires_in"`
	IPAddress string `json:"ip_address"`
}
