BINDIR	= $(CURDIR)/bin
DOCDIR	= $(CURDIR)/doc

build:
	cd $(CURDIR)/src/slurm-qstat && go get slurm-qstat/src/slurm-qstat && go build -o $(CURDIR)/bin/slurm-qstat

all: depend build strip install

depend:
	# go mod will handle dependencies

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/bin
	mkdir -p -m 0755 $(DESTDIR)/usr/share/man/man1

strip: build
	strip --strip-all $(BINDIR)/slurm-qstat

ifneq (, $(shell which upx 2>/dev/null))
	upx -9 $(BINDIR)/slurm-qstat
endif

install: strip destdirs install-bin install-man

install-bin:
	install -m 0755 $(BINDIR)/slurm-qstat $(DESTDIR)/usr/bin

install-man:
	install -m 0644 $(DOCDIR)/slurm-qstat.1 $(DESTDIR)/usr/share/man/man1
	gzip -9 $(DESTDIR)/usr/share/man/man1/slurm-qstat.1

clean:
	/bin/rm -f bin/slurm-qstat

distclean: clean
	rm -rf src/github.com/
	rm -rf pkg/

uninstall:
	/bin/rm -f $(DESTDIR)/usr/bin/slurm-qstat

