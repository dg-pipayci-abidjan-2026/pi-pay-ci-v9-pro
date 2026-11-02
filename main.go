package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const ACHAT = 800
const VENTE = 1200

var tmpl = template.Must(template.New("index").Parse(`
<!DOCTYPE html><html><head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>V25 FORT PiPay CI GO</title>
<style>
body{font-family:Arial;background:#0f172a;color:white;padding:10px;text-align:center}
.card{background:#1e293b;max-width:400px;margin:auto;padding:20px;border-radius:20px}
.grid{display:grid;grid-template-columns:1fr 1fr 1fr;gap:8px;margin:15px 0}
.btn{background:#334155;color:white;border:none;padding:12px;border-radius:10px;font-weight:bold}
.btn.active{background:#f59e0b;color:black}
.btn-big{width:100%;padding:16px;margin:8px 0;border:none;border-radius:12px;font-weight:bold;font-size:16px;cursor:pointer}
.yellow{background:#facc15;color:black}
.green{background:#22c55e;color:black}
input{width:95%;padding:14px;margin:8px 0;border-radius:10px;border:none;font-size:16px}
.res{background:#dcfce7;color:black;padding:15px;border-radius:12px;margin-top:15px;font-weight:bold}
</style></head><body>
<div class="card">
<h2>V25 FORT Go 🔥<br>ALLI & KONAN</h2>
<p><b>Achat 1 Pi=800F | Vente 1 Pi=1200F | Gain 400F</b></p>
<div class="grid">
<button class="btn active">MTN</button><button class="btn">ORANGE</button><button class="btn">MOOV</button>
<button class="btn">WAVE</button><button class="btn">CIE</button><button class="btn">SOTRA</button>
<button class="btn">CANAL+</button><button class="btn">USDT</button><button class="btn">CARTE</button>
<button class="btn" style="grid-column: span 3">SIDECI</button>
</div>
<form method="POST">
<input type="text" name="numero" placeholder="Numero 0574098776" required value="0574098776">
<input type="number" step="0.01" name="montant" placeholder="Ex: 1 ou 50" required>
<button class="btn-big yellow" name="action" value="fcfapi">KONAN: FCFA->Pi 1=1200F</button>
<button class="btn-big green" name="action" value="pifcfa">ALLI: Pi->WAVE 1=800F</button>
</form>
{{if .Res}}<div class="res">{{.Res}}</div>{{end}}
<p style="margin-top:15px;font-size:12px">PiPay CI V25 GO - San Pedro - 0574098776</p>
</div></body></html>
`))

type Data struct { Res template.HTML }

func handler(w http.ResponseWriter, r *http.Request) {
	var res string
	if r.Method == "POST" {
		r.ParseForm()
		num := r.FormValue("numero")
		montStr := r.FormValue("montant")
		act := r.FormValue("action")
		mont, _ := strconv.ParseFloat(montStr, 64)
		if act == "pifcfa" {
			f := mont * ACHAT
			res = fmt.Sprintf("✅ ALLI a envoyé %.2f Pi<br>👉 Tu envoies <b>%.0fF WAVE</b> à %s<br>Calcul: %.2f x 800 = %.0fF<br>Bénéfice: %.0fF", mont, f, num, mont, f, mont*400)
		} else {
			pi := mont / VENTE
			res = fmt.Sprintf("✅ KONAN a envoyé %.0fF (%s)<br>👉 Tu envoies <b>%.2f Pi</b><br>Calcul: %.0f / 1200 = %.2f Pi", mont, num, pi, mont, pi)
		}
	}
	tmpl.Execute(w, Data{Res: template.HTML(res)})
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":10000", nil)
}
`))

**4. Après avoir collé, clique VERT `Commit changes...` > `Commit changes`**

**5. Attends 2 min, ton Render va redémarrer**

Après tu vas voir sur ton site bleu en bas :
- **JAUNE : KONAN FCFA->Pi 1=1200F**
- **VERT : ALLI Pi->WAVE 1=800F**

Teste avec 1 et dis-moi DG ! C'est le VRAI code FORT Go !
