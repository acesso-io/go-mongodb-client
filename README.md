# mango
Mango is a wrapper for MongoDB's official driver created to simplify the configuration of clients and construction of queries and aggregation pipelines.

## Usage

### Creating a new client:
Mango's Client is a simple wrapper for MongoDB's official driver Client which allows the customization of options.

mango.NewClient accepts as a parameter a mango.Options struct. The data in this struct is used to modify you Client as you need. If **nil** is passed as the parameter, mango.DefaultOptions will be used instead.

```
// Create client with DefaultOptions
client := mango.NewClient(nil)
```

mango.Options can also be read directly from a YAML config file. In this case, all you need to do is prepare a config file and call mango.NewOptionsFromConfigFile, passing the file path as its parameter:

```
// Create client with customized options from a YAML file
client := mango.NewClient(
	mango.NewOptionsFromConfigFile("config.yaml"),
)
```

PS: don't forget to connect you client to MongoDB servers after you have created it!

```
if err := client.Connect(); err != nil {
	log.Fatal(errors.Wrap(err, "failed to connect to MongoDB servers"))
}
```

### Creating queries and aggregation pipelines:
Mango has built-in query and aggregation pipelines builders. You can use them to create aggregation stages or simple bson.M blocks:

```
// Use "my_database" as the target database
db := client.Database("my_database")

// Initiate new query builder
query := mango.NewQuery()
query.Equal("_id", uuid.Must(uuid.Parse("548c8a3b-2906-4693-a601-98dcf8225a25")))
query.In("generation", 1, 2, 3)
query.LessThan("timestamp", time.Now().Add(-30*24*time.Hour))

// Make a find query using the BSON from the query builder
cursor, err := db.Collection("my_collection").Find(ctx, query.BSON())
if err != nil {
	log.Fatal(errors.Wrap(err, "failed to get query results"))
}

// Initiate new aggregation builder
aggregation := mango.NewAggregation()

// Use the query built on a $match stage
aggregation.Match(query)

// Make aggregation query using the pipeline from the aggregation builder
cursor, err := db.Collection("my_collection").Aggregate(ctx, aggregation.Pipeline())
if err != nil {
	log.Fatal(errors.Wrap(err, "failed to get query results"))
}
```
