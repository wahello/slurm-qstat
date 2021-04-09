BINDIR	= $(CURDIR)/bin

depend:
	# go mod will handle dependencies

build:
	cd $(CURDIR)/src/slurm-qstat && go build -o $(CURDIR)/bin/slurm-qstat

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/bin

strip: build
	strip --strip-all $(BINDIR)/slurm-qstat

ifneq (, $(shell which upx 2>/dev/null))
	upx -9 $(BINDIR)/slurm-qstat
endif

install: strip destdirs install-bin

install-bin:
	install -m 0755 $(BINDIR)/slurm-qstat $(DESTDIR)/usr/bin

clean:
	/bin/rm -f bin/slurm-qstat

distclean: clean
	rm -rf src/github.com/
	rm -rf pkg/

uninstall:
	/bin/rm -f $(DESTDIR)/usr/bin

all: depend build strip install

