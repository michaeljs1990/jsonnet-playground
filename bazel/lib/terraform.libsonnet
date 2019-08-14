// Library for interacting with terraform. Every resource should be one item of an array which does
// two things for us. Ensures that we don't have any duplicate items and also makes sure that you
// are explicit about what values you are overwritting.
//
//
// tf.Config([
//   tf.Resource('aws_something', 'lol') {
//     bucket: 'lol',
//   } { extra: 'lol' },
//   tf.Resource('aws_something3', 'lol') {
//     bucket: 'lol',
//   },
//   tf.Output('aws_something', 'lol2') {
//     new_var: 'lol',
//   },
// ])

{
  local terraform = self,

  Object(kind, type, name):: {
    kindMeta:: kind,
    typeMeta:: type,
    nameMeta:: name,

    objRef()::
      if self.kindMeta == 'data' then
        std.join('.', [self.kindMeta, self.typeMeta, self.nameMeta])
      else
        std.join('.', [self.typeMeta, self.nameMeta]),

    getAttr(attr)::
      std.join('.', [self.objRef(), attr]),

    getVariable(attr)::
      '${%s}' % [self.getAttr(attr)],
  },

  Output(type, name):: terraform.Object('output', type, name),

  Resource(type, name):: terraform.Object('resource', type, name),

  DataSource(type, name):: terraform.Object('data', type, name),

  // This looks kinda nasty and to be honest it is however this makes everything else super easy.
  // This takes all of your terraform inputs in the form of an array and generates a valid output
  // as well as making sure an error is thrown if anyone defines two resources that are the same.
  Config(array):: {
    [obj.kindMeta]+: {
      [obj.typeMeta]+: {
        [obj.nameMeta]: obj
        for obj in std.filter(function(o) o.kindMeta == obj.kindMeta && o.typeMeta == obj.typeMeta, array)
      }
      for obj in std.set(std.filter(function(o) o.kindMeta == obj.kindMeta, array), function(o) o.typeMeta)
    }
    for obj in std.set(array, function(o) o.kindMeta)
  },
}
