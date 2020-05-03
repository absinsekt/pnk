.DEFAULT_GOAL := dev

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


serve:
	@set -a && source ./.env && set +a; \
cd app; \
reflex -r '\.go$\' -s -- sh -c 'go run pnk.go'


build:
	@cd www; npm run build
