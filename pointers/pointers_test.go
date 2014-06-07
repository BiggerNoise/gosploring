package pointers_test

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

type Dummy struct {
	Value int
}

func ChangeStuff(dolt *Dummy, newValue int) {
	dolt.Value = newValue
}
 
var _ = Describe("Testing some Pointers", func () {
	var deadmeat []Dummy

	BeforeEach(func () {
		deadmeat = make([]Dummy, 0, 4)
		deadmeat = append(deadmeat, Dummy{Value:90})
		deadmeat = append(deadmeat, Dummy{Value: 91})
		deadmeat = append(deadmeat, Dummy{Value: 92})
		deadmeat = append(deadmeat, Dummy{Value: 93})
		ChangeStuff(&deadmeat[3], 89)
	}) 

	It("Stuff was changed", func () {
		Expect(deadmeat[3].Value).To(Equal(89))
	})
	It("Allows me to modify a slice member", func () {
		ChangeStuff(&deadmeat[2], 37)
		Expect(deadmeat[2].Value).To(Equal(37))
	})
	
})