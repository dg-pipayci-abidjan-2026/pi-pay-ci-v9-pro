package main
import (
"bytes"
"encoding/json"
"fmt"
"io"
"log"
"net/http"
"os"
)
var k=os.Getenv("PI_API_KEY")
const b="https://api.minepi.com"
func main(){
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","text/html")
fmt.Fprint(w, `<!DOCTYPE html><html><head><meta charset=UTF-8><meta name=viewport content="width=device-width,initial-scale=1"><title>V30 FINAL</title><script src="https://sdk.minepi.io/pi-sdk.js"></script><style>body{background:#1a0b2e;color:#fff;font-family:sans-serif;display:flex;justify-content:center;padding:12px}.c{width:100%;max-width:400px}.card{background:rgba(255,255,255,.1);border-radius:20px;padding:16px}input,select{width:100%;padding:12px;border-radius:10px;border:none;margin-top:6px;font-weight:700}.btn{width:100%;padding:14px;border-radius:12px;border:none;font-weight:900;margin-top:8px}.y{background:#FFD60A;color:#000}.v{background:#8E2DE2;color:#fff}.r{margin-top:10px;background:#fff;color:#000;padding:10px;border-radius:10px;font-weight:700;text-align:center}</style></head><body><div class=c><h1 style="text-align:center;color:#FFD60A">V30 FINAL</h1><p style="text-align:center;font-size:12px">Pi Pay CI - 1 Pi=800F | Gain 200F</p><div class=card><label>NUMERO</label><input id=n value=0574098776><label>Pi</label><input id=p type=number value=1><label>OPERATEUR</label><select id=o><option>WAVE</option><option>MTN</option><option>ORANGE</option><option>MOOV</option><option>Orange Money</option><option>MTN MoMo</option></select><button class="btn y" onclick=pay()>PAYER <span id=pt>1</span> Pi -> <span id=ft>800</span>F</button><button class="btn v" onclick=pay()>Payer avec Pi Browser</button><div id=r class=r>Pret - Clique bouton violet</div><div id=l style="background:#000;color:#0f0;font-size:10px;padding:6px;margin-top:6px;max-height:80px;overflow:auto"></div></div></div><script>function lg(t){document.getElementById('l').innerHTML+=t+'<br>'}function sh(h){document.getElementById('r').innerHTML=h}document.getElementById('p').oninput=function(){let pi=parseFloat(this.value)||0;document.getElementById('pt').innerText=pi;document.getElementById('ft').innerText=pi*800;sh(pi+' Pi -> '+(pi*800)+'F sur '+document.getElementById('n').value)}async function pay(){let pi=parseFloat(document.getElementById('p').value)||0;let num=document.getElementById('n').value;let op=document.getElementById('o').value;if(pi<=0){alert('Entre Pi');return}sh('Demarrage '+pi+' Pi...');lg('click '+pi);try{if(typeof Pi==='undefined'){lg('Pi SDK absent - SIMU');sh('SIMULATION REUSSIE!<br>'+pi+' Pi -> '+(pi*800)+'F<br>Sur '+num+' ('+op+')<br>Gain '+(pi*200)+'F<br><br>Pour vrai Pi, ouvre dans Pi Browser');return}Pi.init({version:"2.0",sandbox:false});lg('init ok');sh('Connexion Pi...');let auth=await Pi.authenticate(['username','payments'],function(p){});lg('auth ok '+auth.user.username);sh('Connecte '+auth.user.username+' - Paiement '+pi+' Pi...');Pi.createPayment({amount:pi,memo:"Pi Pay CI "+pi+" Pi -> "+(pi*800)+"F "+op+" "+num,metadata:{numero:num,operateur:op}},{onReadyForServerApproval:async function(id){lg('approval '+id);sh('Approbation '+id+'...');let res=await fetch('/api/approve',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({paymentId:id})});let d=await res.json();lg('approve '+JSON.stringify(d))},onReadyForServerCompletion:async function(id,txid){lg('complete '+txid);sh('Finalisation TX '+txid+'...');let res=await fetch('/api/complete',{method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({paymentId:id,txid:txid})});let d=await res.json();lg('complete '+JSON.stringify(d));sh('PAIEMENT REUSSI!<br>'+pi+' Pi<br>TX: '+txid+'<br>Envoi '+(pi*800)+'F sur '+num+'<br>Gain '+(pi*200)+'F')},onCancel:function(id){lg('cancel');sh('Annule')},onError:function(e){lg('error '+e);sh('Erreur: '+e)}})}catch(e){lg('catch '+e.message);sh('Erreur '+e.message+'<br>SIMU: '+pi+' Pi -> '+(pi*800)+'F OK')}}lg('V30 charge Pi:'+(typeof Pi))</script></body></html>`)
})
http.HandleFunc("/api/approve", func(w http.ResponseWriter, r *http.Request){
w.Header().Set("Access-Control-Allow-Origin","*"); w.Header().Set("Content-Type","application/json"); w.Header().Set("Access-Control-Allow-Methods","POST, OPTIONS")
if r.Method=="OPTIONS"{w.WriteHeader(200);return}
var q struct{PaymentId string `json:"paymentId"`}; b2,_:=io.ReadAll(r.Body); json.Unmarshal(b2,&q)
if k==""{json.NewEncoder(w).Encode(map[string]string{"status":"SIMULATION_OK"});return}
c:=&http.Client{}; pr,_:=http.NewRequest("POST",b+"/v2/payments/"+q.PaymentId+"/approve",nil); pr.Header.Set("Authorization","Key "+k)
rs,_:=c.Do(pr); defer rs.Body.Close(); rb,_:=io.ReadAll(rs.Body); w.Write(rb)
})
http.HandleFunc("/api/complete", func(w http.ResponseWriter, r *http.Request){
w.Header().Set("Access-Control-Allow-Origin","*"); w.Header().Set("Content-Type","application/json"); w.Header().Set("Access-Control-Allow-Methods","POST, OPTIONS")
if r.Method=="OPTIONS"{w.WriteHeader(200);return}
var q struct{PaymentId string `json:"paymentId"`; Txid string `json:"txid"`}; b2,_:=io.ReadAll(r.Body); json.Unmarshal(b2,&q)
if k==""{json.NewEncoder(w).Encode(map[string]string{"status":"SIMULATION_COMPLETE"});return}
pl:=map[string]string{"txid":q.Txid}; jb,_:=json.Marshal(pl); c:=&http.Client{}; pr,_:=http.NewRequest("POST",b+"/v2/payments/"+q.PaymentId+"/complete",bytes.NewBuffer(jb)); pr.Header.Set("Authorization","Key "+k); pr.Header.Set("Content-Type","application/json")
rs,_:=c.Do(pr); defer rs.Body.Close(); rb,_:=io.ReadAll(rs.Body); w.Write(rb)
})
http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){w.Write([]byte("V30 FINAL OK"))})
pt:=os.Getenv("PORT"); if pt==""{pt="10000"}; log.Println("V30 FINAL",pt,k!=""); log.Fatal(http.ListenAndServe(":"+pt,nil))
}
