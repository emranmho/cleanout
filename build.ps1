# build.ps1
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o cleanout.exe main.go

$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o cleanout main.go

$env:GOOS = "darwin"
$env:GOARCH = "amd64"
go build -o cleanout-mac main.go

# Clean up
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
