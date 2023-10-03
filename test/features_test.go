package test_test

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

//nolint: unused
type testContext struct {
	err      error
	cmdInput struct {
		parameters string
	}
	cmdResult struct {
		Output string
		Err    error
	}
}

// Errorf is used by the called assertion to report an error and is required to
// make testify assertions work
func (c *testContext) Errorf(format string, args ...interface{}) {
	c.err = fmt.Errorf(format, args...)
}

func (c *testContext) theAppRunsWithParameters(args string) error {
	c.cmdInput.parameters = args
	cmdArgs := strings.Split(args, " ")
	cmd := exec.Command("../abbreviate", cmdArgs...)
	output, err := cmd.CombinedOutput()
	c.cmdResult.Output = string(output)
	c.cmdResult.Err = err

	return nil
}

func (c *testContext) theAppExitsWithoutError() error {
	assert.NoError(c, c.cmdResult.Err)
	return c.err
}

func (c *testContext) theAppExitsWithAnError() error {
	assert.Error(c, c.cmdResult.Err)
	return c.err
}

func (c *testContext) theAppOutputContains(expected string) error {
	expected = strings.ReplaceAll(expected, "\\\"", "\"")
	assert.Contains(c, c.cmdResult.Output, expected)
	return c.err
}

func (c *testContext) theAppOutputContainsExactly(expected string) error {
	expected = strings.ReplaceAll(expected, "\\\"", "\"")
	assert.Equal(c, strings.TrimSuffix(c.cmdResult.Output, "\n"), expected)
	assert.Equal(c, expected, strings.TrimSuffix(c.cmdResult.Output, "\n"))
	return c.err
}

func (c *testContext) theAppOutputDoesNotContain(unexpected string) error {
	unexpected = strings.ReplaceAll(unexpected, "\\\"", "\"")
	assert.NotContains(c, c.cmdResult.Output, unexpected)
	return c.err
}

//nolint: unused
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {})
}

//nolint: unused
func InitializeScenario(ctx *godog.ScenarioContext) {
	tc := testContext{}
	ctx.BeforeScenario(func(*godog.Scenario) {})
	ctx.AfterScenario(func(s *godog.Scenario, err error) {
		if err != nil {
			fmt.Printf(
				"Command line output for \"%s\"\nUsing parameters: %s\n%s",
				s.GetName(),
				tc.cmdInput.parameters,
				tc.cmdResult.Output,
			)
		}
	})
	ctx.Step(`^the app runs with parameters "(.*)"$`, tc.theAppRunsWithParameters)
	ctx.Step(`^the app exits without error$`, tc.theAppExitsWithoutError)
	ctx.Step(`^the app exits with an error$`, tc.theAppExitsWithAnError)
	ctx.Step(`^the app output contains "(.*)"$`, tc.theAppOutputContains)
	ctx.Step(`^the app output contains exactly "(.*)"$`, tc.theAppOutputContainsExactly)
	ctx.Step(`^the app output does not contain "(.*)"$`, tc.theAppOutputDoesNotContain)
}
