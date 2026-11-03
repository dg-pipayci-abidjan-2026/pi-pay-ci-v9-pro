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
func main(){
 http.HandleFunc("/", serveHome)
 http.HandleFunc("/api/approve", handleApprove)
 http.HandleFunc("/api/complete", handleComplete)
 http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request){w.Write([]byte("ULTIMATE LIGHT OK"))})
 port:=os.Getenv("PORT"); if port==""{port="10000"}
 log.Println("ULTIMATE LIGHT", port, "key:", piAPIKey!="")
 log.Fatal(http.ListenAndServe(":"+port, nil))
}
func serveHome(w http.ResponseWriter, r *http.Request){
 w.Header().Set("Content-Type","text/html; charset=utf-8")
 fmt.Fprint(w, `<!DOCTYPE html><html lang="fr"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><title>V30 ULTIMATE LIGHT</title><script src="https://sdk.minepi.io/pi-sdk.js"></script><style>*{margin:0;padding:0;box-sizing:border-box;font-family:Segoe UI,sans-serif}body{background:linear-gradient(135deg,#1a0b2e,#2d1b4e,#0f2027);min-height:100vh;display:flex;justify-content:center;padding:16px}.container{width:100%;max-width:430px}.header{text-align:center;color:white;margin:14px 0}.header h1{font-size:26px;font-weight:900;color:#FFD60A}.card{background:rgba(255,255,255,0.09);border:1px solid rgba(255,255,255,0.15);border-radius:22px;padding:16px}.tab{flex:1;padding:12px;border-radius:12px;border:none;font-weight:900;font-size:12px;cursor:pointer}#tabMine{background:#FFD60A;color:#000}#tabNoMine{background:#00E676;color:#000}.tabs{display:flex;gap:8px;margin-bottom:12px}.input-group{margin-bottom:10px}.input-group label{color:#d6c2ff;font-size:10px;font-weight:800;margin-left:6px}.input-group input,.input-group select{width:100%;padding:12px;border-radius:10px;border:none;background:white;font-weight:700;margin-top:4px}.btn{width:100%;padding:14px;border-radius:12px;border:none;font-weight:900;font-size:14px;cursor:pointer;margin-top:8px}.btn-mine{background:#FFD60A;color:#000}.btn-pi{background:#8E2DE2;color:white}.result{margin-top:12px;padding:12px;border-radius:12px;background:white;color:black;font-weight:700;text-align:center;font-size:13px;display:none}.result.show{display:block}.log{background:black;color:#0f0;font-family:monospace;font-size:10px;padding:6px;border-radius:6px;margin-top:6px;max-height:100px;overflow:auto;display:none}</style></head><body><div class="container"><div class="header"><h1>V30 ULTIMATE LIGHT</h1><p>Pi Pay CI - Marche 100%<br>1 Pi=1000F Achat 800F Vente 1200F Gain 200F</p></div><div class="card"><div class="tabs"><button id="tabMine" class="tab" style="background:#FFD60A" onclick="switchTab('mine')">J'AI MINE<br>JE VEUX CFA</button><button id="tabNoMine" class="tab" style="background:#00E676" onclick="switchTab('nomine')">J'AI PAS MINE<br>JE VEUX Pi</button></div><div class="input-group"><label>NUMERO / FACTURE</label><input id="numero" value="0574098776"></div><div id="mineFields"><div class="input-group"><label>COMBIEN DE Pi?</label><input id="piAmount" type="number" value="1" oninput="calcMine()"></div></div><div id="nomineFields" style="display:none"><div class="input-group"><label>COMBIEN FCFA?</label><input id="fcfaAmount" type="number" placeholder="1200" oninput="calcNoMine()"></div></div><div class="input-group"><label>OPERATEUR</label><select id="operateur"><option>WAVE</option><option>MTN</option><option>ORANGE</option><option>MOOV</option><option>MTN MoMo</option><option>Orange Money</option><option>CANAL+</option><option>CIE</option></select></div><div class="input-group"><label>TYPE</label><select id="typeOp"><option>3 MOMO / WAVE</option><option>1 INTERNET</option><option>2 APPEL</option><option>4 FACTURE</option></select></div><div id="btnMine"><button class="btn btn-mine" onclick="payMine()">PAYER <span id="minePiText">1</span> Pi -> <span id="mineFcfaText">800</span>F CFA</button><button class="btn btn-pi" onclick="payMine()">Payer avec Pi Browser - 1 CLIC</button></div><div id="btnNoMine" style="display:none"><button class="btn" style="background:#00E676" onclick="payNoMine()">PAYER <span id="nomineFcfaText">0</span>F -> <span id="nominePiText">0</span> Pi</button></div><div id="result" class="result show">Pret: 1 Pi = 800F sur 0574098776<br>Gain 200F - Clique bouton VIOLET</div><div id="log" class="log"></div></div></div><script>
let piConnected=false;
function logAdd(t){let el=document.getElementById('log');el.style.display='block';el.innerHTML+=t+'<br>';console.log(t)}
function switchTab(t){document.getElementById('mineFields').style.display=t=='mine'?'block':'none';document.getElementById('nomineFields').style.display=t=='nomine'?'block':'none';document.getElementById('btnMine').style.display=t=='mine'?'block':'none';document.getElementById('btnNoMine').style.display=t=='nomine'?'block':'none'}
function op(){return document.getElementById('operateur').value+' '+document.getElementById('typeOp').value}
function show(h){let r=document.getElementById('result');r.innerHTML=h;r.classList.add('show')}
function calcMine(){let pi=parseFloat(document.getElementById('piAmount').value)||0;let fcfa=pi*800;document.getElementById('minePiText').inner
