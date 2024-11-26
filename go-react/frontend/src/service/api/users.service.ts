import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { UsersServiceClient } from "./generate/users.client";
import { GetUserResponse, IdModel, LoginRequest, RegisterRequest, TokenResponse, UpdateUserRequest } from "./generate/users";

export class UsersService {
    usersService: UsersServiceClient;

    constructor(transport: GrpcWebFetchTransport) {
        this.usersService = new UsersServiceClient(transport);
    }

    async register(registerRequest: RegisterRequest): Promise<TokenResponse> {
        const request = await this.usersService.register(registerRequest);
        return request.response;
    }

    async login(loginRequest: LoginRequest): Promise<TokenResponse> {
        const request = await this.usersService.login(loginRequest);
        return request.response;
    }

    async getUserById(idModel: IdModel): Promise<GetUserResponse> {
        const request = await this.usersService.getUserById(idModel);
        return request.response;
    }

    async updateUser(updateUserRequest: UpdateUserRequest): Promise<GetUserResponse> {
        const request = await this.usersService.updateUser(updateUserRequest);
        return request.response;
    }

    async deleteUser(idModel: IdModel): Promise<IdModel> {
        const request = await this.usersService.deleteUser(idModel);
        return request.response;
    }
}
