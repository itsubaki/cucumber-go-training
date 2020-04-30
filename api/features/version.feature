Feature: get version
  In order to know godog version
  As an API user
  I need to be able to request version

  Scenario: does not allow POST method
    When I send "POST" request to "/version"
    Then the response code should be 405
