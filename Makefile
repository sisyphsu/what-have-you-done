# Go parameters
TARGET=whyd

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
current_dir := $(dir $(mkfile_path))

#UNAME := $(shell uname)
#
#ifeq ($(UNAME), Darwin)
#CFLAGS = -Ofast
#endif
#
#ifeq ($(UNAME), Linux)
#CFLAGS = -O3
#endif

.PHONY: build
all: build package

.ONESHELL:
build:
	cd $(current_dir)
	# compile rcc
	qrc2go $(current_dir)view
	# compile exe
	go build -v -o $(TARGET) main_prod.go

.ONESHELL:
package: build
	cd $(current_dir)
	# create app
	rm -Rf build/*
	mkdir -p build/$(TARGET).app/Contents/{MacOS,Resources}
	cp assets/Info.plist build/$(TARGET).app/Contents/
	cp assets/whyd.icns build/$(TARGET).app/Contents/Resources/
	mv $(TARGET) build/$(TARGET).app/Contents/MacOS/$(TARGET)
	# create dmg
	macdeployqt build/$(TARGET).app -dmg -always-overwrite -qmldir=$(current_dir)view