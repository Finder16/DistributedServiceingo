package main

import(

"log"
"github.com/jaeyeong16/DistributedServiceingo/tree/main/server"


)


func main(){

        srv := server.NewHTTPServer(":8080")
        log.Fatal(srv.ListenAndServe())

}
