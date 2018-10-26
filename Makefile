build:
	go build -o terraform-provider-occm
test:
	terraform init
	terraform plan
clean:
	rm terraform-provider-*
