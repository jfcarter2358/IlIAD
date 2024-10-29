create table if not exists providers_aws_s3_bucket (
    bucket_name serial primary key,
    tags json not null,
)

create table if not exists providers_aws_iam_role (
    role_name serial primary key,
    policy_attachments json not null,
)


create table if not exists providers_aws_iam_policy (
    policy_name serial primary key,
    policy_doc json not null,
)
