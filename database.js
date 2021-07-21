const database = {
  apis: [
    {
      api_id: '',
      product_request_policices: [],
      api_request_policices: [],
      resource_request_policices: [],
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
      product_request_policices: [],
      api_request_policices: [],
      resource_request_policices: [],
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