

Update PATH environment variable
```
export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
```