.PHONY: deploy
deploy:
	@vercel --prod

.PHONY: gen
gen:
	@oapi-codegen -generate types -package openapi ./openapi.yaml > ./internal/openapi/types.gen.go

.PHONY: destory
destory:
	@vercel project rm samplegovercel
