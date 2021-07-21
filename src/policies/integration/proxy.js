const { createProxyMiddleware } = require('http-proxy-middleware');

const setupProxies = (app, routes) => {
    routes.forEach(route => {
        app.use((req, res, next) => {
            console.log("REQUEST FLOW")
            next();
        });

        app.use(route.url, createProxyMiddleware(route.proxy));

        app.use((req, res, next) => {
            console.log("RESPONSE FLOW")
            next();
        });
    })
}

module.exports = setupProxies