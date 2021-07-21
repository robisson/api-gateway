const database = {
  apis: [
    {
      api_id: '',
      product_request_policices: [{
        name: "seller_rate_limit",
        type: "rate_limit",
        properties: {
          windowMs: 15 * 60 * 1000,
          max: 5
        }
      }],
      api_request_policices: [],
      resource_request_policices: [],
      integrations: {
        type: "http",
        properties: {
          target: "http://localhost:3001/sellers",
          changeOrigin: true,
          pathRewrite: {
            [`^/recurryng-payments/v1/sellers`]: '',
          },
        }
      },
      product_response_policices: [],
      api_response_policices: [],
      resource_response_policices: [],
      api_path: '/recurryng-payments/v1/sellers',
      api_regex_pattern: '\\/recurryng-payments\\/v1\\/sellers$',
      method: 'GET',
      enviroments: ['development', 'stagging', 'production']
    },
    {
      api_id: '',
      product_request_policices: [{
        name: "seller_rate_limit",
        type: "rate_limit",
        properties: {
          windowMs: 15 * 60 * 1000,
          max: 5
        }
      }],
      api_request_policices: [],
      resource_request_policices: [],
      integrations: {
        type: "http",
        properties: {
          target: "http://localhost:3001/sellers",
          changeOrigin: true,
          pathRewrite: {
            [`^/recurryng-payments/v1/sellers`]: '',
          },
        }
      },
      product_response_policices: [],
      api_response_policices: [],
      resource_response_policices: [],
      api_path: '/recurryng-payments/v1/sellers/:seller_id',
      api_regex_pattern: '\\/recurryng-payments\\/v1\\/sellers\\/(?:([^\\/]+?))?$',
      method: 'GET',
      enviroments: ['development', 'stagging', 'production']
    },
    {
      api_id: '',
      product_request_policices: [],
      api_request_policices: [],
      resource_request_policices: [],
      integrations: {
        type: "http",
        properties: {
          target: "http://localhost:3001/sellers",
          changeOrigin: true,
          pathRewrite: {
            [`^/recurryng-payments/v1/sellers`]: '',
          },
        }
      },
      product_response_policices: [],
      api_response_policices: [],
      resource_response_policices: [],
      api_path: '/recurryng-payments/v1/sellers/:seller_id/products/:product_id',
      api_regex_pattern: '\\/recurryng-payments\\/v1\\/sellers\\/(?:([^\\/]+?))\\/products\\/(?:([^\\/]+?))?$',
      method: 'GET',
      enviroments: ['development', 'stagging', 'production']
    }
  ]
};

module.exports = database