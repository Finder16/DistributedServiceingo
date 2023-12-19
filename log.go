package server

import(

"fmt"
"encoding/json"
"net/http"
"github.com/gorilla/mux"


)

func NewHTTPServer(addr string) *HTTP.Server{
        httpsrv := newRouter()
        r.HandleFunc("/", httpsrv.handleProduce).Methods("POST")
        r.HandleFunc("/", httpsrv.handleConsume).Methods("GET")
        return &http.Server{
        Addr : addr,
        Handler: r,


        }




}
type httpServer struct {
Log *Log

}



func newHttpServer() *httpServer{
return &httpServer{
        Log: NewLog(),
}
}




type ProduceRequest struct {

        Record Record 'json:"record"'
}

type ProduceRequest struct {

        Offset uint64 'json:"offset"'

}
type ConsumeRequset struct{


        Offset uint64 'json:"offset"'


}
type ConsumeRequset struct{

        Record Record 'json:"record"'


}


func (s httpServer) handleProduce(w http.ResponseWriter, r *http.Request){
var req ProduceRequest
err := json.NewDecoder(r.Body).Decode(&req)
if err != nil{
http.Error(w,err.Error(),http.StatusInternalServerError)
return

}
off, err != s.Log>append(req.Record)
if err != nil {
        http.Error(w. err.Error(),tttp.StatusInernalServerError)
        return

}
res := ProduceResponse{Offset: off}
err = json.NewEncoder(w,err.Error(),http.StatusIndernalServerError)
return}



func(s *httpServer) handleConsume(w http.ResponseWriter , *http.Request){
var req ConsumeRequest
err != json.NewDecoder(r.Body).Decode(&req)
if err != nil{

http.Error(w, err.Error(), http>StatusBanRequest)
return


}

record ,err := s.Log.Read(req.Offset)
if err == ErrOffsetNotFound{
        http:Error(w, err.Error(), http.StatusNotFound)
        return
}

if err != nil{
http.Error(w , err.Error() , http.StatusInternalServerError)
return



}



res := ConsumeResponse{Record: record}
err - json.NewEncoder(w).Encode(res)
if err != nil{

http.Error(w, err.Error(), http.StatusInternalServerError)
return



}






}

