// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Address struct {
	ID           int64          `json:"id"`
	AddressLine1 string         `json:"address_line1"`
	AddressLine2 sql.NullString `json:"address_line2"`
	City         string         `json:"city"`
	State        string         `json:"state"`
	ZipCode      string         `json:"zip_code"`
	Country      string         `json:"country"`
	CreatedAt    time.Time      `json:"created_at"`
}

type ApiAccount struct {
	ID            int64  `json:"id"`
	Username      string `json:"username"`
	IsActive      bool   `json:"is_active"`
	IsAutoRenewal bool   `json:"is_auto_renewal"`
	// ICD|ICD_PRO|APS|APS_TXT|ALL
	ServiceType string `json:"service_type"`
	// DEMO|BASIC|PRO
	PlanName      string    `json:"plan_name"`
	CreditBalance int64     `json:"credit_balance"`
	ActiveAt      time.Time `json:"active_at"`
	LastUseAt     time.Time `json:"last_use_at"`
	ExpiresAt     time.Time `json:"expires_at"`
	CreatedAt     time.Time `json:"created_at"`
}

type Application struct {
	ID               int64         `json:"id"`
	Agent            string        `json:"agent"`
	IsFormal         bool          `json:"is_formal"`
	Is1035Exchange   sql.NullBool  `json:"is_1035_exchange"`
	PrimaryInsuredID int64         `json:"primary_insured_id"`
	JointInsuredID   sql.NullInt64 `json:"joint_insured_id"`
	Carriers         string        `json:"carriers"`
	// non carrier-specific product type
	ProductType   string    `json:"product_type"`
	AppliedAmount int64     `json:"applied_amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type Client struct {
	ID                  int64          `json:"id"`
	Agent               string         `json:"agent"`
	FullName            string         `json:"full_name"`
	AddressID           sql.NullInt64  `json:"address_id"`
	BirthDate           time.Time      `json:"birth_date"`
	DriverLicenseNumber sql.NullString `json:"driver_license_number"`
	DriverLicenseState  sql.NullString `json:"driver_license_state"`
	Email               string         `json:"email"`
	CreatedAt           time.Time      `json:"created_at"`
}

type Document struct {
	ID            int64  `json:"id"`
	ApplicationID int64  `json:"application_id"`
	ClientID      int64  `json:"client_id"`
	DocType       string `json:"doc_type"`
	FileName      string `json:"file_name"`
	FileSize      string `json:"file_size"`
	FileType      string `json:"file_type"`
	// s3 bucket url
	S3Url     string    `json:"s3_url"`
	CreatedAt time.Time `json:"created_at"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type User struct {
	Username          string         `json:"username"`
	HashedPassword    string         `json:"hashed_password"`
	FullName          string         `json:"full_name"`
	Email             string         `json:"email"`
	AddressID         sql.NullInt64  `json:"address_id"`
	Agency            sql.NullString `json:"agency"`
	AppContact        sql.NullString `json:"app_contact"`
	AppContactEmail   sql.NullString `json:"app_contact_email"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreatedAt         time.Time      `json:"created_at"`
}
