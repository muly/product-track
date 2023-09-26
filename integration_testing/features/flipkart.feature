Feature:web scraping
    Scenario:test flipkart product availability
    Given the product url "http://localhost:8006/mock/flipkart_available.html" 
    When i send "post" request to "http://localhost:8006/product" with above product url in body 
    # Then the response should be "{"url":"http://localhost:8006/mock/flipkart_available.html","price":123,"availability":true}"
    Then the response code should be 200 
