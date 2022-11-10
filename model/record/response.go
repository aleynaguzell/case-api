package record

type Response struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Records interface{} `json:"records"`
}

