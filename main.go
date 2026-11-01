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
  fmt.Fprint(w, `<!DOCTYPE html><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>Pi Pay CI V11 - DG Abidjan</title><script src="https://sdk.minepi.com/pi-sdk.js"></script><style>body{font-family:sans-serif;background:#6a1b9a;color:white;text-align:center;padding:20px}button{padding:15px 25px;margin:10px;font-size:18px;border-radius:12px;border:none}.pi{background:#ffeb3b;color:#000;font-weight:bold}h1{font-size:28px}.ok{background:#00e676;color:#000;padding:10px;border-radius:8px;margin:10px}</style></head><body><h1>💜 Pi Pay CI V11 - DG ABIDJAN 1863 Pi</h1><div class="ok">✅ Pi SDK Prêt - Clique Pi - V11 PI REEL LIVE</div><p>Première app Pi vers Orange Money Wave CI</p><button class="pi" onclick="payPi()">💜 Payer avec Pi - 1 Pi Test</button><br><button>Orange Money</button><button>Wave</button><button>MTN</button><button>Moov</button><button>Carte</button><button>USDT</button><button>Cash</button><p id="status"></p><script>const Pi=window.Pi;Pi.init({version:"2.0",sandbox:true});async function payPi(){document.getElementById("status").innerText="Connexion Pi...";try{const auth=await Pi.authenticate(["username","payments"],onIncompletePaymentFound);document.getElementById("status").innerText="Bienvenue "+auth.user.username+" - Paiement...";const payment=await Pi.createPayment({amount:1,memo:"Paiement Pi Pay CI - Test DG",metadata:{service:"PiPayCI",city:"Abidjan"}},{onReadyForServerApproval:function(id){console.log("APPROVAL",id);document.getElementById("status").innerText="Approuvé "+id+" - V11 LIVE!";},onReadyForServerCompletion:function(id,tx){console.log("COMPLETE",id,tx);document.getElementById("status").innerText="✅ Paiement Réussi! TX:"+tx;},onCancel:function(id){document.getElementById("status").innerText="Annulé "+id;},onError:function(e,id){document.getElementById("status").innerText="Erreur "+e;}});}catch(e){document.getElementById("status").innerText="Erreur Pi: "+e.message;}}function onIncompletePaymentFound(p){console.log("INCOMPLETE",p);}</script></body></html>`)
 })

 port := os.Getenv("PORT")
 if port == "" { port = "8080" }
 log.Println("V11 PI REEL LIVE sur port", port)
 log.Fatal(http.ListenAndServe(":"+port, nil))
}
