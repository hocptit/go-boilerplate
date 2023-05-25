const http = require('http');

const targetUrl = 'http://<minikube-ip>'; // Replace with the target URL you want to forward traffic to

const server = http.createServer((req, res) => {
    const options = {
        hostname: new URL(targetUrl).hostname,
        port: 31186, // Replace with the target port if applicable
        path: req.url,
        method: req.method,
        headers: req.headers
    };

    const proxyReq = http.request(options, (proxyRes) => {
        res.writeHead(proxyRes.statusCode, proxyRes.headers);
        proxyRes.pipe(res);
    });

    req.pipe(proxyReq);
});

const port = 3000; // Specify the port on which the proxy server should listen
server.listen(port, () => {
    console.log(`Proxy server running at http://localhost:${port}`);
});