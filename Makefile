clean:
	@rm -rf compiled
	@rm -rf target

compile-darwin:
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o compiled/darwin

compile-win32:
	goversioninfo -icon=Windows/icon_512x512.ico -manifest=Windows/windows.manifest
	GOOS=windows GOARCH=386 go build -ldflags "-s -w -H=windowsgui" -o compiled/win32.exe
	@rm resource.syso

compile-win64:
	goversioninfo -64 -icon=Windows/icon_512x512.ico -manifest=Windows/windows.manifest
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -H=windowsgui" -o compiled/win64.exe
	@rm resource.syso

compile-all: | compile-darwin compile-win32 compile-win64

darwin_path="target/SexyDKPSync.app"
darwin_path_c="$(darwin_path)/Contents"
build-darwin:
	@mkdir -p $(darwin_path_c)/{MacOS,Resources}
	@cp compiled/darwin $(darwin_path_c)/MacOS/SexyDKPSync
	@cp MacOS/Info.plist $(darwin_path_c)/
	@cp MacOS/icon* $(darwin_path_c)/Resources/

build-win32:
	@cp compiled/win32.exe target/SexyDKPSync_Win32.exe

build-win64:
	@cp compiled/win64.exe target/SexyDKPSync_Win64.exe

build-all: | build-darwin build-win32 build-win64

release-darwin: | compile-darwin build-darwin

release-win32: | compile-win32 build-win32

release-win64: | compile-win64 build-win64

release-all: | release-darwin release-win32 release-win64

deploy-local:
	@rm -rf /Applications/World\ of\ Warcraft/_classic_/SexyDKPSync.app/
	@mv $(darwin_path) /Applications/World\ of\ Warcraft/_classic_/

deploy:
	cd target && zip -r7 SexyDKPSync_MacOS.zip SexyDKPSync.app
	cd target && scp SexyDKPSync_* root@192.168.1.55:/root/na_sozvone/sync/
	rm target/SexyDKPSync_MacOS.zip

deploy-fake:
	@rm -rf deploy
	@mkdir -p deploy
	cd target && zip -r7 SexyDKPSync_MacOS.zip SexyDKPSync.app
	cp target/SexyDKPSync_* deploy/