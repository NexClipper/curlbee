# CoolBee
Calling APIs with a YAML instruction.

## Example

```yaml
call:
	clusters: 
		method: "get"
		url: 	"10.30.35.23/clusters"
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
			- "cluster.list"
	  return:
		ok: true
		list: "cluster.list"
```

The instruction above will call 
```
GET 10.30.35.23/clusters?page=1&limit=3
```
with a header X-Auth-Token: {{env.OPEN_STACK_TOKEN}} where the value is from an environment variable OPEN_STACK_TOKEN.
