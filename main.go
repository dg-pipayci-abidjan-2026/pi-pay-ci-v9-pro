package main
import ("log";"net/http";"os")
func main(){
p:=os.Getenv("PORT")
if p==""{p="10000"}
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","text/html; charset=utf-8")
w.Write([]byte("<html><head><meta charset=utf-8><meta name=viewport content='width=device-width'><title>V20 DG 2400F=2Pi</title><style>body{background:#0f0f1a;color:#fff;font-family:Arial;text-align:center;padding:20px}h1{color:#00ff88}.c{background:#1a1a2e;border:2px solid #00ff88;border-radius:20px;padding:20px;max-width:380px;margin:auto}input{width:90%;padding:12px;border-radius:10px;border:none;margin:8px;text-align:center}#r{background:#111;border:1px solid #00ff88;border-radius:10px;padding:12px;color:#00ff88;font-weight:bold}</style></head><body><h1>V20 DG 2400F=2Pi</h1><div class=c><p>MODEL: 0.25 Pi + 200F</p><input id=f placeholder='FCFA 2400' oninput='let v=(parseFloat(this.value)||0);let pi=(v-200)/1000;if(pi<0)pi=0;document.getElementById(\"r\").innerHTML=\"<b>\"+v+\"F -200F=\"+(v-200)+\"F = \"+pi.toFixed(2)+\" Pi</b><br><br>✅ ENVOIE: \"+pi.toFixed(2)+\" Pi\"'><div id=r>2400F => 2 Pi</div><p>San-Pedro 0574098776</p></div></body></html>"))
})
log.Println("V20 DG LIVE "+p)
log.Fatal(http.ListenAndServe(":"+p,nil))
}
