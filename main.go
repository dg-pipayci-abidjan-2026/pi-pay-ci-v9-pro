package main
import ("log";"net/http";"os")
func main(){
p:=os.Getenv("PORT")
if p==""{p="10000"}
http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
w.Write([]byte(`<html><meta charset=utf-8><meta name=viewport content=width=device-width><title>V21</title><style>body{background:#0f0f1a;color:#fff;text-align:center;font-family:Arial;padding:8px}input,button{padding:10px;border-radius:8px;border:2px solid #0f8;width:90%;max-width:320px;margin:4px;font-size:16px}button{background:#00ff88;font-weight:bold}.c{background:#1a1a2e;border:2px solid #00ff88;border-radius:14px;padding:10px;max-width:360px;margin:10px auto}b{color:#00ff88}#rc{background:#fff;color:#000;border-radius:10px;padding:8px;margin-top:8px;display:none;text-align:left;font-size:13px}</style><body><h3 style=color:#00ff88>V21 DG 2400F=2Pi QR+RECU</h3><div class=c><input id=f type=number placeholder="FCFA ex:2400" oninput="c()"><p>Reçois:<b id=p>0 Pi</b> Tiroir:<b id=t>50 Pi</b></p><p>MODEL:<b>0.25Pi+200F</b> Solde:<b>1863 Pi</b></p><img id=q style=display:none;margin:auto><br><button onclick="r()">🧾 Generer Recu</button><div id=rc></div><p style=font-size:11px>13 Svc: OM Wave MTN Pi | San-Pedro 0574098776</p></div><script>function c(){let v=f.value||0,pi=v/1200;p.innerText=pi.toFixed(2)+' Pi';t.innerText=(50+pi*0.25).toFixed(2)+' Pi';let qr=document.getElementById('q');if(v>0){qr.style.display='block';qr.src='https://api.qrserver.com/v1/create-qr-code/?size=120x120&data=PiPayCI-'+v+'F-'+pi.toFixed(2)+'Pi-DG'}else qr.style.display='none'}function r(){let v=f.value||0,pi=v/1200,d=new Date().toLocaleString();rc.style.display='block';rc.innerHTML='<b>RECU PiPay CI V21</b><br>'+d+'<br>Montant:'+v+'F<br>Pi:'+pi.toFixed(2)+' Pi<br>Tiroir:'+t.innerText+'<br>Solde:1863 Pi<br>San-Pedro 0574098776<br>✅2400F=2Pi Merci!'}</script></body></html>`))
})
log.Println("V21 "+p)
log.Fatal(http.ListenAndServe(":"+p,nil))
}
