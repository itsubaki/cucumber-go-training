# cucumber-training

```
$ export GO111MODULE=on
$ go mod init
$ go get github.com/cucumber/godog/cmd/godog@v0.8.1
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