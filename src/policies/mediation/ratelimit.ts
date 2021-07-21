import * as rateLimit from "express-rate-limit";
import { Application } from "express"

const setupRateLimit = (app: Application, route: any) => {
    route.product_request_policices.forEach((policie: any) => {
        if (policie.type === "rate_limit") {
            app.use(route.api_path, rateLimit(policie.properties));
        }
    })
    
    route.api_request_policices.forEach((policie: any) => {
        if (policie.type === "rate_limit") {
            app.use(route.api_path, rateLimit(policie.properties));
        }
    })
    
    route.resource_request_policices.forEach((policie: any) => {
        if (policie.type === "rate_limit") {
            app.use(route.api_path, rateLimit(policie.properties));
        }
    })
}

export default setupRateLimit;
