package main

import (
	"net/http"
	"os"
)

func main() {
	html := `<!DOCTYPE html> <html> <head> <meta charset="utf-8"> <meta name="viewport" content="width=device-width, initial-scale=1"> <title>Pi Pay CI V32 - REAL</title> <script src="https://sdk.minepi.com/pi-sdk.js"></script> <style> body{background:#1a0b2e;color:white;font-family:sans-serif;text-align:center;margin:0;padding:20px} h1{color:#FFD700;font-size:32px;margin-top:10px} .card{background:#2d1b4e;padding:20px;border-radius:20px;max-width:400px;margin:20px auto} input,select{width:100%;padding:15px;border-radius:10px;border:none;margin:10px 0;font-size:18px;font-weight:bold} .btn-yellow{background:#FFD700;color:black;padding:18px;width:100%;border:none;border-radius:15px;font-size:20px;font-weight:bold;margin:10px 0} .btn-violet{background:#7c3aed;color:white;padding:18px;width:100%;border:none;border-radius:15px;font-size:20px;font-weight:bold;margin:10px 0} .result{background:white;color:black;padding:15px;border-radius:10px;margin-top:15px;font-weight:bold} </style> </head> <body> <h1>V32 REAL OK</h1> <p>Pi Pay CI - 1 Pi=800F | Gain 200F | PAIEMENT REEL</p> <div class="card"> NUMERO <input id="numero" value="0544535683"> Pi <input id="pi" type="number" value="2" oninput="calc()"> OPERATEUR <select id="op"><option>WAVE</option><option>ORANGE</option><option>MTN</option><option>MOOV</option></select> <button class="btn-yellow" id="calcBtn" onclick="calc()">PAYER 2 Pi -> 1600F</button> <button class="btn-violet" onclick="payerVraiPi()">Payer avec Pi Browser - REEL</button> <div class="result" id="res">2 Pi -> 1600F sur 0544535683</div> </div> <script> const Pi = window.Pi; Pi.init({ version: "2.0", sandbox: false }); function calc(){ let p = document.getElementById('pi').value || 0; let num = document.getElementById('numero').value; let fcfa = p * 800; document.getElementById('calcBtn').innerText = `PAYER ${p} Pi -> ${fcfa}F`; document.getElementById('res').innerText = `${p} Pi -> ${fcfa}F sur ${num}`; return fcfa; } async function payerVraiPi(){ let piAmount = parseFloat(document.getElementById('pi').value); let numero = document.getElementById('numero').value; let fcfa = piAmount * 800; if(!numero || piAmount <=0){ alert("Numero et Pi requis"); return; } try{ const auth = await Pi.authenticate(['payments'], (payment) => {}); Pi.createPayment({ amount: piAmount, memo: `PiPayCI: ${piAmount} Pi -> ${fcfa}F vers ${numero}`, metadata: { numero: numero, fcfa: fcfa, gain: 200* piAmount } }, { onReadyForServerApproval: (paymentId) => { document.getElementById('res').innerText = `APPROBATION: ${paymentId} - ${piAmount} Pi`; }, onReadyForServerCompletion: (paymentId, txid) => { document.getElementById('res').innerText = `SUCCES REEL: ${piAmount} Pi -> ${fcfa}F sur ${numero} | TX: ${txid}`; alert(`PAIEMENT REEL OK: ${piAmount} Pi -> ${fcfa}F sur ${numero}`); }, onCancel: (paymentId) => { alert("Paiement annulé"); }, onError: (error, payment) => { alert("Erreur Pi: " + error.message); } }); } catch(e){ alert("Ouvre dans Pi Browser pour paiement reel. Erreur: "+e.message); } } calc(); </script> </body> </html>`

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(html))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	http.ListenAndServe(":"+port, nil)
}
