module github.com/baransonmez/coffein/internal/user

go 1.19

require (
	github.com/baransonmez/coffein/kit v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
)

require github.com/julienschmidt/httprouter v1.3.0 // indirect

replace github.com/baransonmez/coffein/kit => ../../kit
