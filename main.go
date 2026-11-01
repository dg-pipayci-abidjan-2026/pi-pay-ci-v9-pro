package main
import ("log";"net/http";"os")
func main(){
p:=os.Getenv("PORT")
if p==""{p="10000"}
http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
w.Write([]byte(`<html><meta charset=utf-8><meta name=viewport content=width=device-width><title>V20 DG</title><style>body{background:#0f0f1a;color:#fff;text-align:center;font-family:Arial;padding:10px}input{padding:12px;border-radius:10px;border:2px solid #00ff88;width:90%;max-width:300px;font-size:18px}.c{background:#1a1a2e;border:2px solid #00ff88;border-radius:16px;padding:12px;max-width:360px;margin:12px auto}b{color:#00ff88}</style><body><h2 style=color:#00ff88>V20 DG 2400F=2Pi PRO</h2><div class=c><input id=f type=number placeholder="FCFA ex:2400" oninput="c()"><p>Tu reçois: <b id=p>0 Pi</b></p><p>MODEL: <b>0.25Pi+200F</b></p><p>Tiroir: <b id=t>50 Pi</b> | Solde: <b>1863 Pi</b></p><p>San-Pedro 0574098776</p><p id=r></p></div><script>function c(){let v=f.value||0;let pi=v/1200;p.innerText=pi.toFixed(2)+' Pi';let tir=50+pi*0.25;t.innerText=tir.toFixed(2)+' Pi';if(v>=2400)r.innerHTML='✅ 2400F=2Pi VALIDÉ<br>🎁 Bonus 10Pi +13 Services';else r.innerHTML='Saisis montant'}</script></body></html>`))
})
log.Println("V20 PRO "+p)
log.Fatal(http.ListenAndServe(":"+p,nil))
}
