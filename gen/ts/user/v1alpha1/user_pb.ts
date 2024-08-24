// @generated by protoc-gen-es v2.0.0 with parameter "target=ts"
// @generated from file user/v1alpha1/user.proto (package user.v1alpha1, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file user/v1alpha1/user.proto.
 */
export const file_user_v1alpha1_user: GenFile = /*@__PURE__*/
  fileDesc("Chh1c2VyL3YxYWxwaGExL3VzZXIucHJvdG8SDXVzZXIudjFhbHBoYTEiHAoOR2V0VXNlclJlcXVlc3QSCgoCaWQYASABKAkiNAoPR2V0VXNlclJlc3BvbnNlEiEKBHVzZXIYASABKAsyEy51c2VyLnYxYWxwaGExLlVzZXIiSAoRQ3JlYXRlVXNlclJlcXVlc3QSIQoEdXNlchgBIAEoCzITLnVzZXIudjFhbHBoYTEuVXNlchIQCghwYXNzd29yZBgCIAEoCSI3ChJDcmVhdGVVc2VyUmVzcG9uc2USIQoEdXNlchgBIAEoCzITLnVzZXIudjFhbHBoYTEuVXNlciJCChFVcGRhdGVVc2VyUmVxdWVzdBIKCgJpZBgBIAEoCRIhCgR1c2VyGAIgASgLMhMudXNlci52MWFscGhhMS5Vc2VyIksKElVwZGF0ZVVzZXJSZXNwb25zZRISCgppc19zdWNjZXNzGAEgASgIEiEKBHVzZXIYAiABKAsyEy51c2VyLnYxYWxwaGExLlVzZXIiPgoYQ2hlY2tVc2VyUGFzc3dvcmRSZXF1ZXN0EhAKCHVzZXJuYW1lGAEgASgJEhAKCHBhc3N3b3JkGAIgASgJIlIKGUNoZWNrVXNlclBhc3N3b3JkUmVzcG9uc2USEgoKaXNfc3VjY2VzcxgBIAEoCBIhCgR1c2VyGAIgASgLMhMudXNlci52MWFscGhhMS5Vc2VyIh8KEURlbGV0ZVVzZXJSZXF1ZXN0EgoKAmlkGAEgASgJIigKEkRlbGV0ZVVzZXJSZXNwb25zZRISCgppc19zdWNjZXNzGAEgASgIInYKBFVzZXISCgoCaWQYASABKAkSEAoIdXNlcm5hbWUYAiABKAkSDgoGdGVsX25vGAMgASgJEhEKCWZpcnN0bmFtZRgEIAEoCRIQCghsYXN0bmFtZRgFIAEoCRINCgVlbWFpbBgGIAEoCRIMCgRyYW5rGAcgASgJMsMDCgxVc2Vyc1NlcnZpY2USSgoHR2V0VXNlchIdLnVzZXIudjFhbHBoYTEuR2V0VXNlclJlcXVlc3QaHi51c2VyLnYxYWxwaGExLkdldFVzZXJSZXNwb25zZSIAElMKCkNyZWF0ZVVzZXISIC51c2VyLnYxYWxwaGExLkNyZWF0ZVVzZXJSZXF1ZXN0GiEudXNlci52MWFscGhhMS5DcmVhdGVVc2VyUmVzcG9uc2UiABJTCgpVcGRhdGVVc2VyEiAudXNlci52MWFscGhhMS5VcGRhdGVVc2VyUmVxdWVzdBohLnVzZXIudjFhbHBoYTEuVXBkYXRlVXNlclJlc3BvbnNlIgASUwoKRGVsZXRlVXNlchIgLnVzZXIudjFhbHBoYTEuRGVsZXRlVXNlclJlcXVlc3QaIS51c2VyLnYxYWxwaGExLkRlbGV0ZVVzZXJSZXNwb25zZSIAEmgKEUNoZWNrVXNlclBhc3N3b3JkEicudXNlci52MWFscGhhMS5DaGVja1VzZXJQYXNzd29yZFJlcXVlc3QaKC51c2VyLnYxYWxwaGExLkNoZWNrVXNlclBhc3N3b3JkUmVzcG9uc2UiAEK8AQoRY29tLnVzZXIudjFhbHBoYTFCCVVzZXJQcm90b1ABWkdnaXRodWIuY29tL3BvcGJlcnJ5LWFpL3BvcGJlcnJ5LXByb3RvL2dlbi9nby91c2VyL3YxYWxwaGExO3VzZXJ2MWFscGhhMaICA1VYWKoCDVVzZXIuVjFhbHBoYTHKAg1Vc2VyXFYxYWxwaGEx4gIZVXNlclxWMWFscGhhMVxHUEJNZXRhZGF0YeoCDlVzZXI6OlYxYWxwaGExYgZwcm90bzM");

/**
 * @generated from message user.v1alpha1.GetUserRequest
 */
export type GetUserRequest = Message<"user.v1alpha1.GetUserRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message user.v1alpha1.GetUserRequest.
 * Use `create(GetUserRequestSchema)` to create a new message.
 */
export const GetUserRequestSchema: GenMessage<GetUserRequest> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 0);

/**
 * @generated from message user.v1alpha1.GetUserResponse
 */
export type GetUserResponse = Message<"user.v1alpha1.GetUserResponse"> & {
  /**
   * @generated from field: user.v1alpha1.User user = 1;
   */
  user?: User;
};

/**
 * Describes the message user.v1alpha1.GetUserResponse.
 * Use `create(GetUserResponseSchema)` to create a new message.
 */
export const GetUserResponseSchema: GenMessage<GetUserResponse> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 1);

/**
 * @generated from message user.v1alpha1.CreateUserRequest
 */
export type CreateUserRequest = Message<"user.v1alpha1.CreateUserRequest"> & {
  /**
   * @generated from field: user.v1alpha1.User user = 1;
   */
  user?: User;

  /**
   * @generated from field: string password = 2;
   */
  password: string;
};

/**
 * Describes the message user.v1alpha1.CreateUserRequest.
 * Use `create(CreateUserRequestSchema)` to create a new message.
 */
export const CreateUserRequestSchema: GenMessage<CreateUserRequest> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 2);

/**
 * @generated from message user.v1alpha1.CreateUserResponse
 */
export type CreateUserResponse = Message<"user.v1alpha1.CreateUserResponse"> & {
  /**
   * @generated from field: user.v1alpha1.User user = 1;
   */
  user?: User;
};

/**
 * Describes the message user.v1alpha1.CreateUserResponse.
 * Use `create(CreateUserResponseSchema)` to create a new message.
 */
export const CreateUserResponseSchema: GenMessage<CreateUserResponse> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 3);

/**
 * @generated from message user.v1alpha1.UpdateUserRequest
 */
export type UpdateUserRequest = Message<"user.v1alpha1.UpdateUserRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: user.v1alpha1.User user = 2;
   */
  user?: User;
};

/**
 * Describes the message user.v1alpha1.UpdateUserRequest.
 * Use `create(UpdateUserRequestSchema)` to create a new message.
 */
export const UpdateUserRequestSchema: GenMessage<UpdateUserRequest> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 4);

/**
 * @generated from message user.v1alpha1.UpdateUserResponse
 */
export type UpdateUserResponse = Message<"user.v1alpha1.UpdateUserResponse"> & {
  /**
   * @generated from field: bool is_success = 1;
   */
  isSuccess: boolean;

  /**
   * @generated from field: user.v1alpha1.User user = 2;
   */
  user?: User;
};

/**
 * Describes the message user.v1alpha1.UpdateUserResponse.
 * Use `create(UpdateUserResponseSchema)` to create a new message.
 */
export const UpdateUserResponseSchema: GenMessage<UpdateUserResponse> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 5);

/**
 * @generated from message user.v1alpha1.CheckUserPasswordRequest
 */
export type CheckUserPasswordRequest = Message<"user.v1alpha1.CheckUserPasswordRequest"> & {
  /**
   * @generated from field: string username = 1;
   */
  username: string;

  /**
   * @generated from field: string password = 2;
   */
  password: string;
};

/**
 * Describes the message user.v1alpha1.CheckUserPasswordRequest.
 * Use `create(CheckUserPasswordRequestSchema)` to create a new message.
 */
export const CheckUserPasswordRequestSchema: GenMessage<CheckUserPasswordRequest> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 6);

/**
 * @generated from message user.v1alpha1.CheckUserPasswordResponse
 */
export type CheckUserPasswordResponse = Message<"user.v1alpha1.CheckUserPasswordResponse"> & {
  /**
   * @generated from field: bool is_success = 1;
   */
  isSuccess: boolean;

  /**
   * @generated from field: user.v1alpha1.User user = 2;
   */
  user?: User;
};

/**
 * Describes the message user.v1alpha1.CheckUserPasswordResponse.
 * Use `create(CheckUserPasswordResponseSchema)` to create a new message.
 */
export const CheckUserPasswordResponseSchema: GenMessage<CheckUserPasswordResponse> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 7);

/**
 * @generated from message user.v1alpha1.DeleteUserRequest
 */
export type DeleteUserRequest = Message<"user.v1alpha1.DeleteUserRequest"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message user.v1alpha1.DeleteUserRequest.
 * Use `create(DeleteUserRequestSchema)` to create a new message.
 */
export const DeleteUserRequestSchema: GenMessage<DeleteUserRequest> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 8);

/**
 * @generated from message user.v1alpha1.DeleteUserResponse
 */
export type DeleteUserResponse = Message<"user.v1alpha1.DeleteUserResponse"> & {
  /**
   * @generated from field: bool is_success = 1;
   */
  isSuccess: boolean;
};

/**
 * Describes the message user.v1alpha1.DeleteUserResponse.
 * Use `create(DeleteUserResponseSchema)` to create a new message.
 */
export const DeleteUserResponseSchema: GenMessage<DeleteUserResponse> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 9);

/**
 * @generated from message user.v1alpha1.User
 */
export type User = Message<"user.v1alpha1.User"> & {
  /**
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * @generated from field: string username = 2;
   */
  username: string;

  /**
   * @generated from field: string tel_no = 3;
   */
  telNo: string;

  /**
   * @generated from field: string firstname = 4;
   */
  firstname: string;

  /**
   * @generated from field: string lastname = 5;
   */
  lastname: string;

  /**
   * @generated from field: string email = 6;
   */
  email: string;

  /**
   * @generated from field: string rank = 7;
   */
  rank: string;
};

/**
 * Describes the message user.v1alpha1.User.
 * Use `create(UserSchema)` to create a new message.
 */
export const UserSchema: GenMessage<User> = /*@__PURE__*/
  messageDesc(file_user_v1alpha1_user, 10);

/**
 * @generated from service user.v1alpha1.UsersService
 */
export const UsersService: GenService<{
  /**
   * @generated from rpc user.v1alpha1.UsersService.GetUser
   */
  getUser: {
    methodKind: "unary";
    input: typeof GetUserRequestSchema;
    output: typeof GetUserResponseSchema;
  },
  /**
   * @generated from rpc user.v1alpha1.UsersService.CreateUser
   */
  createUser: {
    methodKind: "unary";
    input: typeof CreateUserRequestSchema;
    output: typeof CreateUserResponseSchema;
  },
  /**
   * @generated from rpc user.v1alpha1.UsersService.UpdateUser
   */
  updateUser: {
    methodKind: "unary";
    input: typeof UpdateUserRequestSchema;
    output: typeof UpdateUserResponseSchema;
  },
  /**
   * @generated from rpc user.v1alpha1.UsersService.DeleteUser
   */
  deleteUser: {
    methodKind: "unary";
    input: typeof DeleteUserRequestSchema;
    output: typeof DeleteUserResponseSchema;
  },
  /**
   * @generated from rpc user.v1alpha1.UsersService.CheckUserPassword
   */
  checkUserPassword: {
    methodKind: "unary";
    input: typeof CheckUserPasswordRequestSchema;
    output: typeof CheckUserPasswordResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_user_v1alpha1_user, 0);

