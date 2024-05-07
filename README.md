# holidays-api
This code is a Go application that uses the Gin framework to create a REST API for fetching and filtering holidays. Here's a breakdown of the different components:
main Package

    The main package initializes the Gin router and sets up the /holidays endpoint to handle GET requests by invoking handlers.GetHolidays.

handlers Package

    The GetHolidays function is responsible for handling GET requests to /holidays.
        It initializes a HolidayService with a specific API endpoint ("https://api.victorsanmartin.com/feriados/en.json").
        It fetches holidays data using the FetchHolidays method of HolidayService.
        It parses the fetched JSON data into a models.HolidayResponse struct.
        It applies filters based on query parameters (type, start_date, end_date) using the applyFilters function.
        It responds with the filtered holidays in JSON format.

    The applyFilters function filters holidays based on type and date range specified in the query parameters.

services Package

    The HolidayService struct encapsulates the logic for fetching holidays from a specified API endpoint (APIEndpoint).
    The NewHolidayService function creates a new HolidayService instance.
    The FetchHolidays method performs an HTTP GET request to the API endpoint and returns the response body ([]byte) or an error.

How It Works

    When a GET request is made to /holidays endpoint, GetHolidays handler is invoked.
    The handler initializes a HolidayService and fetches holidays data from the specified API endpoint.
    The fetched data is parsed into a HolidayResponse.
    The handler then applies filters based on query parameters (type, start_date, end_date) using the applyFilters function.
    Filtered holidays are returned as a JSON response.

Improvements

    Error handling can be improved to provide more informative error messages.
    Validation of query parameters (start_date, end_date) could be enhanced to ensure they are in the correct format before parsing.
    Adding more robust logging to track requests and errors for debugging purposes.

This setup demonstrates a basic REST API for querying holidays with filtering capabilities in Go using the Gin framework and standard HTTP client.
