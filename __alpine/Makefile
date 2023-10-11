all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./metabypass.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.21.3 x86_64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/metabypass.exe $(DESTDIR)$(PREFIX)/bin/metabypass
	strip $(DESTDIR)$(PREFIX)/bin/metabypass
