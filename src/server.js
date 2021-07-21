require('dotenv').config({ path: __dirname + '/../.env' });

const express = require('express')
const cookieParser = require('cookie-parser')

const app = express();
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

const port = process.env.PORT || 3000;

const { ROUTES } = require("./routes");

app.use((req, res, next) => {
  require("./policies/logging/logging")(app);

  // query by access_token
  require("./policies/security/key-validation")(app);

  // flow
  require('./router')(app, ROUTES);

  next();
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})