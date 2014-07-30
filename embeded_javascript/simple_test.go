package embeded_javascript_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

type RunContext struct{}

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

	It("Allows field functions to be registered and then called - passing argument to call", func() {
		vm := otto.New()
		customFunction := `return field + "-ok"`
		functionId := "custom_XX_232354"

		// Once - compiles the function into the runtime
		vm.Run(fmt.Sprintf("var %s = function(field) { %s };", functionId, customFunction))
		context := &(RunContext{})

		// Each Time
		result, err := vm.Call(functionId, context, "this-is")
		Expect(err).To(BeNil())
		Expect(result.ToString()).To(Equal("this-is-ok"))
	})

	It("Allows record functions to be registered and then called", func() {
		vm := otto.New()
		customFunction := `
			if(record.credit) {
				record.paid = -1 * record.amount
			} else {
				record.paid = record.amount
			}
			record.status = "You're my favorite deputy."
		`
		functionId := "custom_XX_232354"
		functionArgument := fmt.Sprintf("%s_Arg", functionId)
		functionInvocation := fmt.Sprintf("%s(%s);", functionId, functionArgument)
		// Once
		vm.Run(fmt.Sprintf("var %s = function(record) { %s };", functionId, customFunction))

		record := map[string]interface{}{
			"amount": 345.45,
			"credit": true,
		}

		fmt.Println()
		fmt.Println(record)

		// Each Time
		vm.Set(functionArgument, record)
		_, err := vm.Run(functionInvocation)
		fmt.Println(record)
		Expect(err).To(BeNil())
		Expect(record["paid"]).To(Equal(-345.45))
	})

	// Doesn't work.  https://github.com/robertkrimen/otto/pull/92 is a request for clarification
	PIt("Allows record functions to be registered and then called - passing argument to call", func() {
		vm := otto.New()
		customFunction := `
			if(record.credit) {
				record.paid = -1 * record.amount
			} else {
				record.paid = record.amount
			}
			record.status = "You're my favorite deputy."
		`
		functionId := "custom_XX_232354"
		// Once
		vm.Run(fmt.Sprintf("var %s = function(record) { %s };", functionId, customFunction))
		context := &(RunContext{})

		record := map[string]interface{}{
			"amount": 345.45,
			"credit": true,
		}

		fmt.Println()
		fmt.Println(record)

		// Each Time
		_, err := vm.Call(functionId, context, record)
		fmt.Println(record)
		Expect(err).To(BeNil())
		Expect(record["paid"]).To(Equal(-345.45))
	})

	It("Divide numbers by 1000 with Javascript", func() {

		vm := otto.New()
		customFunction := `return field / 1000`
		functionId := "custom_XX_232354"
		functionArgument := fmt.Sprintf("%s_Arg", functionId)
		functionInvocation := fmt.Sprintf("%s(%s);", functionId, functionArgument)
		// Once
		vm.Run(fmt.Sprintf("var %s = function(field) { %s };", functionId, customFunction))

		jsRun := func(b *testing.B) {
			for i := 1; i <= b.N; i++ {
				input := i * 1000
				vm.Set(functionArgument, input)
				vm.Run(functionInvocation)
			}
		}
		jsBench := testing.Benchmark(jsRun)
		fmt.Println("Divide numbers by 1000 with Javascript Run(): ", jsBench)

	})

	It("Divide numbers by 1000 with Javascript Call", func() {
		vm := otto.New()
		customFunction := `return field / 1000`
		functionId := "custom_XX_232354"
		// Once
		vm.Run(fmt.Sprintf("var %s = function(field) { %s };", functionId, customFunction))

		jsRun := func(b *testing.B) {
			for i := 1; i <= b.N; i++ {
				input := i * 1000
				vm.Call(functionId, input)
			}
		}
		jsBench := testing.Benchmark(jsRun)
		fmt.Println("Divide numbers by 1000 with Javascript Call(): ", jsBench)
	})

	It("Divide numbers by 1000 with Go", func() {

		calc := func(field int) int {
			return field / 1000
		}

		run := func(b *testing.B) {
			for i := 1; i <= b.N; i++ {
				input := i * 1000
				calc(input)
			}
		}
		bench := testing.Benchmark(run)
		fmt.Println("Divide numbers by 1000 with Go: ", bench)
	})
	inputKeys := []int{42, 81, 17, 2, 89, 84}

	It("Simple Lookup Table with Javascript", func() {
		vm := otto.New()
		customFunction := `
		switch(field) {
			case 42: return 'answer'
			case 81: return 'monk'
			case 17: return 'r'
			case 2: return 'kids'
			default: return 'whatever'
		}
		`
		functionId := "custom_XX_232354"
		functionArgument := fmt.Sprintf("%s_Arg", functionId)
		functionInvocation := fmt.Sprintf("%s(%s);", functionId, functionArgument)
		// Once
		vm.Run(fmt.Sprintf("var %s = function(field) { %s };", functionId, customFunction))

		jsRun := func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input := inputKeys[i%len(inputKeys)]
				vm.Set(functionArgument, input)
				vm.Run(functionInvocation)
			}
		}
		jsBench := testing.Benchmark(jsRun)
		fmt.Println("Simple Lookup Table with Javascript: ", jsBench)
	})

	It("Simple Record Function with Javascript", func() {
		vm := otto.New()

		customFunction := `
			if(record.credit) {
				record.paid = -1 * record.amount
			} else {
				record.paid = record.amount
			}
			record.status = "You're my favorite deputy."
		`

		functionId := "custom_XX_232354"
		functionArgument := fmt.Sprintf("%s_Arg", functionId)
		functionInvocation := fmt.Sprintf("%s(%s);", functionId, functionArgument)
		// Once
		vm.Run(fmt.Sprintf("var %s = function(record) { %s };", functionId, customFunction))

		jsRun := func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input := map[string]interface{}{
					"amount": 45.45 * float64(i),
					"credit": true,
				}

				vm.Set(functionArgument, input)
				vm.Run(functionInvocation)
			}
		}
		jsBench := testing.Benchmark(jsRun)
		fmt.Println("Simple Record Function with Javascript: ", jsBench)
	})

	It("Simple Lookup Table with Go", func() {
		table := map[int]string{
			42: "answer",
			81: "monk",
			17: "r",
			2:  "kids",
		}

		lookup := func(input int) string {
			if result, present := table[input]; present {
				return result
			}
			return "whatever"
		}

		run := func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				input := inputKeys[i%len(inputKeys)]
				lookup(input)
			}
		}
		bench := testing.Benchmark(run)
		fmt.Println("Simple Lookup Table with Go: ", bench)

	})
})
