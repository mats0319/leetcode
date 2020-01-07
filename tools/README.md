# Tools 
> This file gives brief describe and method to use of each tools.

# test generate
### describe
It can generate a function's test file, but it only works on the same directory-structure as this repo's 'Solutions'.  
### use
```go run test_gen.go -n [number]```  

---
# go mod update
### describe
It can update packages in go.mod to it's master version.
### use
```go run go_mod_update.go```(only works if go.mod can be find in paths: './' or '../')  
```go run go_mod_update.go -dir [absolute dir of go.mod]```

---
# go mod replace(not determined to write it)
### describe
It can update packages in go.mod from 'replace' version to it's 'require' version.  
I want to write to write it because of the well-known reason. But after I use the proxy: 'https://goproxy.io'...
I haven't use replace in go.mod a long time. So I will write it some time I have nothing to do.
### use
I would like to add it to 'go_mod_update' tool, as an extra scene.  ^_^

---
###### Mario
I've been pretending to work hard, but you're really growing up.
