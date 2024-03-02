package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AccountConfirmable struct {
	Value       string    `json:"value" bson:"value"`
	Confirmed   bool      `json:"confirmed" bson:"confirmed"`
	ConfirmedAt time.Time `json:"confirmed_at" bson:"confirmed_at"`
}

type AccountProfile struct {
	Avatar      string `json:"avatar,omitempty" bson:"avatar,omitempty"`
	DisplayName string `json:"display_name,omitempty" bson:"display_name,omitempty"`
}

type AccountMetadata struct {
	Profile   AccountProfile `json:"profile,omitempty" bson:"profile,omitempty"`
	CreatedAt time.Time      `json:"created_at,omitempty" bson:"created_at,omitempty"`
	LastSeen  time.Time      `json:"last_seen_at,omitempty" bson:"last_seen_at,omitempty"`
}

type Account struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Email    AccountConfirmable `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Metadata AccountMetadata    `json:"metadata,omitempty" bson:"metadata,omitempty"`
}
