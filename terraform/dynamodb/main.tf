provider "aws" {
  region = "us-east-1"
}

resource "aws_dynamodb_table" "awsdynamo" {
  name             = "terraform-lock"
  hash_key         = "LockID"
  stream_enabled   = true
  stream_view_type = "NEW_AND_OLD_IMAGES"

  attribute {
    name = "LockID"
    type = "N"
  }

}