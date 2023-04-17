# Mango

`Mango` is a Golang package that acts as a wrapper around the `mongoDB-go-client`. The package is created to simplify access to MongoDB using Golang through abstracting away all the intricacies of context management and error handling.

It currently support a subset of the features available on `mongoDB-go-client` including:

1. Initialize a connection client
2. List all object in a collection
3. Make read query to a collection

## Establish Connection

Connection to a MongoDb cluster can be established though a connection client.

```go
import "gethub.com/cygnusborjigin/mango"

connectionClient := mango.NewAtlasConnection("<connection string>")
```

The returned value `connectionClient` is of type `mango.AtlasConnection`, which should be noted contains a MongoDB client but is not a mongoDB client itself. The mongoDB client can be extracted using the `GetClient()` method.

## Issue Read Query

Read query can be issued though an already created connection client. To simplify the process and ease the learning curve, `mango` has its own query language that is more user friendly.

To issue read query, use the `GetObject()` method of a collection object.

```go
import "gethub.com/cygnusborjigin/mango"

connectionClient := mango.NewAtlasConnection("<connection string>")
collection := mango.NewAtlasCollection(connectionClient, "<database name>", "<collection name>")
res := collection.GetObject(parsedFilters)
```

## Example Query

```json
{
  "filter": [
        {
          "state": "Massachusetts",
          "population": ">2500",
          "returnField": ["state", "city", "population"]
        }
      ]
}
```
