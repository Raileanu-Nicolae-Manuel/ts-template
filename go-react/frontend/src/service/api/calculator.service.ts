import { SumRequest } from "./generate/calculoator";
import { CalculatorServiceClient } from "./generate/calculoator.client";
import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";

export class CalculatorService {
    calculatorService: CalculatorServiceClient;

    constructor(transport: GrpcWebFetchTransport) {
        this.calculatorService = new CalculatorServiceClient(transport);
    }

    async sum(a: number, b: number): Promise<number> {
        const request = await this.calculatorService.sum({firstNumber: a, secondNumber: b});
        return request.response.result;
    }
}