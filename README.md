# dmc-converter

[![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)

API application for converting a list of Hex or RGB values into the closest DMC colors available.

## TODO List
- Error Handling.
- Documentation.
- Clean up code organisation and modules.
- Add swagger docs api somehow.

----
This project has been used as a way to learn the go programming language and its idioms. If you have any issues with the project or wish to contribute, please give reasons why to teach me and help me learn, would be much appreciated.

## Endpoints
### GET
http://localhost:8080/dmc/get <- Get all DMC colours

### POST
http://localhost:8080/dmc/submit <- Submit DMC colours with payload
```
{
    "inputType":"csv",
    "inputs":["1,2,3"]
}
````

csv, hex and space seperated are allowed