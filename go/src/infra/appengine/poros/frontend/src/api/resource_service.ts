// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import { Empty } from './common/empty';
import { FieldMask } from './common/field_mask';
import { rpcClient } from './common/rpc_client';
import { fromJsonTimestamp, isSet } from './common/utils';

/** Performs operations on Resources. */
export interface IResourceService {
  /** Creates the given Resource. */
  create(request: CreateResourceRequest): Promise<ResourceModel>;
  /** Retrieves a Resource for a given unique value. */
  get(request: GetResourceRequest): Promise<ResourceModel>;
  /** Update a single resource in poros. */
  update(request: UpdateResourceRequest): Promise<ResourceModel>;
  /** Deletes the given Resource. */
  delete(request: DeleteResourceRequest): Promise<Empty>;
  /** Lists all Resources. */
  list(request: ListResourcesRequest): Promise<ListResourcesResponse>;
}

export class ResourceService implements IResourceService {
  constructor() {
    this.create = this.create.bind(this);
    this.get = this.get.bind(this);
    this.update = this.update.bind(this);
    this.delete = this.delete.bind(this);
    this.list = this.list.bind(this);
  }

  create = (request: CreateResourceRequest): Promise<ResourceModel> => {
    const data = CreateResourceRequest.toJSON(request);
    const promise = rpcClient.request('poros.Resource', 'Create', data);
    return promise.then((data) => ResourceModel.fromJSON(data));
  };

  get = (request: GetResourceRequest): Promise<ResourceModel> => {
    const data = GetResourceRequest.toJSON(request);
    const promise = rpcClient.request('poros.Resource', 'Get', data);
    return promise.then((data) => ResourceModel.fromJSON(data));
  };

  update = (request: UpdateResourceRequest): Promise<ResourceModel> => {
    const data = UpdateResourceRequest.toJSON(request);
    const promise = rpcClient.request('poros.Resource', 'Update', data);
    return promise.then((data) => ResourceModel.fromJSON(data));
  };

  delete = (request: DeleteResourceRequest): Promise<Empty> => {
    const data = DeleteResourceRequest.toJSON(request);
    const promise = rpcClient.request('poros.Resource', 'Delete', data);
    return promise.then((data) => Empty.fromJSON(data));
  };

  list = (request: ListResourcesRequest): Promise<ListResourcesResponse> => {
    const data = ListResourcesRequest.toJSON(request);
    const promise = rpcClient.request('poros.Resource', 'List', data);
    return promise.then((data) => ListResourcesResponse.fromJSON(data));
  };
}

export interface ResourceModel {
  /** Unique identifier of the resource */
  resourceId: string;
  /** Name of the resource */
  name: string;
  /** Type of the resource */
  type: string;
  /** Description of the resource */
  description: string;
  /** if machine is selected then type of machine */
  image: string;
  /** User who created the record */
  createdBy: string;
  /** Timestamp for the creation of the record */
  createdAt: Date | undefined;
  /** User who modified the record */
  modifiedBy: string;
  /** Timestamp for the last update of the record */
  modifiedAt: Date | undefined;
}

export const ResourceModel = {
  defaultEntity(): ResourceModel {
    return {
      resourceId: '',
      name: '',
      type: 'machine',
      description: '',
      image: '',
      createdBy: '',
      createdAt: undefined,
      modifiedBy: '',
      modifiedAt: undefined,
    };
  },
  fromJSON(object: any): ResourceModel {
    return {
      resourceId: isSet(object.resourceId) ? String(object.resourceId) : '',
      name: isSet(object.name) ? String(object.name) : '',
      type: isSet(object.type) ? String(object.type) : '',
      description: isSet(object.description) ? String(object.description) : '',
      image: isSet(object.image) ? String(object.image) : '',

      createdBy: isSet(object.createdBy) ? String(object.createdBy) : '',
      createdAt: isSet(object.createdAt)
        ? fromJsonTimestamp(object.createdAt)
        : undefined,
      modifiedBy: isSet(object.modifiedBy) ? String(object.modifiedBy) : '',
      modifiedAt: isSet(object.modifiedAt)
        ? fromJsonTimestamp(object.modifiedAt)
        : undefined,
    };
  },

  toJSON(message: ResourceModel): unknown {
    const obj: any = {};
    message.resourceId !== undefined && (obj.resourceId = message.resourceId);
    message.name !== undefined && (obj.name = message.name);
    message.type !== undefined && (obj.type = message.type);
    message.description !== undefined &&
      (obj.description = message.description);
    message.image !== undefined && (obj.image = message.image);
    message.createdBy !== undefined && (obj.createdBy = message.createdBy);
    message.createdAt !== undefined &&
      (obj.createdAt = message.createdAt.toISOString());
    message.modifiedBy !== undefined && (obj.modifiedBy = message.modifiedBy);
    message.modifiedAt !== undefined &&
      (obj.modifiedAt = message.modifiedAt.toISOString());
    return obj;
  },
};

/** Request to create a single resource in ResourceServ */
export interface CreateResourceRequest {
  /** Name of the resource */
  name: string;
  /** Type of the resource */
  type: string;
  /** Description of the resource */
  description: string;
  /** if machine is selected then type of machine */
  image: string;
}

export const CreateResourceRequest = {
  toJSON(message: CreateResourceRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.type !== undefined && (obj.type = message.type);
    message.description !== undefined &&
      (obj.description = message.description);
    message.image !== undefined && (obj.image = message.image);
    return obj;
  },
};

// Request to delete a single resource from poros.
export interface DeleteResourceRequest {
  /** Unique identifier of the resource */
  id: string;
}

export const DeleteResourceRequest = {
  toJSON(message: DeleteResourceRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.name = message.id);
    return obj;
  },
};

/** Gets a Resource. */
export interface GetResourceRequest {
  /**
   * The name of the resource to retrieve.
   * Format: publishers/{publisher}/resources/{resource}
   */
  id: string;
}

export const GetResourceRequest = {
  toJSON(message: GetResourceRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },
};

/** Request to list all resources in poros. */
export interface ListResourcesRequest {
  /** Fields to include on each result */
  readMask: string[] | undefined;
  /** Number of results per page */
  pageSize: number;
  /** Page token from a previous page's ListResourcesResponse */
  pageToken: string;
}

/** Response to ListResourcesRequest. */
export interface ListResourcesResponse {
  /** The result set. */
  resources: ResourceModel[];
  /**
   * A page token that can be passed into future requests to get the next page.
   * Empty if there is no next page.
   */
  nextPageToken: string;
}

export const ListResourcesRequest = {
  toJSON(message: ListResourcesRequest): unknown {
    const obj: any = {};
    message.readMask !== undefined &&
      (obj.readMask = FieldMask.toJSON(FieldMask.wrap(message.readMask)));
    message.pageSize !== undefined &&
      (obj.pageSize = Math.round(message.pageSize));
    message.pageToken !== undefined && (obj.pageToken = message.pageToken);
    return obj;
  },
};

export const ListResourcesResponse = {
  fromJSON(object: any): ListResourcesResponse {
    return {
      resources: Array.isArray(object?.resources)
        ? object.resources.map((e: any) => ResourceModel.fromJSON(e))
        : [],
      nextPageToken: isSet(object.nextPageToken)
        ? String(object.nextPageToken)
        : '',
    };
  },
};

/** Request to update a single resource in poros. */
export interface UpdateResourceRequest {
  /** The existing resource to update in the database. */
  resource: ResourceModel | undefined;
  /** The list of fields to update. */
  updateMask: string[] | undefined;
}

export const UpdateResourceRequest = {
  toJSON(message: UpdateResourceRequest): unknown {
    const obj: any = {};
    message.resource !== undefined &&
      (obj.resource = message.resource
        ? ResourceModel.toJSON(message.resource)
        : undefined);
    message.updateMask !== undefined &&
      (obj.updateMask = FieldMask.toJSON(FieldMask.wrap(message.updateMask)));
    return obj;
  },
};
