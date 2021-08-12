# Create Collection

API to create collection according to the schema specified

## Params

- `ctx` context.Context, context to control API invocation process;

- `collSchema` pointer of entity.Schema, the collection schema definition;

- `shardNum` int32, the shard number of the collection to create; If the `shardNum` is set to 0, default shard number will be used, which is 2

## Response

- `err` error of the creation process (if any), possible error list:

    - ErrClientNotReady, is the client is not connected

    - error for client already exists
    
    - error fo API invocation failed 
