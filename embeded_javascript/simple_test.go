package embeded_javascript_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

var _ = Describe("just a simple test", func() {
	It("Performs the most basic operation possible", func() {
		vm := otto.New()
		vm.Run(`
	    abc = 2 + 2;
	    console.log("The value of abc is " + abc); // 4
	`)
	})

	It("Allows field functions to be registered and then called", func() {
		vm := otto.New()
		customFunction := `return field + "-ok"`
		functionId := "custom_XX_232354"
		functionArgument := fmt.Sprintf("%s_Arg", functionId)
		functionInvocation := fmt.Sprintf("%s(%s);", functionId, functionArgument)
		// Once
		vm.Run(fmt.Sprintf("var %s = function(field) { %s };", functionId, customFunction))

		// Each Time
		vm.Set(functionArgument, "this-is")
		result, _ := vm.Run(functionInvocation)

		Expect(result.ToString()).To(Equal("this-is-ok"))
	})
})
