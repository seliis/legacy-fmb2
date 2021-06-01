targetPath = /mnt/d/fmb

check:
	clear
	@if [ -d $(targetPath) ]; then \
		echo "** Miho: Directory is Exist, Now Calling Clean Command **"; \
		rm -rf $(targetPath); \
		echo "** Miho: Directory Has Been Removed **"; \
	else \
		echo "** Miho: Directory Not Exist, We Good to Go! **"; \
	fi \

build: check
	@npm run build
	@mkdir -p $(targetPath)/document
	@cp -r ./build/* $(targetPath)
	@rm -rf ./build
	@echo "** Miho: Brand New Build Moved to Target Path **";
	@GOOS=windows GOARCH=amd64 go build -o $(targetPath)/server.exe
	@echo "** Miho: Works Done **"