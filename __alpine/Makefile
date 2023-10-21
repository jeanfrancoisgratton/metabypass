all: 
	cd ../src
	CGO_ENABLED=0 go build . -o ./screwUmeta.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.21.3 amd64
	sudo apk update --no-cache

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/screwUmeta.exe $(DESTDIR)$(PREFIX)/bin/screwUmeta
	strip $(DESTDIR)$(PREFIX)/bin/screwUmeta
