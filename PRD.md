# PRD: DID One World Unified Identity Platform


## Product Name


**DID One World**


## Product Category


**Unified Decentralized Identity Lifecycle Management Platform**


## Core Thesis


Every entity in an enterprise should have a **canonical decentralized identity**.


That includes:


* Humans
* AI agents
* APIs
* Applications
* Services
* Skills
* Tools
* Devices
* Robots
* Workloads
* Organizations
* Teams
* Data assets
* Documents
* Credentials
* Physical things
* Digital twins
* Partner systems


DID One World becomes the **identity control plane for everything**.


---


# 1. Executive Summary


DID One World is an enterprise SaaS platform for creating, managing, governing, verifying, and revoking decentralized canonical identities for every entity in an organization.


The platform provides a single system of record for identity lifecycle management across humans, AI agents, APIs, skills, services, devices, things, and organizations.


Instead of managing identities separately across IAM, API gateways, secrets managers, agent platforms, device registries, HR systems, SaaS apps, and workflow tools, DID One World gives each entity a **portable, verifiable, cryptographically controlled decentralized identifier**.


This canonical identity becomes the foundation for:


* Lifecycle management
* Ownership
* Trust
* Authentication
* Authorization
* Delegation
* Credential issuance
* Verification
* Revocation
* Audit
* Compliance
* Interoperability


---


# 2. Vision


DID One World enables the enterprise to answer one critical question:


**"What is this entity, who owns it, what is it allowed to do, what credentials does it hold, what is its current lifecycle state, and can it be trusted?"**


The platform creates a universal identity graph where every entity has:


* A decentralized canonical ID
* A lifecycle state
* An owner
* A controller
* A trust profile
* Credentials
* Relationships
* Delegations
* Permissions
* Policy bindings
* Audit history


---


# 3. Product Positioning


## One-liner


**DID One World is the unified identity lifecycle platform for humans, agents, APIs, skills, and things.**


## Expanded Positioning


DID One World gives every enterprise entity a decentralized canonical identity and manages its full lifecycle from creation to verification, delegation, governance, and revocation.


## Category


**Decentralized Identity Lifecycle Management**


or


**Universal Identity Control Plane**


or


**Canonical Identity Platform for Everything**


---


# 4. Product Definition


DID One World is a production-grade enterprise SaaS platform that provides:


1. **Canonical Identity Registry**
   A single registry for all entity identities.


2. **Decentralized Identifier Management**
   DID creation, resolution, update, rotation, and deactivation.


3. **Lifecycle Management**
   Creation, approval, activation, suspension, review, expiration, revocation, and archival.


4. **Verifiable Credentials**
   Issue, hold, present, verify, renew, and revoke credentials.


5. **Trust Registry**
   Define trusted issuers, verifiers, credential schemas, DID methods, organizations, agents, APIs, and things.


6. **Policy Engine**
   Enforce rules across identity creation, verification, delegation, credential issuance, and lifecycle transitions.


7. **Identity Graph**
   Map relationships between humans, agents, APIs, skills, services, things, and organizations.


8. **Delegation Engine**
   Manage authority transfer between humans, agents, APIs, organizations, and things.


9. **Wallet Infrastructure**
   Wallets for humans, agents, organizations, APIs, and things.


10. **Audit & Compliance Layer**
    Immutable logs, evidence exports, reviews, and reporting.


---


# 5. Supported Identity Types


## 5.1 Human Identity


Represents employees, contractors, customers, developers, admins, partners, auditors, and external users.


Examples:


* Employee
* Contractor
* Customer
* Partner admin
* Developer
* Auditor
* Approver
* Operator


Capabilities:


* Employee credential
* Role credential
* Department credential
* Authorization credential
* Delegation to agent
* Wallet-based presentation
* Access reviews
* Lifecycle sync from HRIS or SCIM


---


## 5.2 Agent Identity


Represents AI agents, autonomous agents, workflow agents, copilots, automation bots, and software agents.


Examples:


* Customer support agent
* Sales agent
* Research agent
* Compliance agent
* Code generation agent
* Security agent
* Finance agent
* Data analysis agent


Capabilities:


* Agent DID
* Agent owner
* Agent sponsor
* Agent purpose
* Model/provider metadata
* Tool access
* Skill bindings
* Delegation from human or organization
* Runtime credential presentation
* Approval workflow
* Revocation and suspension


---


## 5.3 API Identity


Represents APIs, API clients, API products, integration endpoints, and exposed services.


Examples:


* Payments API
* Identity verification API
* Internal HR API
* Partner API
* Data access API
* Agent-facing API


Capabilities:


* API DID
* API ownership
* API trust profile
* Accepted credential policy
* Caller verification
* API credential issuance
* API-to-agent trust
* API lifecycle
* API deprecation and revocation


---


## 5.4 Skill Identity


Represents reusable capabilities that agents, humans, or systems can invoke.


Examples:


* Send email skill
* Run SQL query skill
* Approve invoice skill
* Search knowledge base skill
* Execute payment skill
* Generate legal summary skill
* Book meeting skill
* Deploy code skill


Capabilities:


* Skill DID
* Skill owner
* Skill version
* Required credentials
* Allowed callers
* Risk classification
* Input/output schema
* Execution policy
* Usage audit
* Deprecation and versioning


This is important because in an agentic ecosystem, **skills become permissioned identity-bearing capabilities**, not just functions.


---


## 5.5 Thing Identity


Represents physical devices, IoT devices, robots, sensors, machines, infrastructure, vehicles, chips, and digital twins.


Examples:


* IoT sensor
* Industrial robot
* Medical device
* Vehicle
* Smart lock
* Drone
* Manufacturing machine
* Edge device
* Physical asset
* Digital twin


Capabilities:


* Thing DID
* Manufacturer credential
* Ownership credential
* Maintenance credential
* Firmware credential
* Location credential
* Operational status
* Device lifecycle
* Compromise revocation
* Digital twin linkage


---


## 5.6 Application Identity


Represents SaaS apps, internal apps, microservices, platforms, and enterprise systems.


Examples:


* CRM app
* HR system
* ERP system
* Internal dashboard
* Data warehouse
* Workflow engine
* Agent orchestration platform


Capabilities:


* Application DID
* Environment-specific identity
* Service credentials
* API trust relationships
* Application ownership
* Deployment metadata
* Compliance profile
* Lifecycle state


---


## 5.7 Organization Identity


Represents companies, departments, business units, vendors, partners, issuers, and verifiers.


Examples:


* Enterprise tenant
* Subsidiary
* Department
* Vendor
* Partner organization
* Credential issuer
* Verification authority


Capabilities:


* Organization DID
* Legal identity credentials
* Trust registry membership
* Issuer authority
* Verifier authority
* Delegation rights
* Partner trust framework


---


## 5.8 Data Identity


Represents datasets, documents, models, knowledge bases, records, and data products.


Examples:


* Customer dataset
* Financial report
* Training dataset
* Knowledge base
* Contract
* Model artifact
* Compliance evidence file


Capabilities:


* Data DID
* Provenance credential
* Ownership credential
* Classification credential
* Access policy
* Lineage
* Versioning
* Retention policy
* Verification of origin


---


# 6. Canonical Identity Model


Every entity in DID One World must have a canonical identity record.


## Canonical Identity Record


```json
{
  "canonical_id": "did:web:example.com:entities:123",
  "entity_type": "human | agent | api | skill | thing | application | organization | data",
  "display_name": "string",
  "legal_name": "string",
  "tenant_id": "string",
  "workspace_id": "string",
  "environment": "development | staging | production",
  "owner": "did:web:example.com:users:owner",
  "controller": "did:web:example.com:controllers:controller",
  "sponsor": "did:web:example.com:org:sponsor",
  "status": "draft | pending_approval | active | suspended | expired | revoked | archived",
  "risk_level": "low | medium | high | critical",
  "assurance_level": "basic | verified | trusted | regulated",
  "credentials": [],
  "relationships": [],
  "delegations": [],
  "policies": [],
  "created_at": "ISO-8601",
  "updated_at": "ISO-8601",
  "expires_at": "ISO-8601",
  "last_reviewed_at": "ISO-8601",
  "revoked_at": "ISO-8601"
}
```


---


# 7. Universal Lifecycle Management


Every identity type follows a common lifecycle.


## Lifecycle States


| State            | Meaning                                |
| ---------------- | -------------------------------------- |
| Draft            | Identity record created but incomplete |
| Pending Approval | Awaiting approval before activation    |
| Approved         | Approved but not yet active            |
| Active           | Valid and usable                       |
| Under Review     | Flagged for manual or automated review |
| Suspended        | Temporarily disabled                   |
| Expiring Soon    | Near expiration or review deadline     |
| Expired          | No longer valid                        |
| Deprecated       | Should not be used for new activity    |
| Revoked          | Permanently invalid                    |
| Archived         | Retained for audit and history         |


## Lifecycle Events


The platform must support:


* Create identity
* Validate metadata
* Assign owner
* Assign controller
* Assign policies
* Submit for approval
* Approve identity
* Activate identity
* Issue credentials
* Bind relationships
* Delegate authority
* Review identity
* Suspend identity
* Rotate keys
* Renew credentials
* Expire identity
* Revoke identity
* Archive identity


---


# 8. Universal Identity Graph


DID One World must maintain a graph of relationships between identities.


## Example Relationships


Human owns Agent
Agent uses Skill
Skill calls API
API exposes Data
Thing produces Data
Application verifies Credential
Organization issues Credential
Human delegates authority to Agent
Agent acts on behalf of Organization
Device belongs to Organization
Dataset is derived from Dataset
Credential proves capability of Skill


## Relationship Types


* owns
* controls
* sponsors
* delegates_to
* acts_on_behalf_of
* uses
* exposes
* calls
* verifies
* issues
* holds
* depends_on
* derived_from
* belongs_to
* manages
* approves
* revokes


## Why This Matters


The identity graph allows enterprises to answer:


* Which agents can call this API?
* Which human owns this agent?
* Which skills can this agent use?
* Which credentials prove this thing is trusted?
* Which APIs depend on this credential issuer?
* What identities must be revoked if this key is compromised?
* What agents act on behalf of this business unit?
* What data has this skill accessed?


---


# 9. Verifiable Credentials for Everything


Every identity type can hold credentials.


## Credential Examples


### Human Credentials


* Employee credential
* Role credential
* Department credential
* Training credential
* Approval authority credential


### Agent Credentials


* Agent identity credential
* Agent approval credential
* Agent capability credential
* Agent risk clearance credential
* Agent delegation credential


### API Credentials


* API trust credential
* API authorization credential
* API compliance credential
* API publisher credential


### Skill Credentials


* Skill certification credential
* Skill execution permission credential
* Skill version credential
* Skill safety credential


### Thing Credentials


* Manufacturer credential
* Ownership credential
* Maintenance credential
* Firmware integrity credential
* Location credential
* Compliance credential


### Organization Credentials


* Legal entity credential
* Issuer credential
* Verifier credential
* Partner credential
* Trust framework membership credential


### Data Credentials


* Provenance credential
* Classification credential
* Ownership credential
* Consent credential
* Retention credential


---


# 10. Wallets for Every Entity


The platform must support identity wallets for all major entity types.


| Wallet Type         | Holder                                  |
| ------------------- | --------------------------------------- |
| Human Wallet        | Employee, customer, partner             |
| Agent Wallet        | AI agent or software agent              |
| API Wallet          | API/service endpoint                    |
| Skill Wallet        | Reusable capability or tool             |
| Thing Wallet        | Device, robot, sensor, physical asset   |
| Organization Wallet | Company, department, issuer             |
| Data Wallet         | Dataset, document, model, digital asset |


## Wallet Capabilities


* Store credentials
* Present credentials
* Rotate keys
* Manage status
* Receive credentials
* Prove ownership
* Prove capability
* Prove delegation
* Prove compliance
* Support managed custody
* Support self-custody where applicable


---


# 11. Policy Engine


The policy engine must apply to all identity types.


## Policy Examples


### Human


"Only employees with finance approval credentials can approve payments above $10,000."


### Agent


"Only approved production agents may invoke customer-facing APIs."


### API


"This API only accepts credentials from trusted issuers with regulated assurance level."


### Skill


"Payment execution skill requires human delegation and finance credential."


"Only devices with valid firmware integrity credentials may connect to the production network."


### Data


"Only agents with PII clearance credentials may access customer datasets."


---


# 12. Trust Registry


The Trust Registry defines what the enterprise trusts.


It must manage:


* Trusted humans
* Trusted agents
* Trusted APIs
* Trusted skills
* Trusted things
* Trusted organizations
* Trusted issuers
* Trusted verifiers
* Trusted DID methods
* Trusted credential schemas
* Trusted credential types
* Trusted trust frameworks
* Trusted environments


## Trust Decision Example


Before allowing an agent to use a skill that calls an API, the platform verifies:


1. Is the agent active?
2. Is the agent credential valid?
3. Is the skill trusted?
4. Is the API trusted?
5. Is the delegation valid?
6. Is the human sponsor still active?
7. Is the policy satisfied?
8. Is the credential issuer trusted?
9. Has anything been revoked?
10. Is the action within scope?


---


# 13. Delegation Engine


Delegation is core to the platform.


## Delegation Examples


* Human delegates task authority to agent.
* Organization delegates issuing authority to department.
* Agent delegates subtask to another agent.
* API delegates execution rights to skill.
* Device delegates telemetry rights to gateway.
* Partner organization delegates verification authority to service provider.


## Delegation Record


```json
{
  "delegation_id": "string",
  "delegator": "did:web:enterprise.com:human:123",
  "delegate": "did:web:enterprise.com:agent:456",
  "scope": ["read.customer_record", "summarize.case"],
  "purpose": "customer_support",
  "valid_from": "ISO-8601",
  "valid_until": "ISO-8601",
  "revocable": true,
  "status": "active | suspended | revoked | expired",
  "credential_id": "string"
}
```


---


# 14. Production Platform Modules


## Required Modules


1. **Canonical Identity Registry**
2. **DID Lifecycle Management**
3. **Universal Identity Lifecycle Engine**
4. **Identity Graph**
5. **Verifiable Credential Issuer**
6. **Verifiable Credential Verifier**
7. **Universal Wallet Infrastructure**
8. **Trust Registry**
9. **Policy Engine**
10. **Delegation Engine**
11. **Approval Workflow Engine**
12. **Revocation Engine**
13. **Audit & Compliance Engine**
14. **Admin Console**
15. **Developer Platform**
16. **API Gateway**
17. **Enterprise IAM Integrations**
18. **SIEM & Security Integrations**
19. **Billing & Usage Metering**
20. **Observability Platform**


---


# 15. Core User Flows


## Flow 1: Create Any Identity


1. Admin selects entity type.
2. Platform loads required schema.
3. Admin enters metadata.
4. Platform creates canonical DID.
5. Owner and controller are assigned.
6. Policies are attached.
7. Approval workflow runs.
8. Identity becomes active.
9. Initial credentials are issued.
10. Audit event is recorded.


---


## Flow 2: Human Delegates Authority to Agent


1. Human selects agent.
2. Human selects scope of authority.
3. Policy engine checks if delegation is allowed.
4. Approval workflow runs if required.
5. Delegation credential is issued.
6. Agent wallet receives credential.
7. Agent can now act within delegated scope.
8. All usage is audited.


---


## Flow 3: Agent Uses Skill to Call API


1. Agent requests skill execution.
2. Skill verifies agent identity.
3. Policy engine checks agent credential.
4. Skill checks delegation scope.
5. Skill calls API.
6. API verifies skill and agent credentials.
7. Trust registry confirms all issuers and entities are trusted.
8. Action is allowed or denied.
9. Audit log records full chain.


---


## Flow 4: Thing Proves Trust


1. Device connects to enterprise system.
2. Device presents manufacturer and firmware credentials.
3. System verifies DID and credentials.
4. Trust registry checks issuer.
5. Policy engine checks firmware/version requirements.
6. Device is allowed or blocked.
7. Event is logged.


---


## Flow 5: Revoke a Compromised Identity


1. Security admin selects identity.
2. Admin chooses emergency revocation.
3. Platform revokes credentials.
4. Platform updates lifecycle state.
5. Dependent delegations are revoked.
6. Related graph relationships are flagged.
7. Webhooks notify downstream systems.
8. Future verifications fail.
9. Compliance evidence is generated.


---


# 16. Admin Console Information Architecture


## Main Navigation


* Overview
* Identities
* Humans
* Agents
* APIs
* Skills
* Things
* Applications
* Organizations
* Data Assets
* Credentials
* Wallets
* Trust Registry
* Policies
* Delegations
* Approvals
* Identity Graph
* Audit Logs
* Compliance
* Developers
* Integrations
* Security
* Billing
* Settings


---


# 17. APIs


## Universal Identity APIs


POST   /v1/identities
GET    /v1/identities
GET    /v1/identities/{id}
PATCH  /v1/identities/{id}
POST   /v1/identities/{id}/activate
POST   /v1/identities/{id}/suspend
POST   /v1/identities/{id}/revoke
POST   /v1/identities/{id}/archive


## Type-Specific APIs


/v1/humans
/v1/agents
/v1/apis
/v1/skills
/v1/things
/v1/applications
/v1/organizations
/v1/data-assets


## DID APIs


POST /v1/dids
GET  /v1/dids/{did}
POST /v1/dids/{did}/rotate-key
POST /v1/dids/{did}/deactivate


## Credential APIs


POST /v1/credentials/issue
POST /v1/credentials/verify
POST /v1/credentials/revoke
POST /v1/presentations/create
POST /v1/presentations/verify


## Graph APIs


GET  /v1/graph/entities/{id}
POST /v1/graph/relationships
GET  /v1/graph/impact-analysis/{id}


## Delegation APIs


POST /v1/delegations
GET  /v1/delegations
POST /v1/delegations/{id}/revoke


---


# 18. Key Differentiator


Most identity platforms manage **users**.


Some manage **machines**.


Some manage **agents**.


DID One World manages **everything** using a shared decentralized canonical identity model.


That is the platform wedge.


---


# 19. Updated Product Taglines


Best options:


1. **Identity for Everything**
2. **The Universal Identity Control Plane**
3. **Canonical Identity for Humans, Agents, APIs, Skills, and Things**
4. **Decentralized Identity Lifecycle Management for Everything**
5. **The Trust Layer for the Agentic Enterprise**
6. **One Identity Graph for Every Entity**
7. **DID One World: Verifiable Identity for Everything**


Recommended primary tagline:


> **DID One World: Identity for Everything.**


Recommended enterprise positioning:


> **A unified decentralized identity lifecycle platform for humans, agents, APIs, skills, and things.**


---


# 20. Corrected PRD North Star


The north star is not:


> "Manage AI agent identity."


The north star is:


> **Give every entity in the enterprise a decentralized canonical identity, manage its lifecycle, verify its trust, govern its relationships, and revoke it instantly when trust changes.**


This is the strongest platform-level framing.