.DEFAULT_GOAL := dev


dev:
	@tmux new-session -d 'cd www; npm run dev'; \
tmux split-window -h 'PNK_DEBUG=True go run main.go'; \
tmux attach


build:
	@cd www; npm run build
