# variables
SRV_A=go run user/user.go
SRV_B=go run account/acc.go
SRV_C=go run company/comp.go


run_a:
	$(SRV_A)

run_b:
	$(SRV_B)

run_c:
	$(SRV_C)
