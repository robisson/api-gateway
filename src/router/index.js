const database = require('../../database');

const router = (app, ROUTES) => {
  app.use((req, res, next) => {
    //verify if route exists
    const apiRoute = database.apis.find(api => {
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
      require("../policies/mediation/ratelimit")(app, ROUTES);
      require("../policies/integration/proxy")(app, ROUTES);

      next();
    }
  })
}

module.exports = router