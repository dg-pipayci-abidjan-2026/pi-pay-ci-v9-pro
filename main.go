package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		html := `<!DOCTYPE html>
<html lang="fr">
<head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>Pi Pay CI V9 Pro - DG PiPay</title>
<style>
body{font-family:Arial;background:#0a0e1a;color:white;text-align:center;padding:20px}
.card{background:#1a2332;padding:30px;border-radius:20px;max-width:400px;margin:20px auto;border:2px solid #f59e0b}
h1{color:#f59e0b}.btn{background:#f59e0b;color:black;padding:15px 30px;border:none;border-radius:10px;font-weight:bold;font-size:18px;margin:10px;cursor:pointer;width:90%}
.btn-green{background:#00c851;color:white}
.status{color:#00ff88;margin:20px}
</style></head>
<body>
<h1>💎 Pi Pay CI V9 Pro 💎</h1>
<p>DG PiPay - Abidjan 2026</p>
<div class="card">
<h2>🚀 APP EN LIVE !</h2>
<p class="status">✅ Serveur Go Opérationnel</p>
<p>Payement Pi Network + Mobile Money CI</p>
<button class="btn" onclick="alert('Pi Payment V9 activé !')">💜 Payer avec Pi</button>
<button class="btn btn-green" onclick="alert('Wave CI')">💚 Payer Wave</button>
<button class="btn" style="background:#ff8c00" onclick="alert('Orange Money')">🧡 Orange Money</button>
<p style="margin-top:30px;font-size:12px;opacity:0.7">pi-pay-ci-v9-pro.onrender.com<br>Propulsé par DG-PiPayCI-Abidjan</p>
</div>
</body></html>`
		w.Write([]byte(html))
	})

	log.Println("Pi Pay CI V9 Pro LIVE sur port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
