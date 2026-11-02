package main

import (
"fmt"
"net/http"
)

func main() {
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "text/html; charset=utf-8")
fmt.Fprint(w, `<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<script src="https://sdk.minepi.com/pi-sdk.js"></script>
<script>Pi.init({version:"2.0"})</script>
<style>
body{background:#0f172a;color:white;font-family:sans-serif;text-align:center;padding:10px}
button{background:#1e293b;color:white;border:1px solid #334155;padding:12px;margin:4px;border-radius:10px;width:30%;font-size:12px}
button.active{background:#facc15;color:black;font-weight:bold}
input{width:90%;padding:15px;margin:8px;border-radius:15px;text-align:center;font-size:18px}
.green{background:#22c55e;color:black;font-weight:bold;width:92%;padding:15px;border-radius:15px;font-size:16px;margin:5px}
.qrbox{background:white;padding:10px;display:inline-block;border-radius:10px;margin:10px;border:4px solid #22c55e}
#recu{background:white;color:black;padding:15px;border-radius:10px;margin-top:15px;text-align:left;white-space:pre-line}
</style>
</head>
<body>
<h3>V24 ULTIME PiPay CI 🔥</h3>
<div>
<button onclick="setS('MTN')" id="MTN" class="active">MTN</button>
<button onclick="setS('ORANGE')" id="ORANGE">ORANGE</button>
<button onclick="setS('MOOV')" id="MOOV">MOOV</button>
<button onclick="setS('WAVE')" id="WAVE">WAVE</button>
<button onclick="setS('CIE')" id="CIE">CIE</button>
<button onclick="setS('SOTRA')" id="SOTRA">SOTRA</button>
<button onclick="setS('CANAL')" id="CANAL">CANAL+</button>
<button onclick="setS('USDT')" id="USDT">USDT</button>
<button onclick="setS('CARTE')" id="CARTE">CARTE</button>
<button onclick="setS('SIDECI')" id="SIDECI">SIDECI</button>
</div>
<button onclick="setM('fiat')" id="btnFiat" class="active" style="width:45%">FCFA->Pi</button>
<button onclick="setM('pi')" id="btnPi" style="width:45%">Pi->MTN 50=40000F</button>
<input id="tel" value="0574098776">
<input id="montant" value="50" oninput="calc()">
<div id="info" style="color:#facc15;font-weight:bold;margin:10px;font-size:18px"></div>
<div id="qr" class="qrbox"></div><br>
<button class="green" onclick="payerPi()">Payer avec Pi Browser - Securise</button>
<button class="green" onclick="genRecu()">Generer Recu + QR Noir</button>
<div id="recu"></div>
<script>
let service='MTN'; let mode='pi';
function setS(s){service=s; document.querySelectorAll('button').forEach(b=>b.classList.remove('active')); document.getElementById(s).classList.add('active'); calc()}
function setM(m){mode=m; calc(); document.getElementById('btnFiat').classList.toggle('active', m=='fiat'); document.getElementById('btnPi').classList.toggle('active', m=='pi')}
function calc(){
let mont=parseFloat(document.getElementById('montant').value)||0;
let txt="";
if(mode=='pi'){ let f=Math.round(mont*796); txt='Tu envoies: '+f+'F via MTN | '+mont+' Pi = '+(mont*800)+'F'; }
else { let pi=(mont/800).toFixed(2); txt=pi+' Pi + frais'; }
document.getElementById('info').innerText=txt;
}
async function payerPi(){
try{
const auth=await Pi.authenticate(['payments'], function(){});
alert('Pi Browser Connecte: '+auth.user.username+' - Paiement securise sans 24 mots!');
genRecu();
}catch(e){ alert('Ouvre ce site dans Pi Browser: pi://ci-v9-pro-1.onrender.com'); }
}
function genRecu(){
let tel=document.getElementById('tel').value; let mont=document.getElementById('montant').value;
let info=document.getElementById('info').innerText;
let recu='RECU PiPay CI V24 '+ (mode=='pi'?'Pi->MTN':'FCFA->Pi') +'\n'+new Date().toLocaleString()+'\nService: '+service+'\nTel: '+tel+'\nMontant: '+mont+(mode=='pi'?' Pi':'F')+'\n'+info+'\nMerci! 0574098776 San-Pedro';
document.getElementById('recu').innerText=recu;
document.getElementById('qr').innerHTML='<img src="https://api.qrserver.com/v1/create-qr-code/?size=250x250&color=000000&bgcolor=FFFFFF&data='+encodeURIComponent(recu)+'" style="border:3px solid black">';
calc();
}
calc(); genRecu();
</script>
<p>SIDECI CIE SOTRA STAR LOYE ECOLE MTN ORANGE MOOV WAVE CANAL USDT CARTE QR BOUTIQUE MAGAZINE MARQUES APPEL INTERNET</p>
</body>
</html>
`)
})
http.ListenAndServe(":10000", nil)
}
