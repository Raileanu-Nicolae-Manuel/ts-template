import assert from "assert";
import { VITE_GRPC_URL } from "../env";
import { CalculatorServiceClient } from "./generate/calculoator.client";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";

export class ApiService {
  grpcUrl: string;
  calculatorService: CalculatorServiceClient;
  transport: GrpcWebFetchTransport;
  constructor() {
    console.log(VITE_GRPC_URL);
    // assert(VITE_GRPC_URL, "VITE_GRPC_URL is not defined");
    this.grpcUrl = VITE_GRPC_URL;
    this.transport = new GrpcWebFetchTransport({baseUrl:this.grpcUrl+'/'});
    this.calculatorService = new CalculatorServiceClient(this.transport);
  }
}