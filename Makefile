# note: call scripts from /scripts

all: FlagAgent Master VpnSpy

RunFlagAgent: Dependency
	go run cmd/FlagAgent/10.0.0.1.go

RunVpnSpy: Dependency
	go run cmd/VpnSpy/main.go

RunMaster: Dependency
	go run cmd/Master/main.go


FlagAgents: Dependency
	mkdir -p build/package/FlagAgent/
	find cmd/FlagAgent -type f -name \*.go -exec go build -ldflags "-s -w" -o {}.out {} \;
	find cmd/FlagAgent -type f -name \*.out -exec mv {} build/package/FlagAgent \;

VpnSpy: Dependency
	mkdir -p build/package/VpnSpy/
	go build -o build/package/VpnSpy/VpnSpy cmd/VpnSpy/main.go

Master: Dependency
	mkdir -p build/package/Master/
	go build -o build/package/Master/Master cmd/Master/main.go

Dependency:
	go mod tidy
	go mod vendor

GenKey:
	mkdir -p cmd/FlagAgent/
	mkdir -p internal/pkg/SecretKey/
	mkdir -p configs/AutoGenKey/
	source scripts/.script/bin/activate && python scripts/Configure.py

GenYaml:
	mkdir -p configs/AutoGenYaml/
	source scripts/.script/bin/activate && python scripts/GenYml.py

Clean:
	go mod tidy
	go mod vendor
	go clean -i -r -x
	rm -rf cmd/FlagAgent/*
	rm -rf internal/pkg/SecretKey/*
	rm -rf configs/AutoGenKey
	rm -rf configs/AutoGenYaml
	rm -rf build/*
