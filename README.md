# curlbee
Calling APIs with a YAML instruction.

## Bee Charts
The YAML files instruction for calling API and templating the result is called a bee chart.

### Example

```yaml
call:
  clusters: 
    method: "get"
    url: "10.30.35.23/clusters"
    query:
      page: 1
      limit: 3
    headers:
      - name: "X-Auth-Token"
        value: "{{env.OPEN_STACK_TOKEN}}"
response:
  - if:
      status: 200
      value:
        ok: true
        exists:
          - "$Res.clusters.list"
    return:
      ok: true
      list: "$Res.clusters.list"
```

The instruction above will call 
```
GET 10.30.35.23/clusters?page=1&limit=3
```
with a header X-Auth-Token: {{env.OPEN_STACK_TOKEN}} where the value is from an environment variable OPEN_STACK_TOKEN.

## Templating

```yaml
call:
  clusters: 
    method: "post"
    url: "10.30.35.23/clusters"
    body:
      name: {{.Values.clusterName}}
      description: {{.Values.clusterDescription}}
      origin: {{env.SERVER_ORIGIN}}
    headers:
      - name: "X-Auth-Token"
        value: "{{env.OPEN_STACK_TOKEN}}"
```


## In Command Line
```
$ curlbee cluster.yaml
```

### Passing variables to template
Based on the template above you can pass both clusterName and clusterDescription by running
```
$ $ curlbee cluster.yaml -p clusterName="My Cluster" -p clusterDescription="hello"
```


## As a Service
```
# running it on port 3001
$ curlbee start -p 3001

# pass yaml to the api
POST localhost:3000/call 
```



