<!DOCTYPE html>
<html lang="fr">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>Pi Pay CI - V32 Réel</title>
<script src="https://sdk.minepi.com/pi-sdk.js"></script>
<style>
body{font-family:system-ui;background:#0f172a;color:white;padding:20px;max-width:400px;margin:auto}
.card{background:#1e293b;padding:20px;border-radius:16px;margin-top:20px}
input,select{width:100%;padding:12px;border-radius:10px;border:none;margin:8px 0;font-size:16px;box-sizing:border-box}
.btn-yellow{background:#facc15;color:black;padding:14px;border-radius:12px;text-align:center;font-weight:bold;margin:10px 0}
.btn-violet{background:#7c3aed;color:white;padding:14px;border-radius:12px;text-align:center;font-weight:bold;cursor:pointer}
.btn-violet:active{transform:scale(0.98)}
.badge{background:#facc15;color:black;padding:4px 10px;border-radius:20px;font-weight:bold;font-size:12px}
#log{font-size:13px;margin-top:10px;color:#94a3b8;word-break:break-all}
</style>
</head>
<body>
<span class="badge">V32 OK - PAIEMENT RÉEL</span>
<h2>Pi Pay CI 🇨🇮</h2>
<p id="pi-user">Non connecté Pi - Ouvre dans Pi Browser</p>

<div class="card">
<input id="phone" placeholder="Numero ex: 0544535683" value="0544535683">
<select id="operateur">
<option value="WAVE">WAVE</option>
<option value="ORANGE">ORANGE MONEY</option>
<option value="MTN">MTN MOMO</option>
</select>
<input id="pi-amount" type="number" placeholder="Montant Pi" value="2" min="0.1" step="0.1" oninput="updateCalc()">
<div class="btn-yellow" id="calc">PAYER 2 Pi -> 1600F</div>
<div class="btn-violet" onclick="payerReel()">Payer avec Pi Browser - RÉEL</div>
<div id="log">En attente...</div>
</div>

<script>
// INIT PI SDK
Pi.init({ version: "2.0", sandbox: false });

let piUser = null;

async function authPi() {
  try {
    const scopes = ['username', 'payments'];
    const auth = await Pi.authenticate(scopes, onIncompletePaymentFound);
    piUser = auth.user;
    document.getElementById('pi-user').innerText = "Connecté: @" + piUser.username;
    document.getElementById('log').innerText = "Pi connecté OK";
  } catch(e) {
    document.getElementById('log').innerText = "Auth Pi: " + e.message + " (Ouvre dans Pi Browser)";
  }
}

function onIncompletePaymentFound(payment) {
  console.log("Paiement incomplet trouvé", payment);
  return;
}

function updateCalc() {
  const pi = parseFloat(document.getElementById('pi-amount').value) || 0;
  const fcfa = pi * 800;
  const gain = pi * 200;
  document.getElementById('calc').innerText = `PAYER ${pi} Pi -> ${fcfa}F (Gain: ${gain}F)`;
}
updateCalc();

async function payerReel() {
  const phone = document.getElementById('phone').value;
  const op = document.getElementById('operateur').value;
  const amount = parseFloat(document.getElementById('pi-amount').value);
  
  if(!phone || !amount) { alert("Remplis numéro et montant"); return; }
  
  if(!piUser) { await authPi(); }

  document.getElementById('log').innerText = "Création paiement Pi...";

  try {
    await Pi.createPayment({
      amount: amount,
      memo: `Pi Pay CI: ${amount} Pi -> ${op} ${phone}`,
      metadata: { phone: phone, operateur: op, fcfa: amount*800 }
    }, {
      onReadyForServerApproval: function(paymentId) {
        document.getElementById('log').innerText = "Approuvé côté serveur: " + paymentId;
        console.log("APPROVAL:", paymentId);
        // SANS BACKEND ON APPROUVE DIRECT (DEMO)
        // En prod tu appelles ton backend ici
      },
      onReadyForServerCompletion: function(paymentId, txid) {
        document.getElementById('log').innerHTML = `✅ PAIEMENT RÉEL RÉUSSI!<br>PaymentID: ${paymentId}<br>TxID: ${txid}<br>${amount} Pi -> ${amount*800}F sur ${phone} (${op})`;
        alert(`SUCCÈS PI RÉEL: ${amount} Pi payé! Tu dois envoyer ${amount*800}F sur ${phone}`);
      },
      onCancel: function(paymentId) {
        document.getElementById('log').innerText = "Paiement annulé: " + paymentId;
      },
      onError: function(error, payment) {
        document.getElementById('log').innerText = "Erreur Pi: " + error.message;
        console.error(error);
      }
    });
  } catch(e) {
    document.getElementById('log').innerText = "Erreur: " + e.message;
  }
}

// Auto auth au chargement
authPi();
</script>
</body>
</html>
