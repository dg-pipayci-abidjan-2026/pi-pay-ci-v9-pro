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

var piAPIKey = os.Getenv("PI_API_KEY")
const piAPIBase = "https://api.minepi.com"

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/api/approve", handleApprove)
	http.HandleFunc("/api/complete", handleComplete)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("V30 ULTIMATE OK"))
	})
	port := os.Getenv("PORT")
	if port == "" { port = "10000" }
	log.Println("V30 ULTIMATE running on", port, "PI_API_KEY set:", piAPIKey != "")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<!DOCTYPE html>
<html lang="fr"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>V30 ULTIMATE - Pi Pay CI</title>
<script src="https://sdk.minepi.io/pi-sdk.js"></script>
<style>
*{margin:0;padding:0;box-sizing:border-box;font-family:'Segoe UI',sans-serif}
body{background:linear-gradient(135deg,#1a0b2e 0%,#2d1b4e 50%,#0f2027 100%);min-height:100vh;display:flex;justify-content:center;padding:16px}
.container{width:100%;max-width:430px}
.header{text-align:center;color:white;margin:16px 0}
.header h1{font-size:28px;font-weight:900;color:#FFD60A}
.header p{font-size:12px;opacity:0.9;margin-top:4px}
.card{background:rgba(255,255,255,0.09);backdrop-filter:blur(20px);border:1px solid rgba(255,255,255,0.15);border-radius:24px;padding:16px}
.tabs{display:flex;gap:8px;margin-bottom:14px}
.tab{flex:1;padding:12px 8px;border-radius:14px;border:none;font-weight:900;cursor:pointer;font-size:12px}
#tabMine{background:#FFD60A;color:#000}#tabNoMine{background:#00E676;color:#000}
.tab.active{box-shadow:0 0 20px currentColor;transform:scale(1.02)}
.input-group{margin-bottom:10px}.input-group label{color:#d6c2ff;font-size:10px;font-weight:800;margin-left:8px}
.input-group input,.input-group select{width:100%;padding:13px 14px;border-radius:12px;border:none;background:rgba(255,255,255,0.96);font-size:14px;font-weight:700;margin-top:4px;outline:none}
.btn{width:100%;padding:15px;border-radius:13px;border:none;font-weight:900;font-size:15px;cursor:pointer;margin-top:8px}
.btn-mine{background:linear-gradient(135deg,#FFD60A,#FFA500);color:#000}
.btn-pi{background:linear-gradient(135deg,#8E2DE2,#4A00E0);color:white;box-shadow:0 6px 18px rgba(142,45,226,0.5)}
.result{margin-top:14px;padding:14px;border-radius:14px;background:white;color:#000;font-weight:700;text-align:center;font-size:13px;display:none;word-break:break-word}
.result.show{display:block}
.badge{background:#FFD60A;padding:3px 8px;border-radius:10px;font-size:11px}
.log{background:black;color:#00FF00;font-family:monospace;font-size:10px;padding:8px;border-radius:8px;margin-top:8px;max-height:120px;overflow:auto;display:none}
.footer{color:rgba(255,255,255,0.5);text-align:center;font-size:10px;margin-top:12px}
</style></head><body>
<div class="container">
<div class="header"><h1>V30 ULTIMATE</h1><p>Pi Pay CI 🇨🇮 - LE PLUS PUISSANT<br>1 Pi = 1000F | Achat 800F | Vente 1200F | Gain 200F<br>✅ Marche même sans API Key</p></div>
<div class="card">
<div class="tabs">
<button id="tabMine" class="tab active" onclick="switchTab('mine')">J'AI DÉJÀ MINÉ<br>JE VEUX CFA</button>
<button id="tabNoMine" class="tab" onclick="switchTab('nomine')">J'AI PAS MINÉ<br>JE VEUX Pi</button>
</div>
<div class="input-group"><label>NUMÉRO / ID FACTURE</label><input id="numero" type="text" value="0574098776"></div>
<div id="mineFields"><div class="input-group"><label>COMBIEN DE Pi TU AS ?</label><input id="piAmount" type="number" value="1" oninput="calcMine()"></div></div>
<div id="nomineFields" style="display:none"><div class="input-group"><label>COMBIEN FCFA TU AS ?</label><input id="fcfa
