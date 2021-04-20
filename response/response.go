package response

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Method  string      `json:"method"`
	Data    interface{} `json:"data"`
	Extra   interface{} `json:"type"`
}

func (r *Response) Success(m string, d interface{}) {
	r.Status = 201
	r.Message = m
	r.Data = d
	// r.Extra = reflect.TypeOf(d)
}
func (r *Response) ServerError(m string, d interface{}) {
	r.Status = 400
	r.Data = m
	r.Message =m

}

func (r *Response) InsufficientCredtials(m string, d interface{}) {
	r.Status = 401
	r.Message = m
	r.Data = d
}

func (r *Response) InvalidCrediantials(m string, d interface{}) {
	r.Status = 402
	r.Message = m
	r.Data = d
}

func (r *Response) AlreadyPresent(m string, d interface{}) {
	r.Status = 403
	r.Message = m
	r.Data = d
}
func (r *Response) NosuchDoc(m string, d interface{}) {
	r.Status = 404
	r.Message = m
	r.Data = d
}

func (r *Response) ErrorUpdateDoc(m string, d interface{}) {
	r.Status = 405
	r.Message = m
	r.Data = d
}

func (r *Response) ErrorDeleteDoc(m string, d interface{}) {
	r.Status = 406
	r.Message = m
	r.Data = d
}
