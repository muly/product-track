Feature:web scraping
    Scenario:test amazon product availability
    Given the product url "<product_url>" 
    When i send "<http_method>" request to "<end_point>" with above product url in body 
    Then the response should be "<expected_response_body>"
    Then the response code should be <expected_response_status_code> 

    Examples:
    | name                              | product_url                                       | http_method | end_point                      | expected_response_body                    | expected_response_status_code |
    |test amazon product availability   |http://localhost:8006/mock/amazon_available.html   | post        | http://localhost:8006/product  | amazon_available_product_response.json    | 200                           |

        