build:
	@echo Building app
	@go build -o go-hello

download:
	@echo Download go.mod dependencies
	@go mod download

install-tools: download
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

test-behavior:
	@echo Starting go-hello server
	@./go-hello & echo $$! > server.PID
	@echo Running behavior tests
	@godog . || true
	@echo Shutting down go-hello server
	@kill -s INT `cat server.PID` && rm server.PID

test-security:
	@echo Running security tests
	@gosec .
