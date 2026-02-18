package schema

var InternalSchema = []map[string]interface{}{

	// ============================================================
	// ACCOUNT
	// Source-mappable entity (internal: false)
	// ============================================================
	{
		"entity_name": "Account",
		"description": "Financial account capable of holding or moving value.",
		"internal":    false, // shown on source-to-internal mapping canvas
		"fields": []map[string]interface{}{
			{"name": "account_number", "data_type": "string", "required": false, "internal": false},
			{"name": "account_type", "data_type": "string", "required": false, "internal": false},     // savings, checking, brokerage
			{"name": "status", "data_type": "string", "required": false, "internal": false},           // active, frozen, closed
			{"name": "currency", "data_type": "string", "required": false, "internal": false},
			{"name": "current_balance", "data_type": "float", "required": false, "internal": false},
			{"name": "available_balance", "data_type": "float", "required": false, "internal": false},
			{"name": "opened_at", "data_type": "datetime", "required": false, "internal": false},
			{"name": "closed_at", "data_type": "datetime", "required": false, "internal": false},
			{"name": "country_of_issue", "data_type": "string", "required": false, "internal": false},
			{"name": "owner_customer_id", "data_type": "string", "required": true, "internal": false},  // primary edge anchor to Customer
			{"name": "kyc_verified", "data_type": "bool", "required": true, "internal": false},

			// INTERNAL FRAUD STATE
			{"name": "risk_score", "data_type": "float", "required": false, "internal": true},
			{"name": "risk_level", "data_type": "string", "required": false, "internal": true},
			{"name": "risk_last_evaluated_at", "data_type": "datetime", "required": false, "internal": true},
			{"name": "flag_reason", "data_type": "string", "required": false, "internal": true},  // replaces is_flagged - too vague
			{"name": "flagged_by", "data_type": "string", "required": false, "internal": true},   // system, analyst
		},
	},

	// ============================================================
	// CUSTOMER
	// Source-mappable entity (internal: false)
	// ============================================================
	{
		"entity_name": "Customer",
		"description": "Natural or legal person associated with financial activity.",
		"internal":    false,
		"fields": []map[string]interface{}{
			{"name": "customer_type", "data_type": "string", "required": false, "internal": false}, // individual, business
			{"name": "first_name", "data_type": "string", "required": false, "internal": false},
			{"name": "last_name", "data_type": "string", "required": false, "internal": false},
			{"name": "date_of_birth", "data_type": "datetime", "required": false, "internal": false},
			{"name": "nationality", "data_type": "string", "required": false, "internal": false},          // distinct from residence
			{"name": "country_of_residence", "data_type": "string", "required": false, "internal": false},
			{"name": "onboarding_channel", "data_type": "string", "required": false, "internal": false},  // web, branch, api, partner
			{"name": "onboarding_date", "data_type": "datetime", "required": false, "internal": false},   // account age = strong signal
			{"name": "kyc_verified", "data_type": "bool", "required": true, "internal": false},
			{"name": "is_sanctioned", "data_type": "bool", "required": false, "internal": false},
			{"name": "is_pep", "data_type": "bool", "required": false, "internal": false},

			// INTERNAL
			{"name": "risk_score", "data_type": "float", "required": false, "internal": true},
			{"name": "risk_level", "data_type": "string", "required": false, "internal": true},
			{"name": "behavioral_deviation_score", "data_type": "float", "required": false, "internal": true},
			{"name": "behavioral_deviation_window_days", "data_type": "int", "required": false, "internal": true}, // context for deviation score
			{"name": "risk_last_evaluated_at", "data_type": "datetime", "required": false, "internal": true},
		},
	},

	// ============================================================
	// TRANSACTION
	// Source-mappable entity (internal: false)
	// ============================================================
	{
		"entity_name": "Transaction",
		"description": "Movement of value between entities.",
		"internal":    false,
		"fields": []map[string]interface{}{
			{"name": "transaction_type", "data_type": "string", "required": false, "internal": false}, // transfer, payment, withdrawal
			{"name": "amount", "data_type": "float", "required": true, "internal": false},
			{"name": "currency", "data_type": "string", "required": false, "internal": false},
			{"name": "original_currency", "data_type": "string", "required": false, "internal": false}, // pre-conversion currency
			{"name": "exchange_rate", "data_type": "float", "required": false, "internal": false},
			{"name": "timestamp", "data_type": "datetime", "required": true, "internal": false},
			{"name": "channel", "data_type": "string", "required": false, "internal": false}, // web, mobile, atm, api
			{"name": "authorization_status", "data_type": "string", "required": false, "internal": false},
			{"name": "failure_reason", "data_type": "string", "required": false, "internal": false},      // declined txns are often more informative
			{"name": "sender_account_id", "data_type": "string", "required": true, "internal": false},   // primary graph edge
			{"name": "receiver_account_id", "data_type": "string", "required": true, "internal": false}, // primary graph edge
			{"name": "counterparty_name", "data_type": "string", "required": false, "internal": false},
			{"name": "counterparty_bank_bic", "data_type": "string", "required": false, "internal": false},
			{"name": "merchant_id", "data_type": "string", "required": false, "internal": false},
			{"name": "merchant_category_code", "data_type": "string", "required": false, "internal": false}, // MCC - rich fraud signal
			{"name": "is_international", "data_type": "bool", "required": false, "internal": false},
			{"name": "reversal_of_transaction_id", "data_type": "string", "required": false, "internal": false}, // refund fraud, chargeback
			{"name": "three_ds_status", "data_type": "string", "required": false, "internal": false},            // passed, failed, not_used
			{"name": "session_id", "data_type": "string", "required": false, "internal": false},                 // graph edge to Session
			{"name": "geo_lat", "data_type": "float", "required": false, "internal": false},
			{"name": "geo_long", "data_type": "float", "required": false, "internal": false},
			{"name": "geo_accuracy_meters", "data_type": "float", "required": false, "internal": false}, // GPS vs IP-derived differ significantly

			// INTERNAL FRAUD OUTPUT
			{"name": "risk_score", "data_type": "float", "required": false, "internal": true},
			{"name": "decision", "data_type": "string", "required": false, "internal": true}, // approve, review, block
			{"name": "model_version", "data_type": "string", "required": false, "internal": true},
		},
	},

	// ============================================================
	// DEVICE
	// Source-mappable entity (internal: false)
	// ============================================================
	{
		"entity_name": "Device",
		"description": "Physical or virtual device used to access system.",
		"internal":    false,
		"fields": []map[string]interface{}{
			{"name": "device_fingerprint", "data_type": "string", "required": true, "internal": false},
			{"name": "device_type", "data_type": "string", "required": false, "internal": false},
			{"name": "os", "data_type": "string", "required": false, "internal": false},
			{"name": "browser", "data_type": "string", "required": false, "internal": false},
			{"name": "user_agent", "data_type": "string", "required": false, "internal": false},          // raw UA for fingerprint cross-validation
			{"name": "screen_resolution", "data_type": "string", "required": false, "internal": false},  // fingerprint component
			{"name": "timezone", "data_type": "string", "required": false, "internal": false},            // fingerprint component
			{"name": "language", "data_type": "string", "required": false, "internal": false},            // fingerprint component
			{"name": "is_emulator", "data_type": "bool", "required": false, "internal": false},
			{"name": "is_rooted_or_jailbroken", "data_type": "bool", "required": false, "internal": false},
			{"name": "vpn_detected", "data_type": "bool", "required": false, "internal": false},          // device-level VPN, distinct from IP proxy
			{"name": "first_seen_at", "data_type": "datetime", "required": false, "internal": false},
			{"name": "last_seen_at", "data_type": "datetime", "required": false, "internal": false},

			// INTERNAL
			{"name": "lifetime_trust_score", "data_type": "float", "required": false, "internal": true}, // long-term accumulated trust
			{"name": "risk_score", "data_type": "float", "required": false, "internal": true},           // moment-in-time risk
			{"name": "associated_customer_count", "data_type": "int", "required": false, "internal": true}, // many customers = device farm signal
		},
	},

	// ============================================================
	// IP ADDRESS
	// Source-mappable entity (internal: false)
	// ============================================================
	{
		"entity_name": "IPAddress",
		"description": "Observed IP address.",
		"internal":    false,
		"fields": []map[string]interface{}{
			{"name": "ip_address", "data_type": "string", "required": true, "internal": false},
			{"name": "country", "data_type": "string", "required": false, "internal": false},
			{"name": "city", "data_type": "string", "required": false, "internal": false},
			{"name": "region", "data_type": "string", "required": false, "internal": false},
			{"name": "latitude", "data_type": "float", "required": false, "internal": false}, // for geo-velocity checks
			{"name": "longitude", "data_type": "float", "required": false, "internal": false},
			{"name": "asn", "data_type": "string", "required": false, "internal": false},
			{"name": "isp", "data_type": "string", "required": false, "internal": false},
			{"name": "is_proxy", "data_type": "bool", "required": false, "internal": false},
			{"name": "is_tor_exit_node", "data_type": "bool", "required": false, "internal": false}, // major fraud vector
			{"name": "is_vpn", "data_type": "bool", "required": false, "internal": false},           // distinct from proxy
			{"name": "is_hosting_provider", "data_type": "bool", "required": false, "internal": false},
			{"name": "abuse_confidence_score", "data_type": "float", "required": false, "internal": false}, // from threat intel feeds
			{"name": "first_seen_at", "data_type": "datetime", "required": false, "internal": false},
			{"name": "last_seen_at", "data_type": "datetime", "required": false, "internal": false},

			// INTERNAL
			{"name": "risk_score", "data_type": "float", "required": false, "internal": true},
			{"name": "associated_customer_count", "data_type": "int", "required": false, "internal": true}, // 1 IP : many customers = bot/farm signal
			{"name": "associated_account_count", "data_type": "int", "required": false, "internal": true},
		},
	},

	// ============================================================
	// SESSION
	// Source-mappable entity (internal: false)
	// ============================================================
	{
		"entity_name": "Session",
		"description": "Authenticated interaction window.",
		"internal":    false,
		"fields": []map[string]interface{}{
			{"name": "session_start", "data_type": "datetime", "required": false, "internal": false},
			{"name": "session_end", "data_type": "datetime", "required": false, "internal": false},
			{"name": "session_duration_seconds", "data_type": "float", "required": false, "internal": false},
			{"name": "customer_id", "data_type": "string", "required": true, "internal": false},   // graph edge anchor
			{"name": "account_id", "data_type": "string", "required": false, "internal": false},   // graph edge anchor
			{"name": "device_id", "data_type": "string", "required": false, "internal": false},    // graph edge anchor
			{"name": "ip_address_id", "data_type": "string", "required": false, "internal": false}, // graph edge anchor
			{"name": "authentication_method", "data_type": "string", "required": false, "internal": false},
			{"name": "authentication_success", "data_type": "bool", "required": false, "internal": false},
			{"name": "failed_auth_attempts", "data_type": "int", "required": false, "internal": false}, // credential stuffing detection
			{"name": "user_agent", "data_type": "string", "required": false, "internal": false},

			// INTERNAL
			{"name": "behavioral_anomaly_score", "data_type": "float", "required": false, "internal": true},
			{"name": "risk_score", "data_type": "float", "required": false, "internal": true},
			{"name": "is_suspicious", "data_type": "bool", "required": false, "internal": true}, // pre-decision flag before full risk scoring
		},
	},

	// ============================================================
	// ALERT
	// Internal entity - fraud engine output, not source-mapped
	// ============================================================
	{
		"entity_name": "Alert",
		"description": "Fraud engine output representing a flagged event requiring review or action.",
		"internal":    true,
		"fields": []map[string]interface{}{
			{"name": "alert_type", "data_type": "string", "required": true, "internal": true},              // transaction_fraud, account_takeover, etc.
			{"name": "severity", "data_type": "string", "required": true, "internal": true},                // low, medium, high, critical
			{"name": "status", "data_type": "string", "required": true, "internal": true},                  // open, under_review, escalated, closed
			{"name": "triggered_by_entity_type", "data_type": "string", "required": true, "internal": true},
			{"name": "triggered_by_entity_id", "data_type": "string", "required": true, "internal": true},
			{"name": "assigned_to", "data_type": "string", "required": false, "internal": true},
			{"name": "created_at", "data_type": "datetime", "required": true, "internal": true},
			{"name": "resolved_at", "data_type": "datetime", "required": false, "internal": true},
			{"name": "resolution", "data_type": "string", "required": false, "internal": true}, // true_positive, false_positive, inconclusive
			{"name": "resolution_notes", "data_type": "string", "required": false, "internal": true},
			{"name": "model_version", "data_type": "string", "required": false, "internal": true},
		},
	},

	// ============================================================
	// EXTERNAL ACCOUNT
	// Source-mappable - counterparties outside your institution
	// ============================================================
	{
		"entity_name": "ExternalAccount",
		"description": "Account held at an external institution — counterparty in transactions. Key node for mule network detection.",
		"internal":    false,
		"fields": []map[string]interface{}{
			{"name": "account_number", "data_type": "string", "required": false, "internal": false},
			{"name": "bank_bic", "data_type": "string", "required": false, "internal": false},
			{"name": "bank_name", "data_type": "string", "required": false, "internal": false},
			{"name": "country", "data_type": "string", "required": false, "internal": false},
			{"name": "currency", "data_type": "string", "required": false, "internal": false},
			{"name": "account_holder_name", "data_type": "string", "required": false, "internal": false},

			// INTERNAL
			{"name": "risk_score", "data_type": "float", "required": false, "internal": true},
			{"name": "is_flagged_mule", "data_type": "bool", "required": false, "internal": true},
			{"name": "transaction_count", "data_type": "int", "required": false, "internal": true}, // how many txns seen to/from this account
		},
	},

	// ============================================================
	// PHONE NUMBER
	// First-class graph node — enables identity ring detection
	// ============================================================
	{
		"entity_name": "PhoneNumber",
		"description": "Phone number as a first-class identity node. Shared phone across many customers is a synthetic identity / card farm signal.",
		"internal":    false,
		"fields": []map[string]interface{}{
			{"name": "phone_number", "data_type": "string", "required": true, "internal": false},
			{"name": "country_code", "data_type": "string", "required": false, "internal": false},
			{"name": "carrier", "data_type": "string", "required": false, "internal": false},
			{"name": "line_type", "data_type": "string", "required": false, "internal": false}, // mobile, landline, voip
			{"name": "is_verified", "data_type": "bool", "required": false, "internal": false},
			{"name": "first_seen_at", "data_type": "datetime", "required": false, "internal": false},

			// INTERNAL
			{"name": "associated_customer_count", "data_type": "int", "required": false, "internal": true}, // >1 is a strong signal
			{"name": "risk_score", "data_type": "float", "required": false, "internal": true},
		},
	},

	// ============================================================
	// EMAIL ADDRESS
	// First-class graph node — enables identity ring detection
	// ============================================================
	{
		"entity_name": "EmailAddress",
		"description": "Email address as a first-class identity node. Shared email across accounts signals synthetic identity rings.",
		"internal":    false,
		"fields": []map[string]interface{}{
			{"name": "email_address", "data_type": "string", "required": true, "internal": false},
			{"name": "domain", "data_type": "string", "required": false, "internal": false},
			{"name": "is_disposable", "data_type": "bool", "required": false, "internal": false}, // throwaway email providers
			{"name": "is_verified", "data_type": "bool", "required": false, "internal": false},
			{"name": "first_seen_at", "data_type": "datetime", "required": false, "internal": false},

			// INTERNAL
			{"name": "associated_customer_count", "data_type": "int", "required": false, "internal": true},
			{"name": "risk_score", "data_type": "float", "required": false, "internal": true},
		},
	},

	// ============================================================
	// ADDRESS
	// First-class graph node — enables synthetic identity ring detection
	// ============================================================
	{
		"entity_name": "Address",
		"description": "Physical address as a first-class identity node. Many customers sharing an address is a synthetic identity ring indicator.",
		"internal":    false,
		"fields": []map[string]interface{}{
			{"name": "street_line_1", "data_type": "string", "required": false, "internal": false},
			{"name": "street_line_2", "data_type": "string", "required": false, "internal": false},
			{"name": "city", "data_type": "string", "required": false, "internal": false},
			{"name": "region", "data_type": "string", "required": false, "internal": false},
			{"name": "postal_code", "data_type": "string", "required": false, "internal": false},
			{"name": "country", "data_type": "string", "required": false, "internal": false},
			{"name": "latitude", "data_type": "float", "required": false, "internal": false},
			{"name": "longitude", "data_type": "float", "required": false, "internal": false},
			{"name": "address_type", "data_type": "string", "required": false, "internal": false}, // residential, commercial, po_box

			// INTERNAL
			{"name": "associated_customer_count", "data_type": "int", "required": false, "internal": true},
			{"name": "risk_score", "data_type": "float", "required": false, "internal": true},
		},
	},

	// ============================================================
	// RISK SNAPSHOT
	// Internal only — immutable time-stamped risk evaluation record
	// ============================================================
	{
		"entity_name": "RiskSnapshot",
		"description": "Immutable, time-stamped record of a risk evaluation for any entity. Replaces mutable risk_score fields for auditability and model versioning. Written once, never updated.",
		"internal":    true,
		"fields": []map[string]interface{}{
			{"name": "entity_type", "data_type": "string", "required": true, "internal": true},  // Account, Customer, Transaction, etc.
			{"name": "entity_id", "data_type": "string", "required": true, "internal": true},
			{"name": "risk_score", "data_type": "float", "required": true, "internal": true},
			{"name": "risk_level", "data_type": "string", "required": true, "internal": true},   // low, medium, high, critical
			{"name": "evaluated_at", "data_type": "datetime", "required": true, "internal": true},
			{"name": "model_version", "data_type": "string", "required": true, "internal": true},
			{"name": "triggered_by", "data_type": "string", "required": false, "internal": true}, // what event caused re-evaluation
			{"name": "feature_snapshot", "data_type": "json", "required": false, "internal": true}, // model input features at evaluation time — explainability
			{"name": "evaluation_latency_ms", "data_type": "float", "required": false, "internal": true},
		},
	},
}
