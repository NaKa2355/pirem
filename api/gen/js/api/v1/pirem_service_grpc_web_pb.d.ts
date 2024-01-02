import * as grpcWeb from 'grpc-web';

import * as api_v1_button_pb from '../../api/v1/button_pb';
import * as api_v1_pirem_service_pb from '../../api/v1/pirem_service_pb';
import * as api_v1_device_pb from '../../api/v1/device_pb';
import * as api_v1_empty_pb from '../../api/v1/empty_pb';
import * as api_v1_irdata_pb from '../../api/v1/irdata_pb';
import * as api_v1_remote_pb from '../../api/v1/remote_pb';


export class PiRemServiceClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  sendIr(
    request: api_v1_pirem_service_pb.SendIrRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_pirem_service_pb.SendIrResponse) => void
  ): grpcWeb.ClientReadableStream<api_v1_pirem_service_pb.SendIrResponse>;

  receiveIr(
    request: api_v1_pirem_service_pb.ReceiveIrRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_irdata_pb.IrData) => void
  ): grpcWeb.ClientReadableStream<api_v1_irdata_pb.IrData>;

  listDevices(
    request: api_v1_pirem_service_pb.ListDevicesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_pirem_service_pb.ListDevicesResponse) => void
  ): grpcWeb.ClientReadableStream<api_v1_pirem_service_pb.ListDevicesResponse>;

  getDevice(
    request: api_v1_pirem_service_pb.GetDeviceRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_device_pb.Device) => void
  ): grpcWeb.ClientReadableStream<api_v1_device_pb.Device>;

  createRemote(
    request: api_v1_pirem_service_pb.CreateRemoteRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_remote_pb.Remote) => void
  ): grpcWeb.ClientReadableStream<api_v1_remote_pb.Remote>;

  listRemotes(
    request: api_v1_pirem_service_pb.ListRemotesRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_pirem_service_pb.ListRemotesResponse) => void
  ): grpcWeb.ClientReadableStream<api_v1_pirem_service_pb.ListRemotesResponse>;

  getRemote(
    request: api_v1_pirem_service_pb.GetRemoteRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_remote_pb.Remote) => void
  ): grpcWeb.ClientReadableStream<api_v1_remote_pb.Remote>;

  updateRemote(
    request: api_v1_pirem_service_pb.UpdateRemoteRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<api_v1_empty_pb.Empty>;

  deleteRemote(
    request: api_v1_pirem_service_pb.DeleteRemoteRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<api_v1_empty_pb.Empty>;

  getButton(
    request: api_v1_pirem_service_pb.GetButtonRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_button_pb.Button) => void
  ): grpcWeb.ClientReadableStream<api_v1_button_pb.Button>;

  learnIrData(
    request: api_v1_pirem_service_pb.LearnIrDataRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<api_v1_empty_pb.Empty>;

  pushButton(
    request: api_v1_pirem_service_pb.PushButtonRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_empty_pb.Empty) => void
  ): grpcWeb.ClientReadableStream<api_v1_empty_pb.Empty>;

  getIrData(
    request: api_v1_pirem_service_pb.GetIrDataRequest,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.RpcError,
               response: api_v1_pirem_service_pb.GetIrDataResponse) => void
  ): grpcWeb.ClientReadableStream<api_v1_pirem_service_pb.GetIrDataResponse>;

}

export class PiRemServicePromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; });

  sendIr(
    request: api_v1_pirem_service_pb.SendIrRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_pirem_service_pb.SendIrResponse>;

  receiveIr(
    request: api_v1_pirem_service_pb.ReceiveIrRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_irdata_pb.IrData>;

  listDevices(
    request: api_v1_pirem_service_pb.ListDevicesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_pirem_service_pb.ListDevicesResponse>;

  getDevice(
    request: api_v1_pirem_service_pb.GetDeviceRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_device_pb.Device>;

  createRemote(
    request: api_v1_pirem_service_pb.CreateRemoteRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_remote_pb.Remote>;

  listRemotes(
    request: api_v1_pirem_service_pb.ListRemotesRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_pirem_service_pb.ListRemotesResponse>;

  getRemote(
    request: api_v1_pirem_service_pb.GetRemoteRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_remote_pb.Remote>;

  updateRemote(
    request: api_v1_pirem_service_pb.UpdateRemoteRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_empty_pb.Empty>;

  deleteRemote(
    request: api_v1_pirem_service_pb.DeleteRemoteRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_empty_pb.Empty>;

  getButton(
    request: api_v1_pirem_service_pb.GetButtonRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_button_pb.Button>;

  learnIrData(
    request: api_v1_pirem_service_pb.LearnIrDataRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_empty_pb.Empty>;

  pushButton(
    request: api_v1_pirem_service_pb.PushButtonRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_empty_pb.Empty>;

  getIrData(
    request: api_v1_pirem_service_pb.GetIrDataRequest,
    metadata?: grpcWeb.Metadata
  ): Promise<api_v1_pirem_service_pb.GetIrDataResponse>;

}

