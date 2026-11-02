package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const ACHAT = 800
const VENTE = 1200

var page = `
<!DOCTYPE html><html><head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>V25 FORT</title>
<style>
body{font-family:Arial;background:#0f172a;color:white;padding:10px;text-align:center}
.card{background:#1e293b;max-width:400px;margin:auto;padding:20px;border-radius:20px}
.btn-big{width:100%;padding:16px;margin:8px 0;border:none;border-radius:12px;font-weight:bold;font-size:16px}
.yellow{background:#facc15;color:black}
.green{background:#22c55e;color:black}
input{width:95%;padding:14px;margin:8px 0;border-radius:10px;border:none;font-size:16px}
.res{background:#dcfce7;color:black;padding:15px;border-radius:12px;margin-top:15px;font-weight:bold}
</style></head><body>
<div class="card">
<h2>V25 FORT 🔥 ALLI & KONAN</h2>
<p><b>1 Pi Achat 800F | 1 Pi Vente 1200F | Gain 400F</b></p>
<form method="POST">
<input type="text" name="numero" placeholder="Numero" value="0574098776" required>
<input type="number" step="0.01" name="montant" placeholder="Montant" required>
<button class="btn-big yellow" name="action" value="fcfapi">KONAN FCFA->Pi 1=1200F</button>
<button class="btn-big green" name="action" value="pifcfa">ALLI Pi->WAVE 1=800F</button>
</form>
{{if .}}<div class="res">{{.}}</div>{{end}}
</div></body></html>
`

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("p").Parse(page))
	var msg string
	if r.Method == "POST" {
		r.ParseForm()
		num := r.FormValue("numero")
		mont, _ := strconv.ParseFloat(r.FormValue("montant"), 64)
		act := r.FormValue("action")
		if act == "pifcfa" {
			f := mont * ACHAT
			msg = fmt.Sprintf("✅ ALLI %.2f Pi -> Envoie %.0fF WAVE a %s | Gain %.0fF", mont, f, num, mont*400)
		} else {
			pi := mont / VENTE
			msg = fmt.Sprintf("✅ KONAN %.0fF (%s) -> Envoie %.2f Pi | Pi = %.0f / 1200", mont, num, pi, mont)
		}
	}
	tmpl.Execute(w, template.HTML(msg))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":10000", nil)
}
