// @generated by protobuf-ts 2.9.4
// @generated from protobuf file "calculoator.proto" (package "calculator", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { CalculatorService } from "./calculoator";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { SumResponse } from "./calculoator";
import type { SumRequest } from "./calculoator";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service calculator.CalculatorService
 */
export interface ICalculatorServiceClient {
    /**
     * @generated from protobuf rpc: Sum(calculator.SumRequest) returns (calculator.SumResponse);
     */
    sum(input: SumRequest, options?: RpcOptions): UnaryCall<SumRequest, SumResponse>;
}
/**
 * @generated from protobuf service calculator.CalculatorService
 */
export class CalculatorServiceClient implements ICalculatorServiceClient, ServiceInfo {
    typeName = CalculatorService.typeName;
    methods = CalculatorService.methods;
    options = CalculatorService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: Sum(calculator.SumRequest) returns (calculator.SumResponse);
     */
    sum(input: SumRequest, options?: RpcOptions): UnaryCall<SumRequest, SumResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<SumRequest, SumResponse>("unary", this._transport, method, opt, input);
    }
}
