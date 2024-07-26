package packets

type ClientMetaData struct {
	HandShakeResponse
	Id string `json:"id"`
}

type HandShakeResponse struct {
	Username     string `json:"username"`
	System       string `json:"system"`
	Architecture string `json:"architecture"`
	Privilages   string `json:"privialges"`
}

type HandShake struct {
	Id string `json:"id"`
}
