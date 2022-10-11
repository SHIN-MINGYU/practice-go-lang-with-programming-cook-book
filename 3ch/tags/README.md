# Go's structure tag and basic reflection

- what is reflection?

  whatever you dont know the type of object, reflection let you access to the object

- reflection's useful apply way is to use struct tag

  struct tag is just a string key-value pair

- reflect package is used complicated process more than basic interface conversion

- this example serialize to parsable format all string fields after create to serialize format from struct value

## working flow

1.  if field is string, the field is serialized or deserialized

2.  if not string, ignore

3.  if field's tag have serialize key, that key return serialization/deserialization environment

4.  not allow overlap

5.  if structur tag is not setting, use field.Name

6.  structure tag's value is "-", ignore whatever the field is string
