//gitdb:doc:begin
//gitdb:meta:name
//foobar
//gitdb:field:tf
module "foobar_bucket" {
    source = "some/terraform/registry"
    version = "some_version"
    
    bucket_name = "foobar"
}

//gitdb:doc:end
