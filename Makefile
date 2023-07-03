.PHONY: serve

.prepare:
	go install github.com/cespare/reflex@latest
serve:
	cd golang&&reflex -c reflex.conf