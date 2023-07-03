package domain

type Auth struct {
	Domain   string `json:"domain"`
	Ldap     bool   `json:"ldap"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthRequest struct {
	Domain   string `json:"domain" binding:"required"`
	Ldap     bool   `json:"ldap" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
