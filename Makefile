# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
live/templ:
	go run github.com/a-h/templ/cmd/templ@latest generate --watch --proxy="http://localhost:8282" --open-browser=false -v

# run air to detect any go file changes to re-build and re-run the server.
live/server:
	go run github.com/air-verse/air@latest \
	--build.cmd "go build -o tmp/bin/main" --build.bin "tmp/bin/main" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# run tailwindcss to generate the styles.css bundle in watch mode.
live/tailwind:
	npm run style

# run esbuild to generate the index.js bundle in watch mode.
live/esbuild:
	npm run build

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
live/sync_assets:
	go run github.com/air-verse/air@latest  \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "assets" \
	--build.include_ext "js,css"

# start all 5 watch processes in parallel.
live: 
	make -j5 live/templ live/server live/tailwind live/esbuild live/sync_assets

migrateup:
	migrate -path db/migrations -database "sqlite3://db/test.db" -verbose up

prod:
	mkdir prod
	templ generate
	go build -o prod/main
	npm run prod
	cp -r assets prod/
	migrate -path db/migrations -database "sqlite3://prod/clipboard.db" -verbose up

clean:
	rm -rf prod tmp view/*_templ.go assets/*.{cs,j}s
