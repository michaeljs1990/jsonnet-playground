local tf = import 'lib/terraform.libsonnet';

tf.Config([
  tf.Resource('aws_something', 'lol') {
    bucket: 'lol',
  } { extra: 'lol' },
  tf.Resource('aws_something3', 'lol') {
    bucket: 'lol',
  },
])
