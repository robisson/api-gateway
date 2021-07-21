import * as rateLimit from "express-rate-limit";
import { Application } from "express"

const setupRateLimit = (app: Application, routes: any) => {
    routes.forEach((r: any) => {
        if (r.rateLimit) {
            app.use(r.url, rateLimit(r.rateLimit));
        }
    })
}

module.exports = setupRateLimit
