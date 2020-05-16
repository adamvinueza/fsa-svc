# fsa-svc

REST microservice for recognizing strings.

## API

```
GET /fsa

Returns the list of available finite-state automata.

Sample response:

[
    {
        "id": 1,
        "name": "evens",
        "description": "recognizes an even number of characters from the set [A-Za-z0-9]",
        "examples": {
            "accept": [
                "aa",
                "abcdef",
                "qxrFF8"
            ],
            "reject": [
                "a",
                "",
                "123ab"
            ]
        }
    }
]

GET /fsa/{fsa-id}?test=test_string

Tests whether a string is accepted by the specified finite-state automaton.

Query parameter: test
Value: a string to test against the fsa

```
