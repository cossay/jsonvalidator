# jsonvalidator
A simple GO validation library for validating raw JSON documents.

## Dependencies
```
github.com/tidwall/gjson
```
``` 
github.com/asaskevich/govalidator
 ```
 ```
 github.com/google/go-cmp/cmp
 ```

## Installation
```
go get -u github.com/cossay/jsonvalidator
```

## Example

``` 
package main

import (
	"encoding/json"
	"fmt"

	validator "github.com/cossay/jsonvalidator"
)

func main() {

	data := []byte(`{
		"name": "    ",
		"age": 100,
		"school": {
			"name": "UCC"
		},
		"friends": [
			{
				"name": "Kwame",
				"age": "5"
			},
			{
				"name": "Mansa",
				"age": "3"
			},
			{
				"name": "Mansa",
				"age": "3"
			}
		]
	}`)

	constraints := map[string][]validator.Constraint{
		"name": []validator.Constraint{
			validator.Required("Name is required"),
		},
		"age": []validator.Constraint{
			validator.Required("Age is required"),
			validator.TypeInt("Age must be a integer"),
			validator.EqualTo(10, "Age must be equal to 10"),
		},
		"school": []validator.Constraint{
			validator.Required("School is required"),
			validator.TypeObject("School must be an object"),
			validator.Object(map[string][]validator.Constraint{
				"name": []validator.Constraint{
					validator.Required("School name required"),
				},
			}),
		},
		"friends": []validator.Constraint{
			validator.Required("Friends required"),
			validator.ArrayMinLength(1, "A minimum of 1 friend required."),
			validator.ArrayMaxLength(30, "A maximum of 2 friends required."),
			validator.ArrayLength(3, "Exactly 3 friends required."),
			validator.Array([]validator.Constraint{
				validator.Required("Friend is required"),
				validator.TypeObject("Friend must be an object"),
				validator.Object(map[string][]validator.Constraint{
					"name": []validator.Constraint{
						validator.Required("Friend name required"),
					},
					"age": []validator.Constraint{
						validator.Required("Age is required"),
						validator.TypeString("Age must be a string."),
					},
				}),
			}),
		},
	}
	
	v := validator.NewValidator()

	if e := v.CheckDocument(data); nil != e {
		panic(e) 
	}

	fmt.Println(v.Validate(data, constraints))
}

```
