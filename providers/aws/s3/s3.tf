module "{{ .Name }}_bucket" {
    source = "some/terraform/registry"
    version = "some_version"
    
    bucket_name = "{{ .Name }}"
}
