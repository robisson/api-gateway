import { createProxyMiddleware } from 'http-proxy-middleware';
import { Application } from "express"

const setupProxies = (app: Application, route: any) => {
    app.use(
        route.api_path,
        createProxyMiddleware(route.integrations.properties)
    );
}

export default setupProxies;