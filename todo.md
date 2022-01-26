
The produce database is a single, in memory array of data and supports reads and writes  {cm:2022-01-25T20:22:54}
Supports adding more than one new produce item at a time  {cm:2022-01-25T20:22:57}
The produce includes name, produce code, and unit price  {cm:2022-01-25T20:22:55}
The produce name is alphanumeric and case insensitive  {cm:2022-01-25T20:22:56}
The produce codes are sixteen characters long, with dashes separating each four character group  {cm:2022-01-25T20:22:58}
The produce codes are alphanumeric and case insensitive  {cm:2022-01-25T20:22:59}
The produce unit price is a number with up to 2 decimal places ()  {cm:2022-01-25T22:23:01}
Error handling (GET nonexistent produce)  {cm:2022-01-25T20:23:02}
Error handling (bad POST payload, etc.)  {cm:2022-01-25T20:23:03}
The API is RESTful  {cm:2022-01-25T20:23:04}
The functionality is tested and verifiable without manually exercising the endpoints  {cm:2022-01-25T20:23:05}
The API supports adding and deleting individual produce. You can also get any produce in the database.   {cm:2022-01-25T20:23:06}
Produce Code is unique (no apple or pear have the same code) ()