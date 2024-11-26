import assert from "assert";
import { VITE_GRPC_URL } from "../env";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { UsersService } from "./users.service";

export class ApiService {
  grpcUrl: string;
  transport: GrpcWebFetchTransport;
  usersService: UsersService;
  constructor() {
    // assert(VITE_GRPC_URL, "VITE_GRPC_URL is not defined");
    this.grpcUrl = VITE_GRPC_URL;
    this.transport = new GrpcWebFetchTransport({baseUrl:this.grpcUrl+'/'});
    this.usersService = new UsersService(this.transport);
  }
}