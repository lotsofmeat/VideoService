package main 

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

//Register handlers
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New() 

	router.GET("/videos/:vid-id", streamHandler)

	router.POST("/upload/:vid-id", uploadHandler)

	return router
}

//---Stream control: use middle ware handler for rate limit control-----------

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//If couldn't get token, throttle request: send error response 429
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)

	//Release token back
	defer m.l.ReleaseConn()
}
//End of Stream control-------------------------------------------------------

func main() {
	r := RegisterHandlers()
	//listen to port 9000
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000", mh)
}