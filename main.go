package main
import ("log";"net/http";"os")
func main(){
p:=os.Getenv("PORT")
if p==""{p="10000"}
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","text/html; charset=utf-8")
w.Write([]byte(idx))
})
http.HandleFunc("/api/pi/approve", func(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","application/json")
w.Write([]byte(`{"status":"approved","v18":"2400F=2Pi"}`))
})
http.HandleFunc("/api/pi/complete", func(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","application/json")
w.Write([]byte(`{"status":"completed","v18":"DG 1863 Pi"}`))
})
log.Println("DG V18 2400F=2Pi LIVE",p)
log.Fatal(http.ListenAndServe(":"+p,nil))
}
const idx=`<!DOCTYPE html><html lang="fr"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><title>PI PAY CI V18 - 2400F=2Pi - DG Model</title><script src="https://sdk.minepi.com/pi-sdk.js"></script><script src="https://cdn.jsdelivr.net/npm/qrcodejs@1.0.0/qrcode.min.js"></script><style>body{font-family:Arial;background:#0f0f1a;color:white;text-align:center;padding:12px;margin:0}h1{color:#00ff88;margin:6px;font-size:22px}.card{background:#1a1a2e;border-radius:20px;padding:16px;margin:10px auto;max-width:480px;border:2px solid #00ff88}.btn{width:100%;padding:14px;border:none;border-radius:12px;font-weight:bold;cursor:pointer;margin:5px 0;font-size:15px}.open{background:#00ff88;color:#000}.close{background:#ff3344;color:#fff}.pay{background:#7B2CBF;color:#fff;font-size:17px}.mode{display:flex;gap:8px;margin:8px 0}.mode button{flex:1;padding:12px;border-radius:10px;border:2px solid #333;background:#2a2a4a;color:#fff;cursor:pointer;font-weight:bold}.mode button.active{border-color:#00ff88;background:#00ff88;color:#000}.service{display:inline-block;background:#3a3a6a;padding:8px 10px;border-radius:8px;margin:3px;cursor:pointer;font-size:11px;border:2px solid transparent}.service.active{border-color:#00ff88;background:#00ff88;color:#000}.service-qr{border-color:#facc15;color:#facc15}.service-qr.active{background:#facc15;color:#000}.tiroir{background:#2a2a4a;padding:10px;border-radius:12px;margin:8px 0;font-size:12px}.calcBox{background:#111;border:1px solid #00ff88;border-radius:12px;padding:12px;margin:10px 0;color:#00ff88;font-weight:bold;font-size:13px}input{width:92%;padding:11px;border-radius:10px;border:none;margin:6px 0;text-align:center;font-size:15px}.qrBox{background:#fff;padding:12px;border-radius:16px;margin:12px auto;max-width:340px;display:none}#qrcode{display:flex;justify-content:center;margin:10px auto}</style></head><body><h1>🔐 PI PAY CI V18 - MODEL DG</h1><p style="color:#facc15;font-size:12
