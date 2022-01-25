
The produce database is a single, in memory array of data and supports reads and writes {cm:2022-01-25}
Supports adding more than one new produce item at a time {cm:2022-01-25}
The produce includes name, produce code, and unit price {cm:2022-01-25}
The produce name is alphanumeric and case insensitive {cm:2022-01-25}
The produce codes are sixteen characters long, with dashes separating each four character group {cm:2022-01-25}
The produce codes are alphanumeric and case insensitive {cm:2022-01-25}
The produce unit price is a number with up to 2 decimal places
Error handling (GET nonexistent produce) {cm:2022-01-25T17:26:15}
Error handling (bad POST payload, etc.)
The API is RESTful {cm:2022-01-25}
The functionality is tested and verifiable without manually exercising the endpoints
The API supports adding and deleting individual produce. You can also get any produce in the database. {cm:2022-01-25}
Produce Code is unique (no apple or pear have the same code)