all:
	@cd ~/Development/trueblocks-core/build && make sdk
	@cd typescript && make
#	@cd python && make

publish:
	@cd typescript && yarn release
#	@cd python && make
