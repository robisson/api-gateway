import { Application, Request, Response, NextFunction } from "express"

const database = require('../../database');

const router = (app: Application) => {
  app.use((req: Request, res: Response, next: NextFunction) => {
    //verify if route exists
    const apiRoute = database.apis.find((api: any) => {
      if (req.url.match(api.api_regex_pattern) !== null && api.method == req.method) {
        return true;
      }
      return false;
    })

    if (apiRoute === undefined) {
      res
        .status(404)
        .json({ message: "Not Found" })
    } else {
      //get all pocicies, products, apis and resources
      const ratelimit = require("../policies/mediation/ratelimit").default;
      ratelimit(app, apiRoute);

      if (apiRoute.integrations.type === "http") {
        const proxy = require("../policies/integration/proxy").default;
        proxy(app, apiRoute);
      }

      next();
    }
  })
}

export default router;