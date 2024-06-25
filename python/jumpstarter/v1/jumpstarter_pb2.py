# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: jumpstarter/v1/jumpstarter.proto
# Protobuf Python Version: 5.27.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    1,
    '',
    'jumpstarter/v1/jumpstarter.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n jumpstarter/v1/jumpstarter.proto\x12\x0ejumpstarter.v1\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\"\xe8\x01\n\x0fRegisterRequest\x12\x12\n\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x43\n\x06labels\x18\x02 \x03(\x0b\x32+.jumpstarter.v1.RegisterRequest.LabelsEntryR\x06labels\x12\x41\n\rdevice_report\x18\x03 \x03(\x0b\x32\x1c.jumpstarter.v1.DeviceReportR\x0c\x64\x65viceReport\x1a\x39\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"\xa1\x02\n\x0c\x44\x65viceReport\x12\x1f\n\x0b\x64\x65vice_uuid\x18\x01 \x01(\tR\ndeviceUuid\x12\x31\n\x12parent_device_uuid\x18\x02 \x01(\tH\x00R\x10parentDeviceUuid\x88\x01\x01\x12)\n\x10\x64river_interface\x18\x03 \x01(\tR\x0f\x64riverInterface\x12@\n\x06labels\x18\x04 \x03(\x0b\x32(.jumpstarter.v1.DeviceReport.LabelsEntryR\x06labels\x1a\x39\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x42\x15\n\x13_parent_device_uuid\"\x12\n\x10RegisterResponse\"8\n\nByeRequest\x12\x12\n\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x16\n\x06reason\x18\x02 \x01(\tR\x06reason\"\x0f\n\rListenRequest\"\\\n\x0eListenResponse\x12\'\n\x0frouter_endpoint\x18\x01 \x01(\tR\x0erouterEndpoint\x12!\n\x0crouter_token\x18\x02 \x01(\tR\x0brouterToken\"!\n\x0b\x44ialRequest\x12\x12\n\x04uuid\x18\x01 \x01(\tR\x04uuid\"Z\n\x0c\x44ialResponse\x12\'\n\x0frouter_endpoint\x18\x01 \x01(\tR\x0erouterEndpoint\x12!\n\x0crouter_token\x18\x02 \x01(\tR\x0brouterToken\"\x7f\n\x12\x41uditStreamRequest\x12\x12\n\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x1f\n\x0b\x64\x65vice_uuid\x18\x02 \x01(\tR\ndeviceUuid\x12\x1a\n\x08severity\x18\x03 \x01(\tR\x08severity\x12\x18\n\x07message\x18\x04 \x01(\tR\x07message\"\xec\x01\n\x11GetReportResponse\x12\x12\n\x04uuid\x18\x01 \x01(\tR\x04uuid\x12\x45\n\x06labels\x18\x02 \x03(\x0b\x32-.jumpstarter.v1.GetReportResponse.LabelsEntryR\x06labels\x12\x41\n\rdevice_report\x18\x03 \x03(\x0b\x32\x1c.jumpstarter.v1.DeviceReportR\x0c\x64\x65viceReport\x1a\x39\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"\x85\x01\n\x11\x44riverCallRequest\x12\x1f\n\x0b\x64\x65vice_uuid\x18\x01 \x01(\tR\ndeviceUuid\x12#\n\rdriver_method\x18\x02 \x01(\tR\x0c\x64riverMethod\x12*\n\x04\x61rgs\x18\x03 \x03(\x0b\x32\x16.google.protobuf.ValueR\x04\x61rgs\"R\n\x12\x44riverCallResponse\x12\x1b\n\tcall_uuid\x18\x01 \x01(\tR\x08\x63\x61llUuid\x12\x1f\n\x0bjson_result\x18\x02 \x01(\tR\njsonResult\"\x8e\x01\n\x1aStreamingDriverCallRequest\x12\x1f\n\x0b\x64\x65vice_uuid\x18\x01 \x01(\tR\ndeviceUuid\x12#\n\rdriver_method\x18\x02 \x01(\tR\x0c\x64riverMethod\x12*\n\x04\x61rgs\x18\x03 \x03(\x0b\x32\x16.google.protobuf.ValueR\x04\x61rgs\"[\n\x1bStreamingDriverCallResponse\x12\x1b\n\tcall_uuid\x18\x01 \x01(\tR\x08\x63\x61llUuid\x12\x1f\n\x0bjson_result\x18\x02 \x01(\tR\njsonResult\"j\n\x11LogStreamResponse\x12\x1f\n\x0b\x64\x65vice_uuid\x18\x01 \x01(\tR\ndeviceUuid\x12\x1a\n\x08severity\x18\x02 \x01(\tR\x08severity\x12\x18\n\x07message\x18\x03 \x01(\tR\x07message2\xf8\x02\n\x11\x43ontrollerService\x12M\n\x08Register\x12\x1f.jumpstarter.v1.RegisterRequest\x1a .jumpstarter.v1.RegisterResponse\x12\x39\n\x03\x42ye\x12\x1a.jumpstarter.v1.ByeRequest\x1a\x16.google.protobuf.Empty\x12I\n\x06Listen\x12\x1d.jumpstarter.v1.ListenRequest\x1a\x1e.jumpstarter.v1.ListenResponse0\x01\x12\x41\n\x04\x44ial\x12\x1b.jumpstarter.v1.DialRequest\x1a\x1c.jumpstarter.v1.DialResponse\x12K\n\x0b\x41uditStream\x12\".jumpstarter.v1.AuditStreamRequest\x1a\x16.google.protobuf.Empty(\x01\x32\xea\x02\n\x0f\x45xporterService\x12\x46\n\tGetReport\x12\x16.google.protobuf.Empty\x1a!.jumpstarter.v1.GetReportResponse\x12S\n\nDriverCall\x12!.jumpstarter.v1.DriverCallRequest\x1a\".jumpstarter.v1.DriverCallResponse\x12p\n\x13StreamingDriverCall\x12*.jumpstarter.v1.StreamingDriverCallRequest\x1a+.jumpstarter.v1.StreamingDriverCallResponse0\x01\x12H\n\tLogStream\x12\x16.google.protobuf.Empty\x1a!.jumpstarter.v1.LogStreamResponse0\x01\x42\xcd\x01\n\x12\x63om.jumpstarter.v1B\x10JumpstarterProtoP\x01ZLgithub.com/jumpstarter-dev/jumpstarter-protocol/jumpstarter/v1;jumpstarterv1\xa2\x02\x03JXX\xaa\x02\x0eJumpstarter.V1\xca\x02\x0eJumpstarter\\V1\xe2\x02\x1aJumpstarter\\V1\\GPBMetadata\xea\x02\x0fJumpstarter::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'jumpstarter.v1.jumpstarter_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\022com.jumpstarter.v1B\020JumpstarterProtoP\001ZLgithub.com/jumpstarter-dev/jumpstarter-protocol/jumpstarter/v1;jumpstarterv1\242\002\003JXX\252\002\016Jumpstarter.V1\312\002\016Jumpstarter\\V1\342\002\032Jumpstarter\\V1\\GPBMetadata\352\002\017Jumpstarter::V1'
  _globals['_REGISTERREQUEST_LABELSENTRY']._loaded_options = None
  _globals['_REGISTERREQUEST_LABELSENTRY']._serialized_options = b'8\001'
  _globals['_DEVICEREPORT_LABELSENTRY']._loaded_options = None
  _globals['_DEVICEREPORT_LABELSENTRY']._serialized_options = b'8\001'
  _globals['_GETREPORTRESPONSE_LABELSENTRY']._loaded_options = None
  _globals['_GETREPORTRESPONSE_LABELSENTRY']._serialized_options = b'8\001'
  _globals['_REGISTERREQUEST']._serialized_start=112
  _globals['_REGISTERREQUEST']._serialized_end=344
  _globals['_REGISTERREQUEST_LABELSENTRY']._serialized_start=287
  _globals['_REGISTERREQUEST_LABELSENTRY']._serialized_end=344
  _globals['_DEVICEREPORT']._serialized_start=347
  _globals['_DEVICEREPORT']._serialized_end=636
  _globals['_DEVICEREPORT_LABELSENTRY']._serialized_start=287
  _globals['_DEVICEREPORT_LABELSENTRY']._serialized_end=344
  _globals['_REGISTERRESPONSE']._serialized_start=638
  _globals['_REGISTERRESPONSE']._serialized_end=656
  _globals['_BYEREQUEST']._serialized_start=658
  _globals['_BYEREQUEST']._serialized_end=714
  _globals['_LISTENREQUEST']._serialized_start=716
  _globals['_LISTENREQUEST']._serialized_end=731
  _globals['_LISTENRESPONSE']._serialized_start=733
  _globals['_LISTENRESPONSE']._serialized_end=825
  _globals['_DIALREQUEST']._serialized_start=827
  _globals['_DIALREQUEST']._serialized_end=860
  _globals['_DIALRESPONSE']._serialized_start=862
  _globals['_DIALRESPONSE']._serialized_end=952
  _globals['_AUDITSTREAMREQUEST']._serialized_start=954
  _globals['_AUDITSTREAMREQUEST']._serialized_end=1081
  _globals['_GETREPORTRESPONSE']._serialized_start=1084
  _globals['_GETREPORTRESPONSE']._serialized_end=1320
  _globals['_GETREPORTRESPONSE_LABELSENTRY']._serialized_start=287
  _globals['_GETREPORTRESPONSE_LABELSENTRY']._serialized_end=344
  _globals['_DRIVERCALLREQUEST']._serialized_start=1323
  _globals['_DRIVERCALLREQUEST']._serialized_end=1456
  _globals['_DRIVERCALLRESPONSE']._serialized_start=1458
  _globals['_DRIVERCALLRESPONSE']._serialized_end=1540
  _globals['_STREAMINGDRIVERCALLREQUEST']._serialized_start=1543
  _globals['_STREAMINGDRIVERCALLREQUEST']._serialized_end=1685
  _globals['_STREAMINGDRIVERCALLRESPONSE']._serialized_start=1687
  _globals['_STREAMINGDRIVERCALLRESPONSE']._serialized_end=1778
  _globals['_LOGSTREAMRESPONSE']._serialized_start=1780
  _globals['_LOGSTREAMRESPONSE']._serialized_end=1886
  _globals['_CONTROLLERSERVICE']._serialized_start=1889
  _globals['_CONTROLLERSERVICE']._serialized_end=2265
  _globals['_EXPORTERSERVICE']._serialized_start=2268
  _globals['_EXPORTERSERVICE']._serialized_end=2630
# @@protoc_insertion_point(module_scope)
