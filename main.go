package main
import (
"fmt"
"log"
"net/http"
"os"
)
func main(){
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","text/html")
fmt.Fprint(w, `<!DOCTYPE html><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>V31 OK</title><style>body{background:#1a0b2e;color:#fff;font-family:sans-serif;display:flex;justify-content:center;padding:12px}.c{width:100%;max-width:400px}.card{background:rgba(255,255,255,.1);border-radius:20px;padding:16px}input,select{width:100%;padding:12px;border-radius:10px;border:none;margin-top:6px;font-weight:700;font-size:18px}.btn{width:100%;padding:16px;border-radius:12px;border:none;font-weight:900;margin-top:10px;font-size:18px}.y{background:#FFD60A;color:#000}.v{background:#8E2DE2;color:#fff}.r{margin-top:12px;background:#fff;color:#000;padding:12px;border-radius:10px;font-weight:700;text-align:center;font-size:16px;min-height:50px}</style></head><body><div class=c><h1 style="text-align:center;color:#FFD60A">V31 OK</h1><p style="text-align:center">Pi Pay CI - 1 Pi=800F | Gain 200F</p><div class=card><label>NUMERO</label><input id=num value="0544535683"><label>Pi</label><input id=pi type=number value=2 oninput="upd()"><label>OPERATEUR</label><select id=op><option>WAVE</option><option>MTN</option><option>ORANGE</option><option>MOOV</option></select><button class="btn y" onclick="doPay()">PAYER <span id=pt>2</span> Pi -> <span id=ft>1600</span>F</button><button class="btn v" onclick="doPay()">Payer avec Pi Browser</button><div id=res class=r>Pret - Clique violet pour tester</div></div></div><script>function upd(){let p=parseFloat(document.getElementById('pi').value)||0;document.getElementById('pt').innerText=p;document.getElementById('ft').innerText=p*800;document.getElementById('res').innerHTML=p+' Pi -> '+(p*800)+'F sur '+document.getElementById('num').value}function doPay(){let p=parseFloat(document.getElementById('pi').value)||0;let n=document.getElementById('num').value;let o=document.getElementById('op').value;let f=p*800;let g=p*200;let msg='✅ SIMULATION REUSSIE!<br><br>'+p+' Pi -> '+f+'F<br>Numero: '+n+'<br>Operateur: '+o+'<br>Gain: '+g+'F<br><br>Pour vrai paiement Pi, ouvre dans Pi Browser';document.getElementById('res').innerHTML=msg;alert('SIMULATION OK: '+p+' Pi -> '+f+'F sur '+n)}upd();</script></body></html>`)
})
http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){w.Write([]byte("V31 OK"))})
pt:=os.Getenv("PORT"); if pt==""{pt="10000"}; log.Println("V31",pt); log.Fatal(http.ListenAndServe(":"+pt,nil))
}
