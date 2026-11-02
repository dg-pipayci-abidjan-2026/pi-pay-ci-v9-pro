package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msg := ""
	numeroVal := "0574098776"

	if r.Method == "POST" {
		r.ParseForm()
		numero := r.FormValue("numero")
		montantStr := r.FormValue("montant")
		operateur := r.FormValue("operateur")
		serviceType := r.FormValue("servicetype")
		action := r.FormValue("action")
		numeroVal = numero

		var montant float64
		fmt.Sscan(montantStr, &montant)

		if operateur == "" {
			msg = "❌ Choisis MTN / ORANGE / MOOV !"
		} else if serviceType == "" {
			msg = "❌ Choisis: Internet / Appel / MoMo !"
		} else if montant <= 0 {
			msg = "❌ Mets un montant !"
		} else if action == "konan" {
			pi := montant / 1200
			msg = fmt.Sprintf("✅ KONAN: Client envoie %.0fF via %s %s (%s) | Tu envoies %.2f Pi", montant, operateur, serviceType, numero, pi)
		} else {
			fcfa := montant * 800
			gain := montant * 400
			msg = fmt.Sprintf("✅ ALLI: %.2f Pi -> Tu paies %.0fF pour %s %s au %s | Gain %.0fF", montant, fcfa, operateur, serviceType, numero, gain)
		}
	}

	html := fmt.Sprintf(`
<html>
<head><meta name="viewport" content="width=device-width,initial-scale=1">
<style>
body{background:#0f172a;color:white;font-family:sans-serif;display:flex;justify-content:center;padding:12px;margin:0}
.box{background:#1e293b;padding:20px
