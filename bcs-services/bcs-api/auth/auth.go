

// Package auth xxx
package auth

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// BcsAuth xxx
type BcsAuth interface {
	GetToken(header http.Header) (*Token, error)
	Allow(token *Token, action Action, resource Resource) (bool, error)
}

// Action xxx
type Action string

const (
	// ActionManage xxx
	ActionManage Action = "cluster-manager"
	// ActionRead xxx
	ActionRead Action = "cluster-readonly"

	// TokenDefaultExpireTime xxx
	TokenDefaultExpireTime = 2 * time.Hour
	// TokenRandomLength xxx
	TokenRandomLength = 10
)

// Token xxx
type Token struct {
	Token      string    `json:"token"`
	Username   string    `json:"username"`
	Message    string    `json:"message"`
	ExpireTime time.Time `json:"expire_time"`

	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// Sign xxx
func (t *Token) Sign(duration time.Duration) {
	if duration == 0 {
		duration = TokenDefaultExpireTime
	}
	t.ExpireTime = time.Now().Add(duration)
}

// Generate xxx
func (t *Token) Generate() {
	t.Token = fmt.Sprintf("%d_%s", time.Now().Unix(), randomString())
}

// Resource xxx
type Resource struct {
	ClusterID string `json:"cluster_id"`
	Namespace string `json:"namespace"`
}

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString() string {
	b := make([]rune, TokenRandomLength)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
