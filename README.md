# cucumber-training-golang

 - the coverage only applied to the main package, not include sub package.

```
$ go test -v --godog.format=pretty --godog.random -race -coverprofile=coverage.out -covermode=atomic
Feature: get version
  In order to know godog version
  As an API user
  I need to be able to request version

  Scenario: does not allow POST method       # features/version.feature:6
    When I send "POST" request to "/version" # api_test.go:20 -> *apiFeature
    Then the response code should be 405     # api_test.go:45 -> *apiFeature
    And the response should match json:      # api_test.go:52 -> *apiFeature
      """
      {
        "error": "Method not allowed"
      }
      """
      TODO: write pending definition

  Scenario: should get version number       # features/version.feature:16
    When I send "GET" request to "/version" # api_test.go:20 -> *apiFeature
    Then the response code should be 200    # api_test.go:45 -> *apiFeature
    And the response should match json:     # api_test.go:52 -> *apiFeature
      """
      {
        "version": "v0.5.3"
      }
      """
      TODO: write pending definition

2 scenarios (2 pending)
6 steps (4 passed, 2 pending)
3.380655ms

Randomized with seed: 35929
testing: warning: no tests to run
PASS
coverage: 75.0% of statements
ok      github.com/itsubaki/cucumber-training-golang/api        0.088s
```

```
$ godog
Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Eat 5 out of 12          # features/godogs.feature:7
    Given there are 12 godogs
    When I eat 5
    Then there should be 7 remaining

1 scenarios (1 undefined)
3 steps (3 undefined)
127.634µs

You can implement step definitions for undefined steps with these snippets:

func thereAreGodogs(arg1 int) error {
	return godog.ErrPending
}

func iEat(arg1 int) error {
	return godog.ErrPending
}

func thereShouldBeRemaining(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	s.Step(`^I eat (\d+)$`, iEat)
	s.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
}
```

```
$ godog
Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario: Eat 5 out of 12          # features/godogs.feature:6
    Given there are 12 godogs        # godogs_test.go:10 -> thereAreGodogs
    When I eat 5                     # godogs_test.go:14 -> iEat
    Then there should be 7 remaining # godogs_test.go:22 -> thereShouldBeRemaining

  Scenario: Eat 2 out of 10          # features/godogs.feature:11
    Given there are 10 godogs        # godogs_test.go:10 -> thereAreGodogs
    When I eat 2                     # godogs_test.go:14 -> iEat
    Then there should be 8 remaining # godogs_test.go:22 -> thereShouldBeRemaining

2 scenarios (2 passed)
6 steps (6 passed)
329.122µs
```