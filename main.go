from flask import Flask, render_template_string, request
import qrcode, io, base64
app = Flask(__name__)
ACHAT = 800
VENTE = 1200

HTML = """
<!DOCTYPE html><html><head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>V25 ALLI-KONAN PiPay CI</title>
<style>
body{font-family:Arial;background:#0f172a;color:white;padding:10px;text-align:center}
.card{background:#1e293b;max-width:400px;margin:auto;padding:20px;border-radius:20px}
.grid{display:grid;grid-template-columns:1fr 1fr 1fr;gap:8px;margin:15px 0}
.btn{background:#334155;color:white;border:none;padding:12px;border-radius:10px;font-weight:bold}
.btn.active{background:#f59e0b;color:black}
.btn-big{width:100%;padding:16px;margin:8px 0;border:none;border-radius:12px;font-weight:bold;font-size:15px;cursor:pointer}
.yellow{background:#facc15;color:black}
.green{background:#22c55e;color:black}
input{width:95%;padding:14px;margin:8px 0;border-radius:10px;border:none;font-size:16px}
.res{background:#dcfce7;color:black;padding:15px;border-radius:12px;margin-top:15px;font-weight:bold}
.qr{background:white;padding:10px;border-radius:15px;display:inline-block;margin:10px}
</style></head><body>
<div class="card">
<h2>V25 ULTIME PiPay CI 🔥<br>ALLI & KONAN</h2>
<p>Achat: 1 Pi = 800F | Vente: 1 Pi = 1200F | Gain: 400F/Pi</p>
<div class="grid">
<button class="btn active">MTN</button><button class="btn">ORANGE</button><button class="btn">MOOV</button>
<button class="btn">WAVE</button><button class="btn">CIE</button><button class="btn">SOTRA</button>
<button class="btn">CANAL+</button><button class="btn">USDT</button><button class="btn">CARTE</button>
<button class="btn" style="grid-column: span 3">SIDECI</button>
</div>
<form method="POST">
<input type="text" name="numero" placeholder="Numéro: 0574098776" required value="0574098776">
<input type="number" step="0.01" name="montant" placeholder="Ex: 1 ou 50 ou 1200" required>
<button class="btn-big yellow" name="action" value="fcfapi">KONAN: FCFA->Pi 1=1200F</button>
<button class="btn-big green" name="action" value="pifcfa">ALLI: Pi->WAVE 1=800F</button>
</form>
{% if res %}<div class="res">{{ res|safe }}
{% if qr %}<div class="qr"><img src="{{ qr }}" width="240"></div><br>QR à scanner dans Pi Browser{% endif %}</div>{% endif %}
<p style="margin-top:15px;font-size:12px">Lien Pi: https://pi-pay-ci-v9-pro-1.onrender.com</p>
</div></body></html>
"""

def make_qr(t):
    img=qrcode.make(t); b=io.BytesIO(); img.save(b,format='PNG'); return "data:image/png;base64,"+base64.b64encode(b.getvalue()).decode()

@app.route("/",methods=["GET","POST"])
def idx():
    res=None;qr=None
    if request.method=="POST":
        num=request.form.get("numero"); mont=float(request.form.get("montant")); act=request.form.get("action")
        if act=="pifcfa":
            f= mont * ACHAT
            res=f"✅ ALLI a envoyé {mont} Pi<br>👉 Tu envoies <b>{int(f)}F via WAVE</b> à {num}<br>Calcul: {mont} x 800F = {int(f)}F<br>Tu gardes 400F x {mont} = {int(mont*400)}F bénéfice"
            qr=make_qr(f"pi-pay:{mont}Pi->{int(f)}F to {num}")
        else:
            pi= mont / VENTE
            res=f"✅ KONAN a envoyé {int(mont)}F ({num})<br>👉 Tu envoies <b>{pi:.2f} Pi</b><br>Calcul: {int(mont)}F / 1200F = {pi:.2f} Pi<br>Tu gardes 400F par Pi"
            qr=make_qr(f"fcfa:{int(mont)}->{pi:.2f}Pi to {num}")
    return render_template_string(HTML,res=res,qr=qr)

if __name__=="__main__":
    app.run(host="0.0.0.0",port=10000)
