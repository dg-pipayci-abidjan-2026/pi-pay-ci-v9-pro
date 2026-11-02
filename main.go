package main

import (
"fmt"
"net/http"
)

func main(){
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
msg:=""
num:="0574098776"
if r.Method=="POST"{
r.ParseForm()
numero:=r.FormValue("numero")
mont:=r.FormValue("montant")
op:=r.FormValue("operateur")
typ:=r.FormValue("servicetype")
act:=r.FormValue("action")
num=numero
var m float64
fmt.Sscan(mont,&m)
if op==""{
msg="Choisis operateur"
}else if typ==""{
msg="Choisis type"
}else if act=="konan"{
msg=fmt.Sprintf("KONAN: %.0fF %s %s (%s) -> %.2f Pi",m,op,typ,numero,m/1200)
}else{
msg=fmt.Sprintf("ALLI: %.2f Pi -> %.0fF %s %s %s Gain %.0fF",m,m*800,op,typ,numero,m*400)
}
}
html:=`<html><head><meta name="viewport" content="width=device-width,initial-scale=1"><style>body{background:#0f172a;color:#fff;font-family:sans-serif;display:flex;justify-content:center;padding:12px} .box{background:#1e293b;padding:20px;border-radius:20px;width:100%;max-width:400px} input,select{width:100%;padding:15px;margin:6px 0;border-radius:12px;border:none} .btn{width:100%;padding:16px;margin:6px 0;border:none;border-radius:12px;font-weight:bold} .y{background:#facc15} .g{background:#22c55e} .b{background:#38bdf8} .msg{background:#dcfce7;color:#000;padding:12px;border-radius:10px;margin-top:10px;font-weight:bold;text-align:center}</style></head><body><div class="box"><h2 style="text-align:center">V28 ULTRA ALLI & KONAN</h2><form method="POST"><input name="numero" value="`+num+`" required placeholder="Numero"><input name="montant" type="number" step="any" required placeholder="Montant"><select name="operateur" required><option value="">--Operateur--</option><option value="MTN">MTN</option><option value="ORANGE">ORANGE</option><option value="MOOV">MOOV</option><option value="WAVE">WAVE</option><option value="CANAL">CANAL+</option><option value="CIE">CIE</option><option value="SODECI">SODECI</option><option value="SOTRA">SOTRA</option></select><select name="servicetype" required><option value="">--Type--</option><option value="INTERNET">1 INTERNET</option><option value="APPEL">2 APPEL</option><option value="MOMO">3 MOMO</option><option value="FACTURE">FACTURE</option></select><button name="action" value="konan" class="btn y">KONAN FCFA -> Pi</button><button name="action" value="alli" class="btn g">ALLI Pi -> Service</button><button name="action" value="alli" class="btn b">ALLI Pi -> MOMO Direct</button></form><div class="msg">`+msg+`</div></div></body></html>`
w.Header().Set("Content-Type","text/html")
w.Write([]byte(html))
})
http.ListenAndServe(":10000",nil)
}
