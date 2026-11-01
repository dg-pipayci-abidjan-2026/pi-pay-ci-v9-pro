package main
import ("fmt";"net/http";"os")
func main(){
http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
w.Header().Set("Content-Type","text/html; charset=utf-8")
fmt.Fprint(w,`<html><meta name=viewport content="width=device-width,initial-scale=1"><meta charset=utf-8><style>
body{background:#0a0f1e;color:#fff;font-family:Arial;text-align:center;margin:0;padding:8px}
.box{max-width:430px;margin:auto;background:#1a1f3a;border-radius:18px;padding:10px;border:3px solid #00ff88}
.svc{display:grid;grid-template-columns:1fr 1fr 1fr;gap:6px;margin:8px 0}
.svc div{padding:9px;border-radius:8px;background:#2a2f4f;font-weight:bold;font-size:11px;cursor:pointer;border:1px solid #444}
.sel{background:#00ff88!important;color:#000!important}
.svc2{border:2px solid #ffcc00;border-radius:12px;padding:8px;margin:8px 0}
.svc2 div{display:inline-block;margin:3px;padding:6px 9px;border-radius:8px;background:#000;color:#ffcc00;border:1px solid #ffcc00;font-weight:bold;font-size:11px;cursor:pointer}
input{width:92%;padding:11px;border-radius:10px;border:0;text-align:center;font-size:16px;margin:4px}
button{width:95%;padding:11px;border-radius:10px;border:0;font-weight:bold;margin:4px;cursor:pointer}
.b1{background:#00ff88;color:#000}
.qr{width:180px;height:180px;background:#fff;margin:8px auto;display:none;padding:4px}
.recu{background:#fff;color:#000;border-radius:8px;padding:8px;margin:8px 0;display:none;text-align:left;white-space:pre-line;font-size:12px}
.mode{display:flex;gap:4px;background:#222;padding:4px;border-radius:15px}
.mode div{flex:1;padding:7px;border-radius:12px;cursor:pointer;font-weight:bold;font-size:12px;background:#333}
.active{background:#00ff88!important;color:#000!important}
</style><div class=box><h3 style="color:#00ff88;margin:4px">V23 ULTIME PiPay CI</h3>
<div class=mode><div id=m1 class=active onclick="setMode(1)">FCFA->Pi</div><div id=m2 onclick="setMode(2)">Pi->MTN 50=40000F</div></div>
<div class=svc>
<div onclick="pick(this,'SIDECI')">SIDECI</div><div onclick="pick(this,'CIE')">CIE</div><div onclick="pick(this,'SOTRA')">SOTRA</div>
<div onclick="pick(this,'STAR')">STAR</div><div onclick="pick(this,'LOYE')">LOYE</div><div onclick="pick(this,'ECOLE')">ECOLE</div>
<div class=sel onclick="pick(this,'MTN')">MTN</div><div onclick="pick(this,'ORANGE')">ORANGE</div><div onclick="pick(this,'MOOV')">MOOV</div>
<div onclick="pick(this,'WAVE')">WAVE</div><div onclick="pick(this,'CANAL')">CANAL</div><div onclick="pick(this,'USDT')">USDT</div>
<div onclick="pick(this,'CARTE')">CARTE</div><div onclick="pick(this,'QR BOUTIQUE')">QR BOUTIQUE</div><div onclick="pick(this,'MAGAZINE')">MAGAZINE</div>
<div onclick="pick(this,'MARQUES')">MARQUES</div><div onclick="pick(this,'APPEL')">APPEL</div><div onclick="pick(this,'INTERNET')">INTERNET</div>
</div>
<input id=tel value="0574098776" placeholder
