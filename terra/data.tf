data "aws_iam_policy_document" "aws_ec2_s3_policy" {
  version = "2012-10-17"
  statement {
    sid    = "1"
    effect = "Allow"
    actions = [
      "s3:PutObject",
      "s3:GetObject"
    ]
    resources = [
      "arn:aws:s3:::starwars-mercadolibre/*"
    ]
  }
}
