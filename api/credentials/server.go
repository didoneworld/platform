package main

import (
"crypto/ecdsa"
"crypto/elliptic"
"crypto/rand"
"crypto/sha256"
"crypto/x509"
"encoding/base64"
"encoding/json"
"fmt"
"log"
"net/http"
"time"

"github.com/google/uuid"
"github.com/lestrrat-go/jwx/v2/jwa"
"github.com/lestrrat-go/jwx/v2/jws"
)

// DID Document
type DIDDocument struct {
ID            string   `json:"id"`
Type          string   `json:"@type"`
Controller    string   `json:"controller"`
VerificationMethod []VerificationMethod `json:"verificationMethod"`
Authentication []string `json:"authentication"`
}

type VerificationMethod struct {
ID                 string `json:"id"`
Type               string `json:"type"`
Controller        string `json:"controller"`
PublicKeyJWK      map[string]interface{} `json:"publicKeyJwk"`
}

// Verifiable Credential
type VerifiableCredential struct {
Context           []string      `json:"@context"`
ID               string      `json:"id"`
Type             []string    `json:"type"`
Issuer           string      `json:"issuer"`
IssuanceDateTime string      `json:"issuanceDateTime"`
ExpirationDateTime string   `json:"expirationDateTime"`
CredentialSubject CredentialSubject `json:"credentialSubject"`
Proof            *Proof     `json:"proof,omitempty"`
}

type CredentialSubject struct {
ID       string                 `json:"id"`
Type    string                 `json:"type"`
Claims  map[string]interface{} `json:"claims"`
}

type Proof struct {
Type               string `json:"type"`
Created            string `json:"created"`
ProofPurpose       string `json:"proofPurpose"`
VerificationMethod string `json:"verificationMethod"`
JWS               string `json:"jws"`
}

// API Request/Response
type IssueRequest struct {
HolderDID     string                 `json:"holderDid"`
IssuerDID     string                 `json:"issuerDid"`
CredentialType string                 `json:"credentialType"`
Claims       map[string]interface{} `json:"claims"`
Expiry       string                 `json:"expiry"`
}

type IssueResponse struct {
Credential VerifiableCredential `json:"credential"`
}

type VerifyRequest struct {
Presentation *json.RawMessage `json:"presentation"`
Credential  *json.RawMessage `json:"credential"`
}

type VerifyResponse struct {
Verified bool   `json:"verified"`
Reason   string `json:"reason"`
}

type RevokeRequest struct {
CredentialID string `json:"credentialId"`
}

type ListResponse struct {
Credentials []VerifiableCredential `json:"credentials"`
}

// In-memory storage
var credentials = make(map[string]VerifiableCredential)
var signingKey *ecdsa.PrivateKey

func main() {
// Generate signing key
var err error
signingKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
if err != nil {
log.Fatal(err)
}

// Routes
http.HandleFunc("/health", healthHandler)
http.HandleFunc("/v1/credentials/issue", issueHandler)
http.HandleFunc("/v1/credentials/verify", verifyHandler)
http.HandleFunc("/v1/credentials/revoke", revokeHandler)
http.HandleFunc("/v1/credentials", listHandler)

log.Println("VC Issuer starting on :8001")
log.Fatal(http.ListenAndServe(":8001", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func issueHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

var req IssueRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

// Generate credential ID
vcID := "urn:uuid:" + uuid.New().String()
now := time.Now().UTC().Format(time.RFC3339)

expiry := req.Expiry
if expiry == "" {
expiry = time.Now().Add(365 * 24 * time.Hour).UTC().Format(time.RFC3339)
}

vc := VerifiableCredential{
Context: []string{"https://www.w3.org/2018/credentials/v1", "https://didone.world/contexts/v1"},
ID: vcID,
Type: []string{"VerifiableCredential", req.CredentialType},
Issuer: req.IssuerDID,
IssuanceDateTime: now,
ExpirationDateTime: expiry,
CredentialSubject: CredentialSubject{
ID: req.HolderDID,
Type: req.CredentialType,
Claims: req.Claims,
},
}

// Create proof
proof, err := createProof(vc, req.IssuerDID)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}
vc.Proof = proof

// Store
credentials[vcID] = vc

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(IssueResponse{Credential: vc})
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

var req VerifyRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

var vc VerifiableCredential
if req.Credential != nil {
json.Unmarshal(*req.Credential, &vc)
}

// Basic verification
verified := vc.ID != "" && vc.Proof != nil
reason := "credential verified"

if !verified {
reason = "invalid credential"
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(VerifyResponse{Verified: verified, Reason: reason})
}

func revokeHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

var req RevokeRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

delete(credentials, req.CredentialID)

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(map[string]string{"status": "revoked"})
}

func listHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodGet {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

var list []VerifiableCredential
for _, vc := range credentials {
list = append(list, vc)
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(ListResponse{Credentials: list})
}

func createProof(vc VerifiableCredential, issuerDID string) (*Proof, error) {
// Serialize VC
payload, _ := json.Marshal(vc)

// Sign
signed, err := jws.Sign(payload, jwa.ES256, signingKey)
if err != nil {
return nil, err
}

return &Proof{
Type:               "EcdsaSecp256k1Signature2019",
Created:            time.Now().UTC().Format(time.RFC3339),
ProofPurpose:       "assertionMethod",
VerificationMethod: issuerDID + "#key-1",
JWS:               base64.RawURLEncoding.EncodeToString(signed),
}, nil
}
