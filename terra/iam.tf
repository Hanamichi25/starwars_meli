resource "aws_iam_instance_profile" "ec2_profile_2" {
  name = "ec2-profile-2"
  role = aws_iam_role.role.name
}

resource "aws_iam_role" "role" {
  name = "ec2_role"
  path = "/"

  assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": "sts:AssumeRole",
            "Principal": {
               "Service": "ec2.amazonaws.com"
            },
            "Effect": "Allow",
            "Sid": ""
        }
    ]
}
EOF
}

resource "aws_iam_policy" "aws_ec2_s3_policy" {
  name        = "ec2-to-s3-policy"
  path        = "/"
  description = "This policy grants access to execute S3 from ec2"

  policy = data.aws_iam_policy_document.aws_ec2_s3_policy.json
}

resource "aws_iam_role_policy_attachment" "aws_ec3_attachment_s3" {
  role       = aws_iam_role.role.name
  policy_arn = aws_iam_policy.aws_ec2_s3_policy.arn
}

