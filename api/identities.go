package main

import (
"crypto/ecdsa"
"crypto/elliptic"
"crypto/rand"
"crypto/x509"
"encoding/base64"
"encoding/json"
"log"
"net/http"
"strings"
"time"

"github.com/google/uuid"
)

// W3C DID Core v1.0 (W3C Recommendation)
// https://www.w3.org/TR/did-core/

// DIDDocument (Section 5.2)
type DIDDocument struct {
Context                []string             `json:"@context,omitempty"` // "https://www.w3.org/ns/did/v1"
ID                    string              `json:"id"` // DID URL
AlsoKnownAs           []string            `json:"alsoKnownAs,omitempty"`
Controller            string             `json:"controller,omitempty"`
VerificationMethod    []VerificationMethod `json:"verificationMethod,omitempty"`
Authentication        []interface{}      `json:"authentication,omitempty"`
AssertionMethod       []interface{}      `json:"assertionMethod,omitempty"`
KeyAgreement          []interface{}      `json:"keyAgreement,omitempty"`
CapabilityInvocation []interface{}      `json:"capabilityInvocation,omitempty"`
CapabilityDelegation  []interface{}      `json:"capabilityDelegation,omitempty"`
Service               []Service          `json:"service,omitempty"`
// Metadata
Created string `json:"created,omitempty"`
Updated string `json:"updated,omitempty"`
}

type VerificationMethod struct {
ID                string         `json:"id"`
Type             string         `json:"type"` // "JsonWebKey2020" | "EcdsaSecp256k1VerificationKey2019"
Controller       string         `json:"controller"`
PublicKeyJwk      *JWK           `json:"publicKeyJwk,omitempty"`
PublicKeyMultibase string        `json:"publicKeyMultibase,omitempty"`
}

// JWK for JsonWebKey2020 (RFC 7517)
type JWK struct {
Kty string `json:"kty"` // "EC" | "OKP"
Crv string `json:"crv,omitempty"` // "P-256" | "secp256k1" | "Ed25519"
X   string `json:"x,omitempty"`
Y   string `json:"y,omitempty"`
D   string `json:"d,omitempty"` // private key (only for signing, never exported)
}

// Service (Section 5.3)
type Service struct {
ID              string `json:"id"`
Type            string `json:"type"` // "LinkedDomains" | "DIDCommMessaging" | "VerifiableCredentialService"
ServiceEndpoint string `json:"serviceEndpoint"` // URI
}

// DID Resolution Result (Section 7.1)
type DIDResolutionMetadata struct {
ContentType       string `json:"contentType,omitempty"` // "application/did+json"
Error           string `json:"error,omitempty"` // "invalidDid" | "notFound" | "representationNotSupported"
ResolutionTime  string `json:"resolutionTime,omitempty"`
}

type DIDResolutionResult struct {
Context   string            `json:"@context,omitempty"`
Metadata DIDResolutionMetadata `json:"metadata"`
Document *DIDDocument      `json:"document,omitempty"`
}

// DID List Result
type DIDListResult struct {
Contexts []string   `json:"@context,omitempty"`
Items    []Identity `json:"identities"`
}

// Identity (Internal)
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
Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// Storage
var (
identities    = make(map[string]Identity)
didDocuments  = make(map[string]DIDDocument)
keyPairs     = make(map[string]*ecdsa.PrivateKey)
)

// Handlers

func healthHandler(w http.ResponseWriter, r *http.Request) {
json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

// Resolve DID (Section 3.2)
func resolveHandler(w http.ResponseWriter, r *http.Request) {
did := r.PathValue("did")

doc := findDIDDocument(did)
if doc == nil {
w.Header().Set("Content-Type", "application/did+json")
w.WriteHeader(http.StatusNotFound)
json.NewEncoder(w).Encode(DIDResolutionResult{
Metadata: DIDResolutionMetadata{
Error: "notFound",
},
})
return
}

w.Header().Set("Content-Type", "application/did+json")
json.NewEncoder(w).Encode(DIDResolutionResult{
Context: "https://www.w3.org/ns/did/v1",
Metadata: DIDResolutionMetadata{
ContentType: "application/did+json",
},
Document: doc,
})
}

func createIdentityHandler(w http.ResponseWriter, r *http.Request) {
var req struct {
Type        string   `json:"type"`
DisplayName string  `json:"displayName"`
Owner      string   `json:"owner"`
Controller string  `json:"controller,omitempty"`
Services   []Service `json:"services,omitempty"`
}
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, err.Error(), http.StatusBadRequest)
return
}

// Generate DID using web method
did := "did:web:didone.world:" + req.Type + ":" + uuid.New().String()[:8]

// Generate key pair (P-256 for JWK)
key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
if err != nil {
http.Error(w, "failed to generate key", http.StatusInternalServerError)
return
}

pub := &key.PublicKey
pubB64 := base64.RawURLEncoding.EncodeToString([]byte(pub.X.String() + "." + pub.Y.String()))

controller := req.Controller
if controller == "" {
controller = req.Owner
}

now := time.Now().UTC().Format(time.RFC3339)

// Create verification method with JWK
vm := VerificationMethod{
ID:          did + "#key-1",
Type:        "JsonWebKey2020",
Controller: controller,
PublicKeyJwk: &JWK{
Kty: "EC",
Crv: "P-256",
X:   pubB64,
},
}

// Build DID Document
doc := DIDDocument{
Context: []string{"https://www.w3.org/ns/did/v1"},
ID:      did,
Controller: controller,
VerificationMethod: []VerificationMethod{vm},
Authentication:      []interface{}{did + "#key-1"},
AssertionMethod:    []interface{}{did + "#key-1"},
Service:           req.Services,
Created: now,
Updated: now,
}

identity := Identity{
ID:           uuid.New().String(),
DID:          did,
Type:         req.Type,
DisplayName: req.DisplayName,
Owner:       req.Owner,
Controller:  controller,
Status:      "active",
CreatedAt:   now,
UpdatedAt:   now,
}

identities[did] = identity
didDocuments[did] = doc
keyPairs[did] = key

w.Header().Set("Content-Type", "application/did+json")
w.WriteHeader(http.StatusCreated)
json.NewEncoder(w).Encode(identity)
}

func listIdentitiesHandler(w http.ResponseWriter, r *http.Request) {
q := r.URL.Query()
identityType := q.Get("type")
status := q.Get("status")

var list []Identity
for _, id := range identities {
if identityType != "" && id.Type != identityType { continue }
if status != "" && id.Status != status { continue }
list = append(list, id)
}

w.Header().Set("Content-Type", "application/did+json")
json.NewEncoder(w).Encode(DIDListResult{
Contexts: []string{"https://www.w3.org/ns/did/v1"},
Items:    list,
})
}

func getIdentityHandler(w http.ResponseWriter, r *http.Request) {
did := r.PathValue("did")
doc := findDIDDocument(did)
if doc == nil || doc.ID == "" {
http.Error(w, "identity not found", http.StatusNotFound)
return
}
identity := identities[doc.ID]
json.NewEncoder(w).Encode(identity)
}

func updateIdentityHandler(w http.ResponseWriter, r *http.Request) {
did := findDIDDocument(r.PathValue("did"))
if did == nil || did.ID == "" {
http.Error(w, "identity not found", http.StatusNotFound)
return
}

identity := identities[did.ID]
var req struct {
Status     string `json:"status"`
DisplayName string `json:"displayName,omitempty"`
}
json.NewDecoder(r.Body).Decode(&req)

if req.Status != "" {
identity.Status = req.Status
}
if req.DisplayName != "" {
identity.DisplayName = req.DisplayName
}
identity.UpdatedAt = time.Now().UTC().Format(time.RFC3339)

identities[identity.DID] = identity
didDocuments[identity.DID] = *did
didDocuments[identity.DID].Updated = identity.UpdatedAt

json.NewEncoder(w).Encode(identity)
}

func findDIDDocument(did string) *DIDDocument {
// Try direct lookup
if doc, ok := didDocuments[did]; ok {
return &doc
}
// Try with did: prefix
if doc, ok := didDocuments["did:"+did]; ok {
return &doc
}
cleanDID := strings.TrimPrefix(did, "did:")
if doc, ok := didDocuments["did:"+cleanDID]; ok {
return &doc
}
// Search by suffix
for _, doc := range didDocuments {
if strings.HasSuffix(doc.ID, ":"+cleanDID) {
return &doc
}
}
return nil
}

func main() {
log.Println("DID Registry (W3C DID Core v1.0 (W3C Recommendation) on :8000")

http.HandleFunc("/health", healthHandler)
http.HandleFunc("/v1/identities", listIdentitiesHandler)
http.HandleFunc("/v1/identities", createIdentityHandler)
http.HandleFunc("/v1/identities/{did}", getIdentityHandler)
http.HandleFunc("/v1/identities/{did}", updateIdentityHandler)
http.HandleFunc("/v1/did/{did}", resolveHandler)

log.Fatal(http.ListenAndServe(":8000", nil))
}
