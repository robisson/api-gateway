import * as morgan from "morgan";
import { Application } from "express"

const setupLogging = (app: Application) => {
    app.use(morgan('combined'));
}

module.exports = setupLogging