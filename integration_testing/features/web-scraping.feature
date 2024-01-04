Feature:web scraping
    Scenario:test product availability
    Given test "<name>"
    And the deployed api host "https://smuly-test-ground.ue.r.appspot.com"
    # And the deployed api host "http://localhost:8006"
    And the product url "<mock_product_url>" 
    When i send "<http_method>" request to "<end_point>" with above product url in body 
    Then the response should be "<expected_response_body>"
    Then the response code should be <expected_response_status_code> 

    Examples:
    | name                                | mock_product_url                | http_method | end_point | expected_response_body                    | expected_response_status_code |
    | test amazon product availability    | /mock/amazon_available.html     | get         | /product  | amazon_available_product_response.json    | 200                           |
    | test amazon product unavailability  | /mock/amazon_unavailable.html   | get         | /product  | amazon_unavailable_product_response.json  | 200                           |
    | test flipkart product availability  | /mock/flipkart_available.html   | get         | /product  | flipkart_available_product_response.json  | 200                           |
    | test flipkart product unavailability| /mock/flipkart_unavailable.html | get         | /product  | flipkart_unavailable_product_response.json| 200                           |    
    | test unsupport website              | /mock/unsupported_file.html     | post        | /product  |                                           | 406                           |