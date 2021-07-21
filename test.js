const url = 'api.sensedia.com/recurring-payments/v1/sellers/51acb0cb-3b77-4265-8117-5258864db4fb/products/51acb0cb-3b77-4265-8117-5258864db4fb';

const originalPattern = '/recurring-payments/v1/sellers/:seller_id/products/:product_id';

const paramsIdentified = originalPattern.match(/{([^}]*)}/g);

console.log(paramsIdentified);

const pattern = /\/recurring-payments\/v1\/sellers\/(?:([^\/]+?))\/products\/(?:([^\/]+?))?$/g


console.log(url.match(pattern) !== null);