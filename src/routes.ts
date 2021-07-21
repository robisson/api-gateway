import * as yaml from 'js-yaml'
import * as fs from 'fs';

const loadConfig = (configFile: string, encode: BufferEncoding = 'utf8') => {
    try {
        return yaml.load(fs.readFileSync(configFile, encode));
    } catch (e) {
        console.log(e);
    }
}

const configApi = loadConfig(__dirname + "/../api-gateway.yaml");

const routes = [
    {
        url: '/recurryng-payments/v1/sellers',
        rateLimit: {
            windowMs: 15 * 60 * 1000,
            max: 5
        },
        proxy: {
            target: "http://localhost:3001/sellers",
            changeOrigin: true,
            pathRewrite: {
                [`^/recurryng-payments/v1/sellers`]: '',
            },
        }
    },
    {
        url: '/recurryng-payments/v1/sellers/:seller_id',
        rateLimit: {
            windowMs: 15 * 60 * 1000,
            max: 5
        },
        proxy: {
            target: "http://localhost:3001/sellers",
            changeOrigin: true,
            pathRewrite: {
                [`^/recurryng-payments/v1/sellers`]: '',
            },
        }
    }
]

exports.ROUTES = routes;
