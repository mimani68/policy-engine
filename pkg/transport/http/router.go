package http

import "github.com/mimani68/policy-engine/core"

type Router struct {
	policy core.IPolicy
}

func NewRouter(policyInstance core.IPolicy) *Router {
	return &Router{
		policy: policyInstance,
	}
}
