HOSTNAME=registry.terraform.io
NAMESPACE=joaquimmmagalhaes
NAME=sftpgo
BINARY=terraform-provider-${NAME}
VERSION=0.0.5
OS_ARCH=linux_amd64

default: install

build:
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${BINARY}_v${VERSION}

dev:
	rm .terraform.lock.hcl
	rm -r .terraform
	make install
	terraform init
	terraform plan
