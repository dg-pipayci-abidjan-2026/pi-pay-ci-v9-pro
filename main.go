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
<input id=tel value="0574098776" placeholder="Numero client">
<div id=box1><input id=fcf type=number placeholder="Ex: 10000 FCFA" oninput="calc1()"><div id=res1 style="color:#00ff88;font-weight:bold;font-size:13px">0.00 Pi + 0.00 frais = 0.00 TOTAL</div></div>
<div id=box2 style="display:none"><input id=pi type=number step=0.01 placeholder="Ex: 50 Pi" oninput="calc2()"><div id=res2 style="color:#ffcc00;font-weight:bold;font-size:13px">Tu envoies: 0F via MTN</div></div>
<canvas id=qr class=qr></canvas><button class=b1 onclick="gen()">Generer Recu + QR</button><div id=recu class=recu></div><small style="color:#888">San-Pedro V23 - SIDECI CIE SOTRA STAR LOYE ECOLE MTN ORANGE MOOV WAVE CANAL USDT CARTE QR BOUTIQUE MAGAZINE MARQUES APPEL INTERNET</small></div>
<script src="https://cdn.jsdelivr.net/npm/qrcodejs@1.0.0/qrcode.min.js"></script><script>
let curSvc="MTN",curF=0,curPi=0,mode=1,qrObj=null
function pick(el,name){document.querySelectorAll('.svc div').forEach(d=>d.classList.remove('sel'));el.classList.add('sel');curSvc=name;calc1();calc2()}
function setMode(m){mode=m;document.getElementById('m1').className=m==1?'active':'';document.getElementById('m2').className=m==2?'active':'';document.getElementById('box1').style.display=m==1?'block':'none';document.getElementById('box2').style.display=m==2?'block':'none'}
function calc1(){let f=parseFloat(document.getElementById('fcf').value)||0;let pi=f/1200;let frais=pi*0.25;curF=f;curPi=pi;document.getElementById('res1').innerText=pi.toFixed(2)+" Pi + "+frais.toFixed(2)+" frais = "+(pi+frais).toFixed(2)+" TOTAL";showQR(f+"F="+pi.toFixed(2)+"Pi "+curSvc)}
function calc2(){let p=parseFloat(document.getElementById('pi').value)||0;let net=p*800-200;if(net<0)net=0;curPi=p;curF=net;document.getElementById('res2').innerText="Tu envoies: "+net.toFixed(0)+"F via "+curSvc+" | 50 Pi = 40000F";showQR(p+"Pi="+net+"F "+curSvc)}
function showQR(t){let c=document.getElementById('qr');c.style.display='block';c.innerHTML='';qrObj=new QRCode(c,{text:t+" | "+document.getElementById('tel').value+" | "+new Date().toLocaleString(),width:180,height:180})}
function gen(){if(curF==0){alert("Tape montant DG!");return}let d=new Date().toLocaleString('fr-CI');let b=document.getElementById('recu');b.style.display='block';if(mode==1){b.innerText="RECU PiPay CI V23\n"+d+"\nService: "+curSvc+"\nTel: "+document.getElementById('tel').value+"\nMontant: "+curF+"F\nPi: "+curPi.toFixed(2)+" Pi\nSan-Pedro Merci!"}else{b.innerText="RECU PiPay CI V23 Pi->"+curSvc+"\n"+d+"\nService: "+curSvc+"\nTel: "+document.getElementById('tel').value+"\nRecu: "+curPi+" Pi\nEnvoye: "+curF.toFixed(0)+"F via "+curSvc+"\n50 Pi = 40000F MTN\nMerci!"}window.scrollTo(0,document.body.scrollHeight)}
</script>`)
})
port:=os.Getenv("PORT");if port==""{port="10000"}
http.ListenAndServe(":"+port,nil)
}
