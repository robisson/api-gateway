import * as rateLimit from "express-rate-limit";
import { Application } from "express"

const setupRateLimit = (app: Application, route: any) => {
    [
        "product_request_policices", "api_request_policices", "resource_request_policices"
    ].forEach((policie: any) => {
        route[policie].forEach((policie: any) => {
            if (policie.type === "rate_limit") {
                app.use(route.api_path, rateLimit(policie.properties));
            }
        })
    });
}

export default setupRateLimit;
