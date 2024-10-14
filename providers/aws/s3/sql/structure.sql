create table aws_s3 (
    id serial primary key,
    bucket_name text not null,
    access_rules text,
    versioning boolean,
    readonly_access_accounts text,
    access_log_config text,
    access_tags text,
    aws_profile text
)
