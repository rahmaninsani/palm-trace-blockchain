# Configs
.ONESHELL:
.DEFAULT_GOAL := default

default:
	@echo "CLI Tools by Rahman Insani"

fablo-download:
	@curl -Lf https://github.com/hyperledger-labs/fablo/releases/download/1.1.0/fablo.sh -o ./fablo

fablo-privilege:
	@sudo chmod +x ./fablo

generate-env:
	@python3 generate_env_file.py