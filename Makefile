codegen:
	@echo "${GREEN}Generating templ code...${NC}"
	templ generate
	@echo "${GREEN}Generating sqlc code...${NC}"
	sqlc generate
	@echo "${GREEN}Generating Tailwind CSS code...${NC}"
	npx tailwindcss -c ./tailwind.config.js -i ./ui/css/main.css -o ./public/main.css --minify

watch:
	npx tailwindcss -c ./tailwind.config.js -i ./ui/css/main.css -o ./public/main.css --watch