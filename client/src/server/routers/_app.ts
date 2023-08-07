import { credentials } from "@grpc/grpc-js";
import { procedure, router } from '../trpc'
import { Stock, TradingClient } from "~/grpc/trading";

const host = process.env.GRPC_HOST || '127.0.0.1';
const port = process.env.GRPC_PORT || '9095';
const creds = process.env.GRPC_INSECURE === 'true' ?
    credentials.createInsecure() :
    credentials.createSsl(Buffer.from(process.env.GRPC_CACERT || '', 'base64'));
const opts = process.env.GRPC_HOST_OVERRIDE ? {
    'grpc.ssl_target_name_override': process.env.GRPC_HOST_OVERRIDE,
    'grpc.default_authority': process.env.GRPC_HOST_OVERRIDE
} : undefined;

console.log('GRPC_HOST: ', host);
console.log('GRPC_PORT: ', port);
console.log('GRPC_INSECURE: ', process.env.GRPC_INSECURE || 'false');
console.log('GRPC_HOST_OVERRIDE: ', process.env.GRPC_HOST_OVERRIDE || '');
console.log('GRPC_OPTIONS: ', opts);

const trading = new TradingClient(`${host}:${port}`, creds, opts);

export const appRouter = router({
    getStockList: procedure.query(async (): Promise<Stock[]> => {
        const stocks: Promise<Stock[]> = new Promise((resolve, reject) => {
            trading.getStockList({}, (err, stockListResp) => {
                if (err) {
                    console.error(err);
                    reject(err);
                    return;
                }

                resolve(stockListResp.stockList);
            });
        });
        return stocks;
    }),
});

// export type definition of API
export type AppRouter = typeof appRouter;
