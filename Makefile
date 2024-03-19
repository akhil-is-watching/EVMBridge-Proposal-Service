build:
	@go build -o bin/indexer

run: build
	@./bin/indexer --listenRPC https://rpc.ankr.com/polygon_mumbai/622d31bd5dc0851b5d945cfcb9df9b788009f56b060e39b4f0687e5097caf60f --listenChain 80001 --listenAddress 0x46a9B104856cDfFB067D648cCc4fF446eD365200 --foreignRPC https://rpc.ankr.com/fantom_testnet/622d31bd5dc0851b5d945cfcb9df9b788009f56b060e39b4f0687e5097caf60f --foreignChain 4002 --foreignAddress 0x64456b8d0ce2A5ecE06276E8b44608db1ED6B2d1 --privateKey b97ce166cd35b47cb883219a1c42424203f5e65d685b245ff6312086c13965f7


