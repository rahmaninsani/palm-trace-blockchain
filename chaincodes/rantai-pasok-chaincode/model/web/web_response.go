package web

type WebResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty" metadata:",optional"`
	Data    interface{} `json:"data,omitempty" metadata:",optional"`
}
