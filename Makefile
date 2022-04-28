profile:
	curl -o profile.out http://localhost:8080/debug/pprof/profile\?seconds\=20

heap:
	curl -o heap.out http://localhost:8080/debug/pprof/heap\?seconds\=20

profile-web:
	go tool pprof -http=":8000" profile.out

heap-web:
	go tool pprof -http=":8000" heap.out

ballast-ps:
	ps -eo pmem,comm,pid,maj_flt,min_flt,rss,vsz --sort -rss | numfmt --header --to=iec --field 4-5 | numfmt --header --from-unit=1024 --to=iec --field 6-7 | column -t | egrep "[b]allast|[P]ID"
