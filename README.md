# Restful Policy Managment


## Example

```bash
curlp http://localhost:3000/validate --data '{"realm":"org.risk.support", "name":"ali"}'
```

## Build 

```bash
go build -o policy main.go
chmod +x policy
```