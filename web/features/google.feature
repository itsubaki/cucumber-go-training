Feature:
  In order to get cheese information
  As an user
  I need to search at google.com

  Scenario: should get google.com
    When I send "GET" request to "https://google.com"
    Then the response code should be 200
