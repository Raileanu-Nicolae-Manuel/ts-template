import type { User } from "@/types/users"
import { RequestBuilder } from "../request"

class Users {
    private readonly request: RequestBuilder
    constructor(private readonly path: string, private readonly headers: Headers) {
        this.request = new RequestBuilder(path+"/users", headers)
    }

    async getUsers() {
        const {status, data, message} = await this.request.get<User[]>({})
        if (status !== 200) {
            return [];
        }
        return data?? [];
    }
}