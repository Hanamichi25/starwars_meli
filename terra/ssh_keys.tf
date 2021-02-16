resource "aws_key_pair" "key_class_file" {
  key_name   = "class_file"
  public_key = file("ec2-key-drmartin.pub")
}
