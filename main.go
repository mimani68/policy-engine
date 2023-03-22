package main

import (
	"github.com/mimani68/policy-engine/core"
	"github.com/mimani68/policy-engine/pkg/transport/http"
)

func main() {
	// Make policy
	p := core.NewPolicyMaganer()
	p.SetPolicy("org.risk.support", "name", "ali")
	p.SetPolicy("org.risk.support", "organization", "risk")

	// Load policy from DB
	// conn, err := db.Connection()
	// p.LoadPolicy(conn)

	// Serve policy
	http.InitHttpServer(3000, p)
}
