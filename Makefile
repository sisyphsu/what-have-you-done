# Go parameters
TARGET=whyd

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

all: build

build:
	# flush dependencies
	go get
	# compile rcc
	qrc2go $(current_dir)/view
	# compile exe
	go build -o $(TARGET) -v
