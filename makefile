gen.Mock:
	mockery --all --dir internal --keeptree
	mockery --all --dir pkg --keeptree