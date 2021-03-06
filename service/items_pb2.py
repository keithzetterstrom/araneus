# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: items.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='items.proto',
  package='protoItems',
  syntax='proto3',
  serialized_options=b'Z\014./protoItems',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0bitems.proto\x12\nprotoItems\"+\n\x08ItemGRPC\x12\x11\n\tSignature\x18\x01 \x03(\x04\x12\x0c\n\x04Text\x18\x02 \x01(\t2I\n\x0bItemService\x12:\n\x0cSetSignature\x12\x14.protoItems.ItemGRPC\x1a\x14.protoItems.ItemGRPCB\x0eZ\x0c./protoItemsb\x06proto3'
)




_ITEMGRPC = _descriptor.Descriptor(
  name='ItemGRPC',
  full_name='protoItems.ItemGRPC',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Signature', full_name='protoItems.ItemGRPC.Signature', index=0,
      number=1, type=4, cpp_type=4, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Text', full_name='protoItems.ItemGRPC.Text', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=27,
  serialized_end=70,
)

DESCRIPTOR.message_types_by_name['ItemGRPC'] = _ITEMGRPC
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ItemGRPC = _reflection.GeneratedProtocolMessageType('ItemGRPC', (_message.Message,), {
  'DESCRIPTOR' : _ITEMGRPC,
  '__module__' : 'items_pb2'
  # @@protoc_insertion_point(class_scope:protoItems.ItemGRPC)
  })
_sym_db.RegisterMessage(ItemGRPC)


DESCRIPTOR._options = None

_ITEMSERVICE = _descriptor.ServiceDescriptor(
  name='ItemService',
  full_name='protoItems.ItemService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=72,
  serialized_end=145,
  methods=[
  _descriptor.MethodDescriptor(
    name='SetSignature',
    full_name='protoItems.ItemService.SetSignature',
    index=0,
    containing_service=None,
    input_type=_ITEMGRPC,
    output_type=_ITEMGRPC,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_ITEMSERVICE)

DESCRIPTOR.services_by_name['ItemService'] = _ITEMSERVICE

# @@protoc_insertion_point(module_scope)
