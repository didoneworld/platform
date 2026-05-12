package main

import (
"crypto/ecdsa"
"crypto/elliptic"
"crypto/rand"
"crypto/sha256"
"crypto/x509"
"encoding/base64"
"encoding/json"
"log"
"net/http"
"strings"
"time"

"github.com/google/uuid"
)

// W3C DID Core Specification
// https://www.w3.org/TR/did-core/

// DID URL
type DID string

func (d DID) Method() string {
parts := strings.Split(string(d), ":")
if len(parts) >= 2 {
return parts[1]
}
return ""
}

func (d DID) ID() string {
parts := strings.Split(string(d), ":")
if len(parts) >= 3 {
return parts[2]
}
return ""
}

// DID Document (W3C spec section 5.2)
type DIDDocument struct {
Context           []string             `json:"@context"`
ID                string              `json:"id"`
AlsoKnownAs       []string            `json:"alsoKnownAs,omitempty"`
Controller       string             `json:"controller,omitempty"`
VerificationMethod []VerificationMethod `json:"verificationMethod,omitempty"`
Authentication    []VerificationMethod `json:"authentication,omitempty"`
AssertionMethod  []VerificationMethod `json:"assertionMethod,omitempty"`
KeyAgreement    []VerificationMethod `json:"keyAgreement,omitempty"`
CapabilityInvocation []VerificationMethod `json:"capabilityInvocation,omitempty"`
CapabilityDelegation []VerificationMethod `json:"capabilityDelegation,omitempty"`
Service         []Service            `json:"service,omitempty"`
Created         string              `json:"created"`
Updated         string              `json:"updated"`
Proof           *Proof             `json:"proof,omitempty"`
}

type VerificationMethod struct {
ID           string `json:"id"`
Type        string `json:"type"` // "EcdsaSecp256k1VerificationKey2019" | "JsonWebKey2020"
Controller  string `json:"controller"`
PublicKeyJwk *JWK  `json:"publicKeyJwk,omitempty"`
PublicKeyMultibase string `json:"publicKeyMultibase,omitempty"`
}

type JWK struct {
Kty string `json:"kty"` // "EC"
Crv string `json:"crv"` // "P-256" | "secp256k1"
X   string `json:"x"`
Y   string `json:"y"`
}

type Service struct {
ID              string `json:"id"`
Type            string `json:"type"` // "LinkedService" | "DIDCommMessaging"
ServiceEndpoint string `json:"serviceEndpoint"`
}

type Proof struct {
Type               string `json:"type"`
Created           string `json:"created"`
VerificationMethod string `json:"verificationMethod"`
ProofPurpose      string `json:"proofPurpose"`
ProofValue        string `json:"proofValue"`
}

// DID Core Representation (section 6)
type DIDResolutionMetadata struct {
ContentType string `json:"contentType,omitempty"`
Error       string `json:"error,omitempty"`
}

type DIDResolutionResult struct {
Context    string     `json:"@context"`
Metadata  DIDResolutionMetadata `json:"metadata"`
Document  *DIDDocument `json:"document,omitempty"`
}

type DIDListResponse struct {
Context    string      `json:"@context"`
Identities []Identity `json:"identities"`
}

// Identity (our internal representation)
type Identity struct {
ID           string                 `json:"id"`
DID          string                 `json:"did"`
Type         string                 `json:"type"` // human, agent, api, skill, thing, application, organization
DisplayName string                 `json:"displayName"`
Owner        string                 `json:"owner"`
Controller  string                 `json:"controller"`
Status      string                 `json:"status"` // draft, active, suspended, revoked
CreatedAt   string                 `json:"createdAt"`
UpdatedAt   string                 `json:"updatedAt"`
Metadata    map[string]interface{}   `json:"metadata,omitempty"`
}

// Storage
var identities = make(map[string]Identity)
var didDocuments = make(map[string]DIDDocument)

// Handlers

func healthHandler(w http.ResponseWriter, r *http.Request) {
json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

// Resolve DID (W3C spec section 3.2)
func resolveHandler(w http.ResponseWriter, r *http.Request) {
did := r.PathValue("did")

// Find identity
identity, ok := identities[did]
if !ok {
// Try with did: prefix stripped
cleanDID := strings.TrimPrefix(did, "did:")
identity, ok = identities["did:"+cleanDID]
if !ok {
// Try lookup by ID portion
for _, id := range identities {
if strings.HasSuffix(id.DID, ":"+cleanDID) {
identity = id
ok = true
break
}
}
}
}

if !ok {
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusNotFound)
json.NewEncoder(w).Encode(DIDResolutionResult{
Context: "https://www.w3.org/ns/did/v1",
Metadata: DIDResolutionMetadata{
Error: "notFound",
},
})
return
}

doc := toDIDDocument(identity)

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(DIDResolutionResult{
Context: "https://www.w3.org/ns/did/v1",
Metadata: DIDResolutionMetadata{
ContentType: "application/did+json",
},
Document: &doc,
})
}

func createIdentityHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

var req struct {
Type        string `json:"type"`
DisplayName string `json:"displayName"`
Owner       string `json:"owner"`
Controller string `json:"controller,omitempty"`
AlsoKnownAs []string `json:"alsoKnownAs,omitempty"`
Services   []Service `json:"services,omitempty"`
}
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

// Generate DID using web method (W3C spec section 3.1)
uuidStr := uuid.New().String()[:8]
did := "did:web:didone.world:" + req.Type + ":" + uuidStr

// Generate key pair
key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
if err != nil {
http.Error(w, "failed to generate key", http.StatusInternalServerError)
return
}

pubBytes, _ := x509.MarshalPKIXPublicKey(&key.PublicKey)
pubKeyBase64 := base64.RawURLEncoding.EncodeToString(pubBytes)

controller := req.Controller
if controller == "" {
controller = req.Owner
}

now := time.Now().UTC().Format(time.RFC3339)

// Create verification method
vm := VerificationMethod{
ID:          did + "#key-1",
Type:        "EcdsaSecp256k1VerificationKey2019",
Controller: controller,
PublicKeyMultibase: "z" + pubKeyBase64, // multibase format
}

// Build DID Document
doc := DIDDocument{
Context: []string{"https://www.w3.org/ns/did/v1"},
ID:       did,
Controller: controller,
VerificationMethod: []VerificationMethod{vm},
Authentication: []VerificationMethod{vm},
Created: now,
Updated: now,
}

// Add services
for _, svc := range req.Services {
svc.ID = did + "#" + svc.ID
doc.Service = append(doc.Service, svc)
}

identity := Identity{
ID:           uuid.New().String(),
DID:          did,
Type:         req.Type,
DisplayName: req.DisplayName,
Owner:       req.Owner,
Controller: controller,
Status:      "active",
CreatedAt:  now,
UpdatedAt:  now,
Metadata:   map[string]interface{}{"alsoKnownAs": req.AlsoKnownAs},
}

identities[did] = identity
didDocuments[did] = doc

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(identity)
}

func listIdentitiesHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodGet {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

q := r.URL.Query()
identityType := q.Get("type")
status := q.Get("status")

var list []Identity
for _, id := range identities {
if identityType != "" && id.Type != identityType {
continue
}
if status != "" && id.Status != status {
continue
}
list = append(list, id)
}

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(DIDListResponse{
Context:    "https://www.w3.org/ns/did/v1",
Identities: list,
})
}

func getIdentityHandler(w http.ResponseWriter, r *http.Request) {
did := r.PathValue("did")
did = strings.TrimPrefix(did, "did:")

identity, ok := identities["did:"+did]
if !ok {
identity, ok = identities[did]
}
if !ok {
// Search by suffix
for _, id := range identities {
if strings.HasSuffix(id.DID, ":"+did) {
identity = id
ok = true
break
}
}
}

if !ok {
http.Error(w, "identity not found", http.StatusNotFound)
return
}

json.NewEncoder(w).Encode(identity)
}

func updateIdentityHandler(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPatch {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

did := r.PathValue("did")
did = strings.TrimPrefix(did, "did:")

identity, ok := identities["did:"+did]
if !ok {
identity, ok = identities[did]
if !ok {
for _, id := range identities {
if strings.HasSuffix(id.DID, ":"+did) {
identity = id
ok = true
break
}
}
}
}

if !ok {
http.Error(w, "identity not found", http.StatusNotFound)
return
}

var req struct {
Status string `json:"status"`
DisplayName string `json:"displayName,omitempty"`
}
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

if req.Status != "" {
identity.Status = req.Status
}
if req.DisplayName != "" {
identity.DisplayName = req.DisplayName
}
identity.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
identities[identity.DID] = identity

json.NewEncoder(w).Encode(identity)
}

// Helper functions

func toDIDDocument(identity Identity) DIDDocument {
return didDocuments[identity.DID]
}

func main() {
log.Println("DID Registry starting on :8000")
log.Println("Following W3C DID Core Specification")

http.HandleFunc("/health", healthHandler)
http.HandleFunc("/v1/identities", listIdentitiesHandler)
http.HandleFunc("/v1/identities", createIdentityHandler)
http.HandleFunc("/v1/identities/{did}", getIdentityHandler)
http.HandleFunc("/v1/identities/{did}", updateIdentityHandler)
// DID resolution endpoint (W3C spec)
http.HandleFunc("/v1/did/{did}", resolveHandler)

log.Fatal(http.ListenAndServe(":8000", nil))
}
