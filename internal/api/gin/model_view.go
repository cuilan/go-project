package gin

// ---------------- request ----------------

type UserRegisterModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ---------------- response ----------------

type HealthResp struct {
	Satellite string `json:"satellite"`
}
