resource "aws_launch_configuration" "ec2_go_config" {
  name                 = "web_config"
  image_id             = var.ami_id
  instance_type        = "t2.micro"
  user_data            = filebase64("userdata.sh")
  security_groups      = [aws_security_group.sg_main.id, aws_security_group.sg_http.id]
  key_name             = aws_key_pair.key_class_file.key_name
  iam_instance_profile = aws_iam_instance_profile.ec2_profile_2.id

}

resource "aws_autoscaling_group" "asg" {
  availability_zones   = ["us-east-2a", "us-east-2b", "us-east-2c"]
  desired_capacity     = 2
  max_size             = 3
  min_size             = 2
  launch_configuration = aws_launch_configuration.ec2_go_config.name
}

resource "aws_lb_target_group" "lb_target_group" {
  name        = "lb-target-group"
  port        = 80
  protocol    = "HTTP"
  target_type = "instance"
  vpc_id      = var.vpc_id
  health_check {
    port    = "80"
    matcher = "200-299"
  }
}

resource "aws_lb" "lb" {
  name               = "lb"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.sg_main.id, aws_security_group.sg_http.id]
  subnets            = [var.subnet_a, var.subnet_b, var.subnet_c]

  tags = {
    Environment = "lb"
  }
}

resource "aws_lb_listener" "backend" {
  load_balancer_arn = aws_lb.lb.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.lb_target_group.arn
  }
}

resource "aws_autoscaling_attachment" "asg_attachment_bar" {
  autoscaling_group_name = aws_autoscaling_group.asg.id
  alb_target_group_arn   = aws_lb_target_group.lb_target_group.arn
}
