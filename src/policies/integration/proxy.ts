import { createProxyMiddleware } from 'http-proxy-middleware';
import { Application, Request, Response, NextFunction } from "express"

const setupProxies = (app: Application, routes: any) => {
    routes.forEach((route: any) => {
        app.use((req: Request, res: Response, next: NextFunction) => {
            console.log("REQUEST FLOW")
            next();
        });

        app.use(route.url, createProxyMiddleware(route.proxy));

        app.use((req: Request, res: Response, next: NextFunction) => {
            console.log("RESPONSE FLOW")
            next();
        });
    })
}

export default setupProxies;