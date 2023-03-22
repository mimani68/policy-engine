package core

type PolicyState struct {
	Realm string
	Title string
	Value string
}

type policy struct {
	App    string
	Policy []PolicyState
}

func NewPolicyMaganer() IPolicy {
	return &policy{}
}

func (p *policy) LoadPolicy(db string, table string) {

}

func (p *policy) SetPolicy(realm, title, value string) bool {
	p.Policy = append(p.Policy, PolicyState{
		Realm: realm,
		Title: title,
		Value: value,
	})
	return true
}

func (p *policy) CheckPolicy(realm, title, value string) bool {
	for _, policy := range p.Policy {
		if policy.Realm == realm && policy.Title == title && policy.Value == value {
			return true
		}
	}
	return false
}
