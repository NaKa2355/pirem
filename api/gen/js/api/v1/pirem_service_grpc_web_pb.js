/**
 * @fileoverview gRPC-Web generated client stub for pirem
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.12.4
// source: api/v1/pirem_service.proto


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var api_v1_device_pb = require('../../api/v1/device_pb.js')

var api_v1_irdata_pb = require('../../api/v1/irdata_pb.js')

var api_v1_remote_pb = require('../../api/v1/remote_pb.js')

var api_v1_button_pb = require('../../api/v1/button_pb.js')

var api_v1_empty_pb = require('../../api/v1/empty_pb.js')
const proto = {};
proto.pirem = require('./pirem_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pirem.PiRemServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.pirem.PiRemServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname.replace(/\/+$/, '');

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.SendIrRequest,
 *   !proto.pirem.SendIrResponse>}
 */
const methodDescriptor_PiRemService_SendIr = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/SendIr',
  grpc.web.MethodType.UNARY,
  proto.pirem.SendIrRequest,
  proto.pirem.SendIrResponse,
  /**
   * @param {!proto.pirem.SendIrRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pirem.SendIrResponse.deserializeBinary
);


/**
 * @param {!proto.pirem.SendIrRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.SendIrResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.SendIrResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.sendIr =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/SendIr',
      request,
      metadata || {},
      methodDescriptor_PiRemService_SendIr,
      callback);
};


/**
 * @param {!proto.pirem.SendIrRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.SendIrResponse>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.sendIr =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/SendIr',
      request,
      metadata || {},
      methodDescriptor_PiRemService_SendIr);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.ReceiveIrRequest,
 *   !proto.pirem.IrData>}
 */
const methodDescriptor_PiRemService_ReceiveIr = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/ReceiveIr',
  grpc.web.MethodType.UNARY,
  proto.pirem.ReceiveIrRequest,
  api_v1_irdata_pb.IrData,
  /**
   * @param {!proto.pirem.ReceiveIrRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  api_v1_irdata_pb.IrData.deserializeBinary
);


/**
 * @param {!proto.pirem.ReceiveIrRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.IrData)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.IrData>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.receiveIr =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/ReceiveIr',
      request,
      metadata || {},
      methodDescriptor_PiRemService_ReceiveIr,
      callback);
};


/**
 * @param {!proto.pirem.ReceiveIrRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.IrData>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.receiveIr =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/ReceiveIr',
      request,
      metadata || {},
      methodDescriptor_PiRemService_ReceiveIr);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.ListDevicesRequest,
 *   !proto.pirem.ListDevicesResponse>}
 */
const methodDescriptor_PiRemService_ListDevices = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/ListDevices',
  grpc.web.MethodType.UNARY,
  proto.pirem.ListDevicesRequest,
  proto.pirem.ListDevicesResponse,
  /**
   * @param {!proto.pirem.ListDevicesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pirem.ListDevicesResponse.deserializeBinary
);


/**
 * @param {!proto.pirem.ListDevicesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.ListDevicesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.ListDevicesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.listDevices =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/ListDevices',
      request,
      metadata || {},
      methodDescriptor_PiRemService_ListDevices,
      callback);
};


/**
 * @param {!proto.pirem.ListDevicesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.ListDevicesResponse>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.listDevices =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/ListDevices',
      request,
      metadata || {},
      methodDescriptor_PiRemService_ListDevices);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.GetDeviceRequest,
 *   !proto.pirem.Device>}
 */
const methodDescriptor_PiRemService_GetDevice = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/GetDevice',
  grpc.web.MethodType.UNARY,
  proto.pirem.GetDeviceRequest,
  api_v1_device_pb.Device,
  /**
   * @param {!proto.pirem.GetDeviceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  api_v1_device_pb.Device.deserializeBinary
);


/**
 * @param {!proto.pirem.GetDeviceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.Device)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.Device>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.getDevice =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/GetDevice',
      request,
      metadata || {},
      methodDescriptor_PiRemService_GetDevice,
      callback);
};


/**
 * @param {!proto.pirem.GetDeviceRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.Device>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.getDevice =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/GetDevice',
      request,
      metadata || {},
      methodDescriptor_PiRemService_GetDevice);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.CreateRemoteRequest,
 *   !proto.pirem.Remote>}
 */
const methodDescriptor_PiRemService_CreateRemote = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/CreateRemote',
  grpc.web.MethodType.UNARY,
  proto.pirem.CreateRemoteRequest,
  api_v1_remote_pb.Remote,
  /**
   * @param {!proto.pirem.CreateRemoteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  api_v1_remote_pb.Remote.deserializeBinary
);


/**
 * @param {!proto.pirem.CreateRemoteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.Remote)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.Remote>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.createRemote =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/CreateRemote',
      request,
      metadata || {},
      methodDescriptor_PiRemService_CreateRemote,
      callback);
};


/**
 * @param {!proto.pirem.CreateRemoteRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.Remote>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.createRemote =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/CreateRemote',
      request,
      metadata || {},
      methodDescriptor_PiRemService_CreateRemote);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.ListRemotesRequest,
 *   !proto.pirem.ListRemotesResponse>}
 */
const methodDescriptor_PiRemService_ListRemotes = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/ListRemotes',
  grpc.web.MethodType.UNARY,
  proto.pirem.ListRemotesRequest,
  proto.pirem.ListRemotesResponse,
  /**
   * @param {!proto.pirem.ListRemotesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pirem.ListRemotesResponse.deserializeBinary
);


/**
 * @param {!proto.pirem.ListRemotesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.ListRemotesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.ListRemotesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.listRemotes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/ListRemotes',
      request,
      metadata || {},
      methodDescriptor_PiRemService_ListRemotes,
      callback);
};


/**
 * @param {!proto.pirem.ListRemotesRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.ListRemotesResponse>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.listRemotes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/ListRemotes',
      request,
      metadata || {},
      methodDescriptor_PiRemService_ListRemotes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.GetRemoteRequest,
 *   !proto.pirem.Remote>}
 */
const methodDescriptor_PiRemService_GetRemote = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/GetRemote',
  grpc.web.MethodType.UNARY,
  proto.pirem.GetRemoteRequest,
  api_v1_remote_pb.Remote,
  /**
   * @param {!proto.pirem.GetRemoteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  api_v1_remote_pb.Remote.deserializeBinary
);


/**
 * @param {!proto.pirem.GetRemoteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.Remote)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.Remote>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.getRemote =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/GetRemote',
      request,
      metadata || {},
      methodDescriptor_PiRemService_GetRemote,
      callback);
};


/**
 * @param {!proto.pirem.GetRemoteRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.Remote>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.getRemote =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/GetRemote',
      request,
      metadata || {},
      methodDescriptor_PiRemService_GetRemote);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.UpdateRemoteRequest,
 *   !proto.pirem.Empty>}
 */
const methodDescriptor_PiRemService_UpdateRemote = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/UpdateRemote',
  grpc.web.MethodType.UNARY,
  proto.pirem.UpdateRemoteRequest,
  api_v1_empty_pb.Empty,
  /**
   * @param {!proto.pirem.UpdateRemoteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  api_v1_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.pirem.UpdateRemoteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.updateRemote =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/UpdateRemote',
      request,
      metadata || {},
      methodDescriptor_PiRemService_UpdateRemote,
      callback);
};


/**
 * @param {!proto.pirem.UpdateRemoteRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.Empty>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.updateRemote =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/UpdateRemote',
      request,
      metadata || {},
      methodDescriptor_PiRemService_UpdateRemote);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.DeleteRemoteRequest,
 *   !proto.pirem.Empty>}
 */
const methodDescriptor_PiRemService_DeleteRemote = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/DeleteRemote',
  grpc.web.MethodType.UNARY,
  proto.pirem.DeleteRemoteRequest,
  api_v1_empty_pb.Empty,
  /**
   * @param {!proto.pirem.DeleteRemoteRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  api_v1_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.pirem.DeleteRemoteRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.deleteRemote =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/DeleteRemote',
      request,
      metadata || {},
      methodDescriptor_PiRemService_DeleteRemote,
      callback);
};


/**
 * @param {!proto.pirem.DeleteRemoteRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.Empty>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.deleteRemote =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/DeleteRemote',
      request,
      metadata || {},
      methodDescriptor_PiRemService_DeleteRemote);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.GetButtonRequest,
 *   !proto.pirem.Button>}
 */
const methodDescriptor_PiRemService_GetButton = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/GetButton',
  grpc.web.MethodType.UNARY,
  proto.pirem.GetButtonRequest,
  api_v1_button_pb.Button,
  /**
   * @param {!proto.pirem.GetButtonRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  api_v1_button_pb.Button.deserializeBinary
);


/**
 * @param {!proto.pirem.GetButtonRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.Button)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.Button>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.getButton =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/GetButton',
      request,
      metadata || {},
      methodDescriptor_PiRemService_GetButton,
      callback);
};


/**
 * @param {!proto.pirem.GetButtonRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.Button>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.getButton =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/GetButton',
      request,
      metadata || {},
      methodDescriptor_PiRemService_GetButton);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.LearnIrDataRequest,
 *   !proto.pirem.Empty>}
 */
const methodDescriptor_PiRemService_LearnIrData = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/LearnIrData',
  grpc.web.MethodType.UNARY,
  proto.pirem.LearnIrDataRequest,
  api_v1_empty_pb.Empty,
  /**
   * @param {!proto.pirem.LearnIrDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  api_v1_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.pirem.LearnIrDataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.learnIrData =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/LearnIrData',
      request,
      metadata || {},
      methodDescriptor_PiRemService_LearnIrData,
      callback);
};


/**
 * @param {!proto.pirem.LearnIrDataRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.Empty>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.learnIrData =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/LearnIrData',
      request,
      metadata || {},
      methodDescriptor_PiRemService_LearnIrData);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.PushButtonRequest,
 *   !proto.pirem.Empty>}
 */
const methodDescriptor_PiRemService_PushButton = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/PushButton',
  grpc.web.MethodType.UNARY,
  proto.pirem.PushButtonRequest,
  api_v1_empty_pb.Empty,
  /**
   * @param {!proto.pirem.PushButtonRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  api_v1_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.pirem.PushButtonRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.pushButton =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/PushButton',
      request,
      metadata || {},
      methodDescriptor_PiRemService_PushButton,
      callback);
};


/**
 * @param {!proto.pirem.PushButtonRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.Empty>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.pushButton =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/PushButton',
      request,
      metadata || {},
      methodDescriptor_PiRemService_PushButton);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.pirem.GetIrDataRequest,
 *   !proto.pirem.GetIrDataResponse>}
 */
const methodDescriptor_PiRemService_GetIrData = new grpc.web.MethodDescriptor(
  '/pirem.PiRemService/GetIrData',
  grpc.web.MethodType.UNARY,
  proto.pirem.GetIrDataRequest,
  proto.pirem.GetIrDataResponse,
  /**
   * @param {!proto.pirem.GetIrDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.pirem.GetIrDataResponse.deserializeBinary
);


/**
 * @param {!proto.pirem.GetIrDataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.pirem.GetIrDataResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.pirem.GetIrDataResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.pirem.PiRemServiceClient.prototype.getIrData =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/pirem.PiRemService/GetIrData',
      request,
      metadata || {},
      methodDescriptor_PiRemService_GetIrData,
      callback);
};


/**
 * @param {!proto.pirem.GetIrDataRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.pirem.GetIrDataResponse>}
 *     Promise that resolves to the response
 */
proto.pirem.PiRemServicePromiseClient.prototype.getIrData =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/pirem.PiRemService/GetIrData',
      request,
      metadata || {},
      methodDescriptor_PiRemService_GetIrData);
};


module.exports = proto.pirem;
