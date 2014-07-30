package embeded_javascript_test

import (
	"fmt"
	"testing"

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

	makeFloat := func(input interface{}) float64 {
		switch value := input.(type) {
		case float32:
		case float64:
			return value
		default:
			fmt.Printf("unexpected type %T", value)
		}
		return 0.0
	}

	It("Allows record functions to be registered and then called", func() {
		vm := otto.New()
		customFunction := `
			if(record.credit) {
				record.paid = -1 * record.amount
			} else {
				record.paid = record.amount
			}
			record.status = "You're a complete choad."
			delete record.credit
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
		Expect(makeFloat(record["paid"])).To(Equal(-345.45))
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
			var lastResult int64
			for i := 1; i <= b.N; i++ {
				input := i * 1000
				vm.Set(functionArgument, input)
				res, _ := vm.Run(functionInvocation)
				lastResult, _ = res.ToInteger()
			}
			fmt.Println(lastResult)
		}
		jsBench := testing.Benchmark(jsRun)
		fmt.Println(jsBench)

	})
	It("Divide numbers by 1000 with Go", func() {

		calc := func(field int) int {
			return field / 1000
		}

		run := func(b *testing.B) {
			var lastResult int64
			for i := 1; i <= b.N; i++ {
				input := i * 1000
				lastResult = int64(calc(input))
			}
			fmt.Println(lastResult)
		}
		bench := testing.Benchmark(run)
		fmt.Println(bench)
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
			var lastResult string
			for i := 0; i < b.N; i++ {
				input := inputKeys[i%len(inputKeys)]
				vm.Set(functionArgument, input)
				res, _ := vm.Run(functionInvocation)
				lastResult, _ = res.ToString()
			}
			fmt.Println(lastResult)
		}
		jsBench := testing.Benchmark(jsRun)
		fmt.Println(jsBench)
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
			var lastResult string
			for i := 0; i < b.N; i++ {
				input := inputKeys[i%len(inputKeys)]
				lastResult = lookup(input)
			}
			fmt.Println(lastResult)
		}
		bench := testing.Benchmark(run)
		fmt.Println(bench)

	})
})
