import * as dotenv from 'dotenv'
import { Request, Response, NextFunction } from 'express'
import * as express from "express";
import router from './router';
import * as cookieParser from'cookie-parser';

dotenv.config({ path: __dirname + '/../.env' });

const app = express();
const port = process.env.PORT || 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

app.use((req: Request, res: Response, next: NextFunction) => {
  const logging = require("./policies/logging/logging").default;
  logging(app);

  // query by access_token
  const { apiKeyValidation } = require("./policies/security/key-validation");
  apiKeyValidation(app);

  // flow
  router(app);

  next();
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})