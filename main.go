package main
import ("log";"net/http";"os")
func main(){
p:=os.Getenv("PORT")
if p==""{p="10000"}
http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
w.Write([]byte(`<html><body style=background:#0f0f1a;color:#fff;text-align:center;padding:20px><h1 style=color:#00ff88>V20 DG 2400F=2Pi LIVE</h1><p>MODEL 0.25Pi+200F</p><p>San-Pedro 0574098776</p></body></html>`))
})
log.Println("V20 LIVE "+p)
log.Fatal(http.ListenAndServe(":"+p,nil))
}
