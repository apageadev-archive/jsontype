# JSONType

The JSONType package provides the mechanisms to manage strictly typed JSON Documents.


![codecoverage](https://apageadev.github.io/jsontype-coverage-badge.svg)

## Schemas

At the core of the JSONType package is the "Schema". A Schema, which is just a JSON Document, defines how other JSON documents should look. After the schema has been loaded into memory, it can be used to enforce data validation with minimal effort.

An example schema:

```json
{
	"properties": {
		"name": {
			"type": "string",
			"rules": {
				"max_length": 5
			}
		}
	}
}
```

An example document that would satisfy the schema:

```json
{ "name": "Luna" }
```

An example document that would NOT satisfy the schema:

```json
{ "name": "Blueberry" }
```

Even though the JSON format is correct, the value is too long
based on the defined rules for the name property. This is an
example of how a "strictly" typed document can be enforced.

### Primitive Property Types

- string
- number
- boolean
- object
- array
- list

### TODO:

- [] Add Formats from V10 and Gookit Validator
- [] Add Custom Error Types
- [] Support Extending Schemas
- [] Add Benchmarks
- [] Update Readme to Show Usage Examples
- [] Document Available Validation Rules
