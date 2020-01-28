provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region = "${var.aws_region}"
}

resource "aws_security_group" "site_prod" {
  name = "trackhours_prod"
  description = "Site Security Group"

  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = -1
    to_port = -1
    protocol = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 443
    to_port = 443
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = -1
    to_port = -1
    protocol = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags {
    Name = "Site"
  }
}

resource "aws_route53_record" "site_prod" {
  zone_id = "${var.dns_zone_id}"
  name = "${var.dns_site_domain}"
  type = "A"
  ttl = "300"
  records = ["${aws_eip.site_prod.public_ip}"]
}

resource "aws_eip" "site_prod" {
  instance = "${aws_instance.site_prod.id}"
}

resource "aws_instance" "site_prod" {
  ami = "${var.instance_ami}"
  instance_type = "${var.instance_type_site}"
  key_name = "${var.instance_key_name}"
  security_groups = ["${aws_security_group.site_prod.name}"]

  # Static site data
  provisioner "file" {
    source = "volume/"
    destination = "/tmp"

    connection {
      type = "ssh"
      user = "${var.ssh_user}"
      private_key = "${file(var.ssh_private_key)}"
    }
  }

  # Installs initial necessary programs
  provisioner "remote-exec" {
    inline = "${var.script_setup_site}"

    connection {
      type = "ssh"
      user = "${var.ssh_user}"
      private_key = "${file(var.ssh_private_key)}"
    }
  }
}

