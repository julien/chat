SHELL = /bin/bash
NAME = websocket_chat

all: clean $(NAME) run

$(NAME): 
	@echo building $(NAME)
	@go build

run:
	@echo running $(NAME)
	@PORT=3000 ./${NAME}

release:
	@go get

clean:
	@echo cleaning $(NAME)
	@rm -rf $(NAME)

.PHONY: clean $(NAME) run release

