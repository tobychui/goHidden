# goHidden
[![Go Reference](https://pkg.go.dev/badge/github.com/tobychui/goHidden.svg)](https://pkg.go.dev/github.com/tobychui/goHidden)

A Go module to set a folder hidden and check if a folder is hidden

Support Windows and Linux



## Usage

```
//Hide a folder
err := hidden.HideFile("./test")
if err != nil {
	panic(err)
}

//Check if a folder is hidden
isHidden, err := hidden.IsHidden("./test", false)
if err != nil {
	panic(err)
}
```

You can enable parent directories checking on ```IsHidden``` by passing "true" as the last parameter for ```IsHidden```. This will allow```IsHidden``` to check all the  upstream folders to see if the file is located inside a hidden folder. Set this to false for checking only the targeted file / folder.

```
hidden.HideFile("./test")
isHidden, _ := hidden.IsHidden("./test/myfile.txt", true)
//isHidden is true
```



## License

MIT





### 
