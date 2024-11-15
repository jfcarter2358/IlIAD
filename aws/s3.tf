//gitdb:doc:begin
//gitdb:meta:name
//bar
//gitdb:field:tf
module "bar_bucket" {
    source = "some/terraform/registry"
    version = "some_version"
    
    bucket_name = "bar"
}

//gitdb:doc:end
