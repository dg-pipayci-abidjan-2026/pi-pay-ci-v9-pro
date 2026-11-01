package main
import ("log";"net/http";"os")
func main(){
port:=os.Getenv("PORT")
if port==""{port="10000"}
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","text/html; charset=utf-8")
w.Write([]byte(page))
})
log.Println("DG V19 2400F=2Pi LIVE "+port)
log.Fatal(http.ListenAndServe(":"+port,nil))
}
var page=[]byte(`<!DOCTYPE html><html lang="fr"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><title>V19 DG 2400F=2Pi</title><script src="https://sdk.minepi.com/pi-sdk.js"></script><script src="https://cdn.jsdelivr.net/npm/qrcodejs@1.0.0/qrcode.min.js"></script><style>body{font-family:Arial;background:#0f0f1a;color:#fff;text-align:center;padding:12px}h1{color:#00ff88}.card{background:#1a1a2e;border-radius:20px;padding:16px;margin:10px auto;max-width:480px;border:2px solid #00ff88}.btn{width:100%;padding:14px;border:none;border-radius:12px;font-weight:bold;margin:5px 0;cursor:pointer}.pay{background:#7B2CBF;color:#fff}.mode{display:flex;gap:6px}.mode button{flex:1;padding:10px;border-radius:10px;border:2px solid #333;background:#2a2a4a;color:#fff;font-weight:bold}.mode button.active{background:#00ff88;color:#000}.service{display:inline-block;background:#3a3a6a;padding:7px 9px;border-radius:8px;margin:3px;font-size:11px;cursor:pointer}.service.active{background:#00ff88;color:#000}.qrBox{background:#fff;padding:12px;border-radius:16px;margin:12px auto;max-width:340px;display:none}#qrcode{display:flex;justify-content:center}.calc{background:#111;border:1px solid #00ff88;border-radius:12px;padding:10px;margin:10px 0;color:#00ff88;font-weight:bold}input{width:92%;padding:10px;border-radius:10px;border:none;margin:6px 0;text-align:center}</style></head><body>
