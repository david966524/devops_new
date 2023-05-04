package cloudflaremodel

import (
	"time"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

// Zone describes a Cloudflare zone.
type Zone struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// DevMode contains the time in seconds until development expires (if
	// positive) or since it expired (if negative). It will be 0 if never used.
	DevMode           int                 `json:"development_mode"`
	OriginalNS        []string            `json:"original_name_servers"`
	OriginalRegistrar string              `json:"original_registrar"`
	OriginalDNSHost   string              `json:"original_dnshost"`
	CreatedOn         time.Time           `json:"created_on"`
	ModifiedOn        time.Time           `json:"modified_on"`
	NameServers       []string            `json:"name_servers"`
	Owner             cloudflare.Owner    `json:"owner"`
	Permissions       []string            `json:"permissions"`
	Plan              cloudflare.ZonePlan `json:"plan"`
	PlanPending       cloudflare.ZonePlan `json:"plan_pending,omitempty"`
	Status            string              `json:"status"`
	Paused            bool                `json:"paused"`
	Type              string              `json:"type"`
	Host              struct {
		Name    string
		Website string
	} `json:"host"`
	VanityNS        []string            `json:"vanity_name_servers"`
	Betas           []string            `json:"betas"`
	DeactReason     string              `json:"deactivation_reason"`
	Meta            cloudflare.ZoneMeta `json:"meta"`
	Account         cloudflare.Account  `json:"account"`
	VerificationKey string              `json:"verification_key"`
}
