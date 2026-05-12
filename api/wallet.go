package main

import (
"encoding/json"
"fmt"
"log"
"net/http"
"time"

"github.com/google/uuid"
)

// Wallet stores credentials
type Wallet struct {
ID           string      `json:"id"`
OwnerDID     string      `json:"ownerDid"`
Credentials []Credential `json:"credentials"`
CreatedAt   string      `json:"createdAt"`
UpdatedAt  string      `json:"updatedAt"`
}

type Credential struct {
ID               string                 `json:"id"`
Type             []string               `json:"type"`
Issuer           string                `json:"issuer"`
CredentialSubject map[string]interface{} `json:"credentialSubject"`
IssuedAt         string                `json:"issuedAt"`
ExpiresAt        string                `json:"expiresAt"`
// Raw VC for presentation
Raw map[string]interface{} `json:"-"`
}

// In-memory storage
var wallets = make(map[string]Wallet)

// Handlers

func healthHandler(w http.ResponseWriter, r *http.Request) {
json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func createWalletHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

var req struct {
OwnerDID string `json:"ownerDid"`
}
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

wallet := Wallet{
ID:         "wallet:" + uuid.New().String(),
OwnerDID:  req.OwnerDID,
CreatedAt:  time.Now().UTC().Format(time.RFC3339),
UpdatedAt:  time.Now().UTC().Format(time.RFC3339),
}
wallets[wallet.ID] = wallet

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(wallet)
}

func getWalletHandler(w http.ResponseWriter, r *http.Request) {
walletID := r.PathValue("id")
wallet, ok := wallets[walletID]
if !ok {
http.Error(w, "wallet not found", http.StatusNotFound)
return
}

json.NewEncoder(w).Encode(wallet)
}

func addCredentialHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

walletID := r.PathValue("id")
wallet, ok := wallets[walletID]
if !ok {
http.Error(w, "wallet not found", http.StatusNotFound)
return
}

var cred Credential
if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

cred.ID = uuid.New().String()
wallet.Credentials = append(wallet.Credentials, cred)
wallet.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
wallets[walletID] = wallet

json.NewEncoder(w).Encode(wallet)
}

func listCredentialsHandler(w http.ResponseWriter, r *http.Request) {
walletID := r.PathValue("id")
wallet, ok := wallets[walletID]
if !ok {
http.Error(w, "wallet not found", http.StatusNotFound)
return
}

json.NewEncoder(w).Encode(map[string]interface{}{
"credentials": wallet.Credentials,
})
}

func presentCredentialHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

walletID := r.PathValue("id")
wallet, ok := wallets[walletID]
if !ok {
http.Error(w, "wallet not found", http.StatusNotFound)
return
}

var req struct {
CredentialID string   `json:"credentialId"`
Audience    string   `json:"audience"`
Challenge   string   `json:"challenge"`
}
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

// Find credential
var cred Credential
for _, c := range wallet.Credentials {
if c.ID == req.CredentialID {
cred = c
break
}
}

if cred.ID == "" {
http.Error(w, "credential not found", http.StatusNotFound)
return
}

// Create VP
vp := map[string]interface{}{
"@context":           []string{"https://www.w3.org/2018/credentials/v1"},
"type":             []string{"VerifiablePresentation"},
"verifiableCredential": []interface{}{cred.Raw},
"holder":           wallet.OwnerDID,
"audience":         req.Audience,
"challenge":       req.Challenge,
"created":         time.Now().UTC().Format(time.RFC3339),
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(vp)
}

func main() {
log.Println("Wallet API starting on :3000")

http.HandleFunc("/health", healthHandler)
http.HandleFunc("/v1/wallets", createWalletHandler)
http.HandleFunc("/v1/wallets/{id}", getWalletHandler)
http.HandleFunc("/v1/wallets/{id}/credentials", addCredentialHandler)
http.HandleFunc("/v1/wallets/{id}/credentials", listCredentialsHandler)
http.HandleFunc("/v1/wallets/{id}/present", presentCredentialHandler)

log.Fatal(http.ListenAndServe(":3000", nil))
}
