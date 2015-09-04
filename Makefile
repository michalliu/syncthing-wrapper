.PHONY : all build clean mips mips_test

SYNCTHING_WORKSPACE=$(shell pwd)/src/github.com/syncthing/syncthing/Godeps/_workspace
DUTOOL=$(SYNCTHING_WORKSPACE)/src/github.com/calmh/du

# the order is important! otherwise go lint will report not found, dont know why
export GOPATH:=$(SYNCTHING_WORKSPACE):$(shell pwd)

#GO_DISABLE_SETFINALIZER will remove the annoying message
#runtime.SetFinalizer not implemented for MIPS arch
#export GO_DISABLE_SETFINALIZER=1

# VERY IMPORTENT STAGING_DIR MUST BE SET
# staging dir is set in .bashrc
#export STAGING_DIR=

# install golang to /usr/local/go

# helpful makefile sample
# https://github.com/iancmcc/home-directory/blob/master/makefile

# mips CrossCompiler base dir
MIPS_GCC_DIR:=$(STAGING_DIR)/toolchain-mipsel_24kec+dsp_gcc-5.2.0_eglibc-2.19/bin
# mips target specific tool
MIPS_GO_TOOL:=$(shell pwd)/mips/bin
# mips gcc tool, including gcc and gccgo
# added to PATH ENV if compiling using -compiler gccgo flag
MIPS_GCC:=$(shell pwd)/mips/bin/gcc
# custom go tool fix unsupported mips platform error
# and support cgo enviroment variable
MIPS_GO:=$(MIPS_GO_TOOL)/go
# custom cgo tool supports automatic generate go definations from C 
CGO_PATH:=$(MIPS_GO_TOOL)/cgo
# custom build vet tool
VET_PATH:=$(MIPS_GO_TOOL)/vet
# compiled binary
MIPS_SYNCTHING:=$(shell pwd)/src/github.com/syncthing/syncthing/bin/linux_mipso32/syncthing

all: build_mips

# demo program for mips
# CGO_PATH is read by faux-cgo
# CGO_ENABLED=1 is neccssary for gccgo compiler
mips_demo: mips_demo_godef
	PATH=$(MIPS_GO_TOOL):$(MIPS_GCC):$(PATH) \
	GOOS=linux \
	GOARCH=mipso32 \
	CGO_PATH=$(CGO_PATH) \
	CGO_ENABLED=1 \
	go build -v -compiler gccgo -gccgoflags '-static-libgo' \
	-cgoflags '-debug-gcc=false' \
	mipstest
	scp mipstest root@192.168.2.66:/mnt/sda1

#$1 go sourcefile
#generate godefs for mips target for cross compile
#crosscompile cgo genrate defs for mips arch
#MIPS_GCC mustbe set to generate the correct definations
#$(eval xx:=$(call suffix, $1))
#$(eval xx:=$(call basename, $1))
define GEN_MIPS_GODEF
	# uncomment //type GOSTRUCT C.C_STRUCT
	sed -i 's@^//type\s@type @g' $(1)
	# lookfing for above definations and create a stub
	# using cgo
	PATH=$(MIPS_GO_TOOL):$(MIPS_GCC):$(PATH) \
	CGO_PATH=$(CGO_PATH) \
	GOOS=linux \
	GOARCH=mipso32 \
	go tool cgo -godefs \
	$(1) | tee \
	$(call basename, $(1))_defs$(call suffix, $(1))
	# restore the source file to it's original
	sed -i 's@^type\s@//type @g' $(1)
endef

haha2: haha
	echo $(xx)

haha:
	$(call GEN_MIPS_GODEF, src/mipstest/demo/cgoStatfs.go)

mips_demo_godef: mips_dutool_godef
	$(call GEN_MIPS_GODEF, src/mipstest/demo/cgoStatfs.go)
	$(call GEN_MIPS_GODEF, src/mipstest/demo/cgoStatfs2.go)

mips_dutool_godef:
	$(call GEN_MIPS_GODEF, $(DUTOOL)/diskusage_mips.go)
	sed -i 's@_Ctype_struct___0@Fsid_t_Go@g' $(DUTOOL)/diskusage_mips_defs.go


# demo program for host, just make sure the code is ok
demo: demo_godef
	go build mipstest
	./mipstest

# $(1) source file
# we must use our modified go to load our customed cgo command
# MIPS_GCC MUST REMOVED USE THE SYSTEM GCC TOOLCHAIN
define GEN_GODEF
	# uncomment //type GOSTRUCT C.C_STRUCT
	sed -i 's@^//type\s@type @g' $(1)
	# lookfing for above definations and create a stub
	# using cgo
	PATH=$(MIPS_GO_TOOL):$(PATH) \
	CGO_PATH=$(CGO_PATH) \
	go tool cgo -godefs \
	$(1) | tee \
	$(call basename, $(1))_defs$(call suffix, $(1))
	# restore the source file to it's original
	sed -i 's@^type\s@//type @g' $(1)
endef

demo_godef: dutool_godef
	$(call GEN_GODEF, src/mipstest/demo/cgoStatfs.go)
	$(call GEN_GODEF, src/mipstest/demo/cgoStatfs2.go)

# MIPS_GCC MUST REMOVED USE THE SYSTEM GCC TOOLCHAIN
dutool_godef:
	$(call GEN_GODEF, $(DUTOOL)/diskusage_mips.go)
	sed -i 's@_Ctype_struct___0@Fsid_t_Go@g' $(DUTOOL)/diskusage_mips_defs.go

# ========= the syncthing part =========
# build syncthing for host machine
build: build_tool dutool_godef
	cd src/github.com/syncthing/syncthing; \
	./build

# cross compile syncthing for mips
# set GOOROOT to somewhere doesn't exists to make sure gccgo is using our libraries
build_mips: build_tool mips_dutool_godef
	cd src/github.com/syncthing/syncthing; \
	PATH=$(MIPS_GCC):$(PATH) \
	GOCMD=$(MIPS_GO) \
	GOOS=linux \
	GOARCH=mipso32 \
	CGO_PATH=$(CGO_PATH) \
	VET_PATH=$(VET_PATH) \
	CGO_ENABLED=1 \
	./build -compiler gccgo

# create build tool
build_tool:
	cd src/github.com/syncthing/syncthing; \
	go build build.go

# copy syncthing to board
test: build_mips
	scp $(MIPS_SYNCTHING) root@192.168.2.66:/mnt/sda1

# install dependencies
tools:
	go get -u github.com/golang/lint/golint
	go get -u github.com/pivotal-golang/bytefmt
	go get -u golang.org/x/tools/cmd/vet

# clean
clean:
	cd src/github.com/syncthing/syncthing; \
	rm -vf build; \
	rm -rvf bin; \
	rm -rvf Godeps/_workspace/bin; \
	rm -rvf Godeps/_workspace/pkg

#prepare sysm links
#gccgo is for cross compile golang for mips target
#gcc is for cross compile cgo for mips target
prepare:
	ln -sf $(MIPS_GCC_DIR)/mipsel-openwrt-linux-gnu-gcc \
		$(MIPS_GO_TOOL)/gcc/gcc;
	ln -sf $(MIPS_GCC_DIR)/mipsel-openwrt-linux-gnu-gccgo \
		$(MIPS_GO_TOOL)/gcc/gccgo;
