# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nest one.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0enest one.proto\"\xd9\x01\n\x08\x42uilding\x12\x14\n\x0c\x62uildingName\x18\x01 \x01(\t\x12\x16\n\x0e\x62uildingNumber\x18\x02 \x01(\x05\x12 \n\x06street\x18\x03 \x01(\x0b\x32\x10.Building.Street\x1a}\n\x06Street\x12\x12\n\nstreetName\x18\x01 \x01(\t\x12#\n\x04\x63ity\x18\x02 \x01(\x0b\x32\x15.Building.Street.City\x1a:\n\x04\x43ity\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0f\n\x07zipCode\x18\x02 \x01(\x05\x12\x13\n\x0b\x63ountryName\x18\x03 \x01(\tb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nest one_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _BUILDING._serialized_start=19
  _BUILDING._serialized_end=236
  _BUILDING_STREET._serialized_start=111
  _BUILDING_STREET._serialized_end=236
  _BUILDING_STREET_CITY._serialized_start=178
  _BUILDING_STREET_CITY._serialized_end=236
# @@protoc_insertion_point(module_scope)