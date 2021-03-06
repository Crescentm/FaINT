// +build go1.1,!go1.2

package parser

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/glycerine/goconvey/convey/reporting"
	"github.com/glycerine/goconvey/web/server/contract"
)

func TestParsePackage_NoGoFiles_ReturnsPackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expected_NoGoFiles.PackageName}
	ParsePackageResults(actual, input_NoGoFiles)
	assertEqual(t, expected_NoGoFiles, *actual)
}

func TestParsePackage_NoTestFiles_ReturnsPackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expected_NoTestFiles.PackageName}
	ParsePackageResults(actual, input_NoTestFiles)
	assertEqual(t, expected_NoTestFiles, *actual)
}

func TestParsePacakge_NoTestFunctions_ReturnsPackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expected_NoTestFunctions.PackageName}
	ParsePackageResults(actual, input_NoTestFunctions)
	assertEqual(t, expected_NoTestFunctions, *actual)
}

func TestParsePackage_BuildFailed_ReturnsPackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expected_BuildFailed_InvalidPackageDeclaration.PackageName}
	ParsePackageResults(actual, input_BuildFailed_InvalidPackageDeclaration)
	assertEqual(t, expected_BuildFailed_InvalidPackageDeclaration, *actual)

	actual = &contract.PackageResult{PackageName: expected_BuildFailed_OtherErrors.PackageName}
	ParsePackageResults(actual, input_BuildFailed_OtherErrors)
	assertEqual(t, expected_BuildFailed_OtherErrors, *actual)

	actual = &contract.PackageResult{PackageName: expected_BuildFailed_CantFindPackage.PackageName}
	ParsePackageResults(actual, input_BuildFailed_CantFindPackage)
	assertEqual(t, expected_BuildFailed_CantFindPackage, *actual)
}

func TestParsePackage_OldSchoolWithFailureOutput_ReturnsCompletePackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expectedOldSchool_Fails.PackageName}
	ParsePackageResults(actual, inputOldSchool_Fails)
	assertEqual(t, expectedOldSchool_Fails, *actual)
}

func TestParsePackage_OldSchoolWithSuccessOutput_ReturnsCompletePackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expectedOldSchool_Passes.PackageName}
	ParsePackageResults(actual, inputOldSchool_Passes)
	assertEqual(t, expectedOldSchool_Passes, *actual)
}

func TestParsePackage_OldSchoolWithPanicOutput_ReturnsCompletePackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expectedOldSchool_Panics.PackageName}
	ParsePackageResults(actual, inputOldSchool_Panics)
	assertEqual(t, expectedOldSchool_Panics, *actual)
}

func TestParsePackage_GoConveyOutput_ReturnsCompletePackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expectedGoConvey.PackageName}
	ParsePackageResults(actual, inputGoConvey)
	assertEqual(t, expectedGoConvey, *actual)
}

func TestParsePackage_ActualPackageNameDifferentThanDirectoryName_ReturnsActualPackageName(t *testing.T) {
	actual := &contract.PackageResult{PackageName: strings.Replace(expectedGoConvey.PackageName, "examples", "stuff", -1)}
	ParsePackageResults(actual, inputGoConvey)
	assertEqual(t, expectedGoConvey, *actual)
}

func TestParsePackage_GoConveyOutputMalformed_CausesPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			message := fmt.Sprintf("%v", r)
			if !strings.Contains(message, "bug report") {
				t.Errorf("Should have panicked with a request to file a bug report but we received this error instead: %s", message)
			}
		} else {
			t.Errorf("Should have panicked with a request to file a bug report but we received no error.")
		}
	}()

	actual := &contract.PackageResult{PackageName: expectedGoConvey.PackageName}
	ParsePackageResults(actual, inputGoConvey_Malformed)
}

func TestParsePackage_GoConveyWithRandomOutput_ReturnsPackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expectedGoConvey_WithRandomOutput.PackageName}
	ParsePackageResults(actual, inputGoConvey_WithRandomOutput)
	assertEqual(t, expectedGoConvey_WithRandomOutput, *actual)
}

func TestParsePackage_NestedTests_ReturnsPackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expectedNestedTests.PackageName}
	ParsePackageResults(actual, inputNestedTests)
	assertEqual(t, expectedNestedTests, *actual)
}

func TestParsePacakge_WithExampleFunctions_ReturnsPackageResult(t *testing.T) {
	actual := &contract.PackageResult{PackageName: expectedExampleFunctions.PackageName}
	ParsePackageResults(actual, inputExampleFunctions)
	assertEqual(t, expectedExampleFunctions, *actual)
}

func assertEqual(t *testing.T, expected, actual interface{}) {
	a, _ := json.Marshal(expected)
	b, _ := json.Marshal(actual)
	if string(a) != string(b) {
		t.Errorf(failureTemplate, string(a), string(b))
	}
}

const failureTemplate = "Comparison failed:\n  Expected: %v\n    Actual: %v\n"

const input_NoGoFiles = `can't load package: package github.com/glycerine/goconvey: no Go source files in /Users/matt/Work/Dev/goconvey/src/github.com/glycerine/goconvey`

var expected_NoGoFiles = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey",
	Outcome:     contract.NoGoFiles,
	BuildOutput: input_NoGoFiles,
}

const input_NoTestFiles = `?   	pkg.smartystreets.net/liveaddress-zipapi	[no test files]`

var expected_NoTestFiles = contract.PackageResult{
	PackageName: "pkg.smartystreets.net/liveaddress-zipapi",
	Outcome:     contract.NoTestFiles,
	BuildOutput: input_NoTestFiles,
}

const input_NoTestFunctions = `testing: warning: no tests to run`

var expected_NoTestFunctions = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/webserver/examples",
	Outcome:     contract.NoTestFunctions,
	BuildOutput: input_NoTestFunctions,
}

const input_BuildFailed_InvalidPackageDeclaration = `
can't load package: package github.com/glycerine/goconvey/examples:
bowling_game_test.go:9:1: expected 'package', found 'IDENT' asdf
bowling_game_test.go:10:1: invalid package name _
`

var expected_BuildFailed_InvalidPackageDeclaration = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/examples",
	Outcome:     contract.BuildFailure,
	BuildOutput: strings.TrimSpace(input_BuildFailed_InvalidPackageDeclaration),
}

const input_BuildFailed_CantFindPackage = `
bowling_game.go:3:8: cannot find package "format" in any of:
	/usr/local/go/src/pkg/format (from $GOROOT)
	/Users/mike/work/dev/goconvey/src/format (from $GOPATH)
`

var expected_BuildFailed_CantFindPackage = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/examples",
	Outcome:     contract.BuildFailure,
	BuildOutput: strings.TrimSpace(input_BuildFailed_CantFindPackage),
}

const input_BuildFailed_OtherErrors = `
# github.com/glycerine/goconvey/examples
./bowling_game_test.go:22: undefined: game
./bowling_game_test.go:22: cannot assign to game
./bowling_game_test.go:25: undefined: game
./bowling_game_test.go:28: undefined: game
./bowling_game_test.go:33: undefined: game
./bowling_game_test.go:36: undefined: game
./bowling_game_test.go:41: undefined: game
./bowling_game_test.go:42: undefined: game
./bowling_game_test.go:43: undefined: game
./bowling_game_test.go:46: undefined: game
./bowling_game_test.go:46: too many errors
FAIL	github.com/glycerine/goconvey/examples [build failed]
`

var expected_BuildFailed_OtherErrors = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/examples",
	Outcome:     contract.BuildFailure,
	BuildOutput: strings.TrimSpace(input_BuildFailed_OtherErrors),
}

const inputOldSchool_Passes = `
=== RUN TestOldSchool_Passes
--- PASS: TestOldSchool_Passes (0.02 seconds)
=== RUN TestSkippingTests
--- SKIP: TestSkippingTests (0.00 seconds)
	old_school_test.go:8: blah
=== RUN TestOldSchool_PassesWithMessage
--- PASS: TestOldSchool_PassesWithMessage (0.05 seconds)
	old_school_test.go:10: I am a passing test.
		With a newline.
PASS
ok  	github.com/glycerine/goconvey/webserver/examples	0.018s
`

var expectedOldSchool_Passes = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/webserver/examples",
	Elapsed:     0.018,
	Outcome:     contract.Passed,
	TestResults: []contract.TestResult{
		contract.TestResult{
			TestName: "TestOldSchool_Passes",
			Elapsed:  0.02,
			Passed:   true,
			File:     "",
			Line:     0,
			Message:  "",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "TestSkippingTests",
			Elapsed:  0,
			Passed:   true,
			Skipped:  true,
			File:     "old_school_test.go",
			Line:     8,
			Message:  "old_school_test.go:8: blah",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "TestOldSchool_PassesWithMessage",
			Elapsed:  0.05,
			Passed:   true,
			File:     "old_school_test.go",
			Line:     10,
			Message:  "old_school_test.go:10: I am a passing test.\nWith a newline.",
			Stories:  []reporting.ScopeResult{},
		},
	},
}

const inputOldSchool_Fails = `
=== RUN TestOldSchool_Passes
--- PASS: TestOldSchool_Passes (0.01 seconds)
=== RUN TestOldSchool_PassesWithMessage
--- PASS: TestOldSchool_PassesWithMessage (0.03 seconds)
	old_school_test.go:10: I am a passing test.
		With a newline.
=== RUN TestOldSchool_Failure
--- FAIL: TestOldSchool_Failure (0.06 seconds)
=== RUN TestOldSchool_FailureWithReason
--- FAIL: TestOldSchool_FailureWithReason (0.11 seconds)
	old_school_test.go:18: I am a failing test.
FAIL
exit status 1
FAIL	github.com/glycerine/goconvey/webserver/examples	0.017s
`

var expectedOldSchool_Fails = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/webserver/examples",
	Outcome:     contract.Failed,
	Elapsed:     0.017,
	TestResults: []contract.TestResult{
		contract.TestResult{
			TestName: "TestOldSchool_Passes",
			Elapsed:  0.01,
			Passed:   true,
			File:     "",
			Line:     0,
			Message:  "",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "TestOldSchool_PassesWithMessage",
			Elapsed:  0.03,
			Passed:   true,
			File:     "old_school_test.go",
			Line:     10,
			Message:  "old_school_test.go:10: I am a passing test.\nWith a newline.",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "TestOldSchool_Failure",
			Elapsed:  0.06,
			Passed:   false,
			File:     "",
			Line:     0,
			Message:  "",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "TestOldSchool_FailureWithReason",
			Elapsed:  0.11,
			Passed:   false,
			File:     "old_school_test.go",
			Line:     18,
			Message:  "old_school_test.go:18: I am a failing test.",
			Stories:  []reporting.ScopeResult{},
		},
	},
}

const inputOldSchool_Panics = `
=== RUN TestOldSchool_Panics
--- FAIL: TestOldSchool_Panics (0.02 seconds)
panic: runtime error: index out of range [recovered]
	panic: runtime error: index out of range

goroutine 3 [running]:
testing.func??004()
	/usr/local/go/src/pkg/testing/testing.go:348 +0xcd
github.com/glycerine/goconvey/webserver/examples.TestOldSchool_Panics(0x210292000)
	/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/something_test.go:15 +0xec
testing.tRunner(0x210292000, 0x1b09f0)
	/usr/local/go/src/pkg/testing/testing.go:353 +0x8a
created by testing.RunTests
	/usr/local/go/src/pkg/testing/testing.go:433 +0x86b

goroutine 1 [chan receive]:
testing.RunTests(0x138f38, 0x1b09f0, 0x1, 0x1, 0x1, ...)
	/usr/local/go/src/pkg/testing/testing.go:434 +0x88e
testing.Main(0x138f38, 0x1b09f0, 0x1, 0x1, 0x1b7f60, ...)
	/usr/local/go/src/pkg/testing/testing.go:365 +0x8a
main.main()
	github.com/glycerine/goconvey/webserver/examples/_test/_testmain.go:43 +0x9a
exit status 2
FAIL	github.com/glycerine/goconvey/webserver/examples	0.014s
`

var expectedOldSchool_Panics = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/webserver/examples",
	Elapsed:     0.014,
	Outcome:     contract.Panicked,
	TestResults: []contract.TestResult{
		contract.TestResult{
			TestName: "TestOldSchool_Panics",
			Elapsed:  0.02,
			Passed:   false,
			File:     "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/something_test.go",
			Line:     15,
			Message:  "",
			Error: strings.Replace(`panic: runtime error: index out of range [recovered]
	panic: runtime error: index out of range

goroutine 3 [running]:
testing.func??004()
	/usr/local/go/src/pkg/testing/testing.go:348 +0xcd
github.com/glycerine/goconvey/webserver/examples.TestOldSchool_Panics(0x210292000)
	/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/something_test.go:15 +0xec
testing.tRunner(0x210292000, 0x1b09f0)
	/usr/local/go/src/pkg/testing/testing.go:353 +0x8a
created by testing.RunTests
	/usr/local/go/src/pkg/testing/testing.go:433 +0x86b

goroutine 1 [chan receive]:
testing.RunTests(0x138f38, 0x1b09f0, 0x1, 0x1, 0x1, ...)
	/usr/local/go/src/pkg/testing/testing.go:434 +0x88e
testing.Main(0x138f38, 0x1b09f0, 0x1, 0x1, 0x1b7f60, ...)
	/usr/local/go/src/pkg/testing/testing.go:365 +0x8a
main.main()
	github.com/glycerine/goconvey/webserver/examples/_test/_testmain.go:43 +0x9a`, "\u0009", "\t", -1),
			Stories: []reporting.ScopeResult{},
		},
	},
}

const inputGoConvey_Malformed = `
=== RUN TestPassingStory
>>>>>
{
  "Title": "A passing story",
  "File": "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go",
  "Line": 11,
  "Depth": 0,
  "Assertions": [
    {
      "File": "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go",
      "Line": 10,
      "Failure": "",

      ;aiwheopinen39 n3902n92m

      "Error": null,
      "Skipped": false,
      "StackTrace": "goroutine 3 [running]:\ngithub.com/glycerine/goconvey/webserver/examples.func??001()\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go:10 +0xe3\ngithub.com/glycerine/goconvey/webserver/examples.TestPassingStory(0x210314000)\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go:11 +0xec\ntesting.tRunner(0x210314000, 0x21ab10)\n\u0009/usr/local/go/src/pkg/testing/testing.go:353 +0x8a\ncreated by testing.RunTests\n\u0009/usr/local/go/src/pkg/testing/testing.go:433 +0x86b\n"
    }
  ]
},
<<<<<
--- PASS: TestPassingStory (0.01 seconds)
PASS
ok  	github.com/glycerine/goconvey/webserver/examples	0.019s
`

const inputGoConvey = `
=== RUN TestPassingStory
>>>>>
{
  "Title": "A passing story",
  "File": "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go",
  "Line": 11,
  "Depth": 0,
  "Assertions": [
    {
      "File": "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go",
      "Line": 10,
      "Failure": "",
      "Error": null,
      "Skipped": false,
      "StackTrace": "goroutine 3 [running]:\ngithub.com/glycerine/goconvey/webserver/examples.func??001()\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go:10 +0xe3\ngithub.com/glycerine/goconvey/webserver/examples.TestPassingStory(0x210314000)\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go:11 +0xec\ntesting.tRunner(0x210314000, 0x21ab10)\n\u0009/usr/local/go/src/pkg/testing/testing.go:353 +0x8a\ncreated by testing.RunTests\n\u0009/usr/local/go/src/pkg/testing/testing.go:433 +0x86b\n"
    }
  ]
},
<<<<<
--- PASS: TestPassingStory (0.01 seconds)
PASS
ok  	github.com/glycerine/goconvey/webserver/examples	0.019s
`

var expectedGoConvey = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/webserver/examples",
	Elapsed:     0.019,
	Outcome:     contract.Passed,
	TestResults: []contract.TestResult{
		contract.TestResult{
			TestName: "TestPassingStory",
			Elapsed:  0.01,
			Passed:   true,
			File:     "",
			Line:     0,
			Message:  "",
			Stories: []reporting.ScopeResult{
				reporting.ScopeResult{
					Title: "A passing story",
					File:  "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go",
					Line:  11,
					Depth: 0,
					Assertions: []*reporting.AssertionResult{
						&reporting.AssertionResult{
							File:       "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go",
							Line:       10,
							Failure:    "",
							Error:      nil,
							Skipped:    false,
							StackTrace: "goroutine 3 [running]:\ngithub.com/glycerine/goconvey/webserver/examples.func??001()\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go:10 +0xe3\ngithub.com/glycerine/goconvey/webserver/examples.TestPassingStory(0x210314000)\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/webserver/examples/old_school_test.go:11 +0xec\ntesting.tRunner(0x210314000, 0x21ab10)\n\u0009/usr/local/go/src/pkg/testing/testing.go:353 +0x8a\ncreated by testing.RunTests\n\u0009/usr/local/go/src/pkg/testing/testing.go:433 +0x86b\n",
						},
					},
				},
			},
		},
	},
}

const inputGoConvey_WithRandomOutput = `
=== RUN TestPassingStory
*** Hello, World! (1) ***
*** Hello, World! (2) ***
*** Hello, World! (3) ***>>>>>
{
  "Title": "A passing story",
  "File": "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go",
  "Line": 16,
  "Depth": 0,
  "Assertions": [
    {
      "File": "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go",
      "Line": 14,
      "Failure": "",
      "Error": null,
      "Skipped": false,
      "StackTrace": "goroutine 3 [running]:\ngithub.com/glycerine/goconvey/web/server/testing.func??001()\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go:14 +0x186\ngithub.com/glycerine/goconvey/web/server/testing.TestPassingStory(0x210315000)\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go:16 +0x1b9\ntesting.tRunner(0x210315000, 0x21bb10)\n\u0009/usr/local/go/src/pkg/testing/testing.go:353 +0x8a\ncreated by testing.RunTests\n\u0009/usr/local/go/src/pkg/testing/testing.go:433 +0x86b\n"
    }
  ]
},
<<<<<
*** Hello, World! (4)***
*** Hello, World! (5) ***
>>>>>
{
  "Title": "A passing story",
  "File": "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go",
  "Line": 22,
  "Depth": 0,
  "Assertions": [
    {
      "File": "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go",
      "Line": 20,
      "Failure": "",
      "Error": null,
      "Skipped": false,
      "StackTrace": "goroutine 3 [running]:\ngithub.com/glycerine/goconvey/web/server/testing.func??002()\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go:20 +0x186\ngithub.com/glycerine/goconvey/web/server/testing.TestPassingStory(0x210315000)\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go:22 +0x294\ntesting.tRunner(0x210315000, 0x21bb10)\n\u0009/usr/local/go/src/pkg/testing/testing.go:353 +0x8a\ncreated by testing.RunTests\n\u0009/usr/local/go/src/pkg/testing/testing.go:433 +0x86b\n"
    }
  ]
},
<<<<<
*** Hello, World! (6) ***
--- PASS: TestPassingStory (0.03 seconds)
PASS
ok  	github.com/glycerine/goconvey/web/server/testing	0.024s
`

var expectedGoConvey_WithRandomOutput = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/web/server/testing",
	Elapsed:     0.024,
	Outcome:     contract.Passed,
	TestResults: []contract.TestResult{
		contract.TestResult{
			TestName: "TestPassingStory",
			Elapsed:  0.03,
			Passed:   true,
			File:     "",
			Line:     0,
			Message:  "*** Hello, World! (1) ***\n*** Hello, World! (2) ***\n*** Hello, World! (3) ***\n*** Hello, World! (4)***\n*** Hello, World! (5) ***\n*** Hello, World! (6) ***",
			Stories: []reporting.ScopeResult{
				reporting.ScopeResult{
					Title: "A passing story",
					File:  "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go",
					Line:  16,
					Depth: 0,
					Assertions: []*reporting.AssertionResult{
						&reporting.AssertionResult{
							File:       "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go",
							Line:       14,
							Failure:    "",
							Error:      nil,
							Skipped:    false,
							StackTrace: "goroutine 3 [running]:\ngithub.com/glycerine/goconvey/web/server/testing.func??001()\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go:14 +0x186\ngithub.com/glycerine/goconvey/web/server/testing.TestPassingStory(0x210315000)\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go:16 +0x1b9\ntesting.tRunner(0x210315000, 0x21bb10)\n\u0009/usr/local/go/src/pkg/testing/testing.go:353 +0x8a\ncreated by testing.RunTests\n\u0009/usr/local/go/src/pkg/testing/testing.go:433 +0x86b\n",
						},
					},
				},
				reporting.ScopeResult{
					Title: "A passing story",
					File:  "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go",
					Line:  22,
					Depth: 0,
					Assertions: []*reporting.AssertionResult{
						&reporting.AssertionResult{
							File:       "/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go",
							Line:       20,
							Failure:    "",
							Error:      nil,
							Skipped:    false,
							StackTrace: "goroutine 3 [running]:\ngithub.com/glycerine/goconvey/web/server/testing.func??002()\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go:20 +0x186\ngithub.com/glycerine/goconvey/web/server/testing.TestPassingStory(0x210315000)\n\u0009/Users/mike/work/dev/goconvey/src/github.com/glycerine/goconvey/web/server/testing/go_test.go:22 +0x294\ntesting.tRunner(0x210315000, 0x21bb10)\n\u0009/usr/local/go/src/pkg/testing/testing.go:353 +0x8a\ncreated by testing.RunTests\n\u0009/usr/local/go/src/pkg/testing/testing.go:433 +0x86b\n",
						},
					},
				},
			},
		},
	},
}

/*
Test output for these tests was generated from the following test code:

Old School style tests:

	package examples

	import "testing"

	func TestOldSchool_Passes(t *testing.T) {
		// passes implicitly
	}

	func TestOldSchool_PassesWithMessage(t *testing.T) {
		t.Log("I am a passing test.\nWith a newline.")
	}

	func TestOldSchool_Failure(t *testing.T) {
		t.Fail() // no message
	}

	func TestOldSchool_FailureWithReason(t *testing.T) {
		t.Error("I am a failing test.")
	}

GoConvey style tests:

	package examples

	import (
		. "github.com/glycerine/goconvey/convey"
		"testing"
	)

	func TestPassingStory(t *testing.T) {
		Convey("A passing story", t, func() {
			So("This test passes", ShouldContainSubstring, "pass")
		})
	}

GoConvey style tests with random output:

	package examples

	import (
		"fmt"
		. "github.com/glycerine/goconvey/convey"
		"testing"
	)

	func TestPassingStory(t *testing.T) {
		fmt.Println("*** Hello, World! (1) ***")

		Convey("A passing story", t, func() {
			fmt.Println("*** Hello, World! (2) ***")
			So("This test passes", ShouldContainSubstring, "pass")
			fmt.Println("*** Hello, World! (3) ***")
		})

		Convey("A passing story", t, func() {
			fmt.Println("*** Hello, World! (4)***")
			So("This test passes", ShouldContainSubstring, "pass")
			fmt.Println("*** Hello, World! (5) ***")
		})

		fmt.Println("*** Hello, World! (6) ***")
	}


*/

const inputNestedTests = `
=== RUN TestNestedTests
=== RUN TestNestedTests_Passes
--- PASS: TestNestedTests_Passes (0.02 seconds)
=== RUN TestNestedTests_Failure
--- FAIL: TestNestedTests_Failure (0.06 seconds)
=== RUN TestNestedTests_FailureWithReason
--- FAIL: TestNestedTests_FailureWithReason (0.11 seconds)
	nested_test.go:18: I am a failing test.
=== RUN TestNestedTests_Skipping
--- SKIP: TestNestedTests_Skipping (0.00 seconds)
	nested_test.go:8: blah
=== RUN TestNestedTests_PassesWithMessage
--- PASS: TestNestedTests_PassesWithMessage (0.05 seconds)
	nested_test.go:10: I am a passing test.
		With a newline.
--- FAIL: TestNestedTests (0.25 seconds)
FAIL
exit status 1
FAIL	github.com/glycerine/goconvey/webserver/examples	0.018s
`

var expectedNestedTests = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/webserver/examples",
	Elapsed:     0.018,
	Outcome:     contract.Failed,
	TestResults: []contract.TestResult{
		contract.TestResult{
			TestName: "TestNestedTests",
			Elapsed:  0.25,
			Passed:   false,
			File:     "",
			Line:     0,
			Message:  "",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "TestNestedTests_Passes",
			Elapsed:  0.02,
			Passed:   true,
			File:     "",
			Line:     0,
			Message:  "",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "TestNestedTests_Failure",
			Elapsed:  0.06,
			Passed:   false,
			File:     "",
			Line:     0,
			Message:  "",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "TestNestedTests_FailureWithReason",
			Elapsed:  0.11,
			Passed:   false,
			File:     "nested_test.go",
			Line:     18,
			Message:  "nested_test.go:18: I am a failing test.",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "TestNestedTests_Skipping",
			Elapsed:  0.00,
			Passed:   true,
			Skipped:  true,
			File:     "nested_test.go",
			Line:     8,
			Message:  "nested_test.go:8: blah",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "TestNestedTests_PassesWithMessage",
			Elapsed:  0.05,
			Passed:   true,
			File:     "nested_test.go",
			Line:     10,
			Message:  "nested_test.go:10: I am a passing test.\nWith a newline.",
			Stories:  []reporting.ScopeResult{},
		},
	},
}

const inputExampleFunctions = `
=== RUN Example_Failure
--- FAIL: Example_Failure (0.11 seconds)
got:
actuall output
want:
real output
=== RUN Example_Pass
--- PASS: Example_Pass (0.06 seconds)
FAIL
exit status 1
FAIL	github.com/glycerine/goconvey/webserver/examples	0.18s
`

var expectedExampleFunctions = contract.PackageResult{
	PackageName: "github.com/glycerine/goconvey/webserver/examples",
	Elapsed:     0.18,
	Outcome:     contract.Failed,
	TestResults: []contract.TestResult{
		contract.TestResult{
			TestName: "Example_Failure",
			Elapsed:  0.11,
			Passed:   false,
			File:     "",
			Line:     0,
			Message:  "got:\nactuall output\nwant:\nreal output",
			Stories:  []reporting.ScopeResult{},
		},
		contract.TestResult{
			TestName: "Example_Pass",
			Elapsed:  0.06,
			Passed:   true,
			File:     "",
			Line:     0,
			Message:  "",
			Stories:  []reporting.ScopeResult{},
		},
	},
}
