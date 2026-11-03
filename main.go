package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `
<!DOCTYPE html>
<html lang="fr">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Pi Pay CI V29.1 PRO</title>
<script src="https://sdk.minepi.com/pi-sdk.js"></script>
<style>
*{margin:0;padding:0;box-sizing:border-box;font-family:'Segoe UI',sans-serif}
body{background:linear-gradient(135deg,#1a0b2e 0%,#2d1b4e 50%,#0f2027 100%);min-height:100vh;display:flex;justify-content:center;padding:18px}
.container{width:100%;max-width:420px}
.header{text-align:center;color:white;margin:18px 0}
.header h1{font-size:24px;font-weight:900}
.header p{opacity:0.8;font-size:12px;margin-top:6px;line-height:1.4}
.card{background:rgba(255,255,255,0.08);backdrop-filter:blur(20px);border:1px solid rgba(255,255,255,0.15);border-radius:24px;padding:18px;box-shadow:0 20px 40px rgba(0,0,0,0.4)}
.tabs{display:flex;gap:8px;margin-bottom:16px}
.tab{flex:1;padding:10px 8px;border-radius:14px;border:none;font-weight:800;cursor:pointer;font-size:12px;line-height:1.2;transition:0.2s}
#tabMine{background:#FFD60A;color:#000}
#tabNoMine{background:#00E676;color:#000}
#tabMine.active{box-shadow:0 0 18px #FFD60A;transform:scale(1.02)}
#tabNoMine.active{box-shadow:0 0 18px #00E676;transform:scale(1.02)}
.input-group{margin-bottom:11px}
.input-group label{color:#c9b3ff;font-size:10px;font-weight:800;margin-left:8px}
.input-group input,.input-group select{width:100%;padding:13px 15px;border-radius:13px;border:1px solid rgba(255,255,255,0.15);background:rgba(255,255,255,0.95);font-size:14px;font-weight:600;outline:none;margin-top:3px}
.btn{width:100%;padding:14px;border-radius:13px;border:none;font-weight:900;font-size:14px;cursor:pointer;margin-top:8px}
.btn-mine{background:linear-gradient(135deg,#FFD60A,#FFA500);color:#000}
.btn-nomine{background:linear-gradient(135deg,#00E676,#00C853);color:#000}
.btn-pi{background:linear-gradient(135deg,#8E2DE2,#4A00E0);color:white;margin-top:6px}
.result{margin-top:14px;padding:12px;border-radius:13px;background:rgba(255,255,255,0.95);color:#000;font-weight:700;text-align:center;font-size:13px;display:none;line-height:1.4}
.result.show{display:block}
.badge{background:#FFD60A;padding:2px 7px;border-radius:10px;font-size:10px}
.footer{color:rgba(255,255,255,0.5);text-align:center;font-size:10px;margin-top:12px}
</style>
</head>
<body>
<div class="container">
  <div class="header">
    <h1>V29.1 ULTRA PRO</h1>
    <p>Pi Pay CI 🇨🇮<br>1 Pi = 1000F | Achat 800F | Vente 1200F | Gain 200F</p>
  </div>
  <div class="card">
    <div class="tabs">
      <button id="tabMine" class="tab active" onclick="switchTab('mine')">J'AI DÉJÀ MINÉ<br>JE VEUX CFA</button>
      <button id="tabNoMine" class="tab" onclick="switchTab('nomine')">J'AI PAS MINÉ<br>JE VEUX Pi</button>
    </div>

    <div class="input-group">
      <label>NUMÉRO / ID FACTURE</label>
      <input id="numero" type="text" placeholder="0574098776 / CIE..." value="0574098776">
    </div>

    <div id="mineFields">
      <div class="input-group">
        <label>COMBIEN DE Pi TU AS ?</label>
        <input id="piAmount" type="number" placeholder="Ex: 5" oninput="calcMine()">
      </div>
    </div>

    <div id="nomineFields" style="display:none">
      <div class="input-group">
        <label>COMBIEN FCFA TU AS ?</label>
        <input id="fcfaAmount" type="number" placeholder="Ex: 1200, 5000, 10000" oninput="calcNoMine()">
      </div>
    </div>

    <div class="input-group">
      <label>OPÉRATEUR</label>
      <select id="operateur">
        <option>--Opérateur--</option>
        <option>MTN</option><option>ORANGE</option><option>MOOV</option><option>WAVE</option>
        <option>MTN MoMo</option><option>Orange Money</option><option>Moov Money</option>
        <option>CANAL+</option><option>CIE</option><option>SODECI</option><option>SOTRA</option>
      </select>
    </div>

    <div class="input-group">
      <label>TYPE</label>
      <select id="typeOp">
        <option>--Type--</option>
        <option>1 INTERNET</option><option>2 APPEL</option><option>3 MOMO / WAVE Direct</option><option>4 FACTURE</option>
      </select>
    </div>

    <div id="btnMine">
      <button class="btn btn-mine" onclick="payMine()">💛 PAYER <span id="minePiText">0</span> Pi → <span id="mineFcfaText">0</span>F CFA</button>
      <button class="btn btn-pi" onclick="connectPi()">🔗 Payer avec Pi Browser</button>
    </div>

    <div id="btnNoMine" style="display:none">
      <button class="btn btn-nomine" onclick="payNoMine()">💚 PAYER <span id="nomineFcfaText">0</span>F → <span id="nominePiText">0</span> Pi</button>
      <button class="btn btn-pi" onclick="connectPi()">🔗 Recevoir Pi</button>
    </div>

    <div id="result" class="result"></div>
    <div class="footer">Taux fixe • Gain 200F/Pi • Pi Browser Ready • V29.1</div>
  </div>
</div>

<script>
let tab='mine'
let piConnected=false
function switchTab(t){
  tab=t
  document.getElementById('tabMine').classList.toggle('active',t==='mine')
  document.getElementById('tabNoMine').classList.toggle('active',t==='nomine')
  document.getElementById('mineFields').style.display=t==='mine'?'block':'none'
  document.getElementById('nomineFields').style.display=t==='nomine'?'block':'none'
  document.getElementById('btnMine').style.display=t==='mine'?'block':'none'
  document.getElementById('btnNoMine').style.display=t==='nomine'?'block':'none'
}
function op(){return document.getElementById('operateur').value+' '+document.getElementById('typeOp').value}
function show(html){let r=document.getElementById('result');r.innerHTML=html;r.classList.add('show')}
function calcMine(){
  let pi=parseFloat(document.getElementById('piAmount').value)||0
  let fcfa=pi*800
  document.getElementById('minePiText').innerText=pi
  document.getElementById('mineFcfaText').innerText=fcfa
  if(pi>0) show('💛 J\'AI DÉJÀ MINÉ: '+pi+' Pi<br>Tu reçois <b>'+fcfa+'F CFA</b> sur '+document.getElementById('numero').value+' ('+op()+')<br><span class=badge>Mon gain 200F/Pi = '+(pi*200)+'F</span>')
}
function calcNoMine(){
  let f=parseFloat(document.getElementById('fcfaAmount').value)||0
  let pi=(f/1200).toFixed(2)
  document.getElementById('nomineFcfaText').innerText=f
  document.getElementById('nominePiText').innerText=pi
  if(f>0) show('💚 J\'AI PAS MINÉ: '+f+'F<br>Tu reçois <b>'+pi+' Pi</b> dans wallet Pi<br><span class=badge>Mon gain 200F/Pi</span>')
}
async function connectPi(){
  try{
    if(typeof Pi==='undefined'){alert('Mode test OK. Dans Pi Browser ce sera vrai paiement Pi');piConnected=true;return}
    const auth=await Pi.authenticate(['username','payments'],()=>{})
    piConnected=true; alert('Connecté: '+auth.user.username)
  }catch(e){alert(e)}
}
async function payMine(){
  let pi=parseFloat(document.getElementById('piAmount').value)||0
  if(pi<=0){alert('Entre Pi');return}
  if(typeof Pi!=='undefined' && piConnected){
    try{
      Pi.createPayment({amount:pi,memo:"Echange "+pi+" Pi -> "+(pi*800)+"F "+op(),metadata:{type:"mine_to_cfa"}},{
        onReadyForServerApproval:id=>fetch('/api/approve',{method:'POST',body:JSON.stringify({id})}),
        onReadyForServerCompletion:(id,txid)=>show("✅ PAIEMENT RÉUSSI! "+pi+" Pi TX: "+txid+"<br>Je t'envoie "+(pi*800)+"F sur "+document.getElementById('numero').value+" maintenant!"),
        onCancel:id=>alert('Annulé'), onError:e=>alert(e)
      })
    }catch(e){alert(e)}
  }else{
    show("🔥 SIMULATION: "+pi+" Pi → "+(pi*800)+"F vers "+document.getElementById('numero').value+"<br>Dans Pi Browser = vrai Pi!")
  }
}
function payNoMine(){
  let f=parseFloat(document.getElementById('fcfaAmount').value)||0
  if(f<=0){alert('Entre FCFA');return}
  let pi=(f/1200).toFixed(2)
  show("🟢 Reçu "+f+"F sur WAVE/MoMo.<br>Je t'envoie "+pi+" Pi au wallet Pi.<br>Numéro: "+document.getElementById('numero').value+"<br>✅ Mon gain 200F/Pi validé")
}
</script>
</body>
</html>
	`)
	})
	port := os.Getenv("PORT")
	if port == "" { port = "10000" }
	log.Println("V29.1 with new buttons running on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
