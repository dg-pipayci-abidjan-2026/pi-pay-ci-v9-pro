package main

import (
"log"
"net/http"
)

func main() {
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "text/html; charset=utf-8")
html := `<!DOCTYPE html><html lang="fr"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>Pi Pay CI V10 FULL</title><style>body{font-family:Arial;background:#0a192f;color:white;text-align:center;padding:20px}.card{background:#1a2332;padding:30px;border-radius:20px;max-width:400px;margin:auto;border:2px solid #f59e0b}h1{color:#f59e0b}button{width:100%;padding:15px;margin:8px 0;border:none;border-radius:12px;font-size:18px;font-weight:bold} .pi{background:#f59e0b;color:black}.wave{background:#00c851;color:white}.om{background:#ff9800;color:black}.mtn{background:#ffcc00;color:black}.moov{background:#0066cc;color:white}.cie{background:#ff6f00;color:white}.sodeci{background:#0099ff;color:white}.canal{background:black;color:white;border:1px solid white}.status{color:#00ff88;margin:20px}</style></head><body><h1>💎 Pi Pay CI V10 FULL 💎</h1><p>DG PiPay - Abidjan 2026 - 8 Services</p><div class="card"><h2>🚀 APP EN LIVE !</h2><p class="status">✅ Serveur Go Opérationnel</p><button class="pi">💜 Payer avec Pi</button><button class="wave">💚 Payer Wave</button><button class="om">🧡 Orange Money</button><button class="mtn">💛 MTN MoMo</button><button class="moov">💙 Moov Money</button><button class="cie">⚡ Payer CIE</button><button class="sodeci">💧 Payer SODECI</button><button class="canal">📺 Canal+ CI</button><p style="margin-top:30px;font-size:12px">pi-pay-ci-v9-pro-1.onrender.com<br>V10 FULL - DG</p></div></body></html>`
w.Write([]byte(html))
})
log.Println("V10 FULL LIVE sur :3000")
log.Fatal(http.ListenAndServe(":3000", nil))
}
