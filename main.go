package main

import (
"log"
"net/http"
)

func main() {
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "text/html; charset=utf-8")
html := `<!DOCTYPE html><html lang="fr"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Pi Pay CI V11 - Pi Réel</title><script src="https://sdk.minepi.com/pi-sdk.js"></script><style>body{font-family:Arial;background:#0a192f;color:white;text-align:center;padding:20px}.card{background:#1a2332;padding:25px;border-radius:20px;max-width:420px;margin:auto;border:2px solid #f59e0b}button{width:100%;padding:15px;margin:7px 0;border:none;border-radius:12px;font-size:17px;font-weight:bold;cursor:pointer}.pi{background:#f59e0b;color:black}.wave{background:#00c851;color:white}.om{background:#ff9800;color:black}.mtn{background:#ffcc00;color:black}.moov{background:#0066cc;color:white}.cie{background:#ff6f00;color:white}.sodeci{background:#0099ff;color:white}.canal{background:black;color:white;border:1px solid white}#log{background:#000;color:#00ff88;padding:10px;border-radius:10px;margin-top:15px;font-size:12px;min-height:40px}</style></head><body><h1 style="color:#f59e0b">💎 Pi Pay CI V11 💎</h1><p>Pi Network Paiement Réel - DG 2026</p><div class="card"><h2 id="status">🚀 Connexion Pi...</h2><button class="pi" onclick="payPi()">💜 Payer avec Pi - RÉEL</button><button class="wave" onclick="alert('Wave bientôt')">💚 Payer Wave</button><button class="om" onclick="alert('OM bientôt')">🧡 Orange Money</button><button class="mtn" onclick="alert('MTN bientôt')">💛 MTN MoMo</button><button class="moov" onclick="alert('Moov bientôt')">💙 Moov Money</button><button class="cie" onclick="alert('CIE bientôt')">⚡ Payer CIE</button><button class="sodeci" onclick="alert('SODECI bientôt')">💧 Payer SODECI</button><button class="canal" onclick="alert('Canal+ bientôt')">📺 Canal+ CI</button><div id="log">Initialisation Pi SDK...</div><p style="font-size:10px;margin-top:15px">pi-pay-ci-v9-pro-1.onrender.com<br>V11 Pi Réel - DG Abidjan</p></div><script>
const Pi = window.Pi;
var logDiv = document.getElementById('log');
var statusH2 = document.getElementById('status');
function log(msg){logDiv.innerHTML += '<br>> '+msg; console.log(msg);}
Pi.init({version:"2.0", sandbox:true});
log("Pi SDK Init OK - Sandbox");
statusH2.innerHTML="✅ Pi SDK Prêt - Clique Pi";
async function payPi(){
try{
log("Authentification Pi...");
statusH2.innerHTML="🔐 Connexion Pi Network...";
const auth = await Pi.authenticate(['username','payments'], function(p){});
log("Auth OK: "+auth.user.username);
statusH2.innerHTML="👋 Salut "+auth.user.username;
const payment = await Pi.createPayment({
amount: 1,
memo: "Paiement Pi Pay CI - Test DG",
metadata: {service:"PiPayCI", city:"Abidjan"}
}, {
onReadyForServerApproval: function(paymentId){log("Paiement ID: "+paymentId+" - Approbation serveur..."); fetch('/approve', {method:'POST', body: paymentId});},
onReadyForServerCompletion: function(paymentId, txid){log("Transaction TXID: "+txid+" - Complétion..."); statusH2.innerHTML="✅ Paiement Réussi! TX: "+txid;},
onCancel: function(paymentId){log("Paiement annulé: "+paymentId); statusH2.innerHTML="❌ Annulé";},
onError: function(error, payment){log("Erreur: "+error); statusH2.innerHTML="❌ Erreur Pi";},
});
}catch(e){log("Erreur Auth: "+e.message); statusH2.innerHTML="❌ Ouvre dans Pi Browser"; alert("Ouvre ce lien dans Pi Browser pour payer avec Pi: https://pi-pay-ci-v9-pro-1.onrender.com");}
}
log("Prêt. Ouvre dans Pi Browser pour test réel.");
</script></body></html>`
w.Write([]byte(html))
})
http.HandleFunc("/approve", func(w http.ResponseWriter, r *http.Request){ w.Write([]byte("approved")) })
log.Println("V11 PI REEL LIVE sur :3000")
log.Fatal(http.ListenAndServe(":3000", nil))
}
