Feature: eat godogs
  In order to be happy
  As a hungry gopher
  I need to be able to eat godogs

  Scenario Outline: eating
    Given there are <start> godogs
    When I eat <eat>
    Then there should be <left> remaining

    Examples:
      | start | eat | left |
      |    12 |   5 |    7 |
      |    20 |   5 |   15 |