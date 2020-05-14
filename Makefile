.DEFAULT_GOAL := info

# TODO init
# go get github.com/cespare/reflex


dev:
	@tmux new-session -d 'cd www && npm run dev'; \
tmux split-window -h 'make serve'; \
tmux attach


front:
	@tmux new-session -d 'cd www; npm run dev'; \
tmux split-window -h 'cd app; go run pnk.go'; \
tmux attach


story:
	@cd www; npm run storybook


serve:
	@set -a && source ./.env && set +a; \
cd app; \
reflex -r '\.go$\' -s -- sh -c 'go run pnk.go'


build:
	@cd www; npm run build


test:
	@set -a && source ./.env && set +a; \
cd $(path); \
go test -run $(test) -count=1


db-init:
	@set -a && source ./.env && set +a; \
cd app/models/user; \
go test -run Test_createTable -count=1; \
go test -run Test_createSuperUser -count=1


db-create:
	@set -a && source ./.env && set +a; \
cd app/models/$(mdl); \
go test -run Test_createTable -count=1


db-drop:
	@set -a && source ./.env && set +a; \
cd app/models/$(mdl); \
go test -run Test_dropTable -count=1


info:
	@echo "dev"
	@echo "front"
	@echo "serve"
	@echo "build"
	@echo "test\t\t\tE.g.: db-test path=app/models/order tst=TestGetOrder"
	@echo "db-init"
	@echo "db-create"
	@echo "db-drop"
