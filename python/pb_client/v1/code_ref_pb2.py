#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import sys


def main(argv):
    pass


if __name__ == '__main__':
    main(sys.argv)
# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: v1/code_ref.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from v1 import base_pb2 as v1_dot_base__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='v1/code_ref.proto',
  package='v1',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x11v1/code_ref.proto\x12\x02v1\x1a\rv1/base.proto\"\x80\x01\n\rCodeReference\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04uuid\x18\x02 \x01(\t\x12\x0e\n\x06\x63ommit\x18\x03 \x01(\t\x12\x12\n\nupdated_at\x18\x04 \x01(\t\x12\x0e\n\x06status\x18\x05 \x01(\t\x12\x0f\n\x07git_url\x18\x06 \x01(\t\x12\x10\n\x08is_dirty\x18\x07 \x01(\x08\"n\n\x18\x43odeReferenceBodyRequest\x12(\n\x06\x65ntity\x18\x01 \x01(\x0b\x32\x18.v1.OwnedEntityIdRequest\x12(\n\rCodeReference\x18\x02 \x01(\x0b\x32\x11.v1.CodeReferenceb\x06proto3')
  ,
  dependencies=[v1_dot_base__pb2.DESCRIPTOR,])




_CODEREFERENCE = _descriptor.Descriptor(
  name='CodeReference',
  full_name='v1.CodeReference',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='v1.CodeReference.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uuid', full_name='v1.CodeReference.uuid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='commit', full_name='v1.CodeReference.commit', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='v1.CodeReference.updated_at', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='v1.CodeReference.status', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='git_url', full_name='v1.CodeReference.git_url', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_dirty', full_name='v1.CodeReference.is_dirty', index=6,
      number=7, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=41,
  serialized_end=169,
)


_CODEREFERENCEBODYREQUEST = _descriptor.Descriptor(
  name='CodeReferenceBodyRequest',
  full_name='v1.CodeReferenceBodyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='entity', full_name='v1.CodeReferenceBodyRequest.entity', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='CodeReference', full_name='v1.CodeReferenceBodyRequest.CodeReference', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
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
  serialized_start=171,
  serialized_end=281,
)

_CODEREFERENCEBODYREQUEST.fields_by_name['entity'].message_type = v1_dot_base__pb2._OWNEDENTITYIDREQUEST
_CODEREFERENCEBODYREQUEST.fields_by_name['CodeReference'].message_type = _CODEREFERENCE
DESCRIPTOR.message_types_by_name['CodeReference'] = _CODEREFERENCE
DESCRIPTOR.message_types_by_name['CodeReferenceBodyRequest'] = _CODEREFERENCEBODYREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CodeReference = _reflection.GeneratedProtocolMessageType('CodeReference', (_message.Message,), {
  'DESCRIPTOR' : _CODEREFERENCE,
  '__module__' : 'v1.code_ref_pb2'
  # @@protoc_insertion_point(class_scope:v1.CodeReference)
  })
_sym_db.RegisterMessage(CodeReference)

CodeReferenceBodyRequest = _reflection.GeneratedProtocolMessageType('CodeReferenceBodyRequest', (_message.Message,), {
  'DESCRIPTOR' : _CODEREFERENCEBODYREQUEST,
  '__module__' : 'v1.code_ref_pb2'
  # @@protoc_insertion_point(class_scope:v1.CodeReferenceBodyRequest)
  })
_sym_db.RegisterMessage(CodeReferenceBodyRequest)


# @@protoc_insertion_point(module_scope)