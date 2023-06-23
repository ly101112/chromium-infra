# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/api_proto/features.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from api.api_proto import common_pb2 as api_dot_api__proto_dot_common__pb2
from api.api_proto import features_objects_pb2 as api_dot_api__proto_dot_features__objects__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x61pi/api_proto/features.proto\x12\x08monorail\x1a\x1a\x61pi/api_proto/common.proto\x1a$api/api_proto/features_objects.proto\"<\n\x19ListHotlistsByUserRequest\x12\x1f\n\x04user\x18\x02 \x01(\x0b\x32\x11.monorail.UserRef\"A\n\x1aListHotlistsByUserResponse\x12#\n\x08hotlists\x18\x01 \x03(\x0b\x32\x11.monorail.Hotlist\"?\n\x1aListHotlistsByIssueRequest\x12!\n\x05issue\x18\x02 \x01(\x0b\x32\x12.monorail.IssueRef\"B\n\x1bListHotlistsByIssueResponse\x12#\n\x08hotlists\x18\x01 \x03(\x0b\x32\x11.monorail.Hotlist\"$\n\"ListRecentlyVisitedHotlistsRequest\"J\n#ListRecentlyVisitedHotlistsResponse\x12#\n\x08hotlists\x18\x01 \x03(\x0b\x32\x11.monorail.Hotlist\"\x1c\n\x1aListStarredHotlistsRequest\"B\n\x1bListStarredHotlistsResponse\x12#\n\x08hotlists\x18\x01 \x03(\x0b\x32\x11.monorail.Hotlist\"G\n\x1aGetHotlistStarCountRequest\x12)\n\x0bhotlist_ref\x18\x02 \x01(\x0b\x32\x14.monorail.HotlistRef\"1\n\x1bGetHotlistStarCountResponse\x12\x12\n\nstar_count\x18\x01 \x01(\r\"P\n\x12StarHotlistRequest\x12)\n\x0bhotlist_ref\x18\x02 \x01(\x0b\x32\x14.monorail.HotlistRef\x12\x0f\n\x07starred\x18\x03 \x01(\x08\")\n\x13StarHotlistResponse\x12\x12\n\nstar_count\x18\x01 \x01(\r\">\n\x11GetHotlistRequest\x12)\n\x0bhotlist_ref\x18\x01 \x01(\x0b\x32\x14.monorail.HotlistRef\"8\n\x12GetHotlistResponse\x12\"\n\x07hotlist\x18\x01 \x01(\x0b\x32\x11.monorail.Hotlist\"\xa5\x01\n\x17ListHotlistItemsRequest\x12)\n\x0bhotlist_ref\x18\x02 \x01(\x0b\x32\x14.monorail.HotlistRef\x12(\n\npagination\x18\x03 \x01(\x0b\x32\x14.monorail.Pagination\x12\x0b\n\x03\x63\x61n\x18\x04 \x01(\r\x12\x11\n\tsort_spec\x18\x05 \x01(\t\x12\x15\n\rgroup_by_spec\x18\x06 \x01(\t\"@\n\x18ListHotlistItemsResponse\x12$\n\x05items\x18\x01 \x03(\x0b\x32\x15.monorail.HotlistItem\"\xae\x01\n\x14\x43reateHotlistRequest\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0f\n\x07summary\x18\x03 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x01(\t\x12&\n\x0b\x65\x64itor_refs\x18\x05 \x03(\x0b\x32\x11.monorail.UserRef\x12&\n\nissue_refs\x18\x06 \x03(\x0b\x32\x12.monorail.IssueRef\x12\x12\n\nis_private\x18\x07 \x01(\x08\"\x17\n\x15\x43reateHotlistResponse\"\'\n\x17\x43heckHotlistNameRequest\x12\x0c\n\x04name\x18\x02 \x01(\t\")\n\x18\x43heckHotlistNameResponse\x12\r\n\x05\x65rror\x18\x01 \x01(\t\"u\n\x1fRemoveIssuesFromHotlistsRequest\x12*\n\x0chotlist_refs\x18\x02 \x03(\x0b\x32\x14.monorail.HotlistRef\x12&\n\nissue_refs\x18\x03 \x03(\x0b\x32\x12.monorail.IssueRef\"\"\n RemoveIssuesFromHotlistsResponse\"~\n\x1a\x41\x64\x64IssuesToHotlistsRequest\x12*\n\x0chotlist_refs\x18\x02 \x03(\x0b\x32\x14.monorail.HotlistRef\x12&\n\nissue_refs\x18\x03 \x03(\x0b\x32\x12.monorail.IssueRef\x12\x0c\n\x04note\x18\x04 \x01(\t\"\x1d\n\x1b\x41\x64\x64IssuesToHotlistsResponse\"\xac\x01\n\x1aRerankHotlistIssuesRequest\x12)\n\x0bhotlist_ref\x18\x01 \x01(\x0b\x32\x14.monorail.HotlistRef\x12&\n\nmoved_refs\x18\x02 \x03(\x0b\x32\x12.monorail.IssueRef\x12&\n\ntarget_ref\x18\x03 \x01(\x0b\x32\x12.monorail.IssueRef\x12\x13\n\x0bsplit_above\x18\x04 \x01(\x08\"\x1d\n\x1bRerankHotlistIssuesResponse\"\x7f\n\x1dUpdateHotlistIssueNoteRequest\x12)\n\x0bhotlist_ref\x18\x02 \x01(\x0b\x32\x14.monorail.HotlistRef\x12%\n\tissue_ref\x18\x03 \x01(\x0b\x32\x12.monorail.IssueRef\x12\x0c\n\x04note\x18\x04 \x01(\t\" \n\x1eUpdateHotlistIssueNoteResponse\"A\n\x14\x44\x65leteHotlistRequest\x12)\n\x0bhotlist_ref\x18\x01 \x01(\x0b\x32\x14.monorail.HotlistRef\"\x17\n\x15\x44\x65leteHotlistResponse2\xc8\x0b\n\x08\x46\x65\x61tures\x12\x61\n\x12ListHotlistsByUser\x12#.monorail.ListHotlistsByUserRequest\x1a$.monorail.ListHotlistsByUserResponse\"\x00\x12\x64\n\x13ListHotlistsByIssue\x12$.monorail.ListHotlistsByIssueRequest\x1a%.monorail.ListHotlistsByIssueResponse\"\x00\x12|\n\x1bListRecentlyVisitedHotlists\x12,.monorail.ListRecentlyVisitedHotlistsRequest\x1a-.monorail.ListRecentlyVisitedHotlistsResponse\"\x00\x12\x64\n\x13ListStarredHotlists\x12$.monorail.ListStarredHotlistsRequest\x1a%.monorail.ListStarredHotlistsResponse\"\x00\x12\x64\n\x13GetHotlistStarCount\x12$.monorail.GetHotlistStarCountRequest\x1a%.monorail.GetHotlistStarCountResponse\"\x00\x12L\n\x0bStarHotlist\x12\x1c.monorail.StarHotlistRequest\x1a\x1d.monorail.StarHotlistResponse\"\x00\x12I\n\nGetHotlist\x12\x1b.monorail.GetHotlistRequest\x1a\x1c.monorail.GetHotlistResponse\"\x00\x12[\n\x10ListHotlistItems\x12!.monorail.ListHotlistItemsRequest\x1a\".monorail.ListHotlistItemsResponse\"\x00\x12R\n\rCreateHotlist\x12\x1e.monorail.CreateHotlistRequest\x1a\x1f.monorail.CreateHotlistResponse\"\x00\x12[\n\x10\x43heckHotlistName\x12!.monorail.CheckHotlistNameRequest\x1a\".monorail.CheckHotlistNameResponse\"\x00\x12s\n\x18RemoveIssuesFromHotlists\x12).monorail.RemoveIssuesFromHotlistsRequest\x1a*.monorail.RemoveIssuesFromHotlistsResponse\"\x00\x12\x64\n\x13\x41\x64\x64IssuesToHotlists\x12$.monorail.AddIssuesToHotlistsRequest\x1a%.monorail.AddIssuesToHotlistsResponse\"\x00\x12\x64\n\x13RerankHotlistIssues\x12$.monorail.RerankHotlistIssuesRequest\x1a%.monorail.RerankHotlistIssuesResponse\"\x00\x12m\n\x16UpdateHotlistIssueNote\x12\'.monorail.UpdateHotlistIssueNoteRequest\x1a(.monorail.UpdateHotlistIssueNoteResponse\"\x00\x12R\n\rDeleteHotlist\x12\x1e.monorail.DeleteHotlistRequest\x1a\x1f.monorail.DeleteHotlistResponse\"\x00\x42)Z\'infra/monorailv2/api/api_proto;monorailb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.api_proto.features_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\'infra/monorailv2/api/api_proto;monorail'
  _LISTHOTLISTSBYUSERREQUEST._serialized_start=108
  _LISTHOTLISTSBYUSERREQUEST._serialized_end=168
  _LISTHOTLISTSBYUSERRESPONSE._serialized_start=170
  _LISTHOTLISTSBYUSERRESPONSE._serialized_end=235
  _LISTHOTLISTSBYISSUEREQUEST._serialized_start=237
  _LISTHOTLISTSBYISSUEREQUEST._serialized_end=300
  _LISTHOTLISTSBYISSUERESPONSE._serialized_start=302
  _LISTHOTLISTSBYISSUERESPONSE._serialized_end=368
  _LISTRECENTLYVISITEDHOTLISTSREQUEST._serialized_start=370
  _LISTRECENTLYVISITEDHOTLISTSREQUEST._serialized_end=406
  _LISTRECENTLYVISITEDHOTLISTSRESPONSE._serialized_start=408
  _LISTRECENTLYVISITEDHOTLISTSRESPONSE._serialized_end=482
  _LISTSTARREDHOTLISTSREQUEST._serialized_start=484
  _LISTSTARREDHOTLISTSREQUEST._serialized_end=512
  _LISTSTARREDHOTLISTSRESPONSE._serialized_start=514
  _LISTSTARREDHOTLISTSRESPONSE._serialized_end=580
  _GETHOTLISTSTARCOUNTREQUEST._serialized_start=582
  _GETHOTLISTSTARCOUNTREQUEST._serialized_end=653
  _GETHOTLISTSTARCOUNTRESPONSE._serialized_start=655
  _GETHOTLISTSTARCOUNTRESPONSE._serialized_end=704
  _STARHOTLISTREQUEST._serialized_start=706
  _STARHOTLISTREQUEST._serialized_end=786
  _STARHOTLISTRESPONSE._serialized_start=788
  _STARHOTLISTRESPONSE._serialized_end=829
  _GETHOTLISTREQUEST._serialized_start=831
  _GETHOTLISTREQUEST._serialized_end=893
  _GETHOTLISTRESPONSE._serialized_start=895
  _GETHOTLISTRESPONSE._serialized_end=951
  _LISTHOTLISTITEMSREQUEST._serialized_start=954
  _LISTHOTLISTITEMSREQUEST._serialized_end=1119
  _LISTHOTLISTITEMSRESPONSE._serialized_start=1121
  _LISTHOTLISTITEMSRESPONSE._serialized_end=1185
  _CREATEHOTLISTREQUEST._serialized_start=1188
  _CREATEHOTLISTREQUEST._serialized_end=1362
  _CREATEHOTLISTRESPONSE._serialized_start=1364
  _CREATEHOTLISTRESPONSE._serialized_end=1387
  _CHECKHOTLISTNAMEREQUEST._serialized_start=1389
  _CHECKHOTLISTNAMEREQUEST._serialized_end=1428
  _CHECKHOTLISTNAMERESPONSE._serialized_start=1430
  _CHECKHOTLISTNAMERESPONSE._serialized_end=1471
  _REMOVEISSUESFROMHOTLISTSREQUEST._serialized_start=1473
  _REMOVEISSUESFROMHOTLISTSREQUEST._serialized_end=1590
  _REMOVEISSUESFROMHOTLISTSRESPONSE._serialized_start=1592
  _REMOVEISSUESFROMHOTLISTSRESPONSE._serialized_end=1626
  _ADDISSUESTOHOTLISTSREQUEST._serialized_start=1628
  _ADDISSUESTOHOTLISTSREQUEST._serialized_end=1754
  _ADDISSUESTOHOTLISTSRESPONSE._serialized_start=1756
  _ADDISSUESTOHOTLISTSRESPONSE._serialized_end=1785
  _RERANKHOTLISTISSUESREQUEST._serialized_start=1788
  _RERANKHOTLISTISSUESREQUEST._serialized_end=1960
  _RERANKHOTLISTISSUESRESPONSE._serialized_start=1962
  _RERANKHOTLISTISSUESRESPONSE._serialized_end=1991
  _UPDATEHOTLISTISSUENOTEREQUEST._serialized_start=1993
  _UPDATEHOTLISTISSUENOTEREQUEST._serialized_end=2120
  _UPDATEHOTLISTISSUENOTERESPONSE._serialized_start=2122
  _UPDATEHOTLISTISSUENOTERESPONSE._serialized_end=2154
  _DELETEHOTLISTREQUEST._serialized_start=2156
  _DELETEHOTLISTREQUEST._serialized_end=2221
  _DELETEHOTLISTRESPONSE._serialized_start=2223
  _DELETEHOTLISTRESPONSE._serialized_end=2246
  _FEATURES._serialized_start=2249
  _FEATURES._serialized_end=3729
# @@protoc_insertion_point(module_scope)
