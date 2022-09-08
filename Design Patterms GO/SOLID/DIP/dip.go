package main

import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// storing relationships - it is a low level module
type Relationships struct {
	relations []Info
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}
	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{parent, Parent, child})
	r.relations = append(r.relations, Info{child, Child, parent})
}

// research on the data - how to do it?
// it is highlevel module - operate on data
type Research struct {
	//break DIP
	// relationship
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	// relations := r.relationship.relations
	// for _, rel := range relations {
	// 	if rel.from.name == "John" && rel.relationship == Parent {
	// 		fmt.Println("John has a child called ", rel.to.name)
	// 	}
	// }

	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called ", p.name)
	}
}
func main() {
	parent := Person{"John"}
	child := Person{"Chris"}
	child1 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child)
	relationships.AddParentAndChild(&parent, &child1)
	// r := Research{relationship: relationships}
	r := Research{&relationships}
	r.Investigate()
}
