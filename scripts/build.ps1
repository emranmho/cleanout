# script/build.ps1

$mainPath = "..\main.go"

Write-Host "Building for Windows (amd64)..."
$env:GOOS = "windows"
$env:GOARCH = "amd64"
go build -o "..\cleanout.exe" $mainPath

Write-Host "Building for Linux (amd64)..."
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o "..\cleanout-linux" $mainPath

Write-Host "Building for macOS (amd64)..."
$env:GOOS = "darwin"
$env:GOARCH = "amd64"
go build -o "..\cleanout-mac" $mainPath

# Clean up environment variables
Remove-Item Env:GOOS
Remove-Item Env:GOARCH

Write-Host "✅ Cross-platform builds completed in root folder."
