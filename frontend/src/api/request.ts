export type RequestBuilderParams = {
    path: string
    params: { [key: string]: string|number|boolean }
    headers: Headers
}

export type RequestBuilderBody<T> = {
    body: T
}

export class RequestBuilder {
    private readonly baseUrl: string
    private readonly headers: Headers
    constructor(baseUrl: string, headers: Headers) {
        this.baseUrl = baseUrl
        this.headers = headers
    }

    async get<T>(params: Partial<RequestBuilderParams>): Promise<{status: number, data: T|null, message: string}> {
        return this.requestWithOutBody<T>({...params, method: 'GET'})
    }

    async post<T>(params: RequestBuilderBody<T> & Partial<RequestBuilderParams>): Promise<{status: number, data: T|null, message: string}> {
        return this.requestWithBody<T>({...params, method: 'POST'})
    }

    async put<T>(params: RequestBuilderBody<T> & Partial<RequestBuilderParams>): Promise<{status: number, data: T|null, message: string}> {
        return this.requestWithBody<T>({...params, method: 'PUT'})
    }

    async delete(params: Partial<RequestBuilderParams>): Promise<{status: number, data: null, message: string}> {
        return this.requestWithOutBody<null>({...params, method: 'DELETE'})
    }

    async patch<T>(params: RequestBuilderBody<T> & Partial<RequestBuilderParams>): Promise<{status: number, data: T|null, message: string}> {
        return this.requestWithBody({...params, method: 'PATCH'})
    }

    private async requestWithBody<T>(params: RequestBuilderBody<T> & Partial<RequestBuilderParams> & {method: string}) {
        const response = await fetch(`${this.baseUrl}${computedPath(params.path??'', params.params??{})}`, {
            method: params.method,
            body: JSON.stringify(params.body),
            headers: computedHeaders(this.headers, params.headers??new Headers())
        })
        if (!response.ok) {
            return {status: response.status, data: null, message: await response.text()}
        }
        return {status: response.status, data: await response.json() as T, message: ""}
    }

    private async requestWithOutBody<T>(params: Partial<RequestBuilderParams> & {method: string}) {
        const response = await fetch(`${this.baseUrl}${computedPath(params.path??'', params.params??{})}`, {
            method: params.method,
            headers: computedHeaders(this.headers, params.headers??new Headers())
        });
        if (!response.ok) {
            return {status: response.status, data: null, message: await response.text()}
        }
        return {status: response.status, data: await response.json() as T, message: ""}
    }
}

const computedPath = (path: string, params: { [key: string]: string|number|boolean }) => {
    return `${path}${Object.entries(params).length > 0 ? 
        `?${Object.entries(params).map(([key, value]) => `${key}=${value}`).join('&')}` : ''}`
}

const computedHeaders = (RootHeaders: Headers, headers: Headers) => {
    return new Headers([...RootHeaders, ...headers, ...new Headers({credentials: 'include'})]);
}
