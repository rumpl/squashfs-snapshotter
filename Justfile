# Build the project
build:
    go build main.go

# Create a test squashfs image
image:
    hack/mkimage.sh