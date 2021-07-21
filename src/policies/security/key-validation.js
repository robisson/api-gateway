const axios = require('axios');
const { unauthenticatedResponse } = require('../../utils/response-functions');

const validateToken = async (access_token, res) => {
  const tokenValidationendpoint = process.env.TOKEN_VALIDATION_ENDPOINT;

  try {
    const { data } = await axios.get(tokenValidationendpoint, {
      headers: {
        'Authorization': access_token
      }
    })

    const client_id = data.preferred_username.replace('service-account-', '');

    return client_id;
  } catch (error) {
    if (error.response.status == 401) {
      unauthenticatedResponse(res, "Invalid access_token");
    }
  }
}

const apiKeyValidation = async (app) => {
  app.use(async (req, res, next) => {
    const { authorization: access_token } = req.headers;

    if (access_token === undefined) {
      unauthenticatedResponse(res, "access_token not found. An access_token is required");
    }

    const client_id = await validateToken(access_token, res);

    if (client_id !== undefined) {
      req.client_id = client_id;
      next();
    }

  })
}

module.exports = apiKeyValidation