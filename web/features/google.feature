Feature:
  In order to get cheese information
  As an user
  I need to search at www.google.co.jp

  Scenario: should get google.com page
    When I send "GET" request to "https://www.google.co.jp"
    Then the response code should be 200
    Then the page title is "Google"
    Then I found "search" form
    Then I found "Google 検索" button

  Scenario: should get some cheese pages
    Given I input "cheese" to "search" form
    When I push "Google 検索" button
    Then the response code should be 200
    Then the page title is "cheese - Google 検索"
